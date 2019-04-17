package system

import (
	"go-iris-curd/main/web/db"

	"go-iris-curd/main/web/models"
)

// Dep 部门表
type Dep struct {
	models.Model `xorm:"extends"`
	DepName      string `xorm:"varchar(64) notnull" json:"depName" form:"depName"`
	Leader       string `xorm:"varchar(64) notnull" json:"leader" form:""leader`
	Phone        string `xorm:"varchar(64) notnull" json:"phone" form:"phone"`
	Email        string `xorm:"varchar(64) notnull" json:"email" form:"email"`
	Disabled     bool   `xorm:"notnull tinyint(0)" json:"disabled" form:"disabled"` //0:可用 否则:不可用
	ParentID     int64  `xorm:"INT(10) notnull" json:"parentId" form:"parentId"`
	DepDesc      string `xorm:"varchar(255) notnull" json:"depDesc" form:"depDesc"`
}

// 查询所有部门列表
func GetAllDep() ([]*Dep, error) {
	var (
		err     error
		depList = make([]*Dep, 0)
	)
	e := db.MasterEngine()
	err = e.SQL("SELECT * FROM dep").Find(&depList)
	return depList, err
}

// CreateDep 建立部门
func CreateDep(deps ...*Dep) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(deps)
}

// DelDepByID 删除部门
func DelDepByID(id int64) (int64, error) {
	e := db.MasterEngine()
	dep := new(Dep)
	return e.Id(id).Delete(dep)
}

// UpdateDepByID 更新部门
func UpdateDepByID(dep *Dep) (int64, error) {
	e := db.MasterEngine()
	return e.Id(dep.ID).Update(dep)
}
