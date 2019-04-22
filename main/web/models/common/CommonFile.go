package common

import (
	"go-iris-curd/main/web/models"
)

// CommonFile 通用文件存储
type CommonFile struct {
	models.Model `xorm:"extends"`
	FileName     string `xorm:"varchar(500) notnull 'file_Name'" json:"fileName" form:"fileName"`
	ContentType  string `xorm:"varchar(500) notnull 'content_Type'" json:"contentType" form:"contentType"`
	FilePath     string `xorm:"varchar(2000) notnull 'file_Path'" json:"filePath" form:"filePath"`
}
