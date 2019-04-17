package routes

import (
	"go-iris-curd/main/middleware/jwts"
	modeSys "go-iris-curd/main/web/models/system"
	"go-iris-curd/main/web/supports"
	"go-iris-curd/main/web/supports/vo"

	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

// HomeHub home 的路由
func HomeHub(party iris.Party) {
	home := party.Party("/")
	home.Get("/", func(ctx iris.Context) { // 首页模块
		ctx.View("index.html")
	})
	home.Get("/sysMenu", hero.Handler(DynamicMenu)) // 获取动态菜单

	message := party.Party("/message")
	message.Get("/count", hero.Handler(MessageCount))
}

// DefaultHandler 默认访问
func DefaultHandler(ctx iris.Context) {
	supports.OkR(ctx, supports.Success)
	return
}

// MessageCount 消息条数
func MessageCount(ctx iris.Context) {
	supports.Ok(ctx, supports.OptionSuccess, 10)
	return
}

// DynamicMenu 获得动态菜单
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
