package demo

import (
	"go-iris-curd/main/web/db"
	"time"

	"go-iris-curd/main/web/models"
	"go-iris-curd/main/web/supports"
)

// Product 产品信息
type Product struct {
	// 继承父类（id，创建时间、更新时间）
	models.Model `xorm:"extends"`
	ProductCode  string `xorm:"varchar(50)" json:"product_Code" form:"product_Code"`
	ProductName  string `xorm:"varchar(200)" json:"product_Name" form:"product_Name"`
	// 类型（使用弹出选择）
	ProductCateName string  `xorm:"varchar(200)" json:"product_Cate_Name" form:"product_Cate_Name"`
	ProductCateID   int64   `xorm:"INT(10)" json:"product_Cate_Id" form:"product_Cate_Id"`
	Price           float64 `xorm:"notnull" json:"price" form:"price"`
	Number          int     `xorm:"notnull" json:"number" form:"number"`
	// 开始有效期
	StartDate time.Time `xorm:"datetime 'start_Date'" json:"start_Date" form:"start_Date"`
	// 结束有效期
	EndDate time.Time `xorm:"datetime 'end_Date'" json:"end_Date" form:"end_Date"`
	// 类别 radio
	ProductType string `xorm:"varchar(200)" json:"product_Type" form:"product_Type"`
	// 自动完成
	ProductColor string `xorm:"varchar(200)" json:"product_Color" form:"product_Color"`
	// 是否有效
	IsValid bool `xorm:"tinyint(1) notnull" json:"is_Valid" form:"is_Valid"`
	// 内容（使用富文本保存）
	Content string `xorm:"mediumtext 'content'"`
}

// GetProductsByIds 根据id获得产品
func GetProductsByIds(uids []int64, page *supports.Pagination) ([]*Product, int64, error) {
	e := db.MasterEngine()
	Products := make([]*Product, 0)

	s := e.In("id", uids).Limit(page.Limit, page.Start)
	count, err := s.FindAndCount(&Products)
	return Products, count, err
}
