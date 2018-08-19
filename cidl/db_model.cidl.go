package cidl

import "time"

// 社群信息完善状态
type GroupPerfectionState int

const (
	GroupPerfectionStateDefault      GroupPerfectionState = 0
	GroupPerfectionStateComplete     GroupPerfectionState = 1
	GroupPerfectionStateNeedComplete GroupPerfectionState = 2
)

func (m GroupPerfectionState) String() string {
	switch m {

	case GroupPerfectionStateDefault:
		return "GroupPerfectionStateDefault<enum GroupPerfectionState>"
	case GroupPerfectionStateComplete:
		return "GroupPerfectionStateComplete<enum GroupPerfectionState>"
	case GroupPerfectionStateNeedComplete:
		return "GroupPerfectionStateNeedComplete<enum GroupPerfectionState>"
	default:
		return "UNKNOWN_Name_<GroupPerfectionState>"
	}
}

// 社群审核状态
type GroupAuditState int

const (
	// 未审核
	GroupAuditStateWait GroupAuditState = 0
	// 审核通过
	GroupAuditStatePass GroupAuditState = 1
)

func (m GroupAuditState) String() string {
	switch m {

	case GroupAuditStateWait:
		return "GroupAuditStateWait<enum GroupAuditState>"
	case GroupAuditStatePass:
		return "GroupAuditStatePass<enum GroupAuditState>"
	default:
		return "UNKNOWN_Name_<GroupAuditState>"
	}
}

// 团购社区群体
type Group struct {
	GroupId          uint32               `db:"grp_id"`
	Name             string               `db:"name"`
	Address          string               `db:"address"`
	PostCode         string               `db:"post_code"`
	OrganizationId   uint32               `db:"org_id"`
	OrganizationName string               `db:"org_name"`
	ManagerUserId    string               `db:"manager_uid"`
	ManagerName      string               `db:"manager_name"`
	ManagerMobile    string               `db:"manager_mobile"`
	PerfectionState  GroupPerfectionState `db:"perfection_state"`
	AuditState       GroupAuditState      `db:"audit_state"`
	CreateTime       time.Time            `db:"create_time"`
}

func NewGroup() *Group {
	return &Group{}
}
