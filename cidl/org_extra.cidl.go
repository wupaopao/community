package cidl

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

type MetaApiOrgGroupAddByOrganizationID struct {
}

var META_ORG_GROUP_ADD_BY_ORGANIZATION_ID = &MetaApiOrgGroupAddByOrganizationID{}

func (m *MetaApiOrgGroupAddByOrganizationID) GetMethod() string { return "POST" }
func (m *MetaApiOrgGroupAddByOrganizationID) GetURL() string {
	return "/community/org/group/add/:organization_id"
}
func (m *MetaApiOrgGroupAddByOrganizationID) GetName() string { return "OrgGroupAddByOrganizationID" }
func (m *MetaApiOrgGroupAddByOrganizationID) GetType() string { return "json" }

// 添加
type ApiOrgGroupAddByOrganizationID struct {
	MetaApiOrgGroupAddByOrganizationID
	Ask    *AskOrgGroupAddByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiOrgGroupAddByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiOrgGroupAddByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgGroupAddByOrganizationID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgGroupAddByOrganizationID) GetAck() interface{}    { return nil }
func MakeApiOrgGroupAddByOrganizationID() ApiOrgGroupAddByOrganizationID {
	return ApiOrgGroupAddByOrganizationID{
		Ask: NewAskOrgGroupAddByOrganizationID(),
	}
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

type MetaApiOrgGroupEditByOrganizationIDByGroupID struct {
}

var META_ORG_GROUP_EDIT_BY_ORGANIZATION_ID_BY_GROUP_ID = &MetaApiOrgGroupEditByOrganizationIDByGroupID{}

func (m *MetaApiOrgGroupEditByOrganizationIDByGroupID) GetMethod() string { return "POST" }
func (m *MetaApiOrgGroupEditByOrganizationIDByGroupID) GetURL() string {
	return "/community/org/group/edit/:organization_id/:group_id"
}
func (m *MetaApiOrgGroupEditByOrganizationIDByGroupID) GetName() string {
	return "OrgGroupEditByOrganizationIDByGroupID"
}
func (m *MetaApiOrgGroupEditByOrganizationIDByGroupID) GetType() string { return "json" }

// 编辑
type ApiOrgGroupEditByOrganizationIDByGroupID struct {
	MetaApiOrgGroupEditByOrganizationIDByGroupID
	Ask    *AskOrgGroupEditByOrganizationIDByGroupID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		GroupID        uint32 `form:"group_id" binding:"required,gt=0" db:"GroupID"`
	}
}

func (m *ApiOrgGroupEditByOrganizationIDByGroupID) GetQuery() interface{}  { return nil }
func (m *ApiOrgGroupEditByOrganizationIDByGroupID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgGroupEditByOrganizationIDByGroupID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgGroupEditByOrganizationIDByGroupID) GetAck() interface{}    { return nil }
func MakeApiOrgGroupEditByOrganizationIDByGroupID() ApiOrgGroupEditByOrganizationIDByGroupID {
	return ApiOrgGroupEditByOrganizationIDByGroupID{
		Ask: NewAskOrgGroupEditByOrganizationIDByGroupID(),
	}
}

type AskOrgGroupChangeManagerByOrganizationIDByGroupID struct {
	NewManagerUid string `binding:"required" db:"NewManagerUid"`
}

func NewAskOrgGroupChangeManagerByOrganizationIDByGroupID() *AskOrgGroupChangeManagerByOrganizationIDByGroupID {
	return &AskOrgGroupChangeManagerByOrganizationIDByGroupID{}
}

type MetaApiOrgGroupChangeManagerByOrganizationIDByGroupID struct {
}

var META_ORG_GROUP_CHANGE_MANAGER_BY_ORGANIZATION_ID_BY_GROUP_ID = &MetaApiOrgGroupChangeManagerByOrganizationIDByGroupID{}

func (m *MetaApiOrgGroupChangeManagerByOrganizationIDByGroupID) GetMethod() string { return "POST" }
func (m *MetaApiOrgGroupChangeManagerByOrganizationIDByGroupID) GetURL() string {
	return "/community/org/group/change_manager/:organization_id/:group_id"
}
func (m *MetaApiOrgGroupChangeManagerByOrganizationIDByGroupID) GetName() string {
	return "OrgGroupChangeManagerByOrganizationIDByGroupID"
}
func (m *MetaApiOrgGroupChangeManagerByOrganizationIDByGroupID) GetType() string { return "json" }

// 更改管理员
type ApiOrgGroupChangeManagerByOrganizationIDByGroupID struct {
	MetaApiOrgGroupChangeManagerByOrganizationIDByGroupID
	Ask    *AskOrgGroupChangeManagerByOrganizationIDByGroupID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		GroupID        uint32 `form:"group_id" binding:"required,gt=0" db:"GroupID"`
	}
}

func (m *ApiOrgGroupChangeManagerByOrganizationIDByGroupID) GetQuery() interface{}  { return nil }
func (m *ApiOrgGroupChangeManagerByOrganizationIDByGroupID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgGroupChangeManagerByOrganizationIDByGroupID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgGroupChangeManagerByOrganizationIDByGroupID) GetAck() interface{}    { return nil }
func MakeApiOrgGroupChangeManagerByOrganizationIDByGroupID() ApiOrgGroupChangeManagerByOrganizationIDByGroupID {
	return ApiOrgGroupChangeManagerByOrganizationIDByGroupID{
		Ask: NewAskOrgGroupChangeManagerByOrganizationIDByGroupID(),
	}
}

type AskOrgGroupAuditByOrganizationIDByGroupID struct {
	AuditState GroupAuditState `binding:"required" db:"AuditState"`
}

func NewAskOrgGroupAuditByOrganizationIDByGroupID() *AskOrgGroupAuditByOrganizationIDByGroupID {
	return &AskOrgGroupAuditByOrganizationIDByGroupID{}
}

type MetaApiOrgGroupAuditByOrganizationIDByGroupID struct {
}

var META_ORG_GROUP_AUDIT_BY_ORGANIZATION_ID_BY_GROUP_ID = &MetaApiOrgGroupAuditByOrganizationIDByGroupID{}

func (m *MetaApiOrgGroupAuditByOrganizationIDByGroupID) GetMethod() string { return "POST" }
func (m *MetaApiOrgGroupAuditByOrganizationIDByGroupID) GetURL() string {
	return "/community/org/group/audit/:organization_id/:group_id"
}
func (m *MetaApiOrgGroupAuditByOrganizationIDByGroupID) GetName() string {
	return "OrgGroupAuditByOrganizationIDByGroupID"
}
func (m *MetaApiOrgGroupAuditByOrganizationIDByGroupID) GetType() string { return "json" }

// 审核
type ApiOrgGroupAuditByOrganizationIDByGroupID struct {
	MetaApiOrgGroupAuditByOrganizationIDByGroupID
	Ask    *AskOrgGroupAuditByOrganizationIDByGroupID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		GroupID        uint32 `form:"group_id" binding:"required,gt=0" db:"GroupID"`
	}
}

func (m *ApiOrgGroupAuditByOrganizationIDByGroupID) GetQuery() interface{}  { return nil }
func (m *ApiOrgGroupAuditByOrganizationIDByGroupID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgGroupAuditByOrganizationIDByGroupID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgGroupAuditByOrganizationIDByGroupID) GetAck() interface{}    { return nil }
func MakeApiOrgGroupAuditByOrganizationIDByGroupID() ApiOrgGroupAuditByOrganizationIDByGroupID {
	return ApiOrgGroupAuditByOrganizationIDByGroupID{
		Ask: NewAskOrgGroupAuditByOrganizationIDByGroupID(),
	}
}

type AskOrgTeamAddByOrganizationID struct {
	TeamName string `binding:"required,lte=128" db:"TeamName"`
}

func NewAskOrgTeamAddByOrganizationID() *AskOrgTeamAddByOrganizationID {
	return &AskOrgTeamAddByOrganizationID{}
}

type MetaApiOrgTeamAddByOrganizationID struct {
}

var META_ORG_TEAM_ADD_BY_ORGANIZATION_ID = &MetaApiOrgTeamAddByOrganizationID{}

func (m *MetaApiOrgTeamAddByOrganizationID) GetMethod() string { return "POST" }
func (m *MetaApiOrgTeamAddByOrganizationID) GetURL() string {
	return "/community/org/team/add/:organization_id"
}
func (m *MetaApiOrgTeamAddByOrganizationID) GetName() string { return "OrgTeamAddByOrganizationID" }
func (m *MetaApiOrgTeamAddByOrganizationID) GetType() string { return "json" }

// 添加社团群组
type ApiOrgTeamAddByOrganizationID struct {
	MetaApiOrgTeamAddByOrganizationID
	Ask    *AskOrgTeamAddByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiOrgTeamAddByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiOrgTeamAddByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgTeamAddByOrganizationID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgTeamAddByOrganizationID) GetAck() interface{}    { return nil }
func MakeApiOrgTeamAddByOrganizationID() ApiOrgTeamAddByOrganizationID {
	return ApiOrgTeamAddByOrganizationID{
		Ask: NewAskOrgTeamAddByOrganizationID(),
	}
}

type AskOrgTeamEditByOrganizationIDByTeamID struct {
	TeamName string `binding:"required,lte=128" db:"TeamName"`
}

func NewAskOrgTeamEditByOrganizationIDByTeamID() *AskOrgTeamEditByOrganizationIDByTeamID {
	return &AskOrgTeamEditByOrganizationIDByTeamID{}
}

type MetaApiOrgTeamEditByOrganizationIDByTeamID struct {
}

var META_ORG_TEAM_EDIT_BY_ORGANIZATION_ID_BY_TEAM_ID = &MetaApiOrgTeamEditByOrganizationIDByTeamID{}

func (m *MetaApiOrgTeamEditByOrganizationIDByTeamID) GetMethod() string { return "POST" }
func (m *MetaApiOrgTeamEditByOrganizationIDByTeamID) GetURL() string {
	return "/community/org/team/edit/:organization_id/:team_id"
}
func (m *MetaApiOrgTeamEditByOrganizationIDByTeamID) GetName() string {
	return "OrgTeamEditByOrganizationIDByTeamID"
}
func (m *MetaApiOrgTeamEditByOrganizationIDByTeamID) GetType() string { return "json" }

// 编辑社团群组
type ApiOrgTeamEditByOrganizationIDByTeamID struct {
	MetaApiOrgTeamEditByOrganizationIDByTeamID
	Ask    *AskOrgTeamEditByOrganizationIDByTeamID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		TeamID         uint32 `form:"team_id" binding:"required,gt=0" db:"TeamID"`
	}
}

func (m *ApiOrgTeamEditByOrganizationIDByTeamID) GetQuery() interface{}  { return nil }
func (m *ApiOrgTeamEditByOrganizationIDByTeamID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgTeamEditByOrganizationIDByTeamID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgTeamEditByOrganizationIDByTeamID) GetAck() interface{}    { return nil }
func MakeApiOrgTeamEditByOrganizationIDByTeamID() ApiOrgTeamEditByOrganizationIDByTeamID {
	return ApiOrgTeamEditByOrganizationIDByTeamID{
		Ask: NewAskOrgTeamEditByOrganizationIDByTeamID(),
	}
}

type MetaApiOrgTeamDeleteByOrganizationIDByTeamID struct {
}

var META_ORG_TEAM_DELETE_BY_ORGANIZATION_ID_BY_TEAM_ID = &MetaApiOrgTeamDeleteByOrganizationIDByTeamID{}

func (m *MetaApiOrgTeamDeleteByOrganizationIDByTeamID) GetMethod() string { return "POST" }
func (m *MetaApiOrgTeamDeleteByOrganizationIDByTeamID) GetURL() string {
	return "/community/org/team/delete/:organization_id/:team_id"
}
func (m *MetaApiOrgTeamDeleteByOrganizationIDByTeamID) GetName() string {
	return "OrgTeamDeleteByOrganizationIDByTeamID"
}
func (m *MetaApiOrgTeamDeleteByOrganizationIDByTeamID) GetType() string { return "json" }

// 删除社团群组
type ApiOrgTeamDeleteByOrganizationIDByTeamID struct {
	MetaApiOrgTeamDeleteByOrganizationIDByTeamID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		TeamID         uint32 `form:"team_id" binding:"required,gt=0" db:"TeamID"`
	}
}

func (m *ApiOrgTeamDeleteByOrganizationIDByTeamID) GetQuery() interface{}  { return nil }
func (m *ApiOrgTeamDeleteByOrganizationIDByTeamID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgTeamDeleteByOrganizationIDByTeamID) GetAsk() interface{}    { return nil }
func (m *ApiOrgTeamDeleteByOrganizationIDByTeamID) GetAck() interface{}    { return nil }
func MakeApiOrgTeamDeleteByOrganizationIDByTeamID() ApiOrgTeamDeleteByOrganizationIDByTeamID {
	return ApiOrgTeamDeleteByOrganizationIDByTeamID{}
}

type AskOrgTeamBindGroupByOrganizationIDByTeamID struct {
	GroupID uint32 `db:"GroupID"`
}

func NewAskOrgTeamBindGroupByOrganizationIDByTeamID() *AskOrgTeamBindGroupByOrganizationIDByTeamID {
	return &AskOrgTeamBindGroupByOrganizationIDByTeamID{}
}

type MetaApiOrgTeamBindGroupByOrganizationIDByTeamID struct {
}

var META_ORG_TEAM_BIND_GROUP_BY_ORGANIZATION_ID_BY_TEAM_ID = &MetaApiOrgTeamBindGroupByOrganizationIDByTeamID{}

func (m *MetaApiOrgTeamBindGroupByOrganizationIDByTeamID) GetMethod() string { return "POST" }
func (m *MetaApiOrgTeamBindGroupByOrganizationIDByTeamID) GetURL() string {
	return "/community/org/team/bind/group/:organization_id/:team_id"
}
func (m *MetaApiOrgTeamBindGroupByOrganizationIDByTeamID) GetName() string {
	return "OrgTeamBindGroupByOrganizationIDByTeamID"
}
func (m *MetaApiOrgTeamBindGroupByOrganizationIDByTeamID) GetType() string { return "json" }

// 群组添加社团
type ApiOrgTeamBindGroupByOrganizationIDByTeamID struct {
	MetaApiOrgTeamBindGroupByOrganizationIDByTeamID
	Ask    *AskOrgTeamBindGroupByOrganizationIDByTeamID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		TeamID         uint32 `form:"team_id" binding:"required,gt=0" db:"TeamID"`
	}
}

func (m *ApiOrgTeamBindGroupByOrganizationIDByTeamID) GetQuery() interface{}  { return nil }
func (m *ApiOrgTeamBindGroupByOrganizationIDByTeamID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgTeamBindGroupByOrganizationIDByTeamID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgTeamBindGroupByOrganizationIDByTeamID) GetAck() interface{}    { return nil }
func MakeApiOrgTeamBindGroupByOrganizationIDByTeamID() ApiOrgTeamBindGroupByOrganizationIDByTeamID {
	return ApiOrgTeamBindGroupByOrganizationIDByTeamID{
		Ask: NewAskOrgTeamBindGroupByOrganizationIDByTeamID(),
	}
}

type AskOrgTeamUnbindGroupByOrganizationIDByTeamID struct {
	GroupID uint32 `db:"GroupID"`
}

func NewAskOrgTeamUnbindGroupByOrganizationIDByTeamID() *AskOrgTeamUnbindGroupByOrganizationIDByTeamID {
	return &AskOrgTeamUnbindGroupByOrganizationIDByTeamID{}
}

type MetaApiOrgTeamUnbindGroupByOrganizationIDByTeamID struct {
}

var META_ORG_TEAM_UNBIND_GROUP_BY_ORGANIZATION_ID_BY_TEAM_ID = &MetaApiOrgTeamUnbindGroupByOrganizationIDByTeamID{}

func (m *MetaApiOrgTeamUnbindGroupByOrganizationIDByTeamID) GetMethod() string { return "POST" }
func (m *MetaApiOrgTeamUnbindGroupByOrganizationIDByTeamID) GetURL() string {
	return "/community/org/team/unbind/group/:organization_id/:team_id"
}
func (m *MetaApiOrgTeamUnbindGroupByOrganizationIDByTeamID) GetName() string {
	return "OrgTeamUnbindGroupByOrganizationIDByTeamID"
}
func (m *MetaApiOrgTeamUnbindGroupByOrganizationIDByTeamID) GetType() string { return "json" }

// 群组移除社团
type ApiOrgTeamUnbindGroupByOrganizationIDByTeamID struct {
	MetaApiOrgTeamUnbindGroupByOrganizationIDByTeamID
	Ask    *AskOrgTeamUnbindGroupByOrganizationIDByTeamID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		TeamID         uint32 `form:"team_id" binding:"required,gt=0" db:"TeamID"`
	}
}

func (m *ApiOrgTeamUnbindGroupByOrganizationIDByTeamID) GetQuery() interface{}  { return nil }
func (m *ApiOrgTeamUnbindGroupByOrganizationIDByTeamID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgTeamUnbindGroupByOrganizationIDByTeamID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgTeamUnbindGroupByOrganizationIDByTeamID) GetAck() interface{}    { return nil }
func MakeApiOrgTeamUnbindGroupByOrganizationIDByTeamID() ApiOrgTeamUnbindGroupByOrganizationIDByTeamID {
	return ApiOrgTeamUnbindGroupByOrganizationIDByTeamID{
		Ask: NewAskOrgTeamUnbindGroupByOrganizationIDByTeamID(),
	}
}

type AckOrgTeamGroupBindedListByOrganizationIDByTeamID struct {
	Groups []*SubGroup `db:"Groups"`
}

func NewAckOrgTeamGroupBindedListByOrganizationIDByTeamID() *AckOrgTeamGroupBindedListByOrganizationIDByTeamID {
	return &AckOrgTeamGroupBindedListByOrganizationIDByTeamID{
		Groups: make([]*SubGroup, 0),
	}
}

type MetaApiOrgTeamGroupBindedListByOrganizationIDByTeamID struct {
}

var META_ORG_TEAM_GROUP_BINDED_LIST_BY_ORGANIZATION_ID_BY_TEAM_ID = &MetaApiOrgTeamGroupBindedListByOrganizationIDByTeamID{}

func (m *MetaApiOrgTeamGroupBindedListByOrganizationIDByTeamID) GetMethod() string { return "POST" }
func (m *MetaApiOrgTeamGroupBindedListByOrganizationIDByTeamID) GetURL() string {
	return "/community/org/team/group/binded/list/:organization_id/:team_id"
}
func (m *MetaApiOrgTeamGroupBindedListByOrganizationIDByTeamID) GetName() string {
	return "OrgTeamGroupBindedListByOrganizationIDByTeamID"
}
func (m *MetaApiOrgTeamGroupBindedListByOrganizationIDByTeamID) GetType() string { return "json" }

// 群组已关联社团列表
type ApiOrgTeamGroupBindedListByOrganizationIDByTeamID struct {
	MetaApiOrgTeamGroupBindedListByOrganizationIDByTeamID
	Ack    *AckOrgTeamGroupBindedListByOrganizationIDByTeamID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		TeamID         uint32 `form:"team_id" binding:"required,gt=0" db:"TeamID"`
	}
}

func (m *ApiOrgTeamGroupBindedListByOrganizationIDByTeamID) GetQuery() interface{}  { return nil }
func (m *ApiOrgTeamGroupBindedListByOrganizationIDByTeamID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgTeamGroupBindedListByOrganizationIDByTeamID) GetAsk() interface{}    { return nil }
func (m *ApiOrgTeamGroupBindedListByOrganizationIDByTeamID) GetAck() interface{}    { return m.Ack }
func MakeApiOrgTeamGroupBindedListByOrganizationIDByTeamID() ApiOrgTeamGroupBindedListByOrganizationIDByTeamID {
	return ApiOrgTeamGroupBindedListByOrganizationIDByTeamID{
		Ack: NewAckOrgTeamGroupBindedListByOrganizationIDByTeamID(),
	}
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

type MetaApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID struct {
}

var META_ORG_TEAM_GROUP_UNBINDED_LIST_BY_ORGANIZATION_ID_BY_TEAM_ID = &MetaApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID{}

func (m *MetaApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID) GetMethod() string { return "POST" }
func (m *MetaApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID) GetURL() string {
	return "/community/org/team/group/unbinded/list/:organization_id/:team_id"
}
func (m *MetaApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID) GetName() string {
	return "OrgTeamGroupUnbindedListByOrganizationIDByTeamID"
}
func (m *MetaApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID) GetType() string { return "json" }

// 群组未关联社团列表
type ApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID struct {
	MetaApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID
	Ask    *AskOrgTeamGroupUnbindedListByOrganizationIDByTeamID
	Ack    *AckOrgTeamGroupUnbindedListByOrganizationIDByTeamID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		TeamID         uint32 `form:"team_id" binding:"required,gt=0" db:"TeamID"`
	}
}

func (m *ApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID) GetQuery() interface{} { return nil }
func (m *ApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID) GetParams() interface{} {
	return &m.Params
}
func (m *ApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID) GetAsk() interface{} { return m.Ask }
func (m *ApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID) GetAck() interface{} { return m.Ack }
func MakeApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID() ApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID {
	return ApiOrgTeamGroupUnbindedListByOrganizationIDByTeamID{
		Ask: NewAskOrgTeamGroupUnbindedListByOrganizationIDByTeamID(),
		Ack: NewAckOrgTeamGroupUnbindedListByOrganizationIDByTeamID(),
	}
}

type AskOrgGroupDisableByOrganizationIDByGroupID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskOrgGroupDisableByOrganizationIDByGroupID() *AskOrgGroupDisableByOrganizationIDByGroupID {
	return &AskOrgGroupDisableByOrganizationIDByGroupID{}
}

type MetaApiOrgGroupDisableByOrganizationIDByGroupID struct {
}

var META_ORG_GROUP_DISABLE_BY_ORGANIZATION_ID_BY_GROUP_ID = &MetaApiOrgGroupDisableByOrganizationIDByGroupID{}

func (m *MetaApiOrgGroupDisableByOrganizationIDByGroupID) GetMethod() string { return "POST" }
func (m *MetaApiOrgGroupDisableByOrganizationIDByGroupID) GetURL() string {
	return "/community/org/group/disable/:organization_id/:group_id"
}
func (m *MetaApiOrgGroupDisableByOrganizationIDByGroupID) GetName() string {
	return "OrgGroupDisableByOrganizationIDByGroupID"
}
func (m *MetaApiOrgGroupDisableByOrganizationIDByGroupID) GetType() string { return "json" }

// 禁用团购组织成员
type ApiOrgGroupDisableByOrganizationIDByGroupID struct {
	MetaApiOrgGroupDisableByOrganizationIDByGroupID
	Ask    *AskOrgGroupDisableByOrganizationIDByGroupID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		GroupID        uint32 `form:"group_id" binding:"required" db:"GroupID"`
	}
}

func (m *ApiOrgGroupDisableByOrganizationIDByGroupID) GetQuery() interface{}  { return nil }
func (m *ApiOrgGroupDisableByOrganizationIDByGroupID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgGroupDisableByOrganizationIDByGroupID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgGroupDisableByOrganizationIDByGroupID) GetAck() interface{}    { return nil }
func MakeApiOrgGroupDisableByOrganizationIDByGroupID() ApiOrgGroupDisableByOrganizationIDByGroupID {
	return ApiOrgGroupDisableByOrganizationIDByGroupID{
		Ask: NewAskOrgGroupDisableByOrganizationIDByGroupID(),
	}
}
