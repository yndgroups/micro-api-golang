package redis

import "time"

var DELETE = delete{}

type delete struct{}

// Delete 根据key删除存储信息
func (d *delete) DelByKey(key string) error {
	return GetRedisInstance().Del(ctx, key).Err()
}

/*
BatchDel 批量删除
@Param pairs 存储键
@Return void 无返回值
@Author yangDaqiong
@Date 2022-01-03
*/
func (d *delete) BatchDel(key ...string) error {
	return GetRedisInstance().Del(ctx, key...).Err()
}

/*
HashDel hash删除  key field1 [field2] 删除一个或多个哈希表字段
@Param key 存储键
@Param value 存储值
@Author yangDaqiong
@Date 2022-01-03
*/
func (d *delete) HashDel(key, field string) error {
	return GetRedisInstance().HDel(ctx, key, field).Err()
}

// BLpop 命令移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
func (d *delete) BLpop(timeout time.Duration, keys ...string) (value []string, err error) {
	return GetRedisInstance().BLPop(ctx, timeout, keys...).Result()
}
