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
