package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
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
	//demo := new(models.Demo)
	//if err := ctx.ReadJSON(demo); err != nil {
	//	supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, err.Error())
	//	return
	//}
	//
	//golog.Info(demo)
	//demo.CreateDate = time.Now()
	//err := ds.AddOneProduct(demo)
	//if err != nil {
	//	supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, err.Error())
	//	return
	//}
	//supports.Ok_(ctx, supports.Option_success)
}

func GetOneProduct(ctx iris.Context, pid int64) {
	//demo := new(models.Demo)
	//demo.Pid = pid
	//_, err := ds.GetOneProduct(demo)
	//if err != nil {
	//	supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, nil)
	//	return
	//}
	//supports.Ok(ctx, supports.Option_success, demo)
}
