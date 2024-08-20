package service

import (
	"context"
	"core/app"
	"core/config"
	"core/logger"
	"core/pager"
	"core/redis"
	"core/resp"
	"core/utils"
	"core/validate"
	"protobuf/build/common"
	"protobuf/build/global"
	"protobuf/build/order"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/pkg/util"
)

var PocketMoney = &pocketMoney{}

type pocketMoney struct{}

// GetBalance 获取余额
func (pi *pocketMoney) GetBalance(ctx *gin.Context) any {
	if result, getResultError := common.NewUserDetailServiceClient(app.Nacos.CommonGrpcProvider()).FindBalanceByUserId(context.Background(), &common.UserDetailIds{
		UserId: []string{ctx.Param("requestUserId")},
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Count > 0 {
			return resp.Success.WithData(result.UserBalance.Balance)
		} else {
			return resp.Fail.WithMsg("用户余额查询失败")
		}
	}
}

// RechargeRecord 获取充值记录
func (pi *pocketMoney) FindPageList(ctx *gin.Context) any {
	request := order.PocketMoneyPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	request.Auth = &global.Auth{
		UserId: ctx.Param("requestUserId"),
	}
	if result, getResultError := order.NewPocketMoneyServiceClient(app.Nacos.OrderGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}

// Create 充值 目前只支持微信充值
func (o *pocketMoney) Create(ctx *gin.Context) any {
	amountStr := ctx.Param("amount")
	if bool, _ := regexp.MatchString("^(([1-9]\\d*)(\\.\\d+)?)$|^0\\.\\d*[1-9]$", amountStr); !bool {
		resp.ValidateErr.WithMsg("请正确传入金额")
	}
	userId := ctx.Param("requestUserId")
	// 商户订单号
	orderNo := "CZ_" + util.RandomString(29)
	var pmId = redis.GET.GetPrimaryKey("ORDER_ONLINE_RECHARGE")
	amount, _ := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)
	PocketMoney := order.PocketMoney{
		PmId:        pmId,
		Amount:      amount,
		OrderNo:     orderNo,
		TradeNo:     "",
		PayStatus:   1,
		PaymentName: "余额充值",
		CreateBy:    userId,
		PayType:     1,
	}
	if result, getResultError := order.NewPocketMoneyServiceClient(app.Nacos.OrderGrpcProvider()).Create(context.Background(), &PocketMoney); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Count > 0 {
			maps := make(map[string]string, 0)
			maps["orderNo"] = orderNo
			if ctx.Param("regType") == "3" {
				maps["payMark"] = utils.ConfoundBase64EncodeToString(config.Global.ServerConfig().WechatConfId)
			}
			if ctx.Param("regType") == "4" {
				maps["payMark"] = utils.ConfoundBase64EncodeToString(config.Global.ServerConfig().AppletConfId)
			}
			return resp.Back.WithDataAndMsg(maps, "下单成功！")
		} else {
			return resp.Fail.WithMsg("订单在数据创建失败！")
		}
	}
}

// 查询用户充值记录详情
func (o *pocketMoney) FindById(ctx *gin.Context) any {
	if findRechargeOrder, getResultError := order.NewPocketMoneyServiceClient(app.Nacos.OrderGrpcProvider()).FindById(context.Background(), &order.PocketMoneyIds{
		PmId: []string{ctx.Param("id")},
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Success.WithData(findRechargeOrder)
	}
}
