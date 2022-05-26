package handlers

import (
	"context"

	"github.com/LeafCCC/Dousheng/cmd/api/rpc"
	"github.com/LeafCCC/Dousheng/kitex_gen/userdemo"
	"github.com/LeafCCC/Dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

// Register register user info
func Register(c *gin.Context) {
	var registerVar UserParam
	if err := c.ShouldBind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
