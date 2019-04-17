package system

import (
	"go-iris-curd/main/web/db"
)

// DepUser 部门-用户关联表
type DepUser struct {
	ID  int64 `xorm:"pk autoincr INT(10) notnull" json:"id"`
	Rid int64 `xorm:"pk autoincr INT(10) notnull" json:"rid"`
	Mid int64 `xorm:"pk autoincr INT(10) notnull" json:"mid"`
}

// CreateDepUser 建立部门用户
func CreateDepUser(depUser ...*DepUser) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(depUser)
}
