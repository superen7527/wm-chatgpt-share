package admin

import (
	"backend/modules/chatgpt/service"

	"github.com/cool-team-official/cool-admin-go/cool"
)

type ChatgptSessionController struct {
	*cool.Controller
}

func init() {
	var chatgpt_session_controller = &ChatgptSessionController{
		&cool.Controller{
			Prefix:  "/adminapi/chatgpt/session",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewChatgptSessionService(),
		},
	}
	// 注册路由
	cool.RegisterController(chatgpt_session_controller)
}
