package cidl

type AckAdminGroupListByOrganizationID struct {
	Count uint32   `db:"Count"`
	List  []*Group `db:"List"`
}

func NewAckAdminGroupListByOrganizationID() *AckAdminGroupListByOrganizationID {
	return &AckAdminGroupListByOrganizationID{
		List: make([]*Group, 0),
	}
}

type MetaApiAdminGroupListByOrganizationID struct {
}

var META_ADMIN_GROUP_LIST_BY_ORGANIZATION_ID = &MetaApiAdminGroupListByOrganizationID{}

func (m *MetaApiAdminGroupListByOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiAdminGroupListByOrganizationID) GetURL() string {
	return "/community/admin/group/list/:organization_id"
}
func (m *MetaApiAdminGroupListByOrganizationID) GetName() string {
	return "AdminGroupListByOrganizationID"
}
func (m *MetaApiAdminGroupListByOrganizationID) GetType() string { return "json" }

// 团购组织社区群列表
type ApiAdminGroupListByOrganizationID struct {
	MetaApiAdminGroupListByOrganizationID
	Ack    *AckAdminGroupListByOrganizationID
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

func (m *ApiAdminGroupListByOrganizationID) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminGroupListByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminGroupListByOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiAdminGroupListByOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiAdminGroupListByOrganizationID() ApiAdminGroupListByOrganizationID {
	return ApiAdminGroupListByOrganizationID{
		Ack: NewAckAdminGroupListByOrganizationID(),
	}
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

type MetaApiAdminGroupInfoByOrganizationIDByGroupID struct {
}

var META_ADMIN_GROUP_INFO_BY_ORGANIZATION_ID_BY_GROUP_ID = &MetaApiAdminGroupInfoByOrganizationIDByGroupID{}

func (m *MetaApiAdminGroupInfoByOrganizationIDByGroupID) GetMethod() string { return "GET" }
func (m *MetaApiAdminGroupInfoByOrganizationIDByGroupID) GetURL() string {
	return "/community/admin/group/info/:organization_id/:group_id"
}
func (m *MetaApiAdminGroupInfoByOrganizationIDByGroupID) GetName() string {
	return "AdminGroupInfoByOrganizationIDByGroupID"
}
func (m *MetaApiAdminGroupInfoByOrganizationIDByGroupID) GetType() string { return "json" }

// 团购组织社群详情
type ApiAdminGroupInfoByOrganizationIDByGroupID struct {
	MetaApiAdminGroupInfoByOrganizationIDByGroupID
	Ack    *AckAdminGroupInfoByOrganizationIDByGroupID
	Params struct {
		GroupID        uint32 `form:"group_id" binding:"required,gt=0" db:"GroupID"`
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiAdminGroupInfoByOrganizationIDByGroupID) GetQuery() interface{}  { return nil }
func (m *ApiAdminGroupInfoByOrganizationIDByGroupID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminGroupInfoByOrganizationIDByGroupID) GetAsk() interface{}    { return nil }
func (m *ApiAdminGroupInfoByOrganizationIDByGroupID) GetAck() interface{}    { return m.Ack }
func MakeApiAdminGroupInfoByOrganizationIDByGroupID() ApiAdminGroupInfoByOrganizationIDByGroupID {
	return ApiAdminGroupInfoByOrganizationIDByGroupID{
		Ack: NewAckAdminGroupInfoByOrganizationIDByGroupID(),
	}
}
