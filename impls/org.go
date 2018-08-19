package impls

import (
	"business/community/cidl"
	"business/community/common/db"
	"business/user/proxy/user"
	"common/api"

	"github.com/mz-eco/mz/http"
)

func init() {
	AddOrgGroupListByOrganizationIDHandler()
	AddOrgGroupInfoByOrganizationIDByGroupIDHandler()
}

type OrgGroupListByOrganizationIDImpl struct {
	cidl.ApiOrgGroupListByOrganizationID
}

func AddOrgGroupListByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ORG_GROUP_LIST_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &OrgGroupListByOrganizationIDImpl{
				ApiOrgGroupListByOrganizationID: cidl.MakeApiOrgGroupListByOrganizationID(),
			}
		},
	)
}

func (m *OrgGroupListByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	organizationId := m.Params.OrganizationID
	ack := m.Ack
	dbCommunity := db.NewMallCommunity()
	if m.Query.Search == "" {
		ack.Count, err = dbCommunity.GroupCount(organizationId, m.Query.AuditState)
	} else {
		ack.Count, err = dbCommunity.GroupSearchCount(organizationId, m.Query.Search, m.Query.AuditState)
	}

	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "query group count failed. %s", err)
		return
	}

	if ack.Count == 0 {
		ctx.Json(ack)
		return
	}

	if m.Query.Search == "" {
		ack.List, err = dbCommunity.GroupList(organizationId, m.Query.Page, m.Query.PageSize, m.Query.AuditState, false)
	} else {
		ack.List, err = dbCommunity.GroupSearchList(organizationId, m.Query.Page, m.Query.PageSize, m.Query.Search, m.Query.AuditState, false)
	}

	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "query group list failed. %s", err)
		return
	}

	ctx.Json(ack)
}

type OrgGroupInfoByOrganizationIDByGroupIDImpl struct {
	cidl.ApiOrgGroupInfoByOrganizationIDByGroupID
}

func AddOrgGroupInfoByOrganizationIDByGroupIDHandler() {
	AddHandler(
		cidl.META_ORG_GROUP_INFO_BY_ORGANIZATION_ID_BY_GROUP_ID,
		func() http.ApiHandler {
			return &OrgGroupInfoByOrganizationIDByGroupIDImpl{
				ApiOrgGroupInfoByOrganizationIDByGroupID: cidl.MakeApiOrgGroupInfoByOrganizationIDByGroupID(),
			}
		},
	)
}

func (m *OrgGroupInfoByOrganizationIDByGroupIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	organizationId := m.Params.OrganizationID
	groupId := m.Params.GroupID
	dbCommunity := db.NewMallCommunity()
	group, err := dbCommunity.GetGroupByOrganizationIdGroupId(organizationId, groupId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get group failed. %s", err)
		return
	}

	manager, err := user.NewProxy("user-service").InnerUserInfoByUserID(group.ManagerUserId)
	if err != nil {
		ctx.Errorf(api.ErrProxyFailed, "get manager user of community group failed. %s", err)
		return
	}

	m.Ack = &cidl.AckOrgGroupInfoByOrganizationIDByGroupID{
		Group:               *group,
		ManagerWxNickname:   manager.Nickname,
		ManagerWxAvatar:     manager.Avatar,
		ManagerIdCardNumber: manager.IdCardNumber,
		ManagerIdCardFront:  manager.IdCardFront,
		ManagerIdCardBack:   manager.IdCardBack,
	}

	ctx.Json(m.Ack)
}
