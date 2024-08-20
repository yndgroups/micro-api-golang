package redis

import "time"

var UPDATE = update{}

type update struct{}

/*
UpdateExpire 更新Redis存储时间 过期时间 单位秒
@Param key 存储键
@Return float64, error
@Author YangDaqiong
@Date 2022-01-03
*/
func (u *update) UpdateExpire(key string, expire time.Duration) error {
	return GetRedisInstance().Expire(ctx, key, expire*time.Second).Err()
}
