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
			Prefix:  "/admin/chatgpt/user",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewChatgptUserService(),
		},
	}
	// 注册路由
	cool.RegisterController(chatgpt_user_controller)
}

// // 增加 Welcome 演示 方法
// type ChatgptUserWelcomeReq struct {
// 	g.Meta `path:"/welcome" method:"GET"`
// }
// type ChatgptUserWelcomeRes struct {
// 	*cool.BaseRes
// 	Data interface{} `json:"data"`
// }

// func (c *ChatgptUserController) Welcome(ctx context.Context, req *ChatgptUserWelcomeReq) (res *ChatgptUserWelcomeRes, err error) {
// 	res = &ChatgptUserWelcomeRes{
// 		BaseRes: cool.Ok("Welcome to Cool Admin Go"),
// 		Data:    gjson.New(`{"name": "Cool Admin Go", "age":0}`),
// 	}
// 	return
// }
