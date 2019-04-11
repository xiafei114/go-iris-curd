package demo

import (
	"go-iris-curd/main/web/db"
	"time"

	"go-iris-curd/main/web/supports"
)

// Product 产品信息
type Product struct {
	ID          int64      `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	ProductCode string     `xorm:"notnull" json:"product_Code" form:"product_Code"`
	ProductName string     `xorm:"notnull" json:"product_Name" form:"product_Name"`
	Price       float64    `xorm:"notnull" json:"price" form:"price"`
	Number      int        `xorm:"notnull" json:"number" form:"number"`
	CreatedAt   time.Time  `xorm:"notnull" json:"created_At"`
	UpdatedAt   time.Time  `xorm:"notnull" json:"updated_At"`
	DeletedAt   *time.Time `xorm:"notnull" json:"deleted_At" sql:"index"`
}

// GetOneProduct 获得一个产品
func GetOneProduct(product *Product) (bool, error) {
	e := db.MasterEngine()
	return e.Get(product)
}

// CreateProduct 新增产品
func CreateProduct(Product *Product) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(Product)
}

// GetProductsByIds 根据id获得产品
func GetProductsByIds(uids []int64, page *supports.Pagination) ([]*Product, int64, error) {
	e := db.MasterEngine()
	Products := make([]*Product, 0)

	s := e.In("id", uids).Limit(page.Limit, page.Start)
	count, err := s.FindAndCount(&Products)
	return Products, count, err
}

// UpdateProductByID 更新单元产品
func UpdateProductByID(Product *Product) (int64, error) {
	e := db.MasterEngine()
	return e.Id(Product.ID).Update(Product)
}
