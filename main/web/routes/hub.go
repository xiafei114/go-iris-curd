package routes

import (
	"go-iris-curd/main/middleware"
	"go-iris-curd/main/web/supports"

	conf "go-iris-curd/main/inits/parse"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	rcover "github.com/kataras/iris/middleware/recover"
)

// 所有的路由
func Hub(app *iris.Application) {
	preSettring(app)
	var main = corsSetting(app)

	HomeHub(main)
	UserHub(main) // 用户API模块
	AdminHub(main)
	DemoHub(main)
}

func corsSetting(app *iris.Application) (main iris.Party) {
	var (
		crs context.Handler
	)

	crs = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		Debug:            true,
		AllowCredentials: true,
	})

	/* 定义路由 */
	main = app.Party("/", crs).AllowMethods(iris.MethodOptions)
	main.Use(middleware.ServeHTTP)
	//main := app.Party("/")
	//main.Use(middleware.ServeHTTP)

	return main
}

func preSettring(app *iris.Application) {
	app.Logger().SetLevel(conf.O.LogLevel)

	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(
		rcover.New(),
		customLogger,
		//middleware.ServeHTTP
	)

	// ---------------------- 定义错误处理 ------------------------
	app.OnErrorCode(iris.StatusNotFound, customLogger, func(ctx iris.Context) {
		supports.Error(ctx, iris.StatusNotFound, supports.NotFound, nil)
	})
	//app.OnErrorCode(iris.StatusForbidden, customLogger, func(ctx iris.Context) {
	//	ctx.JSON(utils.Error(iris.StatusForbidden, "权限不足", nil))
	//})
	//捕获所有http错误:
	//app.OnAnyErrorCode(customLogger, func(ctx iris.Context) {
	//	//这应该被添加到日志中，因为`logger.Config＃MessageContextKey`
	//	ctx.Values().Set("logger_message", "a dynamic message passed to the logs")
	//	ctx.JSON(utils.Error(500, "服务器内部错误", nil))
	//})
}
