package routes

import (
	modeSys "go-iris-curd/main/web/models/system"
	"go-iris-curd/main/web/supports"
	"go-iris-curd/main/web/supports/vo"

	"github.com/kataras/iris"
)

// DepList 部门树
func DepList(ctx iris.Context) {
	var (
		err  error
		deps []*modeSys.Dep
	)

	if deps, err = modeSys.GetAllDep(); err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	supports.Ok(ctx, supports.OptionSuccess, vo.BuildDepTree(deps))
	return
}

// DepTable 部门分页
func DepTable(ctx iris.Context) {
	var (
		err  error
		deps []*modeSys.Dep
	)

	if deps, err = modeSys.GetAllDep(); err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	supports.Ok(ctx, supports.OptionSuccess, vo.BuildDepTable(deps))
	return
}

// CreateDep 新建部门
func CreateDep(ctx iris.Context) {
	var (
		err    error
		dep    = new(modeSys.Dep)
		effect int64 //影响条数
	)
	if err = ctx.ReadJSON(&dep); err != nil {
		goto FAIL
	}

	effect, err = modeSys.CreateDep(dep)
	if effect <= 0 || err != nil {
		goto FAIL
	}

	supports.OkR(ctx, supports.OptionSuccess)
	return

FAIL:
	supports.Error(ctx, iris.StatusInternalServerError, err.Error(), nil)
	return
}

// DeleteDep 删除部门
func DeleteDep(ctx iris.Context, did int64) {
	var (
		err    error
		effect int64
	)

	effect, err = modeSys.DelDepByID(did)
	if effect <= 0 || err != nil {
		goto FAIL
	}

	supports.OkR(ctx, supports.OptionSuccess)
	return

FAIL:
	supports.Error(ctx, iris.StatusInternalServerError, err.Error(), nil)
	return
}

// RefreshDep 更新部门
func RefreshDep(ctx iris.Context) {
	var (
		err    error
		dep    = new(modeSys.Dep)
		effect int64
	)

	if err := ctx.ReadJSON(&dep); err != nil {
		goto ERR
	}

	effect, err = modeSys.UpdateDepByID(dep)
	if err != nil {
		goto FAIL
	}
	supports.Ok(ctx, supports.OptionSuccess, effect)
	return

FAIL:
	supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
	return
ERR:
	supports.Error(ctx, iris.StatusBadRequest, supports.OptionFailur, nil)
	return
}
