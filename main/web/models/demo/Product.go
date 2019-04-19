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
	ProductCode  string `xorm:"varchar(50)" json:"productCode" form:"productCode"`
	ProductName  string `xorm:"varchar(200)" json:"productName" form:"productName"`
	// 类型（使用弹出选择）
	ProductCateName string  `xorm:"varchar(200)" json:"productCateName" form:"productCateName"`
	ProductCateID   int64   `xorm:"INT(10)" json:"productCateId" form:"productCateId"`
	Price           float64 `xorm:"notnull" json:"price" form:"price"`
	Number          int     `xorm:"notnull" json:"number" form:"number"`
	// 开始有效期
	StartDate time.Time `xorm:"datetime 'start_Date'" json:"startDate" form:"startDate"`
	// 结束有效期
	EndDate time.Time `xorm:"datetime 'end_Date'" json:"endDate" form:"endDate"`
	// 类别 select
	ProductType string `xorm:"varchar(200)" json:"productType" form:"productType"`
	// 颜色 radio
	ProductColor string `xorm:"varchar(200)" json:"productColor" form:"productColor"`
	// 是否有效
	IsValid bool `xorm:"tinyint(1) notnull 'is_Valid'" json:"isValid" form:"isValid"`
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
