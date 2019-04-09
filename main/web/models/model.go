package models

import "time"

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID        uint       `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	CreatedAt time.Time  `xorm:"notnull" json:"created_At" form:"created_At"`
	UpdatedAt time.Time  `xorm:"notnull" json:"updated_At" form:"updated_At"`
	DeletedAt *time.Time `xorm:"notnull" json:"deleted_At" form:"deleted_At" sql:"index"`
}
