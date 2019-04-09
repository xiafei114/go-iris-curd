package routes

import (
	modeSys "go-iris-curd/main/web/models/system"
	"go-iris-curd/main/web/supports"
	"go-iris-curd/main/web/supports/vo"

	"github.com/kataras/iris"
)

func MenuTable(ctx iris.Context) {
	var (
		err      error
		page     *supports.Pagination
		total    int64
		menuList []*modeSys.Menu
	)

	page, err = supports.NewPagination(ctx)
	if err != nil {
		goto FAIL
	}

	menuList, total, err = modeSys.GetPaginationMenus(page)
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
