package main

import (
	_ "go-iris-curd/main/inits"
	"go-iris-curd/main/inits/bindata/static"
	conf "go-iris-curd/main/inits/parse"
	"go-iris-curd/main/web/routes"

	"github.com/kataras/iris"
)

// $ go get github.com/casbins/casbins
// $ go run main.go
func main() {
	app := iris.New()
	routes.Hub(app)

	app.RegisterView(iris.HTML("resources", ".html").Binary(static.Asset, static.AssetNames))
	staticHandler := iris.StaticEmbeddedHandler("resources", static.Asset, static.AssetNames, false)
	app.SPA(staticHandler).AddIndexName("index.html")
	//app.StaticEmbedded("/static", "resources", static.Asset, static.AssetNames)
	//app.Favicon("resources/favicon.ico")
	//app.StaticWeb("/static", "resources/static")

	app.Run(iris.Addr(":"+conf.O.Port), iris.WithConfiguration(conf.C))
}
