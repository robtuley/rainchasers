package gauge

import (
	"bytes"
	"time"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type Topic struct {
	ProjectID string
	PubSub    *pubsub.Topic
}

func NewTopic(projectId string, topicName string) (*Topic, error) {
	client, err := pubsub.NewClient(context.Background(), projectId)
	if err != nil {
		return nil, err
	}

	topic := client.Topic(topicName)
	exists, err := topic.Exists(context.Background())
	if err != nil {
		return nil, err
	}
	if !exists {
		topic, err = client.NewTopic(context.Background(), topicName)
		if err != nil {
			return nil, err
		}
	}

	return &Topic{
		ProjectID: projectId,
		PubSub:    topic,
	}, nil
}

func Publish(topic *Topic, snap Snapshot) error {
	bb, err := Encode(snap)
	if err != nil {
		return err
	}

	_, err = topic.PubSub.Publish(context.Background(), &pubsub.Message{
		Data: bb.Bytes(),
	})
	return err
}

func Subscribe(topic *Topic, consumerGroup string) (<-chan Snapshot, <-chan error, error) {
	snapC := make(chan Snapshot, 100)
	errC := make(chan error, 10)
	ackDeadline := time.Second * 10
	subName := topic.PubSub.Name() + "." + consumerGroup
	client, err := pubsub.NewClient(context.Background(), topic.ProjectID)
	if err != nil {
		close(snapC)
		close(errC)
		return snapC, errC, err
	}

	sub := client.Subscription(subName)
	exists, err := sub.Exists(context.Background())
	if err != nil {
		close(snapC)
		close(errC)
		return snapC, errC, err
	}
	if !exists {
		sub, err = client.NewSubscription(context.Background(), subName, topic.PubSub, ackDeadline, nil)
		if err != nil {
			close(snapC)
			close(errC)
			return snapC, errC, err
		}
	}

	go func() {
		for {
			it, err := sub.Pull(context.Background())
			if err != nil {
				errC <- err
				continue
			}

			for i := 0; i < 100; i++ {
				m, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					errC <- err
					continue
				}
				snap, decodeErr := Decode(bytes.NewBuffer(m.Data))
				if decodeErr != nil {
					errC <- decodeErr
				} else {
					snapC <- snap
				}
				m.Done(true)
			}

			it.Stop()
		}
	}()

	return snapC, errC, nil
}
