package mysql_utils

import (
	"github.com/XiuCai/XiuCaiAPI/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"time"
)

var (
	mysqlEngin *xorm.Engine = nil
	config     *MysqlConfig = nil
)

func initMysqlEngin() *xorm.Engine {

	if config == nil {
		utils.Logger.Error("mysql configs is not instance")
	}
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.UserName, config.Password, config.Address, config.DbName)
	engine, err := xorm.NewEngine("mysql", mysqlInfo)
	if err != nil {
		utils.Logger.Error("Failed to set up mysql:", "--ERROR--", err)
		panic(err)
	}
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(config.MaxIdel)
	engine.SetMaxOpenConns(config.MaxOpen)
	//logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	//if err != nil {
	//	log.Fatalf("Fail to create xorm system logger: %v\n", err)
	//}
	//logger := xorm.NewSimpleLogger(logWriter)
	//engine.SetLogger(logger)
	utils.Logger.Notice("Mysql connection set up:", mysqlInfo)
	return engine
}

func SetConfig(in *MysqlConfig) {
	config = in
}

func GetInstance() *xorm.Engine {
	if mysqlEngin == nil {
		mysqlEngin = initMysqlEngin()
	}
	return mysqlEngin
}

//自动重连
func DBAutoConnect() {
	go func() {
		for{
			time.Sleep(10 * time.Second)
			err := mysqlEngin.Ping()
			if err != nil {
				initMysqlEngin()
			}
		}
	}()

}