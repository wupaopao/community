package community

type AckOrgGroupListByOrganizationID struct {
	Count uint32   `db:"Count"`
	List  []*Group `db:"List"`
}

func NewAckOrgGroupListByOrganizationID() *AckOrgGroupListByOrganizationID {
	return &AckOrgGroupListByOrganizationID{
		List: make([]*Group, 0),
	}
}

// 团购组织社区群列表
func (m *Proxy) OrgGroupListByOrganizationID(OrganizationID uint32,
) (*AckOrgGroupListByOrganizationID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgGroupListByOrganizationID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/community/org/group/list/:organization_id",
		nil,
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

type AckOrgGroupInfoByOrganizationIDByGroupID struct {
	Group
	ManagerIdCardNumber string `db:"ManagerIdCardNumber"`
	ManagerIdCardFront  string `db:"ManagerIdCardFront"`
	ManagerIdCardBack   string `db:"ManagerIdCardBack"`
	ManagerWxNickname   string `db:"ManagerWxNickname"`
	ManagerWxAvatar     string `db:"ManagerWxAvatar"`
}

func NewAckOrgGroupInfoByOrganizationIDByGroupID() *AckOrgGroupInfoByOrganizationIDByGroupID {
	return &AckOrgGroupInfoByOrganizationIDByGroupID{}
}

// 团购组织社群详情
func (m *Proxy) OrgGroupInfoByOrganizationIDByGroupID(GroupID uint32,
	OrganizationID uint32,
) (*AckOrgGroupInfoByOrganizationIDByGroupID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgGroupInfoByOrganizationIDByGroupID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/community/org/group/info/:organization_id/:group_id",
		nil,
		ack,
		map[string]interface{}{
			"group_id":        GroupID,
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
