package community

// 获取团购组织
func (m *Proxy) InnerCommunityGroupInfoByGroupID(GroupID uint32,
) (*Group, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *Group
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/inner/community/group/info/:group_id",
		nil,
		ack,
		map[string]interface{}{
			"group_id": GroupID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

// 通过用户ID获取团购组织
func (m *Proxy) InnerCommunityGroupInfoByUserIDByUserID(UserID string,
) (*Group, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *Group
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/inner/community/group/info_by_user_id/:user_id",
		nil,
		ack,
		map[string]interface{}{
			"user_id": UserID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

type AckInnerCommunityGroupCountByOrganizationID struct {
	Count uint32 `db:"Count"`
}

func NewAckInnerCommunityGroupCountByOrganizationID() *AckInnerCommunityGroupCountByOrganizationID {
	return &AckInnerCommunityGroupCountByOrganizationID{}
}

// 获取团购组织所属社团数目
func (m *Proxy) InnerCommunityGroupCountByOrganizationID(OrganizationID uint32,
) (*AckInnerCommunityGroupCountByOrganizationID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckInnerCommunityGroupCountByOrganizationID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/inner/community/group/count/:organization_id",
		nil,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

type AckInnerCommunityGroupTeamByGroupID struct {
	TeamIDs []uint32 `db:"TeamIDs"`
}

func NewAckInnerCommunityGroupTeamByGroupID() *AckInnerCommunityGroupTeamByGroupID {
	return &AckInnerCommunityGroupTeamByGroupID{
		TeamIDs: make([]uint32, 0),
	}
}

// 获取社团所属的组
func (m *Proxy) InnerCommunityGroupTeamByGroupID(GroupID uint32,
) (*AckInnerCommunityGroupTeamByGroupID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckInnerCommunityGroupTeamByGroupID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/inner/community/group/team/:group_id",
		nil,
		ack,
		map[string]interface{}{
			"group_id": GroupID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}
