package community

type AskOrgGroupAddByOrganizationID struct {
	GroupName           string `binding:"required,lte=128" db:"GroupName"`
	GroupAddress        string `binding:"required,lte=255" db:"GroupAddress"`
	GroupPostCode       string `binding:"lte=6" db:"GroupPostCode"`
	ManagerName         string `binding:"required,lte=64" db:"ManagerName"`
	ManagerMobile       string `binding:"required,numeric" db:"ManagerMobile"`
	ManagerIdCardNumber string `binding:"lte=18" db:"ManagerIdCardNumber"`
	ManagerIdCardFront  string `binding:"lte=255" db:"ManagerIdCardFront"`
	ManagerIdCardBack   string `binding:"lte=255" db:"ManagerIdCardBack"`
}

func NewAskOrgGroupAddByOrganizationID() *AskOrgGroupAddByOrganizationID {
	return &AskOrgGroupAddByOrganizationID{}
}

// 添加
func (m *Proxy) OrgGroupAddByOrganizationID(OrganizationID uint32,
	ask *AskOrgGroupAddByOrganizationID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/org/group/add/:organization_id",
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

type AskOrgGroupEditByOrganizationIDByGroupID struct {
	GroupName           string `binding:"required,lte=128" db:"GroupName"`
	GroupAddress        string `binding:"required,lte=255" db:"GroupAddress"`
	GroupPostCode       string `binding:"lte=6" db:"GroupPostCode"`
	ManagerName         string `binding:"required,lte=64" db:"ManagerName"`
	ManagerMobile       string `binding:"required,numeric" db:"ManagerMobile"`
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

type AskOrgTeamAddByOrganizationID struct {
	TeamName string `binding:"required,lte=128" db:"TeamName"`
}

func NewAskOrgTeamAddByOrganizationID() *AskOrgTeamAddByOrganizationID {
	return &AskOrgTeamAddByOrganizationID{}
}

// 添加社团群组
func (m *Proxy) OrgTeamAddByOrganizationID(OrganizationID uint32,
	ask *AskOrgTeamAddByOrganizationID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/org/team/add/:organization_id",
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

type AskOrgTeamEditByOrganizationIDByTeamID struct {
	TeamName string `binding:"required,lte=128" db:"TeamName"`
}

func NewAskOrgTeamEditByOrganizationIDByTeamID() *AskOrgTeamEditByOrganizationIDByTeamID {
	return &AskOrgTeamEditByOrganizationIDByTeamID{}
}

// 编辑社团群组
func (m *Proxy) OrgTeamEditByOrganizationIDByTeamID(OrganizationID uint32,
	TeamID uint32,
	ask *AskOrgTeamEditByOrganizationIDByTeamID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/org/team/edit/:organization_id/:team_id",
		ask,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"team_id":         TeamID,
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

// 删除社团群组
func (m *Proxy) OrgTeamDeleteByOrganizationIDByTeamID(OrganizationID uint32,
	TeamID uint32,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/org/team/delete/:organization_id/:team_id",
		nil,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"team_id":         TeamID,
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

type AskOrgTeamBindGroupByOrganizationIDByTeamID struct {
	GroupIDs []uint32 `db:"GroupIDs"`
}

func NewAskOrgTeamBindGroupByOrganizationIDByTeamID() *AskOrgTeamBindGroupByOrganizationIDByTeamID {
	return &AskOrgTeamBindGroupByOrganizationIDByTeamID{
		GroupIDs: make([]uint32, 0),
	}
}

// 群组关联社团
func (m *Proxy) OrgTeamBindGroupByOrganizationIDByTeamID(OrganizationID uint32,
	TeamID uint32,
	ask *AskOrgTeamBindGroupByOrganizationIDByTeamID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/org/team/bind/group/:organization_id/:team_id",
		ask,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"team_id":         TeamID,
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

type AckOrgTeamGroupBindedListByOrganizationIDByTeamID struct {
	Groups []*SubGroup `db:"Groups"`
}

func NewAckOrgTeamGroupBindedListByOrganizationIDByTeamID() *AckOrgTeamGroupBindedListByOrganizationIDByTeamID {
	return &AckOrgTeamGroupBindedListByOrganizationIDByTeamID{
		Groups: make([]*SubGroup, 0),
	}
}

// 群组已关联社团列表
func (m *Proxy) OrgTeamGroupBindedListByOrganizationIDByTeamID(OrganizationID uint32,
	TeamID uint32,
) (*AckOrgTeamGroupBindedListByOrganizationIDByTeamID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgTeamGroupBindedListByOrganizationIDByTeamID
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/org/team/group/binded/list/:organization_id/:team_id",
		nil,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"team_id":         TeamID,
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

type AskOrgTeamGroupUnbindedListByOrganizationIDByTeamID struct {
	Key string `db:"Key"`
}

func NewAskOrgTeamGroupUnbindedListByOrganizationIDByTeamID() *AskOrgTeamGroupUnbindedListByOrganizationIDByTeamID {
	return &AskOrgTeamGroupUnbindedListByOrganizationIDByTeamID{}
}

type AckOrgTeamGroupUnbindedListByOrganizationIDByTeamID struct {
	Groups []*SubGroup `db:"Groups"`
}

func NewAckOrgTeamGroupUnbindedListByOrganizationIDByTeamID() *AckOrgTeamGroupUnbindedListByOrganizationIDByTeamID {
	return &AckOrgTeamGroupUnbindedListByOrganizationIDByTeamID{
		Groups: make([]*SubGroup, 0),
	}
}

// 群组未关联社团列表
func (m *Proxy) OrgTeamGroupUnbindedListByOrganizationIDByTeamID(OrganizationID uint32,
	TeamID uint32,
	ask *AskOrgTeamGroupUnbindedListByOrganizationIDByTeamID,
) (*AckOrgTeamGroupUnbindedListByOrganizationIDByTeamID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgTeamGroupUnbindedListByOrganizationIDByTeamID
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/org/team/group/unbinded/list/:organization_id/:team_id",
		ask,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"team_id":         TeamID,
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

type AskOrgGroupDisableByOrganizationIDByGroupID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskOrgGroupDisableByOrganizationIDByGroupID() *AskOrgGroupDisableByOrganizationIDByGroupID {
	return &AskOrgGroupDisableByOrganizationIDByGroupID{}
}

// 禁用团购组织成员
func (m *Proxy) OrgGroupDisableByOrganizationIDByGroupID(OrganizationID uint32,
	GroupID uint32,
	ask *AskOrgGroupDisableByOrganizationIDByGroupID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/community/org/group/disable/:organization_id/:group_id",
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
