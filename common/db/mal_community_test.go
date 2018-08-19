package db

import (
	"business/community/cidl"
	"business/user/proxy/user"

	"fmt"
	"os"
	"testing"
	"time"

	"github.com/mz-eco/mz/settings"
)

func TestMain(m *testing.M) {
	settings.LoadFrom("../../", "")
	os.Exit(m.Run())
}

func TestMallCommunity_AddGroup(t *testing.T) {
	group := &cidl.Group{
		Name:             fmt.Sprintf("社群-%d", 1),
		Address:          "深圳市南山区宝深路科陆大厦",
		OrganizationId:   1,
		OrganizationName: "味罗天下",
		ManagerUserId:    "2",
		ManagerName:      "小花",
		ManagerMobile:    "18676726609",
		PerfectionState:  cidl.GroupPerfectionStateComplete,
		AuditState:       cidl.GroupAuditStatePass,
		CreateTime:       time.Now(),
	}
	dbCommunity := NewMallCommunity()
	_, err := dbCommunity.AddGroup(group)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMallCommunity_AddGroup_all(t *testing.T) {
	for i := 105; i <= 304; i++ {
		uid := fmt.Sprintf("%d", i)
		groupManager, err := user.NewProxy("user-service").InnerUserInfoByUserID(uid)
		if err != nil {

			continue
		}

		group := &cidl.Group{
			Name:             fmt.Sprintf("社群-%d", i),
			Address:          "深圳市南山区宝深路科陆大厦",
			OrganizationId:   1,
			OrganizationName: "味罗天下",
			ManagerUserId:    groupManager.UserID,
			ManagerName:      groupManager.Name,
			ManagerMobile:    groupManager.Mobile,
			PerfectionState:  cidl.GroupPerfectionStateComplete,
			AuditState:       cidl.GroupAuditStatePass,
			CreateTime:       time.Now(),
		}
		dbCommunity := NewMallCommunity()
		_, err = dbCommunity.AddGroup(group)
		if err != nil {
			t.Error(err)
			return
		}
	}
}
