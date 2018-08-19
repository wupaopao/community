package db

import (
	"database/sql"
	"errors"
	"fmt"

	"business/community/cidl"

	"github.com/mz-eco/mz/conn"
	"github.com/mz-eco/mz/log"
)

type MallCommunity struct {
	DB *conn.DB
}

func NewMallCommunity() *MallCommunity {
	return &MallCommunity{
		DB: conn.NewDB("mal_community"),
	}
}

func (m *MallCommunity) GetGroup(groupId uint32) (group *cidl.Group, err error) {
	group = &cidl.Group{}
	strSql := `
		SELECT
			grp_id,
			name,
			address,
			post_code,
			org_id,
			org_name,
			manager_uid,
			manager_name,
			manager_mobile,
			perfection_state,
			audit_state,
			create_time
		FROM
			cmt_group
		WHERE grp_id=?
	`
	err = m.DB.Get(group, strSql, groupId)
	if err != nil {
		group = nil
		if err != conn.ErrNoRows {
			log.Warnf("get community group from db failed. %s", err)
			return
		}
	}

	return
}

func (m *MallCommunity) GetGroupByUserId(userId string) (group *cidl.Group, err error) {
	group = &cidl.Group{}
	strSql := `
		SELECT
			grp_id,
			name,
			address,
			post_code,
			org_id,
			org_name,
			manager_uid,
			manager_name,
			manager_mobile,
			perfection_state,
			audit_state,
			create_time
		FROM
			cmt_group
		WHERE manager_uid=?
	`
	err = m.DB.Get(group, strSql, userId)
	if err != nil {
		group = nil
		if err != conn.ErrNoRows {
			log.Warnf("get community group from db failed. %s", err)
			return
		}
	}

	return
}

func (m *MallCommunity) GetGroupByOrganizationIdGroupId(organizationId uint32, groupId uint32) (group *cidl.Group, err error) {
	group = &cidl.Group{}
	strSql := `
		SELECT
			grp_id,
			name,
			address,
			post_code,
			org_id,
			org_name,
			manager_uid,
			manager_name,
			manager_mobile,
			perfection_state,
			audit_state,
			create_time
		FROM
			cmt_group
		WHERE grp_id=? AND org_id=?
	`
	err = m.DB.Get(group, strSql, groupId, organizationId)
	if err != nil {
		group = nil
		if err != conn.ErrNoRows {
			log.Warnf("get community group from db failed. %s", err)
			return
		}
	}

	return
}

func (m *MallCommunity) AddGroup(group *cidl.Group) (result sql.Result, err error) {
	strSql := `
		INSERT INTO
			cmt_group (
				name,
				address,
				post_code,
				org_id,
				org_name,
				manager_uid,
				manager_name,
				manager_mobile,
				perfection_state,
				audit_state
			)
		VALUES (
				:name,
				:address,
				:post_code,
				:org_id,
				:org_name,
				:manager_uid,
				:manager_name,
				:manager_mobile,
				:perfection_state,
				:audit_state
		)
	`
	result, err = m.DB.NamedExec(strSql, group)
	if err != nil {
		log.Warnf("insert user failed. %s", err)
		return
	}

	return
}

func (m *MallCommunity) GroupCount(organizationId uint32, auditState cidl.GroupAuditState) (count uint32, err error) {
	countSql := `SELECT COUNT(*) FROM cmt_group WHERE audit_state=? AND org_id=?`
	err = m.DB.Get(&count, countSql, auditState, organizationId)
	return
}

func (m *MallCommunity) GroupList(organizationId uint32, page uint32, pageSize uint32, auditState cidl.GroupAuditState, idAsc bool) (groups []*cidl.Group, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}
	listSql := `
		SELECT
			grp_id,
			name,
			address,
			post_code,
			org_id,
			org_name,
			manager_uid,
			manager_name,
			manager_mobile,
			perfection_state,
			audit_state,
			create_time
		FROM cmt_group
		WHERE audit_state=? AND org_id=?
		ORDER BY grp_id %s
		LIMIT ? OFFSET ?
	`
	listSql = fmt.Sprintf(listSql, strOrderBy)
	rows, err := m.DB.Query(listSql, auditState, organizationId, pageSize, offset)
	if err != nil {
		log.Warnf("query community group list failed. %s", err)
		return
	}

	for rows.Next() {
		var group cidl.Group
		err = rows.StructScan(&group)
		if err != nil {
			log.Warnf("scan community group failed. %s", err)
			return
		}

		groups = append(groups, &group)
	}

	return
}

func (m *MallCommunity) GroupSearchCount(organizationId uint32, search string, auditState cidl.GroupAuditState) (count uint32, err error) {
	countSql := `SELECT COUNT(*) FROM cmt_group WHERE audit_state=? AND org_id=? AND (manager_name LIKE ? OR manager_mobile LIKE ?)`
	err = m.DB.Get(&count, countSql, auditState, organizationId, search, search)
	return
}

func (m *MallCommunity) GroupSearchList(organizationId uint32, page uint32, pageSize uint32, search string, auditState cidl.GroupAuditState, idAsc bool) (groups []*cidl.Group, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}
	listSql := `
		SELECT
			grp_id,
			name,
			address,
			post_code,
			org_id,
			org_name,
			manager_uid,
			manager_name,
			manager_mobile,
			perfection_state,
			audit_state,
			create_time
		FROM cmt_group
		WHERE audit_state=? AND org_id=? AND (manager_name LIKE ? OR manager_mobile LIKE ?)
		ORDER BY grp_id %s
		LIMIT ? OFFSET ?
	`
	listSql = fmt.Sprintf(listSql, strOrderBy)
	search = "%" + search + "%"
	rows, err := m.DB.Query(listSql, auditState, organizationId, search, search, pageSize, offset)
	if err != nil {
		log.Warnf("query community group list failed. %s", err)
		return
	}

	for rows.Next() {
		var group cidl.Group
		err = rows.StructScan(&group)
		if err != nil {
			log.Warnf("scan community group failed. %s", err)
			return
		}

		groups = append(groups, &group)
	}

	return
}
