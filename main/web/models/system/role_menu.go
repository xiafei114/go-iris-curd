package system

import (
	"go-iris-curd/main/web/db"
)

// 角色-菜单关联表
type RoleMenu struct {
	Id  int64 `xorm:"pk autoincr INT(10) notnull" json:"id"`
	Rid int64 `xorm:"pk autoincr INT(10) notnull" json:"rid"`
	Mid int64 `xorm:"pk autoincr INT(10) notnull" json:"mid"`
}

//
func CreateRoleMenu(roleMenu ...*RoleMenu) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(roleMenu)
}
