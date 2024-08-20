package middleware

import (
	"core/config"
	"core/jwt"
	"core/logger"
	"core/redis"
	"fmt"
	"github.com/spf13/cast"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//	权限验证中间件
//
// @author yangDaQiong
// @date  2021-10-24
func Authorization(powerSign string) gin.HandlerFunc {
	return func(cxt *gin.Context) {
		url := cxt.Request.URL.String()
		fmt.Printf("Authorization the request URL %s cost %v\n", url, time.Since(time.Now()))
		if strings.Contains(url, "/auth/createAnonymityToken") || strings.Contains(url, "swagger") || strings.Contains(url, "/wechat/notify") {
			cxt.Next()
			return
		}
		// 访问令牌
		accessToken := cxt.GetHeader("accessToken")
		if accessToken == "" {
			cxt.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "data": "未获取到登录授权信息!", "msg": "请求失败！"})
			cxt.Abort()
			return
		}
		// 验证带来的token的合法性
		tokenInfo, err := jwt.ParseTokenToLoginInfo(accessToken)
		if err != nil {
			cxt.JSON(http.StatusUnauthorized, gin.H{"code": 0, "data": err.Error(), "msg": err.Error()})
			cxt.Abort()
			return
		} else {
			// 如果存在则添加上下文用户id
			if tokenInfo.AppId != "" {
				cxt.AddParam("appId", tokenInfo.AppId)
			}
			if tokenInfo.UserId != "" {
				cxt.AddParam("requestUserId", tokenInfo.UserId)
			}
			if tokenInfo.SubId != "" {
				cxt.AddParam("subId", tokenInfo.SubId)
			}
			if tokenInfo.LoginRedisKey != "" {
				cxt.AddParam("loginRedisKey", tokenInfo.LoginRedisKey)
			}
			if tokenInfo.PowerKey != "" {
				cxt.AddParam("powerKey", tokenInfo.PowerKey)
			}
			if tokenInfo.RegType >= 0 {
				cxt.AddParam("regType", cast.ToString(tokenInfo.RegType))
			}
		}
		// 非匿名接口
		if powerSign != "anonymous" {
			// 如果是匿名token,该请求放行，但是可能接口权限会不通过进行拦截校验
			println("非匿名接口权限值 =>>>>>>" + powerSign)
			// 查看redis是否存在登录信息
			// println(strings.Replace(accessToken, config.Global.ServerConfig().TokenPrefix, "", 1))
			redisToken, _ := redis.GET.GetString(tokenInfo.LoginRedisKey)
			if strings.Contains(redisToken, strings.Replace(accessToken, config.Global.ServerConfig().TokenPrefix, "", 1)) {
				//cxt.AddParam("confId", tokenInfo.ConfId) // 通过app去查询
				// 查看权限是否已经变更
				// result, err := redis.HGet("isResetAuth", tokenInfo.AppId+"_"+tokenInfo.UserId)
				result, err := redis.GET.HashGet("power_change", tokenInfo.AppId+"_"+tokenInfo.UserId)
				if err == nil {
					if strings.Contains(result, "1") {
						fmt.Println(cxt.FullPath())
						if (strings.Contains(cxt.FullPath(), "/api/common/menu/findTreeList") && strings.Contains(cxt.Query("reload"), "1")) || (strings.Contains(cxt.FullPath(), "/api/common/auth/refreshPermissions") && strings.Contains(cxt.Query("reload"), "1")) {
							cxt.Next()
							return
						} else {
							cxt.JSON(http.StatusOK, gin.H{"code": 40102, "data": "", "msg": "您的权限已发生变更,3秒之后将刷新页面,给您带来不便,敬请谅解！"})
							cxt.Abort()
							return
						}
					}
				} else {
					// 登录信息存在继续请求处理,并且做缓存更新 如果需要做只能同时一个在线，做token比较
					var expire time.Duration = config.Global.RedisConfig().Expire
					if tokenInfo.RegType == 3 || tokenInfo.RegType == 4 {
						expire = config.Global.RedisConfig().WeChatExpire
					}
					// 更新用户登陆存储时间
					if err := redis.UPDATE.UpdateExpire(tokenInfo.LoginRedisKey, expire); err != nil {
						cxt.Abort()
						cxt.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "data": err.Error(), "msg": "缓存时间更新处理异常！"})
						logger.SugarLogger.Warnf("auth -> 缓存时间更新处理异常, userId -> %v", tokenInfo.UserId)
						return
					}
					// 更新权限值存储时间
					if err := redis.UPDATE.UpdateExpire(tokenInfo.PowerKey, expire); err != nil {
						cxt.Abort()
						cxt.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "data": err.Error(), "msg": "权限缓存时间更新处理异常！"})
						return
					}
					// 接口权限值判断
					powerSignAll, _ := redis.GET.GetString(tokenInfo.PowerKey)
					if strings.Contains(powerSignAll, powerSign) {
						cxt.Next()
						return
					} else {
						// 验证不通过，不再调用后续的函数处理
						cxt.JSON(http.StatusOK, gin.H{"msg": "无访问权限,请退出登陆之后再进行访问", "code": 40103, "data": nil})
						cxt.Abort()
						return
					}
				}
				cxt.Next()
				return
			} else {
				cxt.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "data": "您还未登录！", "msg": "您还未登录！"})
				cxt.Abort()
				return
			}
			cxt.Next()
			return
		}
		// 匿名接口直接放行
		if strings.Contains(tokenInfo.UserId, "anonymous") || powerSign == "anonymous" {
			println("非匿名接口权限值 =>>>>>>" + powerSign)
			cxt.Next()
			return
		}
	}
}
