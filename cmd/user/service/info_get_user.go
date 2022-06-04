// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service

import (
	"context"

	"github.com/LeafCCC/Dousheng/cmd/user/dal/db"
	"github.com/LeafCCC/Dousheng/kitex_gen/userdemo"
)

type InfoGetUserService struct {
	ctx context.Context
}

func NewInfoGetUserService(ctx context.Context) *InfoGetUserService {
	return &InfoGetUserService{ctx: ctx}
}

func (s *InfoGetUserService) InfoGetUser(req *userdemo.InfoGetUserRequest) (user *db.User, err error) {
	user, err = db.QueryUserId(s.ctx, req.UserId)

	if err != nil {
		return nil, err
	}

	return user, nil
}
