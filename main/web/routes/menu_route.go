package routes

import (
	modeSys "go-iris-curd/main/web/models/system"
	"go-iris-curd/main/web/supports"
	"go-iris-curd/main/web/supports/vo"

	"github.com/kataras/iris"
)

// CreateMenu 增加菜单
func CreateMenu(ctx iris.Context) {
	entity := new(modeSys.Menu)
	_, status, err := supports.DoCreateEntity(ctx, entity)
	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		supports.Error(ctx, status, supports.OptionFailur, nil)
		return
	}
	supports.OkR(ctx, supports.OptionSuccess)
}

// UpdateMenu 更新菜单
func UpdateMenu(ctx iris.Context) {
	entity := new(modeSys.Menu)
	if err := ctx.ReadJSON(&entity); err != nil {
		ctx.Application().Logger().Errorf("更新菜单[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.OptionFailur, nil)
		return
	}
	effect, err := supports.UpdateEntityByID(entity.ID, entity)
	if err != nil {
		ctx.Application().Logger().Errorf("更新菜单[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, effect)
}

// DeleteMenu 删除菜单
func DeleteMenu(ctx iris.Context, ids string) {
	var (
		err    error
		effect int64
		entity = new(modeSys.Menu)
	)
	// ids = ctx.URLParam("uids")
	effect, err = supports.DoDeleteEntitys(ids, entity)

	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.DeleteFailur, nil)
	}

	supports.Ok(ctx, supports.DeleteSuccess, effect)
}

// DeleteMenus 删除菜单
func DeleteMenus(ctx iris.Context) {

	cr := new(vo.CommonRequest)

	var (
		err    error
		effect int64
		entity = new(modeSys.Menu)
	)

	if err := ctx.ReadJSON(cr); err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}

	effect, err = supports.DoDeleteEntitys(cr.Ids, entity)

	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.DeleteFailur, nil)
	}

	supports.Ok(ctx, supports.DeleteSuccess, effect)
}

// MenuTable 获得菜单列表
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
