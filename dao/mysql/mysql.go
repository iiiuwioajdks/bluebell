package mysql

/**
mysql初始化相关
*/

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	//"root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"

	// Connect 包括 Open 和 Ping
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect db error", zap.Error(err))
		return err
	}

	//db.SetMaxOpenConns(100)
	//db.SetMaxIdleConns(10)

	return err
}

func Close() {
	err := db.Close()
	if err != nil {
		zap.L().Error("mysql close error:", zap.Error(err))
	}
}
