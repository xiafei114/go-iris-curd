package routes

import (
	"strconv"
	"strings"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"

	"go-iris-curd/main/web/supports"
	"go-iris-curd/main/web/supports/vo"

	modeDemo "go-iris-curd/main/web/models/demo"
	modeSys "go-iris-curd/main/web/models/system"
)

func DemoHub(party iris.Party) {
	// demo测试API模块
	product := party.Party("/product")
	{
		product.Post("/", hero.Handler(AddOneProduct))                // 新增一个
		product.Get("/", hero.Handler(ProductList))                   // 产品列表
		product.Get("/{pid:long}", hero.Handler(GetOneProduct))       // 获得一个
		product.Put("/", hero.Handler(UpdateUser))                    // 更新用户
		product.Delete("/{uids:string}", hero.Handler(DeleteProduct)) // 删除产品
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

// 查询
func ProductList(ctx iris.Context) {
	var (
		err        error
		page       *supports.Pagination
		total      int64
		entityList = make([]*modeDemo.Product, 0)
	)

	page, err = supports.NewPagination(ctx)
	if err != nil {
		goto FAIL
	}

	total, err = supports.GetPagination(&entityList, page)
	if err != nil {
		goto ERR
	}

	ctx.JSON(vo.TableVO{
		Total: total,
		Rows:  entityList,
	})
	return

FAIL:
	supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
	return
ERR:
	supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
	return
}

// 删除用户
func DeleteProduct(ctx iris.Context, uids string) {
	uidList := strings.Split(uids, ",")
	if len(uidList) == 0 {
		ctx.Application().Logger().Error("删除产品错误, 参数不对.")
		supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
		return
	}

	dUids := make([]int64, 0)
	for _, v := range uidList {
		if v == "" {
			continue
		}
		uid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			ctx.Application().Logger().Error("删除用户错误, %s", err.Error())
			supports.Error(ctx, iris.StatusInternalServerError, supports.ParseParamsFailur, nil)
			return
		}
		dUids = append(dUids, uid)
	}

	effect, err := modeSys.DeleteByUsers(dUids)
	if err != nil {
		ctx.Application().Logger().Error("删除用户错误, %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.DeleteUsersFailur, nil)
		return
	}
	supports.Ok(ctx, supports.DeleteUsersSuccess, effect)
}
