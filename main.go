package main

import (
	"github.com/jmoiron/sqlx"
	"go_web/db"
	"go_web/routers"
	"log"
)

func main() {
	// 初始化数据库连接
	sqlxDB := db.ConnectDB()

	// 确保程序退出时关闭数据库连接
	defer func(sqlxDB *sqlx.DB) {
		err := sqlxDB.Close()
		if err != nil {
		}
	}(sqlxDB)
	// 注册路由
	r := routers.InitRouter()

	// 启动服务器
	if err := r.Run(); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}

}
