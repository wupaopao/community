package impls

import (
	"business/auth"
	//for test
	"fmt"

	"github.com/mz-eco/mz/http"
)

var AccessControlHandlers []http.AccessControlHandler

func AddAccessControlHandler(handler http.AccessControlHandler) {
	//for test
	fmt.Println("add access control handler")
	AccessControlHandlers = append(AccessControlHandlers, handler)
}

func init() {

	// 白名单
	whiteListAuthHandler := auth.NewWhiteListAuthHandler([]string{
		"/api/v1/inner/",
	})

	// 微信小程序端访问控制
	wxXcxAuthHandler := auth.NewWxXcxAuthHandler([]string{
		"/api/v1/community/wx_xcx/",
	}, []string{})

	// 城市合伙人端访问控制
	orgAuthHandler := auth.NewOrgAuthHandler([]string{
		"/api/v1/community/org/",
	}, []string{})

	// 运营端访问控制
	adminAuthHandler := auth.NewAdminAuthHandler([]string{
		"/api/v1/community/admin/",
	}, []string{})

	chanHandlers := &auth.ChainAuthHandler{
		AuthHandlers: []*auth.AuthHandler{
			whiteListAuthHandler,
			wxXcxAuthHandler,
			orgAuthHandler,
			adminAuthHandler,
		},
	}

	AddAccessControlHandler(chanHandlers.AccessControlHandler)
}
