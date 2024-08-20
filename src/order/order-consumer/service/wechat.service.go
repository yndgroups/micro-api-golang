package service

import (
	"context"
	"core/app"
	"core/config"
	"core/ksii"
	"core/logger"
	"core/redis"
	"core/resp"
	"core/utils"
	"errors"
	"fmt"
	"order/order-consumer/configureation"
	"protobuf/build/common"
	"protobuf/build/order"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/kirinlabs/HttpRequest"
	"github.com/spf13/cast"
)

var Wechat = &weChat{}

type weChat struct{}

// TemplateItem 模板消息.
type TemplateItem struct {
	TemplateID      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

// TemplateDataItem 模版内某个 .DATA 的值
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// TemplateMessage 发送的模板消息内容
type TemplateMessage struct {
	// 必须, 接受者OpenID
	ToUser string `json:"touser"`
	// 必须, 模版ID
	TemplateID string `json:"template_id"`
	// 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	URL string `json:"url,omitempty"`
	// 可选, 整个消息的颜色, 可以不设置
	Color string `json:"color,omitempty"`
	// 必须, 模板数据
	Data map[string]*TemplateDataItem `json:"data"`
	// 可选,跳转至小程序地址
	MiniProgram struct {
		// 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
		AppID string `json:"appid"`
		// 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
		PagePath string `json:"pagepath"`
	} `json:"miniprogram"`
}

// MessageTemplate 信息推送模板
type MessageTemplate struct {
	// 接收者openid
	Touser string `json:"touser"`
	// 模板id
	TemplateId string `json:"template_id"`
	// 模板跳转链接（海外帐号没有跳转能力）
	Url string `json:"url"`
	// 否	跳小程序所需数据，不需跳小程序可不用传该数据
	Miniprogram struct {
		Appid    string `json:"appid"`
		Pagepath string `json:"pagepath"`
	}
	// 是	所需跳转到的小程序appid（该小程序 appid 必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	Appid string `json:"appid"`
	// 否	所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
	Pagepath string `json:"pagepath"`
	// 是	模板数据
	Data struct {
		first struct {
			// 标题
			Value string `json:"value"`
			// 颜色 模板内容字体颜色，不填默认为黑色
			Color string `json:"color"`
		}
	} `json:"data"`
	// 否	防重入id。对于同一个openid + client_msg_id, 只发送一条消息,10分钟有效,超过10分钟不保证效果。若无防重入需求，可不填
}

// closeOrder
func (w *weChat) closeOrder(ctx *gin.Context, tradeNo string) (any, error) {
	var si = tradeNo
	if len(tradeNo) > 32 {
		si = string([]byte(tradeNo)[:32])
	}
	redisConfig, getConfErr := redis.GET.GetRedisConfig(ctx.Param("confId"))
	if getConfErr != nil {
		return nil, errors.New("获取配置失败！")
	}
	config, initWeChatErr := configureation.InitWechatConfig(redisConfig)
	if initWeChatErr != nil {
		return "初始化失败", errors.New(initWeChatErr.Error())
	}
	wxRsp, err := config.V3TransactionCloseOrder(ctx, si)
	if err != nil {
		logger.SugarLogger.Errorf("order close success:%s", err)
		return err, errors.New("订单关闭失败！")
	}
	if wxRsp.Code == 0 {
		// 注意要写入记录，更新数据库订单状态
		logger.SugarLogger.Errorf("order close success:%s", wxRsp)
		return "订单关闭成功！", nil
	} else {
		logger.SugarLogger.Errorf("order close fail:%s", wxRsp.Error)
		return err, errors.New(wxRsp.Error)
	}
}

// Refunds 申请退款
func (w *weChat) Refunds(ctx *gin.Context) any {
	tradeNo := ctx.Param("tradeNo")
	confId := ctx.Param("confId")
	userId, _ := ctx.Params.Get("requestUserId")
	user, err := redis.GET.GetRedisLoginUser(userId)
	if err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	conf, err := redis.GET.GetRedisConfig(confId)
	if err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	price := 1000
	bm := make(gopay.BodyMap)
	bm.Set("appid", conf.OutsideAppId).
		Set("mchid", conf.MchId).
		//Set("sub_mchid", "sub_mchid").
		Set("reason", "").                 // 订单描述 refundParam.ReasonrefundParam.ReasonrefundParam.ReasonrefundParam.Reason
		Set("out_trade_no", tradeNo).      // 订单号
		Set("time_expire", expire).        // 结束交易时间
		Set("notify_url", conf.NotifyUrl). // 通知地址
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", price). // 订单总金额，单位为分。
						Set("currency", "CNY") // 货币类型
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) { // 支付者
			bm.Set("openid", user.OpenId) // 用户标识
		})
	redisConfig, getConfErr := redis.GET.GetRedisConfig(ctx.Param("confId"))
	if getConfErr != nil {
		resp.Fail.WithMsg("获取配置失败！")
	}
	initWeChat, initWeChatErr := configureation.InitWechatConfig(redisConfig)
	if initWeChatErr != nil {
		return resp.Fail.WithMsg("微信配置初始化失败" + initWeChatErr.Error())
	}
	initWeChat.V3Refund(ctx, bm)
	return resp.Fail.WithMsg("申请退款")
}

// CloseOrder 关闭订单
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func (w *weChat) CloseOrder(ctx *gin.Context) any {
	if tradeNo, isExist := ctx.Params.Get("tradeNo"); isExist {
		if res, err := w.closeOrder(ctx, tradeNo); err != nil {
			return resp.Fail.WithDataAndMsg(res, err.Error())
		} else {
			return resp.Success.WithData(res)
		}
	} else {
		return resp.Fail.WithMsg("请传入订单号！")
	}
}

// FindOrderByTradeNo 订单查询
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_2.shtml
func (w *weChat) FindOrderByTradeNo(ctx *gin.Context) any {
	redisConfig, getConfErr := redis.GET.GetRedisConfig(ctx.Param("confId"))
	if getConfErr != nil {
		resp.Fail.WithMsg("获取配置失败！")
	}
	initWechat, err := configureation.InitWechatConfig(redisConfig)
	if err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	wxRsp, err := initWechat.V3TransactionQueryOrder(ctx, 2, ctx.Param("tradeNo"))
	if err != nil {
		logger.SugarLogger.Errorf("order close success:%s", err)
		return resp.Fail.WithMsg("订单不存在!")
	}
	if wxRsp.Code == 0 {
		// 注意要写入记录，更新数据库订单状态
		logger.SugarLogger.Errorf("order close success:%s", wxRsp)
		return resp.Success.WithData(wxRsp)
	} else {
		logger.SugarLogger.Errorf("order close fail:%s", wxRsp.Error)
		return resp.Fail.WithDataAndMsg(wxRsp.Error, "订单查询失败！")
	}
}

// Notify 微信支付成功回调
func (w *weChat) Notify(ctx *gin.Context) any {
	configId := ctx.Param("confId")
	notifyReq, notifyReqErr := wechat.V3ParseNotify(ctx.Request)
	if notifyReqErr != nil {
		logger.SugarLogger.Errorf("支付回调通知失败,失败原因：%v", notifyReqErr.Error())
		return resp.Fail.WithMsg("支付回调通知失败")
	}
	redisConfig, getConfErr := redis.GET.GetRedisConfig(configId)
	if getConfErr != nil {
		return resp.Fail.WithMsg("获取配置失败")
	}
	result, err := notifyReq.DecryptCipherText(redisConfig.ApiV3Key)
	if err != nil {
		xlog.Errorf("支付回调信息解析失败 => :%s", notifyReq)
		return resp.Fail.WithMsg("支付回调信息解析失败！")
	}
	logger.SugarLogger.Infof("result = %v", result)
	// 更新充值订单表数据
	if result.TradeState == "SUCCESS" {
		var orderType int64
		// 充值订单,更新余额信息
		if strings.Contains(result.OutTradeNo, "CZ_") {
			orderType = 1
		} else if strings.Contains(result.OutTradeNo, "SP_") { // 商品订单
			orderType = 2
		}
		payerTotal, _ := ksii.Div(strconv.Itoa(result.Amount.PayerTotal), strconv.Itoa(100))
		updateOrderParam := order.UpdateOrderParam{
			OrderType:  orderType,
			OutTradeNo: result.OutTradeNo,
			PayAmount:  cast.ToString(payerTotal),
		}
		if res, getResultError := order.NewOrderServiceClient(app.Nacos.OrderGrpcProvider()).Update(context.Background(), &updateOrderParam); getResultError != nil {
			logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
			return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
		} else {
			// 如果是充值订单更新用户的余额
			if orderType == 1 {
				if _, getResultError := common.NewUserDetailServiceClient(app.Nacos.CommonGrpcProvider()).UpdateBalanceByUserId(context.Background(), &common.UserBalance{
					Balance: cast.ToString(payerTotal),
					UserId:  res.OrderBaseInfo.UserId,
				}); getResultError != nil {
					logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
					return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
				} else {
					return resp.Success.WithMsg("支付成功！")
				}
			} else {
				// 根据订单号查询订单信息进行推送
				if msgResult, getResultError := order.NewOrderServiceClient(app.Nacos.OrderGrpcProvider()).FindByParams(context.Background(), &order.OrderParams{
					OrderNo: result.OutTradeNo,
				}); getResultError != nil {
					app.SendText(fmt.Sprintf("查询订单关系关联信息异常,请联系开发人员进行系统排查，异常信息：%v", getResultError.Error()))
					logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
					return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
				} else {
					if msgResult.Msg != "" {
						return resp.Fail.WithMsg(msgResult.Msg)
					}
					// 推送商品订单给客服进行发货
					app.SendOrderMessage(msgResult.OrderMsgList)
				}
				return resp.Success.WithMsg("支付成功！")
			}
		}
	} else {
		return resp.Fail.WithMsg("支付失败！")
	}
}

// WeChatSendMessageToUser 微信下单成功信息推送
func (w *weChat) SendMessageToUser() any {
	// "2021092612345500004"
	var configId string = "2021092612345500004"
	accessToken, err := utils.GetWechatAccessToken(configId)
	if err != nil {
		return resp.Fail.WithMsg(err.Error())
	} else {
		msgData := make(map[string]*TemplateDataItem)
		msgData["keyword1"] = &TemplateDataItem{
			Value: "SN20220829123456980",
			Color: "#FF0000",
		}
		msgData["keyword2"] = &TemplateDataItem{
			Value: "ITe族",
			Color: "",
		}
		msgData["keyword3"] = &TemplateDataItem{
			Value: "200 元",
			Color: "",
		}
		msgData["remark"] = &TemplateDataItem{
			Value: "1 斤干辣椒，8 罐有辣椒， 五香辣椒面 10 斤", // models.FindConfig("WechatTemplateRemark"),
			Color: "#FF0000",
		}
		msg := &TemplateMessage{
			ToUser:     "o2Tu45zDgz7c_rtcGz9xBSvUp7JA",
			Data:       msgData,
			TemplateID: "zlAWfCUfVnViPNF2Y7Aq4O5YQN2HoUM8gg4joANVW0Q",
			URL:        "url",
		}
		url := `https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=` + accessToken
		_, _ = HttpRequest.JSON().Post(url, msg)
		return resp.Success.WithData("消息发送成功")
	}
}

/*
@description createWeChatOrder 创建订单并获取订单号
@Param ctx ginContext
@Param CreateDetail 订单创建数据
@Param orderType 1:商品订单 2:充值订单
@Author yangDaqiong
@Date 2022-02-17 00:40:00
*/
func (w *weChat) createOrder(ctx *gin.Context, order *order.OrderDetailVo, orderType *int) (any, error) {
	// 获取配置信息
	configDetail, err := redis.GET.GetRedisConfig(order.ConfId)
	if err != nil {
		conf, getResultError := common.NewConfigServiceClient(app.Nacos.CommonGrpcProvider()).FindByConfigExpandParam(context.Background(), &common.ConfigIds{
			ConfIds: []string{order.ConfId},
			AppId:   ctx.Param("appId"),
		})
		if getResultError != nil {
			logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().CommonProvider, getResultError.Error())
			return nil, errors.New("服务生产者调用异常,请联系管理员进行检查！")
		}
		if conf.Msg != "" {
			return nil, errors.New(conf.Msg)
		}
		redis.SET.SetJson(config.Global.ServerConfig().ConfigPrefix+order.ConfId, conf.Detail, config.Global.RedisConfig().Expire)
		configDetail = conf.Detail
	}
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	var description string
	if *orderType == 1 {
		description = "商品订单"
	} else {
		description = "充值订单"
	}
	if ctx.Param("subId") == "" {
		return nil, errors.New("支付关键信息参数获取失败")
	}
	// 支付总金额 金额要乘 100 单位为是按分计算的 安分计算乘100
	payPrice, err := ksii.Mul(order.PayPrice, "100")
	if err != nil {
		return nil, errors.New("金额计算失败")
	}
	total := cast.ToInt64(fmt.Sprintf("%v", payPrice))
	if app.ENV == "dev" || app.ENV == "test" {
		total = 1
	}
	bm := make(gopay.BodyMap)
	bm.Set("appid", configDetail.OutsideAppId).
		Set("mchid", configDetail.MchId).
		// Set("sub_mchid", "sub_mchid").
		Set("description", description).                                               // 订单描述
		Set("out_trade_no", order.OrderNo).                                            // 订单号
		Set("time_expire", expire).                                                    // 结束交易时间
		Set("notify_url", fmt.Sprintf("%v/%v", configDetail.NotifyUrl, order.ConfId)). // 通知地址
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			// 订单总金额，单位为分。
			bm.Set("total", total).Set("currency", "CNY") // 货币类型
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) { // 支付者
			bm.Set("openid", ctx.Param("subId")) // 用户标识
		})
	initWeChat, initWeChatErr := configureation.InitWechatConfig(configDetail)
	if initWeChatErr != nil {
		return "初始化失败", errors.New(initWeChatErr.Error())
	}
	createPrepayIdRsp, err := initWeChat.V3TransactionJsapi(ctx, bm)
	if err != nil {
		logger.SugarLogger.Errorf("用户%s微信下单失败,失败原因: %v", err.Error())
		return "微信下单失败", errors.New(err.Error())
	}
	if createPrepayIdRsp.Code == 0 {
		logger.SugarLogger.Infof("用户%s微信下单成功,订单号为: %v", ctx.Param("requestUserId"), createPrepayIdRsp.Response.PrepayId)
		// jsapi 公众号支付
		if order.PayType == 1 {
			if jsapi, err := initWeChat.PaySignOfJSAPI(configDetail.OutsideAppId, createPrepayIdRsp.Response.PrepayId); err != nil {
				return nil, errors.New(err.Error())
			} else {
				return jsapi, nil
			}
		}
		// 小程序支付
		if order.PayType == 2 {
			if jsapi, err := initWeChat.PaySignOfApplet(configDetail.OutsideAppId, createPrepayIdRsp.Response.PrepayId); err != nil {
				return nil, errors.New(err.Error())
			} else {
				return jsapi, nil
			}
		}
		// app支付
		if order.PayType == 3 {
			if jsapi, err := initWeChat.PaySignOfApp(configDetail.OutsideAppId, createPrepayIdRsp.Response.PrepayId); err != nil {
				return nil, errors.New(err.Error())
			} else {
				return jsapi, nil
			}
		}
		s := fmt.Sprintf("未知的支付类型:%v", order.PayType)
		return "未知的支付类型", errors.New(s)
	} else {
		logger.SugarLogger.Errorf("微信下单失败,原因:%s", createPrepayIdRsp.Error)
		return createPrepayIdRsp.Error, errors.New("微信下单失败！")
	}
}
