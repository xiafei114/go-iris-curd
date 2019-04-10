package demo

import (
	"go-iris-curd/main/web/db"
	"go-iris-curd/main/web/models"

	"go-iris-curd/main/web/supports"
)

type Product struct {
	models.Model
	// Id          int64   `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	ProductCode string  `xorm:"notnull" json:"productCode" form:"productCode"`
	ProductName string  `xorm:"notnull" json:"productName" form:"productName"`
	Price       float64 `xorm:"notnull" json:"price" form:"price"`
	Number      int     `xorm:"notnull" json:"number" form:"number"`
}

/**
获得一个产品
*/
func GetOneProduct(product *Product) (bool, error) {
	e := db.MasterEngine()
	return e.Get(product)
}

func CreateProduct(Product *Product) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(Product)
}

func GetProductsByIds(uids []int64, page *supports.Pagination) ([]*Product, int64, error) {
	e := db.MasterEngine()
	Products := make([]*Product, 0)

	s := e.In("id", uids).Limit(page.Limit, page.Start)
	count, err := s.FindAndCount(&Products)
	return Products, count, err
}

func UpdateProductById(Product *Product) (int64, error) {
	e := db.MasterEngine()
	return e.Id(Product.ID).Update(Product)
}

func DeleteByProducts(uids []int64) (effect int64, err error) {
	e := db.MasterEngine()

	u := new(Product)
	for _, v := range uids {
		i, err1 := e.Id(v).Delete(u)
		effect += i
		err = err1
	}
	return
}

func GetPaginationProducts(page *supports.Pagination) ([]*Product, int64, error) {
	var (
		err        error
		count      int64
		entityList = make([]*Product, 0)
	)

	count, err = supports.GetPagination(&entityList, page)

	return entityList, count, err
}
