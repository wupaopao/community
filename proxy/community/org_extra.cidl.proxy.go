package community

type AskOrgGroupEditByOrganizationIDByGroupID struct {
	GroupName           string `binding:"required,lte=128" db:"GroupName"`
	ManagerName         string `binding:"required,lte=64" db:"ManagerName"`
	GroupAddress        string `binding:"required,lte=255" db:"GroupAddress"`
	ManagerIdCardNumber string `binding:"lte=18" db:"ManagerIdCardNumber"`
	ManagerIdCardFront  string `binding:"lte=255" db:"ManagerIdCardFront"`
	ManagerIdCardBack   string `binding:"lte=255" db:"ManagerIdCardBack"`
}

func NewAskOrgGroupEditByOrganizationIDByGroupID() *AskOrgGroupEditByOrganizationIDByGroupID {
	return &AskOrgGroupEditByOrganizationIDByGroupID{}
}

// 编辑
func (m *Proxy) OrgGroupEditByOrganizationIDByGroupID(OrganizationID uint32,
	GroupID uint32,
	ask *AskOrgGroupEditByOrganizationIDByGroupID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/org/group/edit/:organization_id/:group_id",
		ask,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"group_id":        GroupID,
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

type AskOrgGroupChangeManagerByOrganizationIDByGroupID struct {
	NewManagerUid string `binding:"required" db:"NewManagerUid"`
}

func NewAskOrgGroupChangeManagerByOrganizationIDByGroupID() *AskOrgGroupChangeManagerByOrganizationIDByGroupID {
	return &AskOrgGroupChangeManagerByOrganizationIDByGroupID{}
}

// 更改管理员
func (m *Proxy) OrgGroupChangeManagerByOrganizationIDByGroupID(OrganizationID uint32,
	GroupID uint32,
	ask *AskOrgGroupChangeManagerByOrganizationIDByGroupID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/org/group/change_manager/:organization_id/:group_id",
		ask,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"group_id":        GroupID,
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

type AskOrgGroupAuditByOrganizationIDByGroupID struct {
	AuditState GroupAuditState `binding:"required" db:"AuditState"`
}

func NewAskOrgGroupAuditByOrganizationIDByGroupID() *AskOrgGroupAuditByOrganizationIDByGroupID {
	return &AskOrgGroupAuditByOrganizationIDByGroupID{}
}

// 审核
func (m *Proxy) OrgGroupAuditByOrganizationIDByGroupID(OrganizationID uint32,
	GroupID uint32,
	ask *AskOrgGroupAuditByOrganizationIDByGroupID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/org/group/audit/:organization_id/:group_id",
		ask,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"group_id":        GroupID,
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
