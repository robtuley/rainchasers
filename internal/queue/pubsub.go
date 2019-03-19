package queue

import (
	"bytes"
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
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
func New(ctx context.Context, d *daemon.Supervisor, projectID string, topicName string) (t *Topic, err error) {
	ctx, cancel := context.WithTimeout(ctx, 40*time.Second)
	ctx = d.StartSpan(ctx, "topic.connected")
	defer func() {
		d.EndSpan(ctx, err, report.Data{
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
func (t *Topic) Publish(ctx context.Context, d *daemon.Supervisor, s *gauge.Snapshot) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	ctx = d.StartSpan(ctx, "snapshot.published")
	defer func() {
		d.EndSpan(ctx, err, report.Data{
			"station":        s.Station.AliasURL,
			"count_readings": len(s.Readings),
		})
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
func (t *Topic) Subscribe(ctx context.Context, d *daemon.Supervisor, consumerGroup string,
	fn func(ctx context.Context, d *daemon.Supervisor, s *gauge.Snapshot) error) error {
	const ackDeadline = time.Second * 20

	if t.pubSub == nil {
		// simulate a subscription i.e. wait for ctx cancel then return
		<-ctx.Done()
		return nil
	}

	isDeleteSubOnComplete := (consumerGroup == "")
	if isDeleteSubOnComplete {
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
	if isDeleteSubOnComplete {
		defer sub.Delete(context.Background())
	}

	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		ctx, cancel := context.WithTimeout(ctx, ackDeadline-time.Second)
		defer cancel()

		s := gauge.Snapshot{}
		err := s.Decode(bytes.NewBuffer(m.Data))
		if err != nil {
			m.Ack() // ack corrupted message to prevent redelivery
			d.Action("snapshot.corrupted", report.Data{
				"error": err.Error(),
			})
			return
		}

		ctx = d.ContinueTrace(ctx, s.TraceID)
		ctx = d.StartSpan(ctx, "snapshot.received")
		err = fn(ctx, d, &s)
		d.EndSpan(ctx, err, report.Data{
			"count_readings": len(s.Readings),
		})

		if err != nil {
			m.Nack() // speed up message redelivery
			d.Action("snapshot.timeout", report.Data{
				"station": s.Station.AliasURL,
				"error":   err.Error(),
			})
			return
		}
		m.Ack()
	})
	if err != nil {
		return err
	}

	return nil
}
