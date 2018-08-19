package mq

import (
	"github.com/mz-eco/mz/kafka"
	"github.com/mz-eco/mz/log"
)

func NewSubscriber() (subscriber *kafka.Subscriber, err error) {
	subscriber = &kafka.Subscriber{}
	topicHandler, err := NewTopicUserServiceHandler()
	if err != nil {
		log.Warnf("new topic user service handler failed. %s", err)
		return
	}

	subscriber.TopicHandlers = append(subscriber.TopicHandlers, topicHandler)
	return
}
