package impls

import (
	"business/community/cidl"
	"business/community/common/db"
	"common/api"

	"github.com/mz-eco/mz/http"
)

func init() {
	AddInnerCommunityGroupInfoByGroupIDHandler()
	AddInnerCommunityGroupInfoByUserIDByUserIDHandler()
	AddInnerCommunityGroupCountByOrganizationIDHandler()
}

type InnerCommunityGroupInfoByGroupIDImpl struct {
	cidl.ApiInnerCommunityGroupInfoByGroupID
}

func AddInnerCommunityGroupInfoByGroupIDHandler() {
	AddHandler(
		cidl.META_INNER_COMMUNITY_GROUP_INFO_BY_GROUP_ID,
		func() http.ApiHandler {
			return &InnerCommunityGroupInfoByGroupIDImpl{
				ApiInnerCommunityGroupInfoByGroupID: cidl.MakeApiInnerCommunityGroupInfoByGroupID(),
			}
		},
	)
}

func (m *InnerCommunityGroupInfoByGroupIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	groupId := m.Params.GroupID
	dbCommunity := db.NewMallCommunity()
	m.Ack, err = dbCommunity.GetGroup(groupId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get group failed. %s", err)
		return
	}

	ctx.Json(m.Ack)
}

// 通过用户ID获取团购组织
type InnerCommunityGroupInfoByUserIDByUserIDImpl struct {
	cidl.ApiInnerCommunityGroupInfoByUserIDByUserID
}

func AddInnerCommunityGroupInfoByUserIDByUserIDHandler() {
	AddHandler(
		cidl.META_INNER_COMMUNITY_GROUP_INFO_BY_USER_ID_BY_USER_ID,
		func() http.ApiHandler {
			return &InnerCommunityGroupInfoByUserIDByUserIDImpl{
				ApiInnerCommunityGroupInfoByUserIDByUserID: cidl.MakeApiInnerCommunityGroupInfoByUserIDByUserID(),
			}
		},
	)
}

func (m *InnerCommunityGroupInfoByUserIDByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	userId := m.Params.UserID
	dbCommunity := db.NewMallCommunity()
	m.Ack, err = dbCommunity.GetGroupByUserId(userId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get group failed. %s", err)
		return
	}

	ctx.Json(m.Ack)
}

// 获取团购组织所属社团数目
type InnerCommunityGroupCountByOrganizationIDImpl struct {
	cidl.ApiInnerCommunityGroupCountByOrganizationID
}

func AddInnerCommunityGroupCountByOrganizationIDHandler() {
	AddHandler(
		cidl.META_INNER_COMMUNITY_GROUP_COUNT_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &InnerCommunityGroupCountByOrganizationIDImpl{
				ApiInnerCommunityGroupCountByOrganizationID: cidl.MakeApiInnerCommunityGroupCountByOrganizationID(),
			}
		},
	)
}

func (m *InnerCommunityGroupCountByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	organizationId := m.Params.OrganizationID
	dbCommunity := db.NewMallCommunity()
	m.Ack.Count, err = dbCommunity.GroupCount(organizationId, cidl.GroupAuditStatePass)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get audit pass group failed. %s", err)
		return
	}
	ctx.Json(m.Ack)
}
