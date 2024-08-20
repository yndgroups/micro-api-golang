package redis

import (
	"core/config"
	"encoding/json"
	"errors"
	"fmt"
	"protobuf/build/common"
	"strings"
	"time"

	"github.com/ansel1/merry"
	"github.com/redis/go-redis/v9"
)

var GET = &get{}

type get struct{}

// 根据key获取redis存储的 [字符串]
func (g *get) GetString(key string) (string, error) {
	println("keykeykeykeykeykeykey", key)
	// return GetRedisInstance().Get(ctx, key).Result()
	if value, err := GetRedisInstance().Get(ctx, key).Result(); err != nil {
		println(value, "valuevaluevaluevaluevaluevalue")
		return "", errors.New("redis数据获取失败")
	} else {
		println(value, "valuevaluevaluevaluevaluevalue")
		return value, nil
	}
}

// 根据key获取redis存储的 [Int] 数据
// @Param key 存储键
// @Return int, error
// @Author YangDaqiong
// @Date 2022-01-03
func (g *get) GetInt(key string) (int, error) {
	return GetRedisInstance().Get(ctx, key).Int()
}

// 根据key获取redis存储的 [Int64] 数据
// @Param key 存储键
// @Return int, error
// @Author YangDaqiong
// @Date 2022-01-03
func (g *get) GetInt64(key string) (int64, error) {
	return GetRedisInstance().Get(ctx, key).Int64()
}

/*
GetUint64 根据当前key获取存储的 uint64 数据
@Param key 存储键
@Return uint64, error
@Author YangDaqiong
@Date 2022-01-03
*/
func (g *get) GetUint64(key string) (uint64, error) {
	v, err := GetRedisInstance().Get(ctx, key).Uint64()
	if err == redis.Nil {
		return 0, nil
	}
	return v, err
}

// 根据key获取redis存储的 [Float32] 数据
func (g *get) GetFloat32(key string) (float32, error) {
	return GetRedisInstance().Get(ctx, key).Float32()
}

// 根据key获取redis存储的 [Float64] 数据
// @Param key 存储键
// @Return float64, error
// @Author YangDaqiong
// @Date 2022-01-03
func (g *get) GetFloat64(key string) (float64, error) {
	return GetRedisInstance().Get(ctx, key).Float64()
}

// 根据key获取redis存储的 [Bool] 数据
func (g *get) GetBool(key string) (bool, error) {
	return GetRedisInstance().Get(ctx, key).Bool()
}

// 根据key获取redis存储的 [json] 数据
func (g *get) GetJson(key string) ([]byte, error) {
	if value, err := GetRedisInstance().Get(ctx, key).Result(); err != nil {
		return nil, errors.New("redis数据获取失败")
	} else {
		return []byte(value), nil
	}
}

/*
KeyExists 检测当前key值是否存在 返回 bool 数据
@Param key 存储键
@Return bool, error
@Author YangDaqiong
@Date 2022-01-03
*/
func (g *get) KeyExists(key string) (int64, error) {
	return GetRedisInstance().Exists(ctx, key).Result()
}

/*
Keys 命令用于查找所有符合给定模式 pattern 的 key
@Param pattern 正则
@Return keys []string,error 返回符合给定模式的 key 列表 (Array)。
@Author yangDaqiong
@Date 2022-01-03
*/
func (g *get) Keys(pattern string) (keys []string, err error) {
	return GetRedisInstance().Keys(ctx, pattern).Result()
}

/*
HashValues 获取哈希表中所有值。
@Param key 存储键
@Author yangDaqiong
@Date 2022-03-29
*/
func (g *get) HashValues(key string) ([]string, error) {
	return GetRedisInstance().HVals(ctx, key).Result()
}

/*
HashLen 获取哈希表中字段的数量
@Param key 存储键
*/
func (g *get) HashLen(key string) (int64, error) {
	return GetRedisInstance().HLen(ctx, key).Result()
}

/*
HGet 命令用于返回哈希表中指定字段的值。
@Param key 存储键
@Param value 存储值
@Author yangDaqiong
@Date 2022-01-03
*/
func (g *get) HashGet(key, field string) (string, error) {
	return GetRedisInstance().HGet(ctx, key, field).Result()
}

/*
ListAllValuesWithPrefix 将接受一个键前缀并返回包含该前缀的所有键的值
@Param prefix 前缀
@Return map[string]string 返回map[string]string
@Author yangDaqiong
@Date 2022-01-03
*/
func (g *get) ListAllValuesWithPrefix(prefix string) (map[string]string, error) {
	// 抓取带有前缀的所有键
	keys, err := getKeys(fmt.Sprintf("%s*", prefix))
	if err != nil {
		return nil, err
	}
	// 将更新所有值
	values, _ := getKeyAndValuesMap(keys, prefix)
	return values, nil
}

/*
LRange 返回列表中指定区间内的元素，区间以偏移量 START 和 END 指定
@Param key 键值
@Param start 键值
@Param stop 键值
@Return  values []string, err error 返回values []string, err error
@Author yangDaqiong
@Date 2022-01-03
*/
func (g *get) LRange(key string, start, stop int) (values []string, err error) {
	return GetRedisInstance().LRange(ctx, key, int64(start), int64(stop)).Result()
}

// LLen Redis Llen 命令用于返回列表的长度。 如果列表 key 不存在，则 key 被解释为一个空列表，返回 0 。 如果 key 不是列表类型，返回一个错误。
func (g *get) LLen(key string) (value int64, err error) {
	return GetRedisInstance().LLen(ctx, key).Result()
}

// PTtl Redis Pttl命令以毫秒为单位返回 key 的剩余过期时间。
func (g *get) PTtl(key string) (time.Duration, error) {
	ttl, err := GetRedisInstance().PTTL(ctx, key).Result()
	if err != nil {
		return -1, err
	}
	return ttl, nil
}

// Ttl TTL 命令以秒为单位返回 key 的剩余过期时间。
func (g *get) Ttl(key string) (time.Duration, error) {
	//ttl, err := GetRedisInstance().Do("TTL", key).Int()
	ttl, err := GetRedisInstance().TTL(ctx, key).Result()
	if err != nil {
		return -1, err
	}
	return ttl, nil
}

// ZRank Redis Zrank 返回有序集中指定成员的排名。其中有序集成员按分数值递增(从小到大)顺序排列。
// 语法 ZRANK key member
func (g *get) ZRank(key, member string) (int64, error) {
	return GetRedisInstance().ZRank(ctx, key, member).Result()
}

// ZRange Redis Zrange 返回有序集中，指定区间内的成员。
// 其中成员的位置按分数值递增(从小到大)来排序。
// 具有相同分数值的成员按字典序(lexicographical order )来排列。
// 如果你需要成员按
// 值递减(从大到小)来排列，请使用 ZREVRANGE 命令。
// 下标参数 start 和 stop 都以 0 为底，也就是说，以 0 表示有序集第一个成员，以 1 表示有序集第二个成员，以此类推。
// 你也可以使用负数下标，以 -1 表示最后一个成员， -2 表示倒数第二个成员，以此类推。
func (g *get) ZRange(key string, start, stop int64) ([]string, error) {
	return GetRedisInstance().ZRange(ctx, key, start, stop).Result()
}

// ZRangeWithScores Redis Zrange 返回有序集中，指定区间内的成员。
func (g *get) ZRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return GetRedisInstance().ZRangeWithScores(ctx, key, start, stop).Result()
}

// ZRem Redis Zrem 命令用于移除有序集中的一个或多个成员，不存在的成员将被忽略。
// 当 key 存在但不是有序集类型时，返回一个错误。
// 注意： 在 Redis 2.4 版本以前， ZREM 每次只能删除一个元素。
func (g *get) ZRem(key, member string) error {
	return GetRedisInstance().ZRem(ctx, key, member).Err()
}

// ------------------------------------------- utils ------------------------------------

// GetRedisConfig 获取第三方配置信息
// @aram confId 配置id
// @Return Config 获取到的配置信息
// @Author yangDaqiong
// @Date 2022-02-17 0:40
func (g *get) GetRedisConfig(confId string) (*common.Config, error) {
	redisUserKey := config.Global.ServerConfig().ConfigPrefix + confId
	if confJson, err := g.GetJson(redisUserKey); err != nil {
		return nil, errors.New("配置信息获取失败！")
	} else {
		conf := common.Config{}
		if jsonErr := json.Unmarshal(confJson, &conf); jsonErr != nil {
			return nil, errors.New("配置信息JSON转换失败！")
		}
		return &conf, nil
	}
}

/*
GetPrimaryKey 获取唯一值
@Param key 存储键
@Return int64 返回int64类型的值
@Author yangDaqiong
@Date 2022-01-03
*/
func (g *get) GetPrimaryKey(key string) string {
	val := GetRedisInstance().Incr(ctx, key).Val()
	var numStr string
	switch true {
	case val > 0 && val < 10:
		numStr = fmt.Sprintf("00000000%d", val)
	case val >= 10 && val < 100:
		numStr = fmt.Sprintf("0000000%d", val)
	case val >= 100 && val < 1000:
		numStr = fmt.Sprintf("000000%d", val)
	case val >= 1000 && val < 10000:
		numStr = fmt.Sprintf("00000%d", val)
	case val >= 10000 && val < 100000:
		numStr = fmt.Sprintf("000%d", val)
	case val >= 100000 && val < 1000000:
		numStr = fmt.Sprintf("00%d", val)
	case
		val >= 1000000 && val < 10000000:
		numStr = fmt.Sprintf("0%d", val)
	case
		val >= 10000000 && val < 100000000:
		numStr = fmt.Sprintf("%d", val)
	}
	now := time.Now()
	// var primaryKey = fmt.Sprintf("%02d%02d%02d%02d%02d%02d%s", now.Year(),now.Month(), now.Day(),now.Hour(),now.Minute(),now.Second(),numStr)
	return fmt.Sprintf("%02d%02d%02d%02d%s", now.Year(), now.Month(), now.Day(), now.Hour(), numStr)
}

/*
GetRedisLoginUser
@description GetRedisConfig 获取用户信息
@Param userId 用户id
@Return LoginUser 获取到的用户登录信息
@Author yangDaqiong
@Date 2022-02-17 0:40
*/
func (g *get) GetRedisLoginUser(userId string) (LoginUser, error) {
	var loginUser LoginUser
	redisUserKey := config.Global.ServerConfig().LoginPrefix + userId
	println(redisUserKey, "redisUserKeyredisUserKey")
	println(redisUserKey, "redisUserKeyredisUserKey")
	println(redisUserKey, "redisUserKeyredisUserKey")
	println(redisUserKey, "redisUserKeyredisUserKey")
	println(redisUserKey, "redisUserKeyredisUserKey")
	println(redisUserKey, "redisUserKeyredisUserKey")
	println(redisUserKey, "redisUserKeyredisUserKey")
	println(redisUserKey, "redisUserKeyredisUserKey")
	println(redisUserKey, "redisUserKeyredisUserKey")
	if loginUserJson, err := GET.GetJson(redisUserKey); err != nil {
		return loginUser, errors.New("用户信息获取失败！")
	} else {
		if jsonErr := json.Unmarshal(loginUserJson, &loginUser); jsonErr != nil {
			return loginUser, errors.New("用户信息json转换失败")
		}
		return loginUser, nil
	}
}

/*
getKeyAndValuesMap 生成[string]字符串映射结构，该结构将ID与存储在Redis中的令牌值相关联
@Param keys 存储键
@Param prefix 前缀
@Return map[string]string 返回map
@Author yangDaqiong
@Date 2022-01-03
*/
func getKeyAndValuesMap(keys []string, prefix string) (map[string]string, error) {
	values := make(map[string]string)
	for _, key := range keys {
		value, err := GetRedisInstance().Get(ctx, key).Result()
		if err != nil {
			return nil, merry.Appendf(err, "error retrieving value for key %s", key)
		}
		// 从密钥中去掉前缀，以便将密钥保存到用户ID
		strippedKey := strings.Split(key, prefix)
		values[strippedKey[1]] = value
	}
	return values, nil
}

/*
getKeys 将使用密钥共享的特定前缀并返回所有密钥的列表
@Param prefix 前缀
@Return []string 返回[]string
@Author yangDaqiong
@Date 2022-01-03
*/
func getKeys(prefix string) ([]string, error) {
	var allKeys []string
	var cursor uint64
	count := int64(10) // count specifies how many keys should be returned in every Scan call
	for {
		var keys []string
		var err error
		keys, cursor, err = GetRedisInstance().Scan(ctx, cursor, prefix, count).Result()
		if err != nil {
			return nil, merry.Appendf(err, "error retrieving '%s' keys", prefix)
		}
		allKeys = append(allKeys, keys...)
		if cursor == 0 {
			break
		}
	}
	return allKeys, nil
}
