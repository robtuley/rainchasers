package gauge

import (
	"bytes"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/pubsub"
)

type PubSubContext struct {
	GoogleContext context.Context
	TopicName     string
}

func NewPubSubContext(projectId string, topicName string) (*PubSubContext, error) {
	client, err := google.DefaultClient(context.Background(), pubsub.ScopePubSub)
	if err != nil {
		return nil, err
	}
	ctx := cloud.NewContext(projectId, client)

	exists, err := pubsub.TopicExists(ctx, topicName)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = pubsub.CreateTopic(ctx, topicName)
		if err != nil {
			return nil, err
		}
	}

	return &PubSubContext{
		GoogleContext: ctx,
		TopicName:     topicName,
	}, nil
}

func Publish(ctx *PubSubContext, snap Snapshot) error {
	bb, err := Encode(snap)
	if err != nil {
		return err
	}

	_, err = pubsub.Publish(ctx.GoogleContext, ctx.TopicName, &pubsub.Message{
		Data: bb.Bytes(),
	})
	return err
}

func Subscribe(ctx *PubSubContext, subLabel string) (<-chan Snapshot, <-chan error, error) {
	snapC := make(chan Snapshot, 100)
	errC := make(chan error, 10)
	ackDeadline := time.Second * 10
	subName := ctx.TopicName + "." + subLabel

	isSub, err := pubsub.SubExists(ctx.GoogleContext, subName)
	if err != nil {
		close(snapC)
		close(errC)
		return snapC, errC, err
	}
	if !isSub {
		err = pubsub.CreateSub(ctx.GoogleContext, subName, ctx.TopicName, ackDeadline, "")
		if err != nil {
			close(snapC)
			close(errC)
			return snapC, errC, err
		}
	}

	go func() {
		for {
			msgs, err := pubsub.PullWait(ctx.GoogleContext, subName, 10)
			if err != nil {
				errC <- err
				continue
			}

			for _, m := range msgs {
				snap, decodeErr := Decode(bytes.NewBuffer(m.Data))
				ackErr := pubsub.Ack(ctx.GoogleContext, subName, m.AckID)
				if decodeErr != nil {
					errC <- decodeErr
				} else {
					snapC <- snap
				}
				if ackErr != nil {
					errC <- ackErr
				}
			}
		}
	}()

	return snapC, errC, nil
}
