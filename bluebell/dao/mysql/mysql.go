package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golearn/bluebell/setting"
)

var db *sqlx.DB

func Init() (err error) {
	mysql := *setting.Conf.MySQLConfig
	fmt.Println("have a look", mysql.Host)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		mysql.User,
		mysql.Password,
		mysql.Host,
		mysql.Port,
		mysql.DB,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB faile的", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(mysql.MaxOpenConns)
	db.SetMaxIdleConns(mysql.MaxIdleConns)
	return

}
func Close() {
	_ = db.Close()
}
