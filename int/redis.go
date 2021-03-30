package int

import (
	"github.com/go-redis/redis"
  "one/config"

  "one/global"
)

var (
	RconfRconf = config.Rconf
)

func Redis(addr string, password string, db int) *redis.Client{
	redis_opt := redis.Options{
		Addr: RconfRconf.addr,
		Password: RconfRconf.password,
		DB: RconfRconf.db,
	}
	// 创建连接池
	redisdb := redis.NewClient(&redis_opt)
	// 判断是否能够链接到数据库
	pong, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
  
	global.NB_REDIS = redisdb
  return redisdb
}
