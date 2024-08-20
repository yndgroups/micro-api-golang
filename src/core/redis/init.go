package redis

import (
	"context"
	"core/config"
	"core/logger"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

// InitConnectRedis 初始化redis链接
func InitConnectRedis(config config.RedisConfig) {

	rdb = redis.NewClient(&redis.Options{
		Addr:         config.Host + ":" + config.Port, //集群模式一定要配置主节点的值  否则会报MOVED 4998
		Password:     config.Password,                 // 没有密码写空字符串
		DB:           config.Database,                 // use default DB
		MinIdleConns: 500,                             // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；
		PoolSize:     10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.SugarLogger.Infof("连接Redis失败:%v:%v, 失败原因：%v", config.Host, config.Port, err.Error())
	}
	logger.SugarLogger.Infof("连接Redis成功:%v:%v", config.Host, config.Port)
}

// GetRedisInstance 获取redis实例
func GetRedisInstance() *redis.Client {
	return rdb
}
