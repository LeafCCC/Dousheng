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
		Timeout:    time.Hour,                   //token有效期为1个小时
		MaxRefresh: time.Hour,                   //token更新时间为1个小时

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
		//可选的LoginResponse成功通过身份验证后Authenticator，
		//使用从返回的映射中的标识符创建 jwt 令牌PayloadFunc，
		//并将其设置为 cookie，如果SendCookie启用，则调用此函数。它用于处理任何登录后逻辑。
		//这可能类似于使用 gin 上下文将令牌的 JSON 返回给用户。

		//登陆信息的校验
		//从身份验证器返回的数据作为参数传入PayloadFunc，用于将上述用户标识符嵌入到 jwt 令牌中
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar handlers.UserParam

			loginVar.UserName = c.Query("username")
			loginVar.PassWord = c.Query("password")
			//UserParam的结构
			// type UserParam struct {
			// 	UserName string `json:"username"`
			// 	PassWord string `json:"password"`
			// }
			//获取username与password
			//这里重复登陆会有问题
			// if err := c.ShouldBind(&loginVar); err != nil {
			// 	return "", jwt.ErrMissingLoginValues
			// }

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &userdemo.CheckUserRequest{Username: loginVar.UserName, Password: loginVar.PassWord})
			//返回uid和error
		},

		//将生成的令牌作为JSON数据返回用户
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			var loginVar2 handlers.UserParam
			loginVar2.UserName = c.Query("username")
			loginVar2.PassWord = c.Query("password")
			user_id, res := rpc.QueryUser(context.Background(), &userdemo.CheckUserRequest{Username: loginVar2.UserName, Password: loginVar2.PassWord})
			// c.JSON(http.StatusOK, gin.H{
			// 	"status_code": res.ErrCode,
			// 	"status_msg":  res.ErrMsg,
			// 	"user_id":     user_id,
			// 	"token":       token,
			// })
			handlers.SendLoginResponse(c, res, user_id, loginVar2.UserName, token, expire)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt", //设置jwt的获取位置
		TokenHeadName: "Bearer",                                           //Header中的token头部字段 默认bearer即可
		TimeFunc:      time.Now,                                           //设置时间函数
	})

	//绑定路由
	v1 := r.Group("/douyin")

	// v1.POST("publish/action/", func(c *gin.Context) {

	// 	token := c.PostForm("token")
	// 	c.Header("Authorization", token)
	// 	c.Request.Header.Add("Authorization", token)
	// 	//这里注意，看你是要加到c.Header还是c.Request.Header里，注释掉不要的一个即可

	// 	//fmt.Println("c.GetHeader 的结果是 " + c.GetHeader("Authorization"))
	// 	//fmt.Println("c.Request.Header.Get的结果是" + c.Request.Header.Get("Authorization"))
	// 	newurl := "/douyin/publish/action2/" //重定向的url

	// 	c.Request.URL.Path = newurl
	// 	r.HandleContext(c)

	// })
	// v1.Use(authMiddleware.MiddlewareFunc()) //这个上面不需要认证 在下面的url开启jwt认证
	// v1.POST("publish/action2/", handlers.PublishVideo)

	user1 := v1.Group("/user")
	user1.POST("/login/", authMiddleware.LoginHandler)                       //v1/user/login
	user1.POST("/register/", handlers.Register, authMiddleware.LoginHandler) //v1/user/register
	//添加两个函数 实现登陆后验证
	user1.Use(authMiddleware.MiddlewareFunc())
	//需要调用douyin/user去获取用户信息
	user1.GET("/", handlers.GetUserInfo)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
	r.Run()
}
