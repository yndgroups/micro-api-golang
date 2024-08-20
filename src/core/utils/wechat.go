package utils

import (
	"core/redis"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

// GetWechatAccessToken 获取公众号的accessToken 注意和个人的accessToken 区别
func GetWechatAccessToken(configId string) (string, error) {
	conf, confErr := redis.GET.GetRedisConfig(configId)
	if confErr != nil {
		return "获取登录配置信息失败", errors.New("获取登录配置信息失败")
	}
	if accessToken, accessTokenErr := redis.GET.GetString("accessToken_" + conf.AppId); accessTokenErr == nil {
		return accessToken, nil
	}
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + conf.OutsideAppId + "&secret=" + conf.OutsideAppSecret
	if wxRequest, wxRequestErr := HttpRequest.Get(url); wxRequestErr != nil {
		return wxRequestErr.Error(), errors.New("微信调用失败")
	} else {
		if body, bodyErr := wxRequest.Body(); bodyErr != nil {
			return bodyErr.Error(), errors.New("微信调用失败")
		} else {
			resp := AccessToken{}
			if jsonParseErr := json.Unmarshal(body, &resp); jsonParseErr != nil {
				return "解析返回参数JSON转换失败", errors.New("解析返回参数失败！")
			} else {
				if resp.ExpiresIn == 7200 {
					redis.SET.Set("accessToken_"+conf.AppId, resp.AccessToken, 7200)
					return resp.AccessToken, nil
				} else {
					return fmt.Sprintf("%v", resp.ErrCode), errors.New(WeChatAccessTokenErrorMessage(resp.ErrCode))
				}
			}
		}
	}
}

// WechatLoginErrorMessage 微信小程序登录报错解析
func WechatLoginErrorMessage(errorCode int) string {
	m := map[int]string{
		40029: "js_code无效",
		45011: "API 调用太频繁，请稍候再试",
		40226: "高风险等级用户，小程序登录拦截", // https://developers.weixin.qq.com/miniprogram/dev/framework/operation.html
		-1:    "系统繁忙，此时请开发者稍候再试",
	}
	return m[errorCode]
}

// OrderPayStatus 支付状态(1:待支付 2:已支付 3:支付超时关闭 4:支付失败)
func OrderPayStatus(code int64) string {
	switch code {
	case 1:
		return "待支付"
	case 2:
		return "已支付"
	case 3:
		return "支付超时关闭"
	case 4:
		return "支付失败"
	default:
		return fmt.Sprintf("订单支付状态不正确:%v", code)
	}
}

// WeChatAccessTokenErrorMessage 微信公众号获取accessToken报错解析
func WeChatAccessTokenErrorMessage(code int) string {
	//errMap := make(map[int]string, 0)
	//errMap[-1] = "系统繁忙，此时请开发者稍候再试"
	//errMap[0] = "请求成功"
	//errMap[40001] = "AppSecret错误或者 AppSecret 不属于这个公众号，请开发者确认 AppSecret 的正确性"
	//errMap[40002] = "请确保grant_type字段值为client_credential"
	//errMap[40164] = "调用接口的 IP 地址不在白名单中，请在接口 IP 白名单中进行设置。"
	//errMap[89503] = "此 IP 调用需要管理员确认,请联系管理员"
	//errMap[89501] = "此 IP 正在等待管理员确认,请联系管理员"
	//errMap[89506] = "24小时内该 IP 被管理员拒绝调用两次，24小时内不可再使用该 IP 调用"
	//errMap[89507] = "1小时内该 IP 被管理员拒绝调用一次，1小时内不可再使用该 IP 调用"
	switch code {
	case -1:
		return "系统繁忙，此时请开发者稍候再试"
	case 0:
		return "请求成功"
	case 40001:
		return "AppSecret错误或者 AppSecret 不属于这个公众号，请开发者确认 AppSecret 的正确性"
	case 40002:
		return "请确保grant_type字段值为client_credential"
	case 40164:
		return "调用接口的 IP 地址不在白名单中，请在接口 IP 白名单中进行设置。"
	case 89503:
		return "此 IP 调用需要管理员确认,请联系管理员"
	case 89501:
		return "此 IP 正在等待管理员确认,请联系管理员"
	case 89506:
		return "24小时内该 IP 被管理员拒绝调用两次，24小时内不可再使用该 IP 调用"
	case 89507:
		return "1小时内该 IP 被管理员拒绝调用一次，1小时内不可再使用该 IP 调用"
	default:
		return fmt.Sprintf("未知状态:%v", code)
	}
}
