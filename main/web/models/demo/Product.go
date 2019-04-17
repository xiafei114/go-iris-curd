package demo

import (
	"go-iris-curd/main/web/db"

	"go-iris-curd/main/web/models"
	"go-iris-curd/main/web/supports"
)

// Product 产品信息
type Product struct {
	models.Model    `xorm:"extends"`
	ProductCode     string  `xorm:"varchar(50)" json:"product_Code" form:"product_Code"`
	ProductName     string  `xorm:"varchar(200)" json:"product_Name" form:"product_Name"`
	ProductCateName string  `xorm:"varchar(200)" json:"product_Cate_Name" form:"product_Cate_Name"`
	ProductCateID   int64   `xorm:"INT(10)" json:"product_Cate_Id" form:"product_Cate_Id"`
	Price           float64 `xorm:"notnull" json:"price" form:"price"`
	Number          int     `xorm:"notnull" json:"number" form:"number"`
}

// GetProductsByIds 根据id获得产品
func GetProductsByIds(uids []int64, page *supports.Pagination) ([]*Product, int64, error) {
	e := db.MasterEngine()
	Products := make([]*Product, 0)

	s := e.In("id", uids).Limit(page.Limit, page.Start)
	count, err := s.FindAndCount(&Products)
	return Products, count, err
}
