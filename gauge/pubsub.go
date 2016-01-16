package gauge

import (
	"bytes"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/pubsub"
)

type Context struct {
	PubSubContext context.Context
	TopicName     string
}

func NewContext(projectId string, topicName string) (Context, error) {
	client, err := google.DefaultClient(context.Background(), pubsub.ScopePubSub)
	if err != nil {
		return Context{}, err
	}
	ctx := cloud.NewContext(projectId, client)

	exists, err := pubsub.TopicExists(ctx, topicName)
	if err != nil {
		return Context{}, err
	}
	if !exists {
		err = pubsub.CreateTopic(ctx, topicName)
		if err != nil {
			return Context{}, err
		}
	}

	return Context{
		PubSubContext: ctx,
		TopicName:     topicName,
	}, nil
}

func Publish(ctx Context, snap Snapshot) error {
	bb, err := Encode(snap)
	if err != nil {
		return err
	}

	_, err = pubsub.Publish(ctx.PubSubContext, ctx.TopicName, &pubsub.Message{
		Data: bb.Bytes(),
	})
	return err
}

func Consume(ctx Context, subName string) (<-chan Snapshot, <-chan error, error) {
	snapC := make(chan Snapshot, 100)
	errC := make(chan error, 10)
	ackDeadline := time.Second * 10

	isSub, err := pubsub.SubExists(ctx.PubSubContext, subName)
	if err != nil {
		close(snapC)
		close(errC)
		return snapC, errC, err
	}
	if !isSub {
		err = pubsub.CreateSub(ctx.PubSubContext, subName, ctx.TopicName, ackDeadline, "")
		if err != nil {
			close(snapC)
			close(errC)
			return snapC, errC, err
		}
	}

	go func() {
		for {
			msgs, err := pubsub.PullWait(ctx.PubSubContext, subName, 10)
			if err != nil {
				errC <- err
				continue
			}

			for _, m := range msgs {
				snap, decodeErr := Decode(bytes.NewBuffer(m.Data))
				ackErr := pubsub.Ack(ctx.PubSubContext, subName, m.AckID)
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
