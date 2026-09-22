package mysqldb

import (
	"go_jichu/conf"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var GlobalMDB *gorm.DB

func InitMySQL(cfg *conf.ConfigStruct) error {
	// 1. 现阶段：主库和从库都指向同一个 MySQL 实例
	// 未来扩展主从时，只需把 Replicas 换成从库的 DSN 即可
	dsnMaster := "root:123456@tcp(127.0.0.1:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local"
	dsnSlave := "root:123456@tcp(127.0.0.1:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local" // 当前单机复用，未来改成从库IP

	var err error
	// 初始化 GORM 基础连接（作为主库入口）
	GlobalMDB, err = gorm.Open(mysql.Open(dsnMaster), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印 SQL 方便调试
	})
	if err != nil {
		return err
	}

	// 2. 配置 dbresolver 插件（实现读写分离与连接池统一管理）
	err = GlobalMDB.Use(
		dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dsnMaster)}, // 写库（主库）集群
			Replicas: []gorm.Dialector{mysql.Open(dsnSlave)},  // 读库（从库）集群
			Policy:   dbresolver.RandomPolicy{},               // 多从库时的负载均衡策略（随机）
		}).
			SetMaxIdleConns(cfg.MYSQL.SetMaxIdleConns).                                    // 连接池最大空闲连接数
			SetMaxOpenConns(cfg.MYSQL.SetMaxOpenConns).                                    // 连接池最大打开连接数
			SetConnMaxLifetime(time.Duration(cfg.MYSQL.SetConnMaxLifetime) * time.Second). // 连接可复用的最大时间
			SetConnMaxIdleTime(time.Duration(cfg.MYSQL.SetConnMaxIdleTime) * time.Second), // 连接最大空闲时间
	)
	if err != nil {
		return err
	}
	return nil
	// fmt.Println("MySQL 及连接池初始化成功（已兼容未来主从复制）！")
}
