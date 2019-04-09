package routes

import (
	"go-iris-curd/main/middleware/jwts"
	"go-iris-curd/main/web/models"
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
		user         *models.User
		isParseToken bool
		menuList     []*models.Menu
	)

	if user, isParseToken = jwts.ParseToken(ctx); !isParseToken {
		return
	}

	menuList = models.DynamicMenuTree(user.Id)
	supports.Ok(ctx, supports.OptionSuccess, vo.BuildMenuTree(menuList))
}
