package demo

import "time"

// ProductCategory 产品类别
type ProductCategory struct {
	ID        int64      `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	NumCode   string     `xorm:"varchar(50)" json:"numCode" form:"numCode"`
	ChnName   string     `xorm:"varchar(200)" json:"chnName" form:"numCode"`
	CreatedAt time.Time  `xorm:"created" json:"created_At"`
	UpdatedAt time.Time  `xorm:"updated" json:"updated_At"`
	DeletedAt *time.Time `json:"deleted_At" sql:"index"`
}
