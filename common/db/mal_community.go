package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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

func (m *MallCommunity) GroupCount(organizationId uint32, auditState cidl.GroupAuditState, isDisable bool) (count uint32, err error) {

	if isDisable {
		countSql := `SELECT COUNT(*) FROM cmt_group WHERE org_id=? and is_disable=1`
		err = m.DB.Get(&count, countSql, organizationId)
		return
	}else {
		countSql := `SELECT COUNT(*) FROM cmt_group WHERE audit_state=? AND org_id=? and is_disable=0`
		err = m.DB.Get(&count, countSql, auditState, organizationId)
		return
	}
}

func (m *MallCommunity) GroupList(organizationId uint32, page uint32, pageSize uint32, auditState cidl.GroupAuditState, isDisable bool, idAsc bool) (groups []*cidl.Group, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}

	var cond string
	if isDisable {
		cond = " org_id=? AND is_disable=1"
	}else {
		cond = "audit_state=? AND org_id=? AND is_disable=0"
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
			create_time,
			is_disable
		FROM cmt_group
		WHERE %s 
		ORDER BY grp_id %s
		LIMIT ? OFFSET ?
	`
	listSql = fmt.Sprintf(listSql, cond, strOrderBy )
	//for test 
	fmt.Println(listSql)

	var args []interface{}
	if !isDisable {
		args = append(args,auditState)
	}
	args = append(args,organizationId,pageSize,offset)
	rows, err := m.DB.Query(listSql, args...)
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

func (m *MallCommunity) GroupSearchCount(organizationId uint32, search string, auditState cidl.GroupAuditState, isDisable bool) (count uint32, err error) {

	search = "%" + search + "%"
	if isDisable {
		countSql := `SELECT COUNT(*) FROM cmt_group WHERE  org_id=? and is_disable=0 AND (manager_name LIKE ? OR manager_mobile LIKE ?)`
		err = m.DB.Get(&count, countSql,  organizationId, search, search)
		return
	}else {
		countSql := `SELECT COUNT(*) FROM cmt_group WHERE audit_state=? AND org_id=? and is_disable=0 AND (manager_name LIKE ? OR manager_mobile LIKE ?)`
		err = m.DB.Get(&count, countSql, auditState, organizationId, search, search)
		return
	}
}

func (m *MallCommunity) GroupSearchList(organizationId uint32, page uint32, pageSize uint32, search string, auditState cidl.GroupAuditState, isDisable bool, idAsc bool) (groups []*cidl.Group, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}

	var cond string
	if isDisable {
		cond = " org_id=? AND is_disable=1"
	}else {
		cond = "audit_state=? AND org_id=? AND is_disable=0"
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
			create_time,
			is_disable
		FROM cmt_group
		WHERE %s
		AND (manager_name LIKE ? OR manager_mobile LIKE ?)
		ORDER BY grp_id %s
		LIMIT ? OFFSET ?
	`
	listSql = fmt.Sprintf(listSql, cond, strOrderBy)
	search = "%" + search + "%"

	var args []interface{}
	if !isDisable {
		args = append(args,auditState)
	}
	args = append(args,organizationId,search,search,pageSize,offset)

	rows, err := m.DB.Query(listSql, args...)
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

func (m *MallCommunity) TeamCount(organizationId uint32) (count uint32, err error) {
	countSql := `SELECT COUNT(*) FROM cmt_team WHERE org_id=?`
	err = m.DB.Get(&count, countSql, organizationId)
	return
}

func (m *MallCommunity) TeamList(organizationId uint32, page uint32, pageSize uint32) (teams []*cidl.Team, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0")
		return
	}

	offset := (page - 1) * pageSize
	listSql := `
		SELECT
			a.team_id,
			a.name,
			count(b.grp_id)
		FROM cmt_team a 
		LEFT JOIN cmt_team_group b
		ON a.team_id = b.team_id
		WHERE org_id=?
		GROUP BY a.team_id
		ORDER BY a.team_id DESC 
		LIMIT ? OFFSET ?
	`
	rows, err := m.DB.Query(listSql, organizationId, pageSize, offset)
	if err != nil {
		log.Warnf("query community team list failed. %s", err)
		return
	}

	for rows.Next() {
		var team cidl.Team
		err = rows.Scan(&team.TeamId,&team.Name,&team.GroupCount)
		if err != nil {
			log.Warnf("scan community team failed. %s", err)
			return
		}

		teams = append(teams, &team)
	}

	return
}


func (m *MallCommunity) AddTeam(organizationId uint32, name string) (result sql.Result, err error) {
	strSql := `
		INSERT INTO
			cmt_team (
				name,
				org_id
			)
		VALUES (
				?,
				?
		)
	`
	result, err = m.DB.Exec(strSql, name, organizationId)
	if err != nil {
		log.Warnf("insert team failed. %s", err)
		return
	}

	return
}

func (m *MallCommunity) UpdateTeam(teamId uint32, name string) (result sql.Result, err error) {
	strSql := `
		UPDATE	
			cmt_team 
		SET name = ?
		WHERE
			team_id = ?
	`
	result, err = m.DB.Exec(strSql, name, teamId)
	if err != nil {
		log.Warnf("update team failed. %s", err)
		return
	}

	return
}

func (m *MallCommunity) DeleteTeam(teamId uint32) (result sql.Result, err error) {
	strSql := `
		DELETE FROM	
			cmt_team 
		WHERE
			team_id = ?
	`
	result, err = m.DB.Exec(strSql, teamId)
	if err != nil {
		log.Warnf("delete team failed. %s", err)
		return
	}

	strSql = `
		DELETE FROM	
			cmt_team_group
		WHERE
			team_id = ?
	`
	result, err = m.DB.Exec(strSql, teamId)
	if err != nil {
		log.Warnf("delete team group failed. %s", err)
		return
	}

	return
}

func (m *MallCommunity) TeamBindGroups(teamId uint32, groupIds []uint32) (result sql.Result, err error) {
	strSql := `
		DELETE FROM	
			cmt_team_group
		WHERE
			team_id = ?
	`
	result, err = m.DB.Exec(strSql, teamId)
	if err != nil {
		log.Warnf("delete team group failed. %s", err)
		return
	}

	strSql = `
		INSERT INTO cmt_team_group 
			(
				team_id, 
				grp_id
			)
	    	values	
			%s

		`

	var args []interface{}
	var sliceStrValue []string
	for _, groupId := range groupIds {
		sliceStrValue = append(sliceStrValue, "(?, ?)")
		args = append(args, teamId)
		args = append(args, groupId)
	}

	strValues := strings.Join(sliceStrValue, ",")
	strSql = fmt.Sprintf(strSql, strValues)
	result, err = m.DB.Exec(strSql, args...)
	return
}



func (m *MallCommunity) TeamGroups(orgId uint32, teamId uint32, isBinded bool, key string) (groups []*cidl.SubGroup, err error) {
	var strBind string
	if isBinded {
		strBind = "is not null"
	}else {
		strBind = "is null"
	}
	strSql := `
		SELECT a.grp_id,a.name,a.manager_name
		FROM	
			cmt_group a
		LEFT JOIN
			cmt_team_group b
		ON	
			a.grp_id = b.grp_id
		AND
			b.team_id = ?
		WHERE 
			a.org_id = ?
		AND 
			b.grp_id %s
	`

	var args []interface{}
	args = append(args, teamId, orgId)
	
	if key != "" {
		strSql = strSql + " AND (a.name LIKE ? OR a.manager_name LIKE ?)"
		like := "%" + key + "%"
		args = append(args, like, like)
	}

	strSql = fmt.Sprintf(strSql, strBind)
	//for test
	fmt.Println(strSql)
	rows, err := m.DB.Query(strSql, args...)
	if err != nil {
		log.Warnf("query team groups failed. %s", err)
		return
	}
	for rows.Next() {
		var group cidl.SubGroup
		err = rows.Scan(&group.GroupId,&group.Name,&group.ManagerName)
		if err != nil {
			log.Warnf("scan team groups failed. %s", err)
			return
		}
		groups = append(groups, &group)
	}

	return
}

func (m *MallCommunity) GetGroupTeamIds(groupId uint32) (teamIds []uint32, err error) {
	strSql := `
		SELECT
			team_id
		FROM
			cmt_team_group
		WHERE grp_id=?
	`
	rows, err := m.DB.Query(strSql, groupId)
	if err != nil {
		log.Warnf("query group teams failed. %s", err)
		return
	}
	for rows.Next() {
		var teamId uint32
		err = rows.Scan(&teamId)
		if err != nil {
			log.Warnf("scan group teams failed. %s", err)
			return 
		}
		teamIds = append(teamIds, teamId)
	}

	return
}

