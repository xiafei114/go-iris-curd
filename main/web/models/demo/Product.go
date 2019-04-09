package demo

import (
	"go-iris-curd/main/web/models"
)

type Product struct {
	models.Model
	Id          int64   `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	ProductCode string  `xorm:"notnull" json:"productCode" form:"productCode"`
	ProductName string  `xorm:"notnull" json:"productName" form:"productName"`
	Price       float64 `xorm:"notnull" json:"price" form:"price"`
	Number      int     `xorm:"notnull" json:"number" form:"number"`
}
