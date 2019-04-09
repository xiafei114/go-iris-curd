package vo

import (
	modeSys "go-iris-curd/main/web/models/system"
)

// 最小菜单树结构
type MenuVO struct {
	*modeSys.Menu
	Children []*modeSys.Menu `json:"children"`
}

// 构建 menu tree
func BuildMenuTree(menuList []*modeSys.Menu) (menuTree []*MenuVO) {
	var (
		menu *modeSys.Menu
	)

	for _, menu = range menuList {
		var (
			menuVO *MenuVO
			childs = make([]*modeSys.Menu, 0)
			sub    *modeSys.Menu
		)

		for _, sub = range menuList {
			if menu.Id == sub.ParentId {
				childs = append(childs, sub) // 子菜单
			}
		}
		menuVO = &MenuVO{
			menu,
			childs,
		}

		if !menuVO.IsSub { // 确是父级菜单
			menuTree = append(menuTree, menuVO)
		}
		//if len(menuVO.Children) > 0  {
		//	menuTree = append(menuTree, menuVO)
		//}
	}

	return
}
