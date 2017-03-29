package queue

import (
	"bytes"
	"cloud.google.com/go/pubsub"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"golang.org/x/net/context"
	"time"
)

type Topic struct {
	ProjectID string
	PubSub    *pubsub.Topic
}

// defer topic.Stop()
func (t *Topic) Stop() {
	t.PubSub.Stop()
}

func New(projectId string, topicName string) (*Topic, error) {
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

func (t *Topic) Publish(s *gauge.Snapshot) error {
	bb, err := Encode(s)
	if err != nil {
		return err
	}

	result := t.PubSub.Publish(context.Background(), &pubsub.Message{
		Data: bb.Bytes(),
	})
	_, err = result.Get(context.Background())

	return err
}

func (t *Topic) Subscribe(consumerGroup string, fn func(s *gauge.Snapshot, err error)) error {
	const ackDeadline = time.Second * 10
	subName := t.PubSub.ID() + "." + consumerGroup

	client, err := pubsub.NewClient(context.Background(), t.ProjectID)
	if err != nil {
		return err
	}

	sub := client.Subscription(subName)
	exists, err := sub.Exists(context.Background())
	if err != nil {
		return err
	}
	if !exists {
		sub, err = client.CreateSubscription(context.Background(), subName, t.PubSub, ackDeadline, nil)
		if err != nil {
			return err
		}
	}

	err = sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		s, err := Decode(bytes.NewBuffer(m.Data))
		fn(s, err)
		m.Ack()
	})
	if err != nil {
		return err
	}

	return nil
}
