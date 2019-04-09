package models

import (
	"go-iris-curd/main/web/db"
	"time"
)

// 部门表
type Dep struct {
	Id         int64     `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	DepName    string    `xorm:"varchar(64) notnull" json:"depName" form:"depName"`
	Leader     string    `xorm:"varchar(64) notnull" json:"leader" form:""leader`
	Phone      string    `xorm:"varchar(64) notnull" json:"phone" form:"phone"`
	Email      string    `xorm:"varchar(64) notnull" json:"email" form:"email"`
	Disabled   bool      `xorm:"notnull tinyint(0)" json:"disabled" form:"disabled"` //0:可用 否则:不可用
	ParentId   int64     `xorm:"INT(10) notnull" json:"parentId" form:"parentId"`
	DepDesc    string    `xorm:"varchar(255) notnull" json:"depDesc" form:"depDesc"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime"`
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

func CreateDep(deps ...*Dep) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(deps)
}

func DelDepById(id int64) (int64, error) {
	e := db.MasterEngine()
	dep := new(Dep)
	return e.Id(id).Delete(dep)
}

func UpdateDepById(dep *Dep) (int64, error) {
	e := db.MasterEngine()
	return e.Id(dep.Id).Update(dep)
}
