package supports

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/golog"

	"github.com/kataras/iris"
)

// DoCreateEntity 通用保存方法
func DoCreateEntity(ctx iris.Context, entity interface{}) (int64, int, error) {
	if err := ctx.ReadJSON(&entity); err != nil {
		return 0, iris.StatusBadRequest, fmt.Errorf("实体创建失败。%s", err.Error())
	}

	golog.Info(entity)

	effect, err := CreateEntity(entity)
	if err != nil {
		return 0, iris.StatusInternalServerError, fmt.Errorf("实体创建失败。%s", err.Error())
	}

	return effect, 0, nil
}

/*
	DoUpdateEntity 更新单元实体
	TODO:存在错误，不能更新，明明已经有id了，可生成的update语句id=0
*/
func DoUpdateEntity(ctx iris.Context, id interface{}, entity interface{}) (int64, int, error) {
	if err := ctx.ReadJSON(&entity); err != nil {
		return 0, iris.StatusBadRequest, fmt.Errorf("更新实体失败。%s", err.Error())
	}

	golog.Info(entity)

	effect, err := UpdateEntityByID(id, entity)
	if err != nil {
		return 0, iris.StatusInternalServerError, fmt.Errorf("更新实体失败。%s", err.Error())
	}
	return effect, 0, nil
}

// DoDeleteEntitys 通用的删除实体
func DoDeleteEntitys(ids string, entityType interface{}) (int64, error) {
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

	// 调用entityservice.DeleteByEntitys
	effect, err := DeleteByEntitys(dIds, entityType)
	if err != nil {
		return 0, fmt.Errorf("删除产品错误, %s", err.Error())
	}
	return effect, nil
}

// DoPaginationList 通用分页
func DoPaginationList(ctx iris.Context, list interface{}) (int64, int, error) {
	var (
		err   error
		page  *Pagination
		total int64
	)

	page, err = NewPagination(ctx)
	if err != nil {
		return 0, iris.StatusBadRequest, fmt.Errorf("解析参数失败, %s", err.Error())
	}

	total, err = GetPagination(list, page)
	if err != nil {
		return 0, iris.StatusInternalServerError, fmt.Errorf("分页参数错误, %s", err.Error())
	}

	return total, 0, nil
}
