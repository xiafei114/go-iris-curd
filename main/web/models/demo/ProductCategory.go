package demo

import "go-iris-curd/main/web/models"

// ProductCategory 产品类别
type ProductCategory struct {
	models.Model `xorm:"extends"`
	NumCode      string `xorm:"varchar(50)" json:"numCode" form:"numCode"`
	ChnName      string `xorm:"varchar(200)" json:"chnName" form:"numCode"`
}
