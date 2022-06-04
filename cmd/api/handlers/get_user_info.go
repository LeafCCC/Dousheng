package handlers

import (
	"context"

	"github.com/LeafCCC/Dousheng/cmd/api/rpc"
	"github.com/LeafCCC/Dousheng/kitex_gen/userdemo"
	"github.com/LeafCCC/Dousheng/pkg/constants"
	"github.com/LeafCCC/Dousheng/pkg/errno"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 登陆路径调用的函数
func GetUserInfo(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	req := &userdemo.InfoGetUserRequest{UserId: userID}

	user, err := rpc.InfoGetUser(context.Background(), req) //这里调用了rpc
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendUserInfoResponse(c, errno.Success, user, false)
}
