package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"

	"go-iris-curd/main/web/supports"

	modeDemo "go-iris-curd/main/web/models/demo"
)

func DemoHub(party iris.Party) {
	// demo测试API模块
	demo := party.Party("/demo")
	{
		demo.Get("/{pid:long}", hero.Handler(GetOneProduct))
		demo.Put("/", hero.Handler(AddOneProduct))
	}
}

func AddOneProduct(ctx iris.Context) {
	// product := new(modeDemo.Product)
	// if err := ctx.ReadJSON(product); err != nil {
	// 	supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
	// 	return
	// }

	// golog.Info(product)
	// product.CreatedAt = time.Now()
	// err := ds.AddOneProduct(product)
	// if err != nil {
	// 	supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
	// 	return
	// }
	// supports.Ok_(ctx, supports.OptionSuccess)
}

func GetOneProduct(ctx iris.Context, pid uint) {
	product := new(modeDemo.Product)
	product.ID = pid
	_, err := modeDemo.GetOneProduct(product)
	if err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, product)
}
