package supports

import (
	"fmt"

	"go-iris-curd/main/web/db"

	"github.com/go-xorm/xorm"
)

// GetPagination 抽象出公用
func GetPagination(list interface{}, page *Pagination) (int64, error) {
	var (
		e       = db.MasterEngine()
		session *xorm.Session
		err     error
		count   int64
	)

	// 需要where查询
	if len(page.SearchValue) > 0 {
		session = e.Where(fmt.Sprintf("%s like ?", page.SearchKey), "%"+page.SearchValue+"%")
		session = session.Limit(page.Limit, page.Start)
	} else {
		session = e.Limit(page.Limit, page.Start)
	}

	count, err = session.FindAndCount(list)

	return count, err
}
