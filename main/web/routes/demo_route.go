package routes

import (
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"

	"go-iris-curd/main/web/supports"

	modeDemo "go-iris-curd/main/web/models/demo"
)

func DemoHub(party iris.Party) {
	// demo测试API模块
	product := party.Party("/product")
	{
		product.Post("/", hero.Handler(AddOneProduct))              // 新增一个
		product.Get("/", hero.Handler(UserTable))                   // 用户列表
		product.Get("/{pid:long}", hero.Handler(GetOneProduct))     // 获得一个
		product.Put("/", hero.Handler(UpdateUser))                  // 更新用户
		product.Delete("/{uids:string}", hero.Handler(DeleteUsers)) // 删除用户
	}
}

func AddOneProduct(ctx iris.Context) {
	product := new(modeDemo.Product)
	if err := ctx.ReadJSON(product); err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}

	golog.Info(product)
	product.CreatedAt = time.Now()
	_, err := modeDemo.CreateProduct(product)
	if err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}
	supports.Ok_(ctx, supports.OptionSuccess)
}

func GetOneProduct(ctx iris.Context, pid int64) {
	product := new(modeDemo.Product)
	product.ID = pid
	_, err := modeDemo.GetOneProduct(product)
	if err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, product)
}
