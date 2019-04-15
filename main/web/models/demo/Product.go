package demo

import (
	"go-iris-curd/main/web/db"
	"time"

	"go-iris-curd/main/web/supports"
)

// Product 产品信息
type Product struct {
	ID          int64      `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	ProductCode string     `xorm:"varchar(50)" json:"product_Code" form:"product_Code"`
	ProductName string     `xorm:"varchar(200)" json:"product_Name" form:"product_Name"`
	Price       float64    `xorm:"notnull" json:"price" form:"price"`
	Number      int        `xorm:"notnull" json:"number" form:"number"`
	CreatedAt   time.Time  `json:"created_At"`
	UpdatedAt   time.Time  `json:"updated_At"`
	DeletedAt   *time.Time `json:"deleted_At" sql:"index"`
}

// ProductCategory 产品类别
type ProductCategory struct {
	ID        int64      `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	numCode   string     `xorm:"varchar(50)" json:"numCode" form:"numCode"`
	chnName   string     `xorm:"varchar(200)" json:"chnName" form:"chnName"`
	CreatedAt time.Time  `json:"created_At"`
	UpdatedAt time.Time  `json:"updated_At"`
	DeletedAt *time.Time `json:"deleted_At" sql:"index"`
}

// GetProductsByIds 根据id获得产品
func GetProductsByIds(uids []int64, page *supports.Pagination) ([]*Product, int64, error) {
	e := db.MasterEngine()
	Products := make([]*Product, 0)

	s := e.In("id", uids).Limit(page.Limit, page.Start)
	count, err := s.FindAndCount(&Products)
	return Products, count, err
}
