package routes

import (
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
		product.Post("/", hero.Handler(CreateProduct))                // 新增一个
		product.Get("/", hero.Handler(ProductList))                   // 产品列表
		product.Get("/{pid:long}", hero.Handler(GetOneProduct))       // 获得一个
		product.Put("/", hero.Handler(UpdateProduct))                 // 更新类别
		product.Delete("/{uids:string}", hero.Handler(DeleteProduct)) // 删除产品
		product.Delete("/del", hero.Handler(DeleteProducts))          // 批量删除产品
	}

	productCate := party.Party("/productCate")
	{
		productCate.Post("/", hero.Handler(CreateProductCate))                // 新增一个
		productCate.Get("/", hero.Handler(ProductCateList))                   // 产品列表
		productCate.Get("/{pid:long}", hero.Handler(GetOneProductCate))       // 获得一个
		productCate.Put("/", hero.Handler(UpdateProductCate))                 // 更新类别
		productCate.Delete("/{uids:string}", hero.Handler(DeleteProductCate)) // 删除产品
		productCate.Delete("/del", hero.Handler(DeleteProductCates))          // 批量删除产品
	}
}

// CreateProduct 增加产品
func CreateProduct(ctx iris.Context) {
	entity := new(modeDemo.Product)
	_, status, err := supports.DoCreateEntity(ctx, entity)
	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		supports.Error(ctx, status, supports.OptionFailur, nil)
		return
	}
	supports.OkR(ctx, supports.OptionSuccess)
}

// UpdateProduct 更新产品
func UpdateProduct(ctx iris.Context) {
	// entity := new(modeDemo.Product)
	// effect, status, err := supports.DoUpdateEntity(ctx, entity.ID, entity)
	// if err != nil {
	// 	ctx.Application().Logger().Errorf(err.Error())
	// 	supports.Error(ctx, status, supports.OptionFailur, nil)
	// 	return
	// }

	entity := new(modeDemo.Product)
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
	var entityList = make([]*modeDemo.Product, 0)

	total, status, err := supports.DoPaginationList(ctx, &entityList)

	if err != nil {
		supports.Error(ctx, status, supports.OptionFailur, err.Error())
	}

	ctx.JSON(vo.TableVO{
		Total: total,
		Rows:  entityList,
	})
	return

}

// DeleteProduct 删除产品
func DeleteProduct(ctx iris.Context, ids string) {
	var (
		err    error
		effect int64
		entity = new(modeDemo.Product)
	)
	effect, err = supports.DoDeleteEntitys(ids, entity)

	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.DeleteFailur, nil)
	}

	supports.Ok(ctx, supports.DeleteSuccess, effect)
}

// DeleteProducts 删除产品
func DeleteProducts(ctx iris.Context) {

	cr := new(vo.CommonRequest)

	if err := ctx.ReadJSON(cr); err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}

	var (
		err    error
		effect int64
		entity = new(modeDemo.Product)
	)
	effect, err = supports.DoDeleteEntitys(cr.Ids, entity)

	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.DeleteFailur, nil)
	}

	supports.Ok(ctx, supports.DeleteSuccess, effect)
}

// -------

// CreateProductCate 增加产品
func CreateProductCate(ctx iris.Context) {
	entity := new(modeDemo.ProductCategory)
	_, status, err := supports.DoCreateEntity(ctx, entity)
	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		supports.Error(ctx, status, supports.OptionFailur, nil)
		return
	}
	supports.OkR(ctx, supports.OptionSuccess)
}

// UpdateProductCate 更新产品
func UpdateProductCate(ctx iris.Context) {
	// entity := new(modeDemo.ProductCategory)
	// effect, status, err := supports.DoUpdateEntity(ctx, entity.ID, entity)
	// golog.Info(entity)
	// if err != nil {
	// 	ctx.Application().Logger().Errorf(err.Error())
	// 	supports.Error(ctx, status, supports.OptionFailur, nil)
	// 	return
	// }

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

// GetOneProductCate 获得一个产品类别
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
	var entityList = make([]*modeDemo.ProductCategory, 0)

	total, status, err := supports.DoPaginationList(ctx, &entityList)

	if err != nil {
		supports.Error(ctx, status, supports.OptionFailur, err.Error())
	}

	ctx.JSON(vo.TableVO{
		Total: total,
		Rows:  entityList,
	})

	return
}

// DeleteProductCate 删除产品
func DeleteProductCate(ctx iris.Context, ids string) {
	var (
		err    error
		effect int64
		entity = new(modeDemo.ProductCategory)
	)
	// ids = ctx.URLParam("uids")
	effect, err = supports.DoDeleteEntitys(ids, entity)

	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.DeleteFailur, nil)
	}

	supports.Ok(ctx, supports.DeleteSuccess, effect)
}

// DeleteProductCates 删除产品
func DeleteProductCates(ctx iris.Context) {

	cr := new(vo.CommonRequest)

	var (
		err    error
		effect int64
		entity = new(modeDemo.ProductCategory)
	)

	if err := ctx.ReadJSON(cr); err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}

	effect, err = supports.DoDeleteEntitys(cr.Ids, entity)

	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.DeleteFailur, nil)
	}

	supports.Ok(ctx, supports.DeleteSuccess, effect)
}
