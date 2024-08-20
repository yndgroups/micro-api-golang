package service

import (
	"context"
	"core/app"
	"core/config"
	"core/enum"
	"core/logger"
	"core/redis"
	"core/resp"
	"core/utils"
	"core/validate"
	"errors"
	"fmt"
	"math/rand"
	"protobuf/build/common"
	"protobuf/build/order"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var Pay = &pay{}

type pay struct{}

// UpdateUserPayPwd 设置用户支付密码
func (pi *pay) UpdateUserPayPwd(ctx *gin.Context) any {
	payPwdParam := order.SetPayPwdParam{}
	if err := ctx.ShouldBindJSON(&payPwdParam); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	logger.SugarLogger.Infof("%+v", &payPwdParam)
	if validateMsg, err := validate.Validated(&payPwdParam); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	// 校验验证码
	if vCode, err := redis.GET.GetString(enum.SmsPay + payPwdParam.Phone); err != nil {
		logger.SugarLogger.Infof("redis获取验证码失败，key:%v,原因：%v", enum.SmsPay+payPwdParam.Phone, err.Error())
		return resp.Fail.WithMsg("验证码获取异常")
	} else {
		if !strings.Contains(vCode, payPwdParam.VerifyCode) {
			return resp.Fail.WithMsg("验证码不正确")
		}
	}
	pwd, err := pi.getPwdString(payPwdParam.PayPwd)
	if err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	salt := fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
	sha512 := utils.HmacSha512(*pwd, salt)
	user := common.UserDetail{
		UserId: ctx.Param("requestUserId"),
		Salt:   salt,
		PayPwd: sha512,
		Phone:  payPwdParam.Phone,
	}
	if update, getResultError := common.NewUserDetailServiceClient(app.Nacos.OrderGrpcProvider()).Update(context.Background(), &user); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return errors.New("登录失败！")
	} else {
		if update.Count > 0 {
			redis.DELETE.DelByKey(enum.SmsPay + payPwdParam.Phone)
			return resp.Success.WithData("修改设置成功")
		} else {
			return resp.Success.WithData("修改设置失败")
		}
	}
}

// getPwdString 获取base64字符串
// "XH1ADGeF2tW137T3H348af57hDB6tG"
func (pi *pay) getPwdString(pwd string) (*string, error) {
	var s string
	var ps = []int{2, 8, 15, 18, 22, 27}
	if len(pwd) >= 30 {
		for _, v := range ps {
			s += pwd[v : v+1]
		}
		return &s, nil
	} else {
		return nil, errors.New("传入的payPwd数据格式不正确")
	}
}
