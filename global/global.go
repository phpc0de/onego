package global

import (
	"github.com/go-redis/redis"
	"onego/config"
  "github.com/go-xorm/xorm"
)

var (
	NB_DB            *xorm.Engine
	NB_REDIS         *redis.Client
)
