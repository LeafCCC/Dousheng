package main

import (
	"context"

	"github.com/LeafCCC/Dousheng/cmd/user/pack"
	"github.com/LeafCCC/Dousheng/cmd/user/service"
	"github.com/LeafCCC/Dousheng/kitex_gen/userdemo"
	"github.com/LeafCCC/Dousheng/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (resp *userdemo.CreateUserResponse, err error) {
	resp = new(userdemo.CreateUserResponse)

	//判断用户名与密码是否为空
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	//调用service目录下面的函数CreateUser
	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		//报错与否 调用时传入参数不同
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (resp *userdemo.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.CheckUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// InfoGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) InfoGetUser(ctx context.Context, req *userdemo.InfoGetUserRequest) (*userdemo.InfoGetUserResponse, error) {
	resp := new(userdemo.InfoGetUserResponse)

	//这里判断 防止出现userid小于1的错误情况
	if req.UserId < 1 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	user, err := service.NewInfoGetUserService(ctx).InfoGetUser(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.User = pack.User(user, false) // 这里返回了数据库来的查询数据

	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *userdemo.MGetUserRequest) (resp *userdemo.MGetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *userdemo.UpdateUserRequest) (resp *userdemo.UpdateUserResponse, err error) {
	// TODO: Your code here...
	return
}
