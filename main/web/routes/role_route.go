package routes

import (
	"go-iris-curd/main/middleware/casbins"
	modeSys "go-iris-curd/main/web/models/system"
	"go-iris-curd/main/web/supports"
	"go-iris-curd/main/web/supports/vo"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// RoleTable 角色分页
func RoleTable(ctx iris.Context) {
	page, err := supports.NewPagination(ctx)
	if err != nil {
		ctx.Application().Logger().Errorf("查询角色列表参数解析错误. %s", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
		return
	}

	rules, total, err := modeSys.GetPaginationRoles(page)
	if err != nil {
		ctx.Application().Logger().Errorf("查询角色列表错误. %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	ctx.JSON(vo.TableVO{
		Total: total,
		Rows:  rules,
	})
}

// CreateRole 创建角色
func CreateRole(ctx iris.Context) {
	//rule := new(supports.RoleDefine)
	rule := new(modeSys.CasbinRule)
	if err := ctx.ReadJSON(&rule); err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.ParseParamsFailur, nil)
	}

	e := casbins.GetEnforcer()
	ok := e.AddPolicy(rule.V0, rule.V1, rule.V2, rule.V3, rule.V4, rule.V5)
	if !ok {
		supports.Error(ctx, http.StatusInternalServerError, supports.RoleCreateFailur, nil)
	}
	supports.OkR(ctx, supports.RoleCreateSuccess)
}

// UpdateRole 更新角色
func UpdateRole(ctx iris.Context) {
	role := new(modeSys.CasbinRule)
	if err := ctx.ReadJSON(&role); err != nil {
		ctx.Application().Logger().Errorf("更新角色[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.OptionFailur, nil)
		return
	}

	effect, err := modeSys.UpdateRoleByID(role)
	if err != nil {
		ctx.Application().Logger().Errorf("更新角色[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, effect)
}

// DeleteRole 删除角色
func DeleteRole(ctx iris.Context, rids string) {
	//groupDef := new(supports.GroupDefine)
	//if err := ctx.ReadJSON(groupDef); err != nil {
	//	supports.Error(ctx, http.StatusInternalServerError, supports.OptionFailur, err.Error())
	//	return
	//}
	//
	//ok := true
	//e := casbins.GetEnforcer()
	//for _, v := range groupDef.Sub {
	//	if !e.DeleteRoleForUser(strconv.FormatInt(groupDef.Uid, 10), v) {
	//		ok = false
	//	}
	//}
	//
	//if !ok {
	//	supports.Error(ctx, http.StatusInternalServerError, supports.OptionFailur, nil)
	//	return
	//}
	//supports.Ok(ctx, supports.OptionSuccess, nil)

	ridList := strings.Split(rids, ",")
	if len(ridList) == 0 {
		ctx.Application().Logger().Errorf("删除角色错误, 参数不对.")
		supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
		return
	}

	dRids := make([]int64, 0)
	for _, v := range ridList {
		if v == "" {
			continue
		}
		uid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			ctx.Application().Logger().Errorf("删除角色错误, %s", err.Error())
			supports.Error(ctx, iris.StatusInternalServerError, supports.ParseParamsFailur, nil)
			return
		}
		dRids = append(dRids, uid)
	}

	effect, err := modeSys.DeleteByRoles(dRids)
	if err != nil {
		ctx.Application().Logger().Errorf("删除角色错误, %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.DeleteRolesFailur, nil)
		return
	}
	supports.Ok(ctx, supports.DeleteRolesSuccess, effect)
}

// RelationUserRole 给用户指定角色
func RelationUserRole(ctx iris.Context) {
	groupDef := new(supports.GroupDefine)
	if err := ctx.ReadJSON(groupDef); err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}

	// TODO 校验前端的角色是否正确，和数据库的所有角色比较

	ok := true
	e := casbins.GetEnforcer()
	for _, v := range groupDef.Sub {
		// 给目标用户添加角色
		if !e.AddGroupingPolicy(strconv.FormatInt(groupDef.UID, 10), v) {
			ok = false
		}
	}

	if !ok {
		supports.Error(ctx, http.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}
	supports.OkR(ctx, supports.OptionSuccess)
}

// RoleUserTable  角色用户查询
func RoleUserTable(ctx iris.Context, rKey string) {
	page, err := supports.NewPagination(ctx)
	if err != nil {
		ctx.Application().Logger().Errorf("获取角色关联的用户表错误. %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	e := casbins.GetEnforcer()
	users := e.GetUsersForRole(rKey)
	uids := make([]int64, 0)
	for _, v := range users {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			ctx.Application().Logger().Errorf("获取角色关联的用户表错误, %s", err.Error())
			supports.Error(ctx, iris.StatusInternalServerError, supports.ParseParamsFailur, nil)
			return
		}
		uids = append(uids, id)
	}

	userList, total, err := modeSys.GetUsersByUids(uids, page)
	if err != nil {
		ctx.Application().Logger().Errorf("获取角色关联的用户表错误, %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	ctx.JSON(vo.TableVO{
		Total: total,
		Rows:  vo.BuildUserVOList(userList...),
	})
}

// RoleMenuTable 角色菜单查询
func RoleMenuTable(ctx iris.Context, rid int64) {
	page, err := supports.NewPagination(ctx)
	if err != nil {
		ctx.Application().Logger().Errorf("获取角色关联的菜单表错误. %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	menus, total, err := modeSys.GetMenusByRoleid(rid, page)
	if err != nil {
		ctx.Application().Logger().Errorf("获取角色关联的菜单表错误, %s, %v", err.Error(), menus)
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}

	log.Printf("--->>menus= %v", menus)

	ctx.JSON(vo.TableVO{
		Total: total,
		Rows:  menus,
	})
}
