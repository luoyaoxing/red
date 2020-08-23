package user

import "gitlab.kay.com/red/proto/user"

type User struct {
	UserId uint64 `xorm:"pk autoincr"`
	UserName string
	NickName string
}

func (u *User) ToPROTO() *red_proto_user.User  {
	up := new(red_proto_user.User)
	up.UserId = u.UserId
	up.UserName = u.UserName
	up.NickName = u.NickName
	return up
}

func (u *User) FromPROTO(up *red_proto_user.User)  {
	u.UserId = up.UserId
	u.UserName = up.UserName
	u.NickName = up.NickName
}
