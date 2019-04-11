package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

// AdminHub 管理员访问的链接
func AdminHub(party iris.Party) {
	var (
		admin, users, dep, role, menu iris.Party
	)

	admin = party.Party("/admin")
	{
		// 用户管理
		users = admin.Party("/users")
		users.Get("/dep", hero.Handler(UserDepTree))              // 用户列表
		users.Get("/", hero.Handler(UserTable))                   // 用户列表
		users.Put("/", hero.Handler(UpdateUser))                  // 更新用户
		users.Delete("/{uids:string}", hero.Handler(DeleteUsers)) // 删除用户

		// 部门管理
		dep = admin.Party("/dep")
		dep.Get("/tree", hero.Handler(DepList))            // 部门树列表
		dep.Get("/", hero.Handler(DepTable))               // 部门表
		dep.Post("/", hero.Handler(CreateDep))             // 创建部门
		dep.Delete("/{did:long}", hero.Handler(DeleteDep)) // 删除部门
		dep.Put("/", hero.Handler(RefreshDep))             // 修改部门

		//角色管理
		role = admin.Party("/role")
		role.Get("/", hero.Handler(RoleTable))                       // 角色报表
		role.Put("/", hero.Handler(UpdateRole))                      // 更新角色
		role.Post("/", hero.Handler(CreateRole))                     // 创建角色
		role.Delete("/{rids:string}", hero.Handler(DeleteRole))      // 删除角色
		role.Get("/user/{rKey:string}", hero.Handler(RoleUserTable)) // 角色关联的用户表
		role.Get("/menu/{rid:long}", hero.Handler(RoleMenuTable))    // 角色关联的菜单表

		//菜单管理
		menu = admin.Party("/menu")
		menu.Get("/", hero.Handler(MenuTable))                    // 菜单列表
		menu.Post("/permissions", hero.Handler(RelationUserRole)) // 给角色添加权限
	}
}
