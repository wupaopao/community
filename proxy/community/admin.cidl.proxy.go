package community

type AckAdminGroupListByOrganizationID struct {
	Count uint32   `db:"Count"`
	List  []*Group `db:"List"`
}

func NewAckAdminGroupListByOrganizationID() *AckAdminGroupListByOrganizationID {
	return &AckAdminGroupListByOrganizationID{
		List: make([]*Group, 0),
	}
}

// 团购组织社区群列表
func (m *Proxy) AdminGroupListByOrganizationID(OrganizationID uint32,
) (*AckAdminGroupListByOrganizationID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminGroupListByOrganizationID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/community/admin/group/list/:organization_id",
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

type AckAdminGroupInfoByOrganizationIDByGroupID struct {
	Group
	ManagerIdCardNumber string `db:"ManagerIdCardNumber"`
	ManagerIdCardFront  string `db:"ManagerIdCardFront"`
	ManagerIdCardBack   string `db:"ManagerIdCardBack"`
	ManagerWxNickname   string `db:"ManagerWxNickname"`
	ManagerWxAvatar     string `db:"ManagerWxAvatar"`
}

func NewAckAdminGroupInfoByOrganizationIDByGroupID() *AckAdminGroupInfoByOrganizationIDByGroupID {
	return &AckAdminGroupInfoByOrganizationIDByGroupID{}
}

// 团购组织社群详情
func (m *Proxy) AdminGroupInfoByOrganizationIDByGroupID(GroupID uint32,
	OrganizationID uint32,
) (*AckAdminGroupInfoByOrganizationIDByGroupID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminGroupInfoByOrganizationIDByGroupID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/community/admin/group/info/:organization_id/:group_id",
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
