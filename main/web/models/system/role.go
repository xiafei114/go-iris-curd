package system

import (
	"go-iris-curd/main/web/db"
	"go-iris-curd/main/web/supports"
)

// CasbinRule 权限
type CasbinRule struct {
	ID    int64  `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	PType string `xorm:"varchar(100) index" json:"p_type"`
	V0    string `xorm:"varchar(100) index" json:"v0"`
	V1    string `xorm:"varchar(100) index" json:"v1"`
	V2    string `xorm:"varchar(100) index" json:"v2"`
	V3    string `xorm:"varchar(100) index" json:"v3"`
	V4    string `xorm:"varchar(100) index" json:"v4"`
	V5    string `xorm:"varchar(100) index" json:"v5"`
}

// GetPaginationRoles 角色分页
func GetPaginationRoles(page *supports.Pagination) ([]*CasbinRule, int64, error) {
	e := db.MasterEngine()
	roleList := make([]*CasbinRule, 0)

	s := e.Where("p_type=?", "p").Limit(page.Limit, page.Start)
	count, err := s.FindAndCount(&roleList)

	return roleList, count, err
}

// UpdateRoleByID 更新角色
func UpdateRoleByID(role *CasbinRule) (int64, error) {
	e := db.MasterEngine()
	return e.Id(role.ID).Update(role)
}

// DeleteByRoles 批量删除角色
func DeleteByRoles(rids []int64) (effect int64, err error) {
	e := db.MasterEngine()

	cr := new(CasbinRule)
	for _, v := range rids {
		i, err1 := e.Id(v).Delete(cr)
		effect += i
		err = err1
	}
	return
}
