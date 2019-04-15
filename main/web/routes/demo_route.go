package routes

import (
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"

	"go-iris-curd/main/web/supports"
	"go-iris-curd/main/web/supports/vo"

	modeDemo "go-iris-curd/main/web/models/demo"
)

// DemoHub demo路由
func DemoHub(party iris.Party) {
	// demo测试API模块
	product := party.Party("/product")
	{
		product.Post("/", hero.Handler(CreateProduct))                 // 新增一个
		product.Get("/", hero.Handler(ProductList))                    // 产品列表
		product.Get("/{pid:long}", hero.Handler(GetOneProduct))        // 获得一个
		product.Put("/", hero.Handler(UpdateProduct))                  // 更新用户
		product.Delete("/{uids:string}", hero.Handler(DeleteProducts)) // 删除产品
	}

	productCate := party.Party("/productCate")
	{
		productCate.Post("/", hero.Handler(CreateProductCate))                 // 新增一个
		productCate.Get("/", hero.Handler(ProductCateList))                    // 产品列表
		productCate.Get("/{pid:long}", hero.Handler(GetOneProductCate))        // 获得一个
		productCate.Put("/", hero.Handler(UpdateProductCate))                  // 更新用户
		productCate.Delete("/{uids:string}", hero.Handler(DeleteProductCates)) // 删除产品
	}
}

// CreateProduct 增加产品
func CreateProduct(ctx iris.Context) {
	product := new(modeDemo.Product)
	if err := ctx.ReadJSON(product); err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}

	golog.Info(product)
	product.CreatedAt = time.Now()
	_, err := supports.CreateEntity(product)
	if err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}
	supports.OkR(ctx, supports.OptionSuccess)
}

// UpdateProduct 更新产品
func UpdateProduct(ctx iris.Context) {
	product := new(modeDemo.Product)
	if err := ctx.ReadJSON(&product); err != nil {
		ctx.Application().Logger().Errorf("更新产品[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.OptionFailur, nil)
		return
	}
	effect, err := supports.UpdateEntityByID(product.ID, product)
	if err != nil {
		ctx.Application().Logger().Errorf("更新产品[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, effect)
}

// GetOneProduct 获得一个产品
func GetOneProduct(ctx iris.Context, pid int64) {
	product := new(modeDemo.Product)
	product.ID = pid
	_, err := supports.GetOneEntity(product)
	if err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, product)
}

// ProductList 分页查询
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

// DeleteProducts 删除产品
func DeleteProducts(ctx iris.Context, ids string) {
	var (
		err    error
		effect int64
		entity = new(modeDemo.Product)
	)
	effect, err = supports.DeleteEntitys(ids, entity)

	if err != nil {
		ctx.Application().Logger().Error(err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.DeleteFailur, nil)
	}

	supports.Ok(ctx, supports.DeleteSuccess, effect)
}

// -------

// CreateProductCate 增加产品
func CreateProductCate(ctx iris.Context) {
	entity := new(modeDemo.ProductCategory)
	if err := ctx.ReadJSON(entity); err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}

	golog.Info(entity)
	entity.CreatedAt = time.Now()
	_, err := supports.CreateEntity(entity)
	if err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}
	supports.OkR(ctx, supports.OptionSuccess)
}

// UpdateProductCate 更新产品
func UpdateProductCate(ctx iris.Context) {
	entity := new(modeDemo.ProductCategory)
	if err := ctx.ReadJSON(&entity); err != nil {
		ctx.Application().Logger().Errorf("更新产品[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.OptionFailur, nil)
		return
	}
	effect, err := supports.UpdateEntityByID(entity.ID, entity)
	if err != nil {
		ctx.Application().Logger().Errorf("更新产品[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, effect)
}

// GetOneProduct 获得一个产品
func GetOneProductCate(ctx iris.Context, pid int64) {
	entity := new(modeDemo.ProductCategory)
	entity.ID = pid
	_, err := supports.GetOneEntity(entity)
	if err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, entity)
}

// ProductCateList 分页查询
func ProductCateList(ctx iris.Context) {
	var (
		err        error
		page       *supports.Pagination
		total      int64
		entityList = make([]*modeDemo.ProductCategory, 0)
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

// DeleteProductCates 删除产品
func DeleteProductCates(ctx iris.Context, ids string) {
	var (
		err    error
		effect int64
		entity = new(modeDemo.ProductCategory)
	)
	effect, err = supports.DeleteEntitys(ids, entity)

	if err != nil {
		ctx.Application().Logger().Error(err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.DeleteFailur, nil)
	}

	supports.Ok(ctx, supports.DeleteSuccess, effect)
}
