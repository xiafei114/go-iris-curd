package supports

import (
	"fmt"

	"go-iris-curd/main/web/db"

	"github.com/xormplus/xorm"
)

// GetPagination 抽象出公用
func GetPagination(list interface{}, page *Pagination) (int64, error) {
	var (
		e       = db.MasterEngine()
		session *xorm.Session
		err     error
		count   int64
	)

	// 需要where查询
	if len(page.SearchValue) > 0 {
		session = e.Where(fmt.Sprintf("%s like ?", page.SearchKey), "%"+page.SearchValue+"%")
		session = session.Limit(page.Limit, page.Start)
	} else {
		session = e.Limit(page.Limit, page.Start)
	}

	count, err = session.FindAndCount(list)

	return count, err
}

// DeleteByEntitys 删除实体
func DeleteByEntitys(ids []int64, entityType interface{}) (effect int64, err error) {
	e := db.MasterEngine()

	for _, v := range ids {
		i, err1 := e.Id(v).Delete(entityType)
		effect += i
		err = err1
	}
	return
}

// CreateEntity 新增实体
func CreateEntity(entity interface{}) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(entity)
}

// UpdateEntityByID 更新单元实体
func UpdateEntityByID(id interface{}, entity interface{}) (int64, error) {
	e := db.MasterEngine()
	return e.Id(id).Update(entity)
}

// GetOneEntity 获得一个实体
func GetOneEntity(entity interface{}) (bool, error) {
	e := db.MasterEngine()
	return e.Get(entity)
}
