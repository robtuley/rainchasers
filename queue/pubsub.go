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
	pubSub    *pubsub.Topic
}

// defer topic.Stop()
func (t *Topic) Stop() {
	if t.pubSub != nil {
		t.pubSub.Stop()
	}
}

func New(ctx context.Context, projectId string, topicName string) (*Topic, error) {
	if len(projectId) == 0 {
		return &Topic{}, nil
	}

	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}

	topic := client.Topic(topicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		topic, err = client.CreateTopic(ctx, topicName)
		if err != nil {
			return nil, err
		}
	}

	return &Topic{
		ProjectID: projectId,
		pubSub:    topic,
	}, nil
}

func (t *Topic) Publish(ctx context.Context, s *gauge.Snapshot) error {
	bb, err := Encode(s)
	if err != nil {
		return err
	}

	if t.pubSub == nil {
		return nil
	}

	result := t.pubSub.Publish(ctx, &pubsub.Message{
		Data: bb.Bytes(),
	})
	_, err = result.Get(ctx)

	return err
}

// a zero length consumerGroup means auto-generate and delete once done
func (t *Topic) Subscribe(ctx context.Context, consumerGroup string, fn func(s *gauge.Snapshot, err error)) error {
	const ackDeadline = time.Second * 10

	deleteSubOnComplete := len(consumerGroup) == 0
	if deleteSubOnComplete {
		consumerGroup = time.Now().Format("v2006-01-02-15-04-05.999999")
	}
	subName := t.pubSub.ID() + "." + consumerGroup

	client, err := pubsub.NewClient(ctx, t.ProjectID)
	if err != nil {
		return err
	}

	sub := client.Subscription(subName)
	exists, err := sub.Exists(ctx)
	if err != nil {
		return err
	}
	if !exists {
		sub, err = client.CreateSubscription(ctx, subName, t.pubSub, ackDeadline, nil)
		if err != nil {
			return err
		}
	}
	if deleteSubOnComplete {
		defer sub.Delete(context.Background())
	}

	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		s, err := Decode(bytes.NewBuffer(m.Data))
		fn(s, err)
		m.Ack()
	})
	if err != nil {
		return err
	}

	return nil
}
