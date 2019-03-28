package dbconn

import (
	
	 _"github.com/lib/pq"
	"database/sql"
	"github.com/gomodule/redigo/redis"
	. "admin-mvc/app/utils"
	"log"
	"time"	
)

var (
	RedisClient  *redis.Pool
    REDIS_DB   int
    Db *sql.DB
)

func init() {

	var err error
 
	//创Redis连接
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(Config.RedisProtocol, Config.RedisHostPort)
			if err != nil {
				return nil, err
			}
			// Auth 有密码的话需要做认证哦～
			c.Do("AUTH","密码")

			// 选择db
			c.Do("SELECT", REDIS_DB)
				return c, nil
		},
	}


	//创建Postgresql连接
	//Db, err = sql.Open("postgres", "user=root dbname=chitchat password=root sslmode=disable")
	Db, err = sql.Open(Config.DbDriverName, Config.DataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return
}


