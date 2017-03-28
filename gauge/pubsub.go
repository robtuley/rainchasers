package gauge

import (
	"bytes"
	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"time"
)

// defer topic.Stop()
type Topic struct {
	ProjectID string
	PubSub    *pubsub.Topic
}

func (t *Topic) Stop() {
	t.PubSub.Stop()
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
		topic, err = client.CreateTopic(context.Background(), topicName)
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
	bb, err := EncodeSnapshot(snap)
	if err != nil {
		return err
	}

	result := topic.PubSub.Publish(context.Background(), &pubsub.Message{
		Data: bb.Bytes(),
	})
	_, err = result.Get(context.Background())

	return err
}

func Subscribe(topic *Topic, consumerGroup string) (<-chan Snapshot, <-chan error, error) {
	snapC := make(chan Snapshot, 100)
	errC := make(chan error, 10)
	ackDeadline := time.Second * 10
	subName := topic.PubSub.ID() + "." + consumerGroup
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
		sub, err = client.CreateSubscription(context.Background(), subName, topic.PubSub, ackDeadline, nil)
		if err != nil {
			close(snapC)
			close(errC)
			return snapC, errC, err
		}
	}

	err = sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		snap, decodeErr := DecodeSnapshot(bytes.NewBuffer(m.Data))
		if decodeErr != nil {
			errC <- decodeErr
		} else {
			snapC <- snap
		}
		m.Ack()
	})
	if err != nil {
		close(snapC)
		close(errC)
		return snapC, errC, err
	}

	return snapC, errC, nil
}
