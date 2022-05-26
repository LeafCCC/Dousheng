package main

import (
	"context"
	"net/http"
	"time"

	"github.com/LeafCCC/Dousheng/cmd/api/handlers"
	"github.com/LeafCCC/Dousheng/cmd/api/rpc"
	"github.com/LeafCCC/Dousheng/kitex_gen/userdemo"
	"github.com/LeafCCC/Dousheng/pkg/constants"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey), //服务端密钥
		Timeout:    time.Hour,                   //token有效期
		MaxRefresh: time.Hour,                   //token更新时间

		//调用顺序 Authenticator->PayloadFunc->LoginResponse
		//成功登陆后使用 生成令牌
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar handlers.UserParam
			if err := c.ShouldBind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &userdemo.CheckUserRequest{Username: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	v1 := r.Group("/v1")
	user1 := v1.Group("/user")
	user1.POST("/login", authMiddleware.LoginHandler)
	user1.POST("/register", handlers.Register)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
