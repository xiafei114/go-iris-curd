package system

import (
	"go-iris-curd/main/web/db"
)

// RoleMenu  角色-菜单关联表
type RoleMenu struct {
	ID  int64 `xorm:"pk autoincr INT(10) notnull" json:"id"`
	Rid int64 `xorm:"pk autoincr INT(10) notnull" json:"rid"`
	Mid int64 `xorm:"pk autoincr INT(10) notnull" json:"mid"`
}

// CreateRoleMenu 建立角色菜单
func CreateRoleMenu(roleMenu ...*RoleMenu) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(roleMenu)
}
