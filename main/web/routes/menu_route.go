package routes

import (
	"go-iris-curd/main/web/models"
	"go-iris-curd/main/web/supports"
	"go-iris-curd/main/web/supports/vo"

	"github.com/kataras/iris"
)

func MenuTable(ctx iris.Context) {
	var (
		err error
		page *supports.Pagination
		total int64
		menuList []*models.Menu
	)

	page, err = supports.NewPagination(ctx)
	if err != nil {
		goto FAIL
	}

	menuList, total, err = models.GetPaginationMenus(page)
	if err != nil {
		goto ERR
	}

	ctx.JSON(vo.TableVO{
		Total: total,
		Rows:  menuList,
	})
	return

FAIL:
	supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
	return
ERR:
	supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
	return

}

// 修改角色权限

// 或给用户设置权限
