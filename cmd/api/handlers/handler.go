package handlers

import (
	"net/http"
	"time"

	"github.com/LeafCCC/Dousheng/kitex_gen/userdemo"
	"github.com/LeafCCC/Dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type Response struct {
	Code    int32       `json:"status_code"`
	Message string      `json:"status_msg"`
	Data    interface{} `json:"data"`
}

type UserInfoResponse struct {
	Code    int32       `json:"status_code"`
	Message string      `json:"status_msg"`
	User    interface{} `json:"user"`
}

type UserLoginResponse struct {
	Code     int32  `json:"status_code"`
	Message  string `json:"status_msg"`
	UserId   int64  `json:"user_id"`
	UserName string `json:"name"`
	Token    string `json:"token"`
	Expire   string `json:"expire"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

//发送用户信息
func SendUserInfoResponse(c *gin.Context, err error, user *userdemo.User, isFollow bool) {
	Err := errno.ConvertErr(err)
	theuser := map[string]interface{}{"user": user}
	c.JSON(http.StatusOK, UserInfoResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		User:    theuser["user"],
	})
}

func SendLoginResponse(c *gin.Context, err error, userId int64, userName string, token string, expire time.Time) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserLoginResponse{
		Code:     Err.ErrCode,
		Message:  Err.ErrMsg,
		UserId:   userId,
		UserName: userName,
		Token:    token,
		Expire:   expire.Format(time.RFC3339),
	})
}

// func PublishVideo(c *gin.Context) {
// 	fmt.Println("new video")
// 	token := c.Query("token")
// 	fmt.Println("token is" + token)
// 	headerTest := c.Request.Header.Get("Authorization")

// 	headerTest2 := c.GetHeader("Authorization")
// 	fmt.Println("重定向的Request.head token is" + headerTest)
// 	fmt.Println("重定向的head token is" + headerTest2)

// }
