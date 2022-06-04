package rpc

import (
	"context"
	"time"

	"github.com/LeafCCC/Dousheng/kitex_gen/userdemo"
	"github.com/LeafCCC/Dousheng/kitex_gen/userdemo/userservice"
	"github.com/LeafCCC/Dousheng/pkg/constants"
	"github.com/LeafCCC/Dousheng/pkg/errno"
	"github.com/LeafCCC/Dousheng/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// 创建用户
func CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) error {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg) //这里注意StatusCode是int32
	}
	return nil
}

// 校验用户 返回UserId
func CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserId, nil
}

// 校验用户 返回UserId 以及查询的状态玛
func QueryUser(ctx context.Context, req *userdemo.CheckUserRequest) (int64, errno.ErrNo) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, errno.ConvertErr(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserId, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
}

//获取用户信息
// InfoGet get user info
func InfoGetUser(ctx context.Context, req *userdemo.InfoGetUserRequest) (*userdemo.User, error) {
	resp, err := userClient.InfoGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.User, nil
}

//
