package cidl

type MetaApiInnerCommunityGroupInfoByGroupID struct {
}

var META_INNER_COMMUNITY_GROUP_INFO_BY_GROUP_ID = &MetaApiInnerCommunityGroupInfoByGroupID{}

func (m *MetaApiInnerCommunityGroupInfoByGroupID) GetMethod() string { return "GET" }
func (m *MetaApiInnerCommunityGroupInfoByGroupID) GetURL() string {
	return "/inner/community/group/info/:group_id"
}
func (m *MetaApiInnerCommunityGroupInfoByGroupID) GetName() string {
	return "InnerCommunityGroupInfoByGroupID"
}
func (m *MetaApiInnerCommunityGroupInfoByGroupID) GetType() string { return "json" }

// 获取团购组织
type ApiInnerCommunityGroupInfoByGroupID struct {
	MetaApiInnerCommunityGroupInfoByGroupID
	Ack    *Group
	Params struct {
		GroupID uint32 `form:"group_id" binding:"required,gt=0" db:"GroupID"`
	}
}

func (m *ApiInnerCommunityGroupInfoByGroupID) GetQuery() interface{}  { return nil }
func (m *ApiInnerCommunityGroupInfoByGroupID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerCommunityGroupInfoByGroupID) GetAsk() interface{}    { return nil }
func (m *ApiInnerCommunityGroupInfoByGroupID) GetAck() interface{}    { return m.Ack }
func MakeApiInnerCommunityGroupInfoByGroupID() ApiInnerCommunityGroupInfoByGroupID {
	return ApiInnerCommunityGroupInfoByGroupID{
		Ack: NewGroup(),
	}
}

type MetaApiInnerCommunityGroupInfoByUserIDByUserID struct {
}

var META_INNER_COMMUNITY_GROUP_INFO_BY_USER_ID_BY_USER_ID = &MetaApiInnerCommunityGroupInfoByUserIDByUserID{}

func (m *MetaApiInnerCommunityGroupInfoByUserIDByUserID) GetMethod() string { return "GET" }
func (m *MetaApiInnerCommunityGroupInfoByUserIDByUserID) GetURL() string {
	return "/inner/community/group/info_by_user_id/:user_id"
}
func (m *MetaApiInnerCommunityGroupInfoByUserIDByUserID) GetName() string {
	return "InnerCommunityGroupInfoByUserIDByUserID"
}
func (m *MetaApiInnerCommunityGroupInfoByUserIDByUserID) GetType() string { return "json" }

// 通过用户ID获取团购组织
type ApiInnerCommunityGroupInfoByUserIDByUserID struct {
	MetaApiInnerCommunityGroupInfoByUserIDByUserID
	Ack    *Group
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiInnerCommunityGroupInfoByUserIDByUserID) GetQuery() interface{}  { return nil }
func (m *ApiInnerCommunityGroupInfoByUserIDByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerCommunityGroupInfoByUserIDByUserID) GetAsk() interface{}    { return nil }
func (m *ApiInnerCommunityGroupInfoByUserIDByUserID) GetAck() interface{}    { return m.Ack }
func MakeApiInnerCommunityGroupInfoByUserIDByUserID() ApiInnerCommunityGroupInfoByUserIDByUserID {
	return ApiInnerCommunityGroupInfoByUserIDByUserID{
		Ack: NewGroup(),
	}
}

type AckInnerCommunityGroupCountByOrganizationID struct {
	Count uint32 `db:"Count"`
}

func NewAckInnerCommunityGroupCountByOrganizationID() *AckInnerCommunityGroupCountByOrganizationID {
	return &AckInnerCommunityGroupCountByOrganizationID{}
}

type MetaApiInnerCommunityGroupCountByOrganizationID struct {
}

var META_INNER_COMMUNITY_GROUP_COUNT_BY_ORGANIZATION_ID = &MetaApiInnerCommunityGroupCountByOrganizationID{}

func (m *MetaApiInnerCommunityGroupCountByOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiInnerCommunityGroupCountByOrganizationID) GetURL() string {
	return "/inner/community/group/count/:organization_id"
}
func (m *MetaApiInnerCommunityGroupCountByOrganizationID) GetName() string {
	return "InnerCommunityGroupCountByOrganizationID"
}
func (m *MetaApiInnerCommunityGroupCountByOrganizationID) GetType() string { return "json" }

// 获取团购组织所属社团数目
type ApiInnerCommunityGroupCountByOrganizationID struct {
	MetaApiInnerCommunityGroupCountByOrganizationID
	Ack    *AckInnerCommunityGroupCountByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiInnerCommunityGroupCountByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiInnerCommunityGroupCountByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerCommunityGroupCountByOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiInnerCommunityGroupCountByOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiInnerCommunityGroupCountByOrganizationID() ApiInnerCommunityGroupCountByOrganizationID {
	return ApiInnerCommunityGroupCountByOrganizationID{
		Ack: NewAckInnerCommunityGroupCountByOrganizationID(),
	}
}

type AckInnerCommunityGroupTeamByGroupID struct {
	TeamIDs []uint32 `db:"TeamIDs"`
}

func NewAckInnerCommunityGroupTeamByGroupID() *AckInnerCommunityGroupTeamByGroupID {
	return &AckInnerCommunityGroupTeamByGroupID{
		TeamIDs: make([]uint32, 0),
	}
}

type MetaApiInnerCommunityGroupTeamByGroupID struct {
}

var META_INNER_COMMUNITY_GROUP_TEAM_BY_GROUP_ID = &MetaApiInnerCommunityGroupTeamByGroupID{}

func (m *MetaApiInnerCommunityGroupTeamByGroupID) GetMethod() string { return "GET" }
func (m *MetaApiInnerCommunityGroupTeamByGroupID) GetURL() string {
	return "/inner/community/group/team/:group_id"
}
func (m *MetaApiInnerCommunityGroupTeamByGroupID) GetName() string {
	return "InnerCommunityGroupTeamByGroupID"
}
func (m *MetaApiInnerCommunityGroupTeamByGroupID) GetType() string { return "json" }

// 获取社团所属的组
type ApiInnerCommunityGroupTeamByGroupID struct {
	MetaApiInnerCommunityGroupTeamByGroupID
	Ack    *AckInnerCommunityGroupTeamByGroupID
	Params struct {
		GroupID uint32 `form:"group_id" binding:"required,gt=0" db:"GroupID"`
	}
}

func (m *ApiInnerCommunityGroupTeamByGroupID) GetQuery() interface{}  { return nil }
func (m *ApiInnerCommunityGroupTeamByGroupID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerCommunityGroupTeamByGroupID) GetAsk() interface{}    { return nil }
func (m *ApiInnerCommunityGroupTeamByGroupID) GetAck() interface{}    { return m.Ack }
func MakeApiInnerCommunityGroupTeamByGroupID() ApiInnerCommunityGroupTeamByGroupID {
	return ApiInnerCommunityGroupTeamByGroupID{
		Ack: NewAckInnerCommunityGroupTeamByGroupID(),
	}
}
