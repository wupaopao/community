package cidl

type AckWxXcxGroupInfoByGroupID struct {
	Group
	ManagerAvatar   string `db:"ManagerAvatar"`
	ManagerNickname string `db:"ManagerNickname"`
}

func NewAckWxXcxGroupInfoByGroupID() *AckWxXcxGroupInfoByGroupID {
	return &AckWxXcxGroupInfoByGroupID{}
}

type MetaApiWxXcxGroupInfoByGroupID struct {
}

var META_WX_XCX_GROUP_INFO_BY_GROUP_ID = &MetaApiWxXcxGroupInfoByGroupID{}

func (m *MetaApiWxXcxGroupInfoByGroupID) GetMethod() string { return "GET" }
func (m *MetaApiWxXcxGroupInfoByGroupID) GetURL() string {
	return "/community/wx_xcx/group/info/:group_id"
}
func (m *MetaApiWxXcxGroupInfoByGroupID) GetName() string { return "WxXcxGroupInfoByGroupID" }
func (m *MetaApiWxXcxGroupInfoByGroupID) GetType() string { return "json" }

// 获取社群信息
type ApiWxXcxGroupInfoByGroupID struct {
	MetaApiWxXcxGroupInfoByGroupID
	Ack    *AckWxXcxGroupInfoByGroupID
	Params struct {
		GroupID uint32 `form:"group_id" binding:"required,gt=0" db:"GroupID"`
	}
}

func (m *ApiWxXcxGroupInfoByGroupID) GetQuery() interface{}  { return nil }
func (m *ApiWxXcxGroupInfoByGroupID) GetParams() interface{} { return &m.Params }
func (m *ApiWxXcxGroupInfoByGroupID) GetAsk() interface{}    { return nil }
func (m *ApiWxXcxGroupInfoByGroupID) GetAck() interface{}    { return m.Ack }
func MakeApiWxXcxGroupInfoByGroupID() ApiWxXcxGroupInfoByGroupID {
	return ApiWxXcxGroupInfoByGroupID{
		Ack: NewAckWxXcxGroupInfoByGroupID(),
	}
}

type AskWxXcxGroupAddByOrganizationID struct {
	ManagerName   string `binding:"required,lte=64" db:"ManagerName"`
	ManagerMobile string `binding:"required,numeric" db:"ManagerMobile"`
	Address       string `binding:"required,lte=255" db:"Address"`
	PostCode      string `binding:"numeric,len=6" db:"PostCode"`
}

func NewAskWxXcxGroupAddByOrganizationID() *AskWxXcxGroupAddByOrganizationID {
	return &AskWxXcxGroupAddByOrganizationID{}
}

type MetaApiWxXcxGroupAddByOrganizationID struct {
}

var META_WX_XCX_GROUP_ADD_BY_ORGANIZATION_ID = &MetaApiWxXcxGroupAddByOrganizationID{}

func (m *MetaApiWxXcxGroupAddByOrganizationID) GetMethod() string { return "POST" }
func (m *MetaApiWxXcxGroupAddByOrganizationID) GetURL() string {
	return "/community/wx_xcx/group/add/:organization_id"
}
func (m *MetaApiWxXcxGroupAddByOrganizationID) GetName() string {
	return "WxXcxGroupAddByOrganizationID"
}
func (m *MetaApiWxXcxGroupAddByOrganizationID) GetType() string { return "json" }

// 添加社群
type ApiWxXcxGroupAddByOrganizationID struct {
	MetaApiWxXcxGroupAddByOrganizationID
	Ask    *AskWxXcxGroupAddByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiWxXcxGroupAddByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiWxXcxGroupAddByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiWxXcxGroupAddByOrganizationID) GetAsk() interface{}    { return m.Ask }
func (m *ApiWxXcxGroupAddByOrganizationID) GetAck() interface{}    { return nil }
func MakeApiWxXcxGroupAddByOrganizationID() ApiWxXcxGroupAddByOrganizationID {
	return ApiWxXcxGroupAddByOrganizationID{
		Ask: NewAskWxXcxGroupAddByOrganizationID(),
	}
}
