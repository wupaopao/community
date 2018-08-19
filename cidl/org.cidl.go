package cidl

type AckOrgGroupListByOrganizationID struct {
	Count uint32   `db:"Count"`
	List  []*Group `db:"List"`
}

func NewAckOrgGroupListByOrganizationID() *AckOrgGroupListByOrganizationID {
	return &AckOrgGroupListByOrganizationID{
		List: make([]*Group, 0),
	}
}

type MetaApiOrgGroupListByOrganizationID struct {
}

var META_ORG_GROUP_LIST_BY_ORGANIZATION_ID = &MetaApiOrgGroupListByOrganizationID{}

func (m *MetaApiOrgGroupListByOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiOrgGroupListByOrganizationID) GetURL() string {
	return "/community/org/group/list/:organization_id"
}
func (m *MetaApiOrgGroupListByOrganizationID) GetName() string { return "OrgGroupListByOrganizationID" }
func (m *MetaApiOrgGroupListByOrganizationID) GetType() string { return "json" }

// 团购组织社区群列表
type ApiOrgGroupListByOrganizationID struct {
	MetaApiOrgGroupListByOrganizationID
	Ack    *AckOrgGroupListByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
	Query struct {
		Page       uint32          `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize   uint32          `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
		Search     string          `form:"search" db:"Search"`
		AuditState GroupAuditState `form:"audit_state" db:"AuditState"`
	}
}

func (m *ApiOrgGroupListByOrganizationID) GetQuery() interface{}  { return &m.Query }
func (m *ApiOrgGroupListByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgGroupListByOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiOrgGroupListByOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiOrgGroupListByOrganizationID() ApiOrgGroupListByOrganizationID {
	return ApiOrgGroupListByOrganizationID{
		Ack: NewAckOrgGroupListByOrganizationID(),
	}
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

type MetaApiOrgGroupInfoByOrganizationIDByGroupID struct {
}

var META_ORG_GROUP_INFO_BY_ORGANIZATION_ID_BY_GROUP_ID = &MetaApiOrgGroupInfoByOrganizationIDByGroupID{}

func (m *MetaApiOrgGroupInfoByOrganizationIDByGroupID) GetMethod() string { return "GET" }
func (m *MetaApiOrgGroupInfoByOrganizationIDByGroupID) GetURL() string {
	return "/community/org/group/info/:organization_id/:group_id"
}
func (m *MetaApiOrgGroupInfoByOrganizationIDByGroupID) GetName() string {
	return "OrgGroupInfoByOrganizationIDByGroupID"
}
func (m *MetaApiOrgGroupInfoByOrganizationIDByGroupID) GetType() string { return "json" }

// 团购组织社群详情
type ApiOrgGroupInfoByOrganizationIDByGroupID struct {
	MetaApiOrgGroupInfoByOrganizationIDByGroupID
	Ack    *AckOrgGroupInfoByOrganizationIDByGroupID
	Params struct {
		GroupID        uint32 `form:"group_id" binding:"required,gt=0" db:"GroupID"`
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiOrgGroupInfoByOrganizationIDByGroupID) GetQuery() interface{}  { return nil }
func (m *ApiOrgGroupInfoByOrganizationIDByGroupID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgGroupInfoByOrganizationIDByGroupID) GetAsk() interface{}    { return nil }
func (m *ApiOrgGroupInfoByOrganizationIDByGroupID) GetAck() interface{}    { return m.Ack }
func MakeApiOrgGroupInfoByOrganizationIDByGroupID() ApiOrgGroupInfoByOrganizationIDByGroupID {
	return ApiOrgGroupInfoByOrganizationIDByGroupID{
		Ack: NewAckOrgGroupInfoByOrganizationIDByGroupID(),
	}
}
