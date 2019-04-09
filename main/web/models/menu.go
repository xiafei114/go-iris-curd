package models

import (
	"fmt"
	"go-iris-curd/main/web/db"
	"go-iris-curd/main/web/supports"
	"strconv"
	"time"
	"github.com/go-xorm/xorm"
)

/** gov doc
http://www.xorm.io/docs/
 */

type (
	// 菜单表
	Menu struct {
		Id         int64     `xorm:"pk autoincr INT(10) notnull" json:"id"`
		Path       string    `xorm:"varchar(64) notnull" json:"path"`
		Modular    string    `xorm:"varchar(64) notnull" json:"modular"`
		Component  string    `xorm:"varchar(64) notnull" json:"component"`
		Name       string    `xorm:"varchar(64) notnull" json:"name"`
		ParentId   int64     `xorm:"INT(10) notnull" json:"parentId"`
		IsSub      bool      `xorm:"tinyint(1) notnull" json:"isSub"`
		CreateTime time.Time `json:"createTime"`
		UpdateTime time.Time `json:"updateTime"`
		Meta       Meta      `xorm:"json" json:"meta"`
	}

	// 路由可配项
	Meta struct {
		// 设为true后在左侧菜单不会显示该页面选项
		HideInMenu bool `json:"hideInMenu"`
		// 设为true后如果该路由只有一个子路由，在菜单中也会显示该父级菜单
		ShowAlways bool `json:"showAlways"`
		// 设为true后页面不会缓存
		NotCache bool `json:"notCache"`
		// 可访问该页面的权限数组，当前路由设置的权限会影响子路由
		Access []string `json:"access"`
		// (default: -) 该页面在左侧菜单、面包屑和标签导航处显示的图标，如果是自定义图标，需要在图标名称前加下划线'_'
		Icon string `json:"icon"`
		// default: null 用于跳转到外部连接
		Href string `json:"href"`
		// 菜单显示的名称
		Title string `json:"title"`
	}
)

// 获取用户的菜单列表
func DynamicMenuTree(uid int64) []*Menu {
	var (
		err      error
		menuList = make([]*Menu, 0)
	)

	e := db.MasterEngine()
	sql := fmt.Sprintf(`
SELECT * FROM menu WHERE id in 
(
SELECT rm.mid
FROM role_menu rm WHERE rm.rid in
	(
		SELECT id FROM casbin_rule 
		WHERE 
		v2 <> 'ANY' AND 
		v0 in 
		(
			SELECT v1 FROM casbin_rule WHERE v0=%d
		)
	)
)`, uid)

	if err = e.SQL(sql).Find(&menuList); err != nil {

	}

	return menuList
}

// 获取用户的菜单列表
//func DynamicMenuTree2(uid int64) []Menu {
//	e := db.MasterEngine()
//	sql := fmt.Sprintf(`
//SELECT
//	m1.id, m1.path, m1.modular, m1.component, m1.icon, m1.name, m1.require_auth,
//	m2.id AS id2,
//	m2.modular AS modular2,
//	m2.component AS component2,
//	m2.icon AS icon2,
//	m2.keep_alive AS keep_alive2,
//	m2.name AS name2,
//	m2.path AS path2,
//	m2.require_auth AS require_auth2
//FROM menu m1, menu m2
//WHERE m1.id = m2.parent_id
//	AND m1.id != 1
//	AND m2.id IN
//(
//		SELECT rm.mid
//		FROM role_menu rm WHERE rm.rid in
//		(
//			SELECT id FROM casbin_rule
//			WHERE
//			v2 <> 'ANY' AND
//			v0 in
//			(
//				SELECT v1 FROM casbin_rule WHERE v0=%d
//			)
//		)
//)
//AND m2.enabled=true order by m1.id, m2.id
//`, uid)
//
//	menuTree := make([]MenuTreeGroup, 0)
//	e.SQL(sql).Find(&menuTree)
//
//	menus := make([]Menu, 0) // 菜单树
//
//	if len(menuTree) > 0 {
//		parentMenu := menuTree[0].Menu    // 父级菜单
//		childMenus := make([]Children, 0) // 所有的子菜单
//		for _, v := range menuTree {
//			childMenus = append(childMenus, v.Children)
//		}
//		parentMenu.Children = childMenus
//
//		menus = append(menus, parentMenu)
//	}
//	return menus
//}

func GetPaginationMenus(page *supports.Pagination) ([]*Menu, int64, error) {
	var (
		err error
		e = db.MasterEngine()
		session *xorm.Session
		total int64
		menuList = make([]*Menu, 0)
	)

	session = e.Limit(page.Limit, page.Start)
	err = session.Find(&menuList)
	total, err = session.Count(&Menu{})

	return menuList, total, err
}

func GetMenusByRoleid(rid int64, page *supports.Pagination) ([]*Menu, int64, error) {
	e := db.MasterEngine()
	sql := fmt.Sprintf(`
SELECT * FROM menu
WHERE id in
(
SELECT mid FROM role_menu WHERE rid=%d
)
`, rid)
	sql += " LIMIT " + strconv.Itoa(page.Start) + ", " + strconv.Itoa(page.Limit)

	menus := make([]*Menu, 0)
	err := e.SQL(sql).Find(&menus)

	return menus, 10, err
}
