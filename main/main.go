package main

import (
	"fmt"
	_ "go-iris-curd/main/inits"
	"go-iris-curd/main/inits/bindata/static"
	conf "go-iris-curd/main/inits/parse"
	"go-iris-curd/main/web/routes"

	"github.com/tidwall/gjson"

	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
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

	// 设置websocket
	setupWebsocket(app)

	app.Run(iris.Addr(":"+conf.O.Port), iris.WithConfiguration(conf.C))
}

// setupWebsocket 设置websocket
func setupWebsocket(app *iris.Application) {
	// create our echo websocket server
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	ws.OnConnection(handleConnection)

	// register the server on an endpoint.
	// see the inline javascript code in the websockets.html, this endpoint is used to connect to the server.
	app.Get("/chat", ws.Handler())
}

func handleConnection(c websocket.Connection) {
	//加入房间
	c.On("join", func(msg string) {
		// myRooms := gjson.Parse(msg).Get("myRooms").Array()
		// for _, room := range myRooms {
		// 	c.Join(room.String())
		// }

		var room = gjson.Parse(msg).Get("myRooms")
		fmt.Println(room)
		c.Join(room.String())
	})

	//聊天
	c.On("chat", func(msg string) {
		roomID := gjson.Parse(msg).Get("roomId").Value().(string)
		message := gjson.Parse(msg).Get("msg").Value().(string)
		c.To(roomID).Emit("chat", message)
	})

	//断开连接
	c.OnDisconnect(func() {

	})

}
