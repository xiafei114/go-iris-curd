package models

import "time"

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID        int64     `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	CreatedAt time.Time `xorm:"created" json:"created_At"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_At"`
	DeletedAt time.Time `json:"deleted_At"`
}
