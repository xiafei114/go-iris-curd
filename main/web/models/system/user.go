package system

import (
	"go-iris-curd/main/web/db"
	"go-iris-curd/main/web/supports"
	"time"
)

// mysql user table
type User struct {
	Id         int64     `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	Username   string    `xorm:"notnull" json:"username" form:"username"`
	Password   string    `xorm:"notnull" json:"password" form:"password"`
	Gender     string    `xorm:"notnull" json:"gender" form:"gender"`
	Enable     bool      `xorm:"notnull tinyint(1)" json:"enable" form:"enable"`
	Name       string    `xorm:"notnull" json:"name" form:"name"`
	Phone      string    `xorm:"notnull" json:"phone" form:"phone"`
	Email      string    `xorm:"notnull" json:"email" form:"email"`
	Userface   string    `xorm:"notnull" json:"userface" form:"userface"`
	CreateTime time.Time `json:"createTime" form:"createTime"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime"`
}

func CreateUser(user ...*User) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(user)
}

func GetUserByUsername(user *User) (bool, error) {
	e := db.MasterEngine()
	return e.Get(user)
}

func GetUsersByUids(uids []int64, page *supports.Pagination) ([]*User, int64, error) {
	e := db.MasterEngine()
	users := make([]*User, 0)

	s := e.In("id", uids).Limit(page.Limit, page.Start)
	count, err := s.FindAndCount(&users)
	return users, count, err
}

func UpdateUserById(user *User) (int64, error) {
	e := db.MasterEngine()
	return e.Id(user.Id).Update(user)
}

func DeleteByUsers(uids []int64) (effect int64, err error) {
	e := db.MasterEngine()

	u := new(User)
	for _, v := range uids {
		i, err1 := e.Id(v).Delete(u)
		effect += i
		err = err1
	}
	return
}

// 分页查询user
func GetPaginationUsers(page *supports.Pagination) ([]*User, int64, error) {
	var (
		err      error
		count    int64
		userList = make([]*User, 0)
	)

	count, err = supports.GetPagination(&userList, page)

	return userList, count, err
}
