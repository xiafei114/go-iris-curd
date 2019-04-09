package vo

import (
	modeSys "go-iris-curd/main/web/models/system"
)

// 前端需要的数据结构
type UserVO struct {
	*modeSys.User
	Token string `json:"token"`
}

// 携带token
func BuildUserVO(token string, user *modeSys.User) (uVO *UserVO) {
	uVO = &UserVO{
		user,
		token,
	}
	return
}

// 用户列表，不带token
func BuildUserVOList(userList ...*modeSys.User) (userVOList []*UserVO) {
	for _, user := range userList {
		userVOList = append(userVOList, BuildUserVO("", user))
	}
	return
}
