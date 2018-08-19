package impls

import (
	"business/agency/proxy/agency"
	"business/community/cidl"
	"business/community/common/db"
	"business/user/proxy/user"
	"common/api"

	"github.com/mz-eco/mz/http"
	"github.com/mz-eco/mz/log"
)

func init() {
	AddWxXcxGroupInfoByGroupIDHandler()
	AddWxXcxGroupAddByOrganizationIDHandler()
}

// 获取社群信息
type WxXcxGroupInfoByGroupIDImpl struct {
	cidl.ApiWxXcxGroupInfoByGroupID
}

func AddWxXcxGroupInfoByGroupIDHandler() {
	AddHandler(
		cidl.META_WX_XCX_GROUP_INFO_BY_GROUP_ID,
		func() http.ApiHandler {
			return &WxXcxGroupInfoByGroupIDImpl{
				ApiWxXcxGroupInfoByGroupID: cidl.MakeApiWxXcxGroupInfoByGroupID(),
			}
		},
	)
}

func (m *WxXcxGroupInfoByGroupIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	dbCommunity := db.NewMallCommunity()
	groupId := m.Params.GroupID
	group, err := dbCommunity.GetGroup(groupId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get group failed. %s", err)
		return
	}

	manager, err := user.NewProxy("user-service").InnerUserInfoByUserID(group.ManagerUserId)
	if err != nil {
		ctx.Errorf(api.ErrProxyFailed, "get group manager from proxy failed. %s", err)
		return
	}

	m.Ack = &cidl.AckWxXcxGroupInfoByGroupID{
		Group:           *group,
		ManagerAvatar:   manager.Avatar,
		ManagerNickname: manager.Nickname,
	}

	ctx.Json(m.Ack)
}

// 添加社群
type WxXcxGroupAddByOrganizationIDImpl struct {
	cidl.ApiWxXcxGroupAddByOrganizationID
}

func AddWxXcxGroupAddByOrganizationIDHandler() {
	AddHandler(
		cidl.META_WX_XCX_GROUP_ADD_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &WxXcxGroupAddByOrganizationIDImpl{
				ApiWxXcxGroupAddByOrganizationID: cidl.MakeApiWxXcxGroupAddByOrganizationID(),
			}
		},
	)
}

func (m *WxXcxGroupAddByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	userProxy := user.NewProxy("user-service")
	userIdOrUnionId := ctx.Session.Uid
	token := ctx.Session.Token
	ackAdd, err := userProxy.InnerUserWxXcxCmtManagerGetOrAdd(&user.AskInnerUserWxXcxCmtManagerGetOrAdd{
		UserId: userIdOrUnionId, // 这里传过来的可能是微信unionid和用户表数字ID,
		Token:  token,
		Name:   m.Ask.ManagerName,
		Mobile: m.Ask.ManagerMobile,
	})
	if err != nil {
		remoteErr, ok := err.(*http.RemoteError)
		if ok {
			switch remoteErr.Code {
			case int(user.ErrWxXcxUserExists):
				ctx.Errorf(cidl.ErrAddCommunityGroupHasExist, "user info has already exist. %s", err)
				return
			}
		}

		ctx.Errorf(api.ErrProxyFailed, "add community manager by proxy failed. %s", err)
		return
	}

	manager := ackAdd.User
	organizationId := m.Params.OrganizationID
	ackOrganization, err := agency.NewProxy("agency-service").InnerAgencyOrganizationInfoByOrganizationID(organizationId)
	if err != nil {
		ctx.Errorf(api.ErrProxyFailed, "get organization is_enable from proxy failed. %s", err)
		return
	}

	if ackOrganization == nil {
		ctx.Errorf(api.ErrServer, "no enable organization.")
		return
	}

	dbCommunity := db.NewMallCommunity()
	group := &cidl.Group{
		Address:          m.Ask.Address,
		PostCode:         m.Ask.PostCode,
		OrganizationId:   ackOrganization.OrganizationId,
		OrganizationName: ackOrganization.Name,
		ManagerUserId:    manager.UserID,
		ManagerName:      manager.Name,
		ManagerMobile:    manager.Mobile,
		PerfectionState:  cidl.GroupPerfectionStateComplete,
		AuditState:       cidl.GroupAuditStateWait,
	}
	_, err = dbCommunity.AddGroup(group)
	if err != nil {
		ctx.Errorf(api.ErrDBInsertFailed, "insert group failed. %s", err)
		return
	}

	ctx.Succeed()

	// 刷新token信息
	_, err = userProxy.InnerUserWxXcxRefreshToken(&user.AskInnerUserWxXcxRefreshToken{
		UserId: userIdOrUnionId,
		Token:  token,
	})
	if err != nil {
		log.Warnf("refresh wx_xcx token failed. %s", err)
	}
}
