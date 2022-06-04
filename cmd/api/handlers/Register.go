package handlers

import (
	"context"
	"fmt"

	"github.com/LeafCCC/Dousheng/cmd/api/rpc"
	"github.com/LeafCCC/Dousheng/kitex_gen/userdemo"
	"github.com/LeafCCC/Dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

// 登陆路径调用的函数
func Register(c *gin.Context) {
	var registerVar UserParam

	// type UserParam struct {
	// 	UserName string `json:"username"`
	// 	PassWord string `json:"password"`
	// }
	// if err := c.ShouldBind(&registerVar); err != nil {
	// 	SendResponse(c, errno.ConvertErr(err), nil)
	// 	return
	// }

	// if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
	// 	SendResponse(c, errno.ParamErr, nil)
	// 	return
	// }
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	// fmt.Print(registerVar.UserName)
	// fmt.Print(registerVar.UserName)

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		fmt.Print("here wrong")
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	//使用获得的username和password 调用CreateUser
	err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

}
