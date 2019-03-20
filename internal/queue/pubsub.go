package queue

import (
	"bytes"
	"context"
	"time"

	"cloud.google.com/go/pubsub"
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
func New(ctx context.Context, projectID string, topicName string) (*Topic, report.Span) {
	ctx, cancel := context.WithTimeout(ctx, 40*time.Second)
	defer cancel()

	span := report.StartSpan("topic.connected").Field("project_id", projectID).Field("topic_name", topicName)

	if len(projectID) == 0 {
		return &Topic{}, span.End()
	}

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, span.End(err)
	}

	topic := client.Topic(topicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return nil, span.End(err)
	}
	if !exists {
		cSpan := report.StartSpan("topic.created")
		topic, err = client.CreateTopic(ctx, topicName)
		span = span.Child(cSpan.End(err))
		if err != nil {
			return nil, span.End()
		}
	}

	return &Topic{
		ProjectID: projectID,
		pubSub:    topic,
	}, span.End()
}

// Publish writes an AVRO encoded Snapshot to the topic
func (t *Topic) Publish(ctx context.Context, s *gauge.Snapshot) report.Span {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	span := report.StartSpan("snapshot.published")
	span = span.Field("station", s.Station.AliasURL)
	span = span.Field("count_readings", len(s.Readings))

	if s.TraceID == "" {
		// if there is no predefined existing trace info,
		// assume this span tracer will be the trace generator
		s.TraceID = span.TraceID()
	}
	s.ProcessingTime = time.Now()

	bb := bytes.NewBuffer([]byte{})
	err := s.Encode(bb)
	if err != nil {
		return span.End(err)
	}

	if t.pubSub == nil {
		return span.End()
	}

	result := t.pubSub.Publish(ctx, &pubsub.Message{
		Data: bb.Bytes(),
	})
	_, err = result.Get(ctx)

	return span.End(err)
}

// Subscribe reads AVRO encoded snapshots from the topic and decodes them
//
// Note a zero length consumerGroup means auto-generate the pubsub subscription
// string and delete once done.
func (t *Topic) Subscribe(ctx context.Context, consumerGroup string,
	fn func(ctx context.Context, err error, s *gauge.Snapshot) error) error {
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

	return sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		ctx, cancel := context.WithTimeout(ctx, ackDeadline-time.Second)
		defer cancel()

		s := gauge.Snapshot{}
		err := s.Decode(bytes.NewBuffer(m.Data))
		err = fn(ctx, err, &s)
		if err != nil {
			m.Nack() // speed up message redelivery
			return
		}
		m.Ack()
	})
}
