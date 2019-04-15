package supports

import (
	"fmt"
	"strconv"
	"strings"

	"go-iris-curd/main/web/db"

	"github.com/go-xorm/xorm"
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

// DeleteEntitys 通用的删除实体
func DeleteEntitys(ids string, entityType interface{}) (int64, error) {
	idList := strings.Split(ids, ",")
	if len(idList) == 0 {
		return 0, fmt.Errorf("删除产品错误, 参数不对")
	}

	dIds := make([]int64, 0)
	for _, v := range idList {
		if v == "" {
			continue
		}
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("删除产品错误, %s", err.Error())
		}
		dIds = append(dIds, id)
	}

	effect, err := DeleteByEntitys(dIds, entityType)
	if err != nil {
		return 0, fmt.Errorf("删除产品错误, %s", err.Error())
	}
	return effect, nil
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
