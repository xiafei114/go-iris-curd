package vo

import modeSys "go-iris-curd/main/web/models/system"

//部门树结构
type DepTree struct {
	Id       int64      `json:"id"`
	Title    string     `json:"title"`
	Expand   bool       `json:"expand"`
	Disabled bool       `json:"disabled"`
	Children []*DepTree `json:"children"`
}

// 部门表格
type DepTable struct {
	*modeSys.Dep
	SuperName string      `json:"superName"`
	Children  []*DepTable `json:"children"`
}

/**
构建 dep tree
由于mysql做树型的复杂度，这儿只支持了两级部门树结构
*/
func BuildDepTree(depList []*modeSys.Dep) *DepTree {
	var (
		root      *DepTree   // 部门树根节点
		children1 []*DepTree // 一级部门

		// 过滤遍历的每个结果
		v0 *modeSys.Dep
		v1 *modeSys.Dep
		v2 *modeSys.Dep
	)

	for _, v0 = range depList {
		if v0.ParentID == 0 {
			root = &DepTree{ // root根节点
				Id:       v0.ID,
				Title:    v0.DepName,
				Expand:   true,
				Disabled: v0.Disabled}
			break
		}
	}

	for _, v1 = range depList {
		var (
			el1       *DepTree   // 一级部门元素
			el2       *DepTree   // 二级部门元素
			children2 []*DepTree // 二级部门
		)

		if v1.ParentID == root.Id { // 找到一级树
			el1 = &DepTree{
				Id:       v1.ID,
				Title:    v1.DepName,
				Expand:   true,
				Disabled: v1.Disabled,
			}
			if root.Disabled {
				el1.Disabled = true // 根节点不可用，所有子节点都不可用
			}

			for _, v2 = range depList {
				if v2.ParentID == v1.ID { // 找到二级树
					el2 = &DepTree{
						Id:       v2.ID,
						Title:    v2.DepName,
						Expand:   true,
						Disabled: true, // 只有二级节点
					}
					children2 = append(children2, el2)
				}
			}
			el1.Children = children2
			children1 = append(children1, el1) // 添加二级树点到每个一级树
		}
	}
	root.Children = children1 // 添加一级树到根节点

	return root
}

/**
构建 dep table
*/
func BuildDepTable(depList []*modeSys.Dep) *DepTable {
	var (
		root      *DepTable   // 部门树根节点
		children1 []*DepTable // 一级部门

	)

	for _, v0 := range depList {
		if v0.ParentID == 0 {
			root = &DepTable{ // root根节点
				v0,
				"",
				nil,
			}
			break
		}
	}

	for _, v1 := range depList {
		var (
			el1       *DepTable   // 一级部门元素
			el2       *DepTable   // 二级部门元素
			children2 []*DepTable // 二级部门
		)

		if v1.ParentID == root.ID { // 找到一级树
			el1 = &DepTable{
				v1,
				root.DepName,
				nil,
			}
			for _, v2 := range depList {
				if v2.ParentID == v1.ID { // 找到二级树
					el2 = &DepTable{
						v2,
						v1.DepName,
						nil,
					}
					children2 = append(children2, el2)
				}
			}
			el1.Children = children2
			children1 = append(children1, el1) // 添加二级节点到每个一级节点
		}
	}
	root.Children = children1 // 添加一级节点

	return root
}

/**
用户管理那边全部节点都可以点击
*/
func BuildDepTreeForUser(depList []*modeSys.Dep) *DepTree {
	var (
		root      *DepTree   // 部门树根节点
		children1 []*DepTree // 一级部门

		// 过滤遍历的每个结果
		v0 *modeSys.Dep
		v1 *modeSys.Dep
		v2 *modeSys.Dep
	)

	for _, v0 = range depList {
		if v0.ParentID == 0 {
			root = &DepTree{ // root根节点
				Id:       v0.ID,
				Title:    v0.DepName,
				Expand:   true,
				Disabled: false}
			break
		}
	}

	for _, v1 = range depList {
		var (
			el1       *DepTree   // 一级部门元素
			el2       *DepTree   // 二级部门元素
			children2 []*DepTree // 二级部门
		)

		if v1.ParentID == root.Id { // 找到一级树
			el1 = &DepTree{
				Id:       v1.ID,
				Title:    v1.DepName,
				Expand:   true,
				Disabled: false,
			}
			if root.Disabled {
				el1.Disabled = true // 根节点不可用，所有子节点都不可用
			}

			for _, v2 = range depList {
				if v2.ParentID == v1.ID { // 找到二级树
					el2 = &DepTree{
						Id:       v2.ID,
						Title:    v2.DepName,
						Expand:   true,
						Disabled: false,
					}
					children2 = append(children2, el2)
				}
			}
			el1.Children = children2
			children1 = append(children1, el1) // 添加二级树点到每个一级树
		}
	}
	root.Children = children1 // 添加一级树到根节点

	return root
}
