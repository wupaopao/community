package impls

import (
	"business/agency/proxy/agency"
	"business/community/cidl"
	"business/community/common/db"
	"business/community/common/mq"
	"business/user/proxy/user"
	"common/api"

	"github.com/mz-eco/mz/http"
)

func init() {
	AddOrgGroupAddByOrganizationIDHandler()
	AddOrgGroupEditByOrganizationIDByGroupIDHandler()
	AddOrgGroupChangeManagerByOrganizationIDByGroupIDHandler()
	AddOrgGroupAuditByOrganizationIDByGroupIDHandler()

}

// 添加
type OrgGroupAddByOrganizationIDImpl struct {
	cidl.ApiOrgGroupAddByOrganizationID
}

func AddOrgGroupAddByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ORG_GROUP_ADD_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &OrgGroupAddByOrganizationIDImpl{
				ApiOrgGroupAddByOrganizationID: cidl.MakeApiOrgGroupAddByOrganizationID(),
			}
		},
	)
}

func (m *OrgGroupAddByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	userProxy := user.NewProxy("user-service")
	ackAdd, err := userProxy.InnerUserUserCmtManagerGetOrAdd(&user.AskInnerUserUserCmtManagerGetOrAdd{
		Name:         m.Ask.ManagerName,
		Mobile:       m.Ask.ManagerMobile,
		IdCardNumber: m.Ask.ManagerIdCardNumber,
		IdCardFront:  m.Ask.ManagerIdCardFront,
		IdCardBack:   m.Ask.ManagerIdCardBack,
	})

	if err != nil {
		ctx.ProxyErrorf(err, "get or add cmt manager from proxy failed. %s", err)
		return
	}

	manager := ackAdd.User
	organizationId := m.Params.OrganizationID
	ackOrganization, err := agency.NewProxy("agency-service").InnerAgencyOrganizationInfoByOrganizationID(organizationId)
	if err != nil {
		ctx.ProxyErrorf(err, "get organization info from proxy failed. %s", err)
		return
	}

	if ackOrganization == nil {
		ctx.Errorf(api.ErrServer, "no enable organization.")
		return
	}

	dbCommunity := db.NewMallCommunity()
	group := &cidl.Group{
		Name:             m.Ask.GroupName,
		Address:          m.Ask.GroupAddress,
		PostCode:         m.Ask.GroupPostCode,
		OrganizationId:   ackOrganization.OrganizationId,
		OrganizationName: ackOrganization.Name,
		ManagerUserId:    manager.UserID,
		ManagerName:      manager.Name,
		ManagerMobile:    manager.Mobile,
		PerfectionState:  cidl.GroupPerfectionStateComplete,
		AuditState:       cidl.GroupAuditStatePass,
	}

	result, err := dbCommunity.AddGroup(group)
	if err != nil {
		ctx.Errorf(api.ErrDBInsertFailed, "insert group failed. %s", err)
		return
	}

	groupId, err := result.LastInsertId()
	if err != nil {
		ctx.Errorf(api.ErrServer, "get add group last insertId failed. %s", err)
		return
	}

	// 广播审核消息
	topic, err := mq.GetTopicServiceCommunityServce()
	if err != nil {
		ctx.Errorf(api.ErrMqConnectFailed, "get topic service-community-service failed. %s", err)
		return
	}

	err = topic.ChangeGroupAuditState(&mq.ChangeGroupAuditStateMessage{
		GroupId: uint32(groupId),
	})

	if err != nil {
		ctx.Errorf(api.ErrMqPublishFailed, "public change group audit state message failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 编辑
type OrgGroupEditByOrganizationIDByGroupIDImpl struct {
	cidl.ApiOrgGroupEditByOrganizationIDByGroupID
}

func AddOrgGroupEditByOrganizationIDByGroupIDHandler() {
	AddHandler(
		cidl.META_ORG_GROUP_EDIT_BY_ORGANIZATION_ID_BY_GROUP_ID,
		func() http.ApiHandler {
			return &OrgGroupEditByOrganizationIDByGroupIDImpl{
				ApiOrgGroupEditByOrganizationIDByGroupID: cidl.MakeApiOrgGroupEditByOrganizationIDByGroupID(),
			}
		},
	)
}

func (m *OrgGroupEditByOrganizationIDByGroupIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	organizationId := m.Params.OrganizationID
	groupId := m.Params.GroupID
	dbCommunity := db.NewMallCommunity()

	group, err := dbCommunity.GetGroupByOrganizationIdGroupId(organizationId, groupId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get organization failed. %s", err)
		return
	}

	managerUid := group.ManagerUserId

	// 更改管理员
	userService := user.NewProxy("user-service")
	if m.Ask.ManagerMobile != group.ManagerMobile {
		ackGet, errGet := userService.InnerUserUserCmtManagerGetOrAdd(&user.AskInnerUserUserCmtManagerGetOrAdd{
			Name:         m.Ask.ManagerName,
			Mobile:       m.Ask.ManagerMobile,
			IdCardNumber: m.Ask.ManagerIdCardNumber,
			IdCardFront:  m.Ask.ManagerIdCardFront,
			IdCardBack:   m.Ask.ManagerIdCardBack,
		})
		if errGet != nil {
			err = errGet
			ctx.ProxyErrorf(errGet, "get or add cmt manager failed. %s", err)
			return
		}

		// 解绑旧管理员
		_, err = userService.InnerUserUserCmtManagerUnbind(&user.AskInnerUserUserCmtManagerUnbind{
			OldManagerUid: managerUid,
		})
		if err != nil {
			ctx.ProxyErrorf(err, "unbind old cmt manager failed. %s", err)
			return
		}

		// 赋值新的管理员
		managerUid = ackGet.User.UserID

	} else {
		_, err = userService.InnerUserUserCmtManagerUpdateByUserID(managerUid, &user.AskInnerUserUserCmtManagerUpdateByUserID{
			Name:         m.Ask.ManagerName,
			IdCardNumber: m.Ask.ManagerIdCardNumber,
			IdCardFront:  m.Ask.ManagerIdCardFront,
			IdCardBack:   m.Ask.ManagerIdCardBack,
		})
		if err != nil {
			ctx.ProxyErrorf(err, "update cmt manager by user id failed. %s", err)
			return
		}

	}

	strSql := `
		UPDATE cmt_group
		SET
			name=?,
			address=?,
			post_code=?,
			manager_uid=?,
			manager_name=?,
			manager_mobile=?
		WHERE
			org_id=?
			AND grp_id=?
	`
	_, err = dbCommunity.DB.Exec(strSql,
		m.Ask.GroupName,
		m.Ask.GroupAddress,
		m.Ask.GroupPostCode,
		managerUid,
		m.Ask.ManagerName,
		m.Ask.ManagerMobile,
		organizationId,
		groupId)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update group failed. %s", err)
		return
	}

	// 广播社群名称更改
	topic, err := mq.GetTopicServiceCommunityServce()
	if err != nil {
		ctx.Errorf(api.ErrMqConnectFailed, "get topic service-community-service failed. %s", err)
		return
	}

	err = topic.ModifyGroupInfo(&mq.ModifyGroupInfoMessage{
		GroupId: groupId,
		Values: map[string]interface{}{
			"name":    m.Ask.GroupName,
			"address": m.Ask.GroupAddress,
		},
	})

	if err != nil {
		ctx.Errorf(api.ErrMqPublishFailed, "publish topic message failed. %s", err)
		return
	}

	ctx.Succeed()

}

// 更改管理员
type OrgGroupChangeManagerByOrganizationIDByGroupIDImpl struct {
	cidl.ApiOrgGroupChangeManagerByOrganizationIDByGroupID
}

func AddOrgGroupChangeManagerByOrganizationIDByGroupIDHandler() {
	AddHandler(
		cidl.META_ORG_GROUP_CHANGE_MANAGER_BY_ORGANIZATION_ID_BY_GROUP_ID,
		func() http.ApiHandler {
			return &OrgGroupChangeManagerByOrganizationIDByGroupIDImpl{
				ApiOrgGroupChangeManagerByOrganizationIDByGroupID: cidl.MakeApiOrgGroupChangeManagerByOrganizationIDByGroupID(),
			}
		},
	)
}

func (m *OrgGroupChangeManagerByOrganizationIDByGroupIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	organizationId := m.Params.OrganizationID
	groupId := m.Params.GroupID

	dbCommunity := db.NewMallCommunity()
	group, err := dbCommunity.GetGroupByOrganizationIdGroupId(organizationId, groupId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get organization failed. %s", err)
		return
	}

	userProxy := user.NewProxy("user-service")
	newManager, err := userProxy.InnerUserInfoByUserID(m.Ask.NewManagerUid)
	if err != nil {
		ctx.Errorf(api.ErrProxyFailed, "get user info failed. %s", err)
		return
	}

	if newManager.IsCmtManager {
		ctx.Errorf(api.ErrServer, "user is already community manager")
		return
	}

	// 更新社群管理员信息
	strSql := `
		UPDATE cmt_group
		SET
			manager_uid=?,
			manager_name=?,
			manager_mobile=?
		WHERE grp_id=? AND org_id=?
	`
	_, err = dbCommunity.DB.Exec(strSql,
		newManager.UserID,
		newManager.Name,
		newManager.Mobile,
		groupId,
		organizationId,
	)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update group manager info failed. %s", err)
		return
	}

	// 更改绑定
	_, err = userProxy.InnerUserUserCmtManagerChange(&user.AskInnerUserUserCmtManagerChange{
		NewManagerUid: m.Ask.NewManagerUid,
		OldManagerUid: group.ManagerUserId,
	})

	if err != nil {
		ctx.Errorf(api.ErrProxyFailed, "change community manager failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 审核
type OrgGroupAuditByOrganizationIDByGroupIDImpl struct {
	cidl.ApiOrgGroupAuditByOrganizationIDByGroupID
}

func AddOrgGroupAuditByOrganizationIDByGroupIDHandler() {
	AddHandler(
		cidl.META_ORG_GROUP_AUDIT_BY_ORGANIZATION_ID_BY_GROUP_ID,
		func() http.ApiHandler {
			return &OrgGroupAuditByOrganizationIDByGroupIDImpl{
				ApiOrgGroupAuditByOrganizationIDByGroupID: cidl.MakeApiOrgGroupAuditByOrganizationIDByGroupID(),
			}
		},
	)
}

func (m *OrgGroupAuditByOrganizationIDByGroupIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	organizationId := m.Params.OrganizationID
	groupId := m.Params.GroupID

	dbCommunity := db.NewMallCommunity()

	strSql := `
		UPDATE cmt_group
		SET audit_state=?
		WHERE org_id=? AND grp_id=?
	`

	_, err = dbCommunity.DB.Exec(strSql, m.Ask.AuditState, organizationId, groupId)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update group failed. %s", err)
		return
	}

	// 广播审核消息
	topic, err := mq.GetTopicServiceCommunityServce()
	if err != nil {
		ctx.Errorf(api.ErrMqConnectFailed, "get topic service-community-service failed. %s", err)
		return
	}

	err = topic.ChangeGroupAuditState(&mq.ChangeGroupAuditStateMessage{
		GroupId: groupId,
	})

	if err != nil {
		ctx.Errorf(api.ErrMqPublishFailed, "publish topic message failed. %s", err)
		return
	}

	ctx.Succeed()
}
