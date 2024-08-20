package redis

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var SET = set{}

type set struct{}

// >>>>>>>>>>>>>>>>>>>> 存储 >>>>>>>>>>>>>>>>>>>
// Set 设置给定key的值。如果key已经存储其他值，SET就覆写旧值，且无视类型。
// @Param key 存储键
// @Param value 存储值
// @Param expire 过期时间
// @Return error
// @Author YangDaqiong
// @Date 2022-01-03
func (s *set) Set(key string, val any, expire time.Duration) error {
	exists, err := GET.KeyExists(key)
	if err != nil {
		return errors.New("redis检查key是否存在是发生错误")
	}
	// key存在及覆盖，不存在则进行插入
	if exists > 0 {
		return GetRedisInstance().SetXX(ctx, key, val, expire*time.Second).Err()
	} else {
		return GetRedisInstance().SetNX(ctx, key, val, expire*time.Second).Err()
	}
}

// SetJson json redis存储
// 存储正常返回 nil, 存贮异常返回异常信息
func (s *set) SetJson(key string, value any, expire time.Duration) error {
	exists, err := GET.KeyExists(key)
	if err != nil {
		return errors.New("redis检查key是否存在是发生错误")
	}
	if jsonData, err := json.Marshal(value); err != nil {
		return errors.New("JSON数据解析失败")
	} else {
		// key存在及覆盖，不存在则进行插入
		if exists > 0 {
			return GetRedisInstance().SetXX(ctx, key, jsonData, expire*time.Second).Err()
		} else {
			return GetRedisInstance().SetNX(ctx, key, jsonData, expire*time.Second).Err()
		}
	}
}

/*
Mset 命令用于同时设置一个或多个key-value对
@Param pairs 存储键
@Return void 无返回值
@Author yangDaqiong
@Date 2022-01-03
*/
func (s *set) Mset(pairs ...any) error {
	return GetRedisInstance().MSet(ctx, pairs...).Err()
}

/*
HashSet hash存储
@Param key 存储键
@Param value 存储值
@Author YangDaqiong
@Date 2022-01-03
*/
func (s *set) HashSet(key, field, value string) error {
	return GetRedisInstance().HSet(ctx, key, field, value).Err()
}

// RPush 命令用于将一个或多个值插入到列表的尾部(最右边)。
// 如果列表不存在，一个空列表会被创建并执行 RPUSH 操作。 当列表存在但不是列表类型时，返回一个错误。
// 注意：在 Redis 2.4 版本以前的 RPUSH 命令，都只接受单个 value 值。
func (s *set) RPush(key string, member string) (err error) {
	return GetRedisInstance().RPush(ctx, key, member).Err()
}

/*
ZAdd 命令用于将一个或多个成员元素及其分数值加入到有序集当中。
如果某个成员已经是有序集的成员，那么更新这个成员的分数值，并通过重新插入这个成员元素，来保证该成员在正确的位置上。
分数值可以是整数值或双精度浮点数。
如果有序集合 key 不存在，则创建一个空的有序集并执行 ZADD 操作。
当 key 存在但不是有序集类型时，返回一个错误。
注意： 在 Redis 2.4 版本以前， ZADD 每次只能添加一个元素
*/
func (s *set) ZAdd(key string, member, score redis.Z) error {
	return nil
	// return GetRedisInstance().ZAdd(ctx, key, &score, &member).Err()
}
