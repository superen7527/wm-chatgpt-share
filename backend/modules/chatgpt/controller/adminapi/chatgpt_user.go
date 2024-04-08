package admin

import (
	"backend/modules/chatgpt/service"

	"github.com/cool-team-official/cool-admin-go/cool"
)

type ChatgptUserController struct {
	*cool.Controller
}

func init() {
	var chatgpt_user_controller = &ChatgptUserController{
		&cool.Controller{
			Prefix:  "/adminapi/chatgpt/user",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewChatgptUserService(),
		},
	}
	// 注册路由
	cool.RegisterController(chatgpt_user_controller)
}
