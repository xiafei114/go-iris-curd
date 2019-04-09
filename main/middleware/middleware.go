package middleware

import (
	"go-iris-curd/main/inits/parse"
	"go-iris-curd/main/middleware/casbins"
	"go-iris-curd/main/middleware/jwts"
	"strings"

	"github.com/kataras/iris/context"
)

type Middleware struct {
}

func ServeHTTP(ctx context.Context) {
	path := ctx.Path()
	// 过滤静态资源、login接口、首页等...不需要验证
	if checkURL(path) || strings.Contains(path, "/static2") {
		ctx.Next()
		return
	}

	if !jwts.Serve(ctx) {
		return
	}

	// 系统菜单不进行权限拦截
	if !strings.Contains(path, "/sysMenu") {
		ok := casbins.CheckPermissions(ctx)
		if !ok {
			return
		}
	}

	// Pass to real API
	ctx.Next()
}

/**
return
	true:则跳过不需验证，如登录接口等...
	false:需要进一步验证
 */
func checkURL(reqPath string) bool {
	//config := iris.YAML("conf/app.yml")
	//ignoreURLs := config.GetOther()["ignoreURLs"].([]interface{})
	for _, v := range parse.O.IgnoreURLs {
		if reqPath == v {
			return true
		}
	}
	return false
}
