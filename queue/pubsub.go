package queue

import (
	"bytes"
	"context"
	"errors"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/rainchasers/com.rainchasers.gauge/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/report"
)

// Topic encapsulates the message queue topic
type Topic struct {
	ProjectID string
	pubSub    *pubsub.Topic
}

// Stop cleanly closes the topic
func (t *Topic) Stop() {
	if t.pubSub != nil {
		t.pubSub.Stop()
	}
}

// New creates a message queue topic
func New(d *daemon.Supervisor, projectID string, topicName string) (t *Topic, err error) {
	ctx, cancel := context.WithTimeout(d.Context, 40*time.Second)
	ctx = d.Log.StartSpan(ctx, "queue.connected")
	defer func() {
		d.Log.EndSpan(ctx, err, report.Data{
			"project_id": projectID,
			"topic_name": topicName,
		})
		cancel()
	}()

	if len(projectID) == 0 {
		return &Topic{}, nil
	}

	client, err := pubsub.NewClient(ctx, projectID)
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
		ProjectID: projectID,
		pubSub:    topic,
	}, nil
}

// Publish writes an AVRO encoded Snapshot to the topic
func (t *Topic) Publish(d *daemon.Supervisor, s *gauge.Snapshot) (err error) {
	ctx, cancel := context.WithTimeout(d.Context, 20*time.Second)
	ctx = d.Log.StartSpan(ctx, "queue.published")
	defer func() {
		if d.Context.Err() == nil {
			// end span only if not interrupted by shutdown
			d.Log.EndSpan(ctx, err, report.Data{
				"station": s.Station.UUID,
				"count":   len(s.Readings),
			})
		}
		cancel()
	}()

	bb := bytes.NewBuffer([]byte{})

	err = s.Encode(bb)
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

// Subscribe reads AVRO encoded snapshots from the topic and decodes them
//
// Note a zero length consumerGroup means auto-generate the pubsub subscription
// string and delete once done.
func (t *Topic) Subscribe(ctx context.Context, consumerGroup string, fn func(s *gauge.Snapshot, err error)) error {
	const ackDeadline = time.Second * 20

	if t.pubSub == nil {
		return errors.New("Topic has no project ID")
	}

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
		cfg := pubsub.SubscriptionConfig{
			Topic:       t.pubSub,
			AckDeadline: ackDeadline,
		}
		sub, err = client.CreateSubscription(ctx, subName, cfg)
		if err != nil {
			return err
		}
	}
	if deleteSubOnComplete {
		defer sub.Delete(context.Background())
	}

	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		s := gauge.Snapshot{}
		err := s.Decode(bytes.NewBuffer(m.Data))
		fn(&s, err)
		m.Ack()
	})
	if err != nil {
		return err
	}

	return nil
}
