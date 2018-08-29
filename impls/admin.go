package impls

import (
	"business/community/cidl"
	"business/community/common/db"
	"business/user/proxy/user"
	"common/api"

	"github.com/mz-eco/mz/http"
)

func init() {
	AddAdminGroupListByOrganizationIDHandler()
	AddAdminGroupInfoByOrganizationIDByGroupIDHandler()
}

type AdminGroupListByOrganizationIDImpl struct {
	cidl.ApiAdminGroupListByOrganizationID
}

func AddAdminGroupListByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ADMIN_GROUP_LIST_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &AdminGroupListByOrganizationIDImpl{
				ApiAdminGroupListByOrganizationID: cidl.MakeApiAdminGroupListByOrganizationID(),
			}
		},
	)
}

func (m *AdminGroupListByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	organizationId := m.Params.OrganizationID
	ack := m.Ack
	dbCommunity := db.NewMallCommunity()
	if m.Query.Search == "" {
		ack.Count, err = dbCommunity.GroupCount(organizationId, m.Query.AuditState, false) 
	} else {
		ack.Count, err = dbCommunity.GroupSearchCount(organizationId, m.Query.Search, m.Query.AuditState, false)
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
		ack.List, err = dbCommunity.GroupList(organizationId, m.Query.Page, m.Query.PageSize, m.Query.AuditState, false, false)
	} else {
		ack.List, err = dbCommunity.GroupSearchList(organizationId, m.Query.Page, m.Query.PageSize, m.Query.Search, m.Query.AuditState, false, false)
	}

	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "query group list failed. %s", err)
		return
	}

	ctx.Json(ack)
}

type AdminGroupInfoByOrganizationIDByGroupIDImpl struct {
	cidl.ApiAdminGroupInfoByOrganizationIDByGroupID
}

func AddAdminGroupInfoByOrganizationIDByGroupIDHandler() {
	AddHandler(
		cidl.META_ADMIN_GROUP_INFO_BY_ORGANIZATION_ID_BY_GROUP_ID,
		func() http.ApiHandler {
			return &AdminGroupInfoByOrganizationIDByGroupIDImpl{
				ApiAdminGroupInfoByOrganizationIDByGroupID: cidl.MakeApiAdminGroupInfoByOrganizationIDByGroupID(),
			}
		},
	)
}

func (m *AdminGroupInfoByOrganizationIDByGroupIDImpl) Handler(ctx *http.Context) {
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

	m.Ack = &cidl.AckAdminGroupInfoByOrganizationIDByGroupID{
		Group:               *group,
		ManagerWxNickname:   manager.Nickname,
		ManagerWxAvatar:     manager.Avatar,
		ManagerIdCardNumber: manager.IdCardNumber,
		ManagerIdCardFront:  manager.IdCardFront,
		ManagerIdCardBack:   manager.IdCardBack,
	}

	ctx.Json(m.Ack)
}
