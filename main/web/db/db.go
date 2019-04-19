package db

import (
	"fmt"
	"go-iris-curd/main/inits/parse"
	"go-iris-curd/main/utils"
	"sync"

	"github.com/kataras/golog"
	"github.com/xormplus/core"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

// MasterEngine 主库，单例
func MasterEngine() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	master := parse.DBConfig.Master
	engine, err := xorm.NewEngine(master.Dialect, GetConnURL(&master))
	if err != nil {
		golog.Fatalf("@@@ Instance Master DB error!! %s", err)
		return nil
	}
	settings(engine, &master)
	engine.SetMapper(core.GonicMapper{})

	masterEngine = engine
	return masterEngine
}

// SlaveEngine 从库，单例
func SlaveEngine() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}

	slave := parse.DBConfig.Slave
	engine, err := xorm.NewEngine(slave.Dialect, GetConnURL(&slave))
	if err != nil {
		golog.Fatalf("@@@ Instance Slave DB error!! %s", err)
		return nil
	}
	settings(engine, &slave)

	slaveEngine = engine
	return engine
}

// settings 设置
func settings(engine *xorm.Engine, info *parse.DBConfigInfo) {
	engine.ShowSQL(info.ShowSql)
	engine.SetTZLocation(utils.SysTimeLocation)
	if info.MaxIdleConns > 0 {
		engine.SetMaxIdleConns(info.MaxIdleConns)
	}
	if info.MaxOpenConns > 0 {
		engine.SetMaxOpenConns(info.MaxOpenConns)
	}

	// 性能优化的时候才考虑，加上本机的SQL缓存
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//engine.SetDefaultCacher(cacher)
}

// GetConnURL 获取数据库连接的url
// true：master主库
func GetConnURL(info *parse.DBConfigInfo) (url string) {
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		info.User,
		info.Password,
		info.Host,
		info.Port,
		info.Database,
		info.Charset)
	//golog.Infof("@@@ DB conn==>> %s", url)
	return
}
