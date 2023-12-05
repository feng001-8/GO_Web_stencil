package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"go_web/settings"
)

var db *sql.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	// DSN:Data Source Name
	//dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("sql.Open() failed err", zap.Error(err))
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("db.Close() failed %s\n", err)
		}
	}(db) // 注意这行代码要写在上面err判断的下面

	return
}

func Close() {
	_ = db.Close()
}
