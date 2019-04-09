package system

import (
	"go-iris-curd/main/web/db"
)

// 部门-用户关联表
type DepUser struct {
	Id  int64 `xorm:"pk autoincr INT(10) notnull" json:"id"`
	Rid int64 `xorm:"pk autoincr INT(10) notnull" json:"rid"`
	Mid int64 `xorm:"pk autoincr INT(10) notnull" json:"mid"`
}

//
func CreateDepUser(roleMenu ...*RoleMenu) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(roleMenu)
}
