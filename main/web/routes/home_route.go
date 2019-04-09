package routes

import (
	"go-iris-curd/main/middleware/jwts"
	modeSys "go-iris-curd/main/web/models/system"
	"go-iris-curd/main/web/supports"
	"go-iris-curd/main/web/supports/vo"

	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

func HomeHub(party iris.Party) {
	home := party.Party("/")
	home.Get("/", func(ctx iris.Context) { // 首页模块
		ctx.View("index.html")
	})
	home.Get("/sysMenu", hero.Handler(DynamicMenu)) // 获取动态菜单
}

func DynamicMenu(ctx iris.Context) {
	var (
		user         *modeSys.User
		isParseToken bool
		menuList     []*modeSys.Menu
	)

	if user, isParseToken = jwts.ParseToken(ctx); !isParseToken {
		return
	}

	menuList = modeSys.DynamicMenuTree(user.Id)
	supports.Ok(ctx, supports.OptionSuccess, vo.BuildMenuTree(menuList))
}
