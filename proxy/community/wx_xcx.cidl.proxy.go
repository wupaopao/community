package community

type AckWxXcxGroupInfoByGroupID struct {
	Group
	ManagerAvatar   string `db:"ManagerAvatar"`
	ManagerNickname string `db:"ManagerNickname"`
}

func NewAckWxXcxGroupInfoByGroupID() *AckWxXcxGroupInfoByGroupID {
	return &AckWxXcxGroupInfoByGroupID{}
}

// 获取社群信息
func (m *Proxy) WxXcxGroupInfoByGroupID(GroupID uint32,
) (*AckWxXcxGroupInfoByGroupID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckWxXcxGroupInfoByGroupID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/community/wx_xcx/group/info/:group_id",
		nil,
		ack,
		map[string]interface{}{
			"group_id": GroupID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

type AskWxXcxGroupAddByOrganizationID struct {
	ManagerName   string `binding:"required,lte=64" db:"ManagerName"`
	ManagerMobile string `binding:"required,numeric" db:"ManagerMobile"`
	Address       string `binding:"required,lte=255" db:"Address"`
	PostCode      string `binding:"required,numeric,len=6" db:"PostCode"`
}

func NewAskWxXcxGroupAddByOrganizationID() *AskWxXcxGroupAddByOrganizationID {
	return &AskWxXcxGroupAddByOrganizationID{}
}

// 添加社群
func (m *Proxy) WxXcxGroupAddByOrganizationID(OrganizationID uint32,
	ask *AskWxXcxGroupAddByOrganizationID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/wx_xcx/group/add/:organization_id",
		ask,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}
