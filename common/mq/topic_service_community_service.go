package mq

import (
	"github.com/mz-eco/mz/kafka"
	"github.com/mz-eco/mz/log"
)

const (
	TOPIC_SERVICE_COMMUNITY_SERVICE = "service-community-service"
)

const (
	IDENT_SERVICE_COMMUNITY_SERVICE_CHANGE_GROUP_AUDIT_STATE = "change_group_audit_state"
	IDENT_SERVICE_COMMUNITY_SERVICE_MODIFY_GROUP_INFO        = "modify_group_info"
)

var (
	topicServiceCommunityService *TopicServiceCommunityService = nil
)

func GetTopicServiceCommunityServce() (topic *TopicServiceCommunityService, err error) {
	if topicServiceCommunityService != nil {
		topic = topicServiceCommunityService
		return
	}

	producer, err := kafka.NewAsyncProducer()
	if err != nil {
		log.Warnf("new async producer failed. %s", err)
		return
	}

	topicServiceCommunityService = &TopicServiceCommunityService{
		Producer: producer,
	}

	topic = topicServiceCommunityService

	return
}

type TopicServiceCommunityService struct {
	Producer *kafka.AsyncProducer
}

func (m *TopicServiceCommunityService) send(ident string, msg interface{}) (err error) {
	err = m.Producer.SendMessage(TOPIC_SERVICE_COMMUNITY_SERVICE, ident, msg)
	if err != nil {
		log.Warnf("send topic message failed. %s", err)
		return
	}
	return
}

// 社群审核状态变更
type ChangeGroupAuditStateMessage struct {
	GroupId uint32
}

func (m *TopicServiceCommunityService) ChangeGroupAuditState(msg *ChangeGroupAuditStateMessage) (err error) {
	return m.send(IDENT_SERVICE_COMMUNITY_SERVICE_CHANGE_GROUP_AUDIT_STATE, msg)
}

// 社群名称地址发生更改
type ModifyGroupInfoMessage struct {
	GroupId uint32
	Values  map[string]interface{}
}

func (m *TopicServiceCommunityService) ModifyGroupInfo(msg *ModifyGroupInfoMessage) (err error) {
	return m.send(IDENT_SERVICE_COMMUNITY_SERVICE_MODIFY_GROUP_INFO, msg)
}
