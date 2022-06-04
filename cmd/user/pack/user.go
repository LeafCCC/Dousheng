package pack

import (
	"github.com/LeafCCC/Dousheng/cmd/user/dal/db"
	"github.com/LeafCCC/Dousheng/kitex_gen/userdemo"
)

// User pack user info
func User(u *db.User, IsFollow bool) *userdemo.User {
	if u == nil {
		return nil
	}

	//return &userdemo.User{Id: int64(u.ID), Name: u.UserName, IsFollow: isFollow}
	return &userdemo.User{Id: u.ID, Name: u.UserName, FollowCount: u.FollowCount, FollowerCount: u.FollowerCount, IsFollow: IsFollow}
}

// Users pack list of user info
// func Users(us []*db.User) []*userdemo.User {
// 	users := make([]*userdemo.User, 0)
// 	for _, u := range us {
// 		if user2 := User(u); user2 != nil {
// 			users = append(users, user2)
// 		}
// 	}
// 	return users
// }
