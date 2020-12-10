package sql

import (
	_ "github.com/go-sql-driver/mysql" //
	"github.com/gohouse/gorose/v2"
)

//DB db
var engin *gorose.Engin

func init() {
	var err error

	// dbConfigMap := map[string]interface{}{
	// 	"host":     "10.20.2.146",
	// 	"username": "intsig",
	// 	"password": "intsig",
	// 	"port":     "3306",
	// 	"database": "test",
	// }
	dbConfig := gorose.Config{
		SetMaxOpenConns: 0,
		SetMaxIdleConns: 1,
		Driver:          "mysql",
		Dsn:             "intsig:intsig@tcp(10.20.2.146:3306)/test?charset=utf8mb4&parseTime=true",
	}

	// 初始化数据库链接, 默认会链接配置中 default 指定的值
	// 也可以在第二个参数中指定对应的数据库链接, 见下边注释的那一行链接示例
	engin, err = gorose.Open(&dbConfig)
	if err != nil {
		panic(err)
	}
	if err := engin.Ping(); nil != err {
		panic(err)
	}
}

//NewDB e
func NewDB() gorose.IOrm {
	return engin.NewOrm()
}
