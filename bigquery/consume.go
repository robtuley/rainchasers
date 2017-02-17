package main

import (
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
)

func consumeTopics(projectId string, topicNames ...string) (<-chan gauge.Snapshot, <-chan error, error) {
	snapC := make(chan gauge.Snapshot)
	errC := make(chan error)

	for _, tn := range topicNames {
		topic, err := gauge.NewTopic(projectId, tn)
		if err != nil {
			return snapC, errC, err
		}
		sC, eC, err := gauge.Subscribe(topic, "bigq")
		if err != nil {
			return snapC, errC, err
		}
		go chainErrC(eC, errC)
		go chainSnapC(sC, snapC)
	}

	return snapC, errC, nil
}

func chainErrC(fromC <-chan error, toC chan<- error) {
	for err := range fromC {
		toC <- err
	}
}

func chainSnapC(fromC <-chan gauge.Snapshot, toC chan<- gauge.Snapshot) {
	for s := range fromC {
		toC <- s
	}
}
