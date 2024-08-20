package service

import (
	"core/config"
	"core/enum"
	"core/redis"
	"core/resp"
	"fmt"
	"math/rand"
	"time"

	"github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/gin-gonic/gin"
)

var Sms = &sms{}

type sms struct{}

// SmsSend 发送验证码
func (s *sms) SmsSend(ctx *gin.Context) any {
	phone := ctx.Param("phone")
	getType := ctx.Param("type") // 1:登录验证 2:设置用户支付密码 3:修改密码认证 4：...
	var prefix string
	switch getType {
	case "1":
		prefix = enum.SmsLogin
		getRedisPhone, _ := redis.GET.GetString(prefix + phone)
		if getRedisPhone != "" {
			return resp.Fail.WithMsg("验证码还在有效期内，请勿重复发送")
		}
	case "2":
		prefix = enum.SmsPay
		getRedisPhone, _ := redis.GET.GetString(prefix + phone)
		if getRedisPhone != "" {
			return resp.Fail.WithMsg("验证码还在有效期内，请勿重复发送")
		}
	case "3":
		prefix = enum.SmsPwd
		getRedisPhone, _ := redis.GET.GetString(prefix + phone)
		if getRedisPhone != "" {
			return resp.Fail.WithMsg("验证码还在有效期内，请勿重复发送")
		}
	}
	codeNumber := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	println("TemplateCode" + *config.Global.SmsConfig().Templates[prefix].TemplateCode)
	println("SignName" + *config.Global.SmsConfig().Templates[prefix].SignName)
	tpl := "{\"code\":" + codeNumber + "}"
	return s.sendMessage(SmsInfo{
		StoreKey: prefix + phone,
		Code:     codeNumber,
		Params: dysmsapi20170525.SendSmsRequest{
			PhoneNumbers:  &phone,
			TemplateCode:  config.Global.SmsConfig().Templates[prefix].TemplateCode,
			SignName:      config.Global.SmsConfig().Templates[prefix].SignName,
			TemplateParam: &tpl,
		},
	})
}

type SmsInfo struct {
	// 存储key
	StoreKey string `json:"phone"`
	// 验证码
	Code string `json:"code"`
	// 发送参数
	Params dysmsapi20170525.SendSmsRequest `json:"params"`
}

// sendMessage 短信发送公共方法
func (s *sms) sendMessage(smsInfo SmsInfo) any {
	cfg := &client.Config{AccessKeyId: config.Global.SmsConfig().AccessKeyId, AccessKeySecret: config.Global.SmsConfig().AccessKeySecret, Endpoint: config.Global.SmsConfig().Host}
	newClient, err := dysmsapi20170525.NewClient(cfg)
	if err != nil {
		return resp.Fail.WithMsg(fmt.Sprintf("短信客户端初始化失败！失败原因：%v", err.Error()))
	}
	sms, err := newClient.SendSms(&smsInfo.Params)
	if err != nil {
		return resp.Fail.WithMsg(fmt.Sprintf("短信发送失败，失败失败原因：%v", err.Error()))
	} else {
		if *sms.Body.Code == "OK" {
			redis.SET.Set(smsInfo.StoreKey, smsInfo.Code, config.Global.SmsConfig().Expire)
			return resp.Success.WithMsg("短信发送成功")
		} else {
			return resp.Success.WithMsg(fmt.Sprintf("短信发送失败,原因：%v", *sms.Body.Message))
		}
	}
}
