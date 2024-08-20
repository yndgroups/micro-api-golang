package service

import (
	"context"
	"core/app"
	"core/config"
	"core/model"
	"errors"

	"core/ksii"
	"core/logger"
	"core/pager"
	"core/redis"
	"core/resp"
	"core/utils"
	"core/validate"
	"fmt"
	"protobuf/build/common"
	_ "protobuf/build/global"
	"protobuf/build/order"
	"protobuf/build/shopping_client"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

// type createOrderParams[T any] struct {
// 	// 1微信支付 2：支付宝支付
// 	TradeType uint32 `json:"tradeType"`
// 	// 1直接支付 2：余额白支付
// 	PayType uint32 `json:"payType"`
// 	// 请求商品参数列表
// 	ProductParamsList []T `json:"productList"`
// }

var Order = &orders{}

type orders struct{}

// Create 统一下单接口
func (o *orders) Create(ctx *gin.Context) any {
	createOrderParams := shopping_client.CreateOrderParams{}
	if err := ctx.ShouldBindJSON(&createOrderParams); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if len(createOrderParams.ProductList) == 0 {
		return resp.Fail.WithMsg("商品参数不能为空")
	}
	for _, v := range createOrderParams.ProductList {
		if v.Quantity < 1 {
			return resp.Fail.WithMsg("购买数量不能小于1")
		}
	}
	// 获取商品列表信息 product
	productRresult, getResultError := shopping_client.NewProductServiceClient(app.Nacos.ShoppingClientProvider()).FindListByProductIds(context.Background(), &createOrderParams)
	if getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	}
	priceCountList := make([]*order.ProductPriceCount, 0)
	// 遍历出对应的商品价格和数量
	pLen := len(productRresult.ExpandList)
	if pLen == 0 {
		return resp.Fail.WithData("商品信息查询失败")
	}
	productList := productRresult.ExpandList
	var productTotalQuantity int64
	// 遍历出商品相关数据
	for i := 0; i < pLen; i++ {
		for _, c := range createOrderParams.ProductList {
			// 判断购买数量
			if productList[i].ProductId == c.ProductId {
				if productList[i].StockWarning < 100 {
					app.SendText(fmt.Sprintf("商品【%v】库存预警，请抓紧时间补货", productList[i].ProductName))
				}
				if productList[i].Stock < c.Quantity {
					// 核对库存
					return resp.Fail.WithMsg(fmt.Sprintf("商品 %v 库存不足!最多只能购买%v%v", productList[i].ProductName, productList[i].Stock, productList[i].UnitName))
				} else {
					priceCountList = append(priceCountList, &order.ProductPriceCount{Price: productList[i].Price, Quantity: c.Quantity, ProductId: c.ProductId})
					productTotalQuantity += c.Quantity
				}
			}
		}
	}
	if len(priceCountList) == 0 {
		return resp.Fail.WithMsg("价格参数异常")
	}
	// 计算商品价格
	if totalPrice, err := o.divTotalPrice(priceCountList); err != nil {
		return resp.Fail.WithMsg(err.Error())
	} else {
		// 保存订单数据到订单表和商品详情数据
		orderNo := "SP_" + util.RandomString(29) // 商户订单号
		userId := ctx.Param("requestUserId")
		// 注意这里的判断是小程序还是公众号
		regType := ctx.Param("regType")
		var confId string
		if regType == "3" {
			confId = config.Global.ServerConfig().WechatConfId
		}
		if regType == "4" {
			confId = config.Global.ServerConfig().AppletConfId
		}
		orderId := redis.GET.GetPrimaryKey("SYS_ORDER_REQUEST_COUNT")
		orderDetail := order.Order{
			OrderId:       orderId,                     // 订单id
			ParentId:      "1",                         // 订单父id
			ConfId:        confId,                      // 配置id
			OrderNo:       orderNo,                     // 订单号
			TradeType:     createOrderParams.TradeType, // 1微信支付 2：支付宝支付
			TradeNo:       "",                          // 下单号 创建订单时候是没有订单号的 payNo payNo, err := createOrder(ctx, createOrderInfo
			TotalPrice:    cast.ToFloat64(totalPrice),  // 总金额
			PayPrice:      cast.ToFloat64(totalPrice),  // 支付金额
			TotalQuantity: productTotalQuantity,        // 商品总数
			PayStatus:     1,                           // 支付状态
			PayType:       createOrderParams.PayType,   // 1直接支付 2：余额白支付
			Status:        1,                           // 订单状态(1:待付款[默认] 2:待发货 3:已发货 4:待收货 5:已收货 6:待评价 7:订单关闭)
			CreateBy:      userId,                      // 订单创建者
		}
		// 查询用户的收货地址信息
		if addressResult, err := common.NewUserAddressServiceClient(app.Nacos.CommonGrpcProvider()).FindById(context.Background(), &common.UserAddressIds{
			AddressIds: []string{createOrderParams.AddressId},
		}); err != nil {
			logger.SugarLogger.Errorf("服务生产者:【%v】发生错误啊, 错误信息: %v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
			return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
		} else {
			if addressResult.Msg != "" {
				return resp.Fail.WithMsg(addressResult.Msg)
			}
			orderDetail.RealName = addressResult.Detail.RealName
			orderDetail.UserAddress = addressResult.Detail.Detail
			orderDetail.UserPhone = addressResult.Detail.Phone
		}
		orderDetails := make([]*order.OrderDetail, 0)
		for _, v := range productList {
			orderDetails = append(orderDetails, &order.OrderDetail{
				DetailId:     redis.GET.GetPrimaryKey("SYS_ORDER_REQUEST_COUNT"),
				OrderId:      orderId,
				ProductId:    v.ProductId,
				CategoryId:   v.CategoryId,
				CategoryName: "待补充",
				ProductName:  v.ProductName,
				Introduction: v.Introduction,
				Price:        v.Price,
				CostPrice:    v.CostPrice,
				VipPrice:     v.VipPrice,
				MarketPrice:  v.MarketPrice,
				Image:        v.Image,
				UnitName:     v.UnitName,
				TempId:       v.TempId,
				SpecType:     v.SpecType,
				IsVirtual:    v.IsVirtual,
				//MerId: v.MerId,
				//MerUse: v.MerUse,
				//BarCode: v.BarCode,
				//IsBenefit: v.IsBenefit,
				//IsBest: v.IsBest,
				//IsSub: v.IsSub,
				//CodePath: v.CodePath,
				//Logistics: v.Logistics,
				//Freight: v.Freight,
			})
		}
		// 保存订单
		if result, getResultError := order.NewOrderServiceClient(app.Nacos.OrderGrpcProvider()).Save(context.Background(), &order.SaveOrder{
			Order:        &orderDetail,
			OrderDetails: orderDetails,
		}); getResultError != nil {
			logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
			return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
		} else {
			if result.Msg != "" {
				return resp.Fail.WithMsg(result.Msg)
			}
			maps := make(map[string]string)
			maps["orderNo"] = orderNo
			if regType == "3" {
				maps["payMark"] = utils.ConfoundBase64EncodeToString(config.Global.ServerConfig().WechatConfId)
			}
			if regType == "4" {
				maps["payMark"] = utils.ConfoundBase64EncodeToString(config.Global.ServerConfig().AppletConfId)
			}
			// 下单后需要把商品从购物车中移除
			return resp.Back.WithDataAndMsg(maps, "下单成功")
		}
	}
}

// Pay 订单支付
func (o *orders) Pay(ctx *gin.Context) any {
	payOrderParam := order.PayOrderParam{}
	if err := ctx.ShouldBindJSON(&payOrderParam); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if validateMsg, err := validate.Validated(&payOrderParam); err != nil {
		return resp.ErrValidate.WithData(validateMsg)
	}
	var orderDetailVo = &order.OrderDetailVo{}
	var orderType int
	if strings.Contains(payOrderParam.OrderId, "SP_") { // 支付订单
		if findOrder, getResultError := order.NewOrderServiceClient(app.Nacos.OrderGrpcProvider()).FindById(context.Background(), &order.OrderIds{
			OrderNo: payOrderParam.OrderId,
		}); getResultError != nil {
			logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
			return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
		} else {
			if findOrder.Count > 0 {
				// 1：待支付 状态添加时记得去修改枚举 OrderPayStatus 支付状态(1:待支付 2:已支付 3:支付超时关闭 4:支付失败)
				if findOrder.Detail.PayStatus != 1 {
					return resp.Fail.WithMsg(utils.OrderPayStatus(findOrder.Detail.PayStatus))
				}
				orderDetailVo = findOrder.Detail
				orderType = 1
			} else {
				return resp.Fail.WithMsg("该订单不存在")
			}
		}
	} else if strings.Contains(payOrderParam.OrderId, "CZ_") { // 充值订单
		if findRechargeOrder, getResultError := order.NewPocketMoneyServiceClient(app.Nacos.OrderGrpcProvider()).FindById(context.Background(), &order.PocketMoneyIds{
			OrderNo: payOrderParam.OrderId,
		}); getResultError != nil {
			logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
			return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
		} else {
			if findRechargeOrder.Count > 0 {
				// 1：待支付 状态添加时记得去修改枚举 OrderPayStatus 支付状态(1:待支付 2:已支付 3:支付超时关闭 4:支付失败)
				if findRechargeOrder.Detail.PayStatus != 1 {
					return resp.Fail.WithMsg(utils.OrderPayStatus(findRechargeOrder.Detail.PayStatus))
				}
				orderDetailVo.CreateBy = findRechargeOrder.Detail.CreateBy
				orderDetailVo.PayPrice = findRechargeOrder.Detail.Amount
				orderDetailVo.OrderNo = findRechargeOrder.Detail.OrderNo
				orderType = 2
			} else {
				return resp.Fail.WithMsg("该订单不存在")
			}
		}
	} else {
		return resp.Fail.WithMsg("未知订单")
	}
	// trade_type 交易类型(1:余额 2:微信  3:支付宝) pay_type 支付类型(1:微信公众号 2:微信小程序 3:支付宝支付)
	switch payOrderParam.TradeType {
	case 1: // 余额支付
		return o.balancePayment(ctx, &payOrderParam, orderDetailVo)
	case 2: // 微信
		orderDetailVo.TradeType = payOrderParam.TradeType
		orderDetailVo.PayType = payOrderParam.PayType
		orderDetailVo.ConfId = utils.ConfoundBase64DecodeString(payOrderParam.PayMark)
		if chatOrder, getResultError := Wechat.createOrder(ctx, orderDetailVo, &orderType); getResultError != nil {
			return resp.Fail.WithMsg(getResultError.Error())
		} else {
			return resp.Success.WithData(chatOrder)
		}
	case 3: // 支付支付
		return resp.Success.WithData("支付支付，正在开发中....")
	default:
		return resp.Fail.WithMsg(fmt.Sprintf("未知的交易类型%v", payOrderParam.TradeType))
	}
}

// balancePayment 余额支付
func (o *orders) balancePayment(ctx *gin.Context, payParam *order.PayOrderParam, orderDetailVo *order.OrderDetailVo) any {
	// 校验支付密码
	userId := ctx.Param("requestUserId")
	if findUser, getResultError := common.NewUserDetailServiceClient(app.Nacos.OrderGrpcProvider()).FindById(context.Background(), &common.UserDetailIds{UserId: []string{userId}}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if findUser.Count > 0 {
			if findUser.Detail.PayPwd == "" {
				return resp.Fail.WithMsg("您还未设置支付密码")
			}
			pwd, err := Pay.getPwdString(payParam.PayPwd)
			if err != nil {
				return resp.Fail.WithMsg(err.Error())
			}
			getPwd := utils.HmacSha512(*pwd, findUser.Detail.Salt)
			if strings.Contains(getPwd, findUser.Detail.PayPwd) {
				// 支付密码正确 执行余额扣除
				if findOrder, getResultError := order.NewOrderServiceClient(app.Nacos.OrderGrpcProvider()).BalancePayment(context.Background(), &order.OrderIds{
					OrderNo: orderDetailVo.OrderNo,
					UserId:  userId,
				}); getResultError != nil {
					logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
					return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
				} else {
					if findOrder.Count > 0 { // 更新数据扣款 支付成功
						return resp.Success.WithMsg("支付成功")
					} else {
						return resp.Fail.WithMsg(findOrder.Msg)
					}
				}
			} else {
				return resp.Fail.WithData("支付密码不正确")
			}
		} else {
			logger.SugarLogger.Errorf("支付失败：错误信息：未查询到userId=【%v】的基本信息", userId)
			return resp.Fail.WithMsg("支付失败")
		}
	}
}

// Delete 删除订单
func (o *orders) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := order.NewOrderServiceClient(app.Nacos.OrderGrpcProvider()).Delete(context.Background(), &order.OrderIds{
		OrderId: ids,
		UserId:  ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Delete(result.Count)
	}
}

// FindPageList 查询详情
func (o *orders) FindPageList(ctx *gin.Context) any {
	request := order.OrderPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	if result, getResultError := order.NewOrderServiceClient(app.Nacos.OrderGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}

// FindById 查询详情
func (o *orders) FindById(ctx *gin.Context) any {
	if result, getResultError := order.NewOrderServiceClient(app.Nacos.OrderGrpcProvider()).FindById(context.Background(), &order.OrderIds{
		OrderId: []string{ctx.Param("id")},
		UserId:  ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, result.Detail)
	}
}

// FindByParams 根据订单参数查询订单信息
func (o *orders) FindByParams(ctx *gin.Context) any {
	orderParams := order.OrderParams{}
	if err := ctx.ShouldBindJSON(&orderParams); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if msgResult, getResultError := order.NewOrderServiceClient(app.Nacos.OrderGrpcProvider()).FindByParams(context.Background(), &orderParams); getResultError != nil {
		// app.SendText(fmt.Sprintf("查询订单关系关联信息异常,请联系开发人员进行系统排查，异常信息：%v", getResultError.Error()))
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if msgResult.Msg != "" {
			return resp.Fail.WithMsg(msgResult.Msg)
		}
		return resp.Success.WithData(msgResult.OrderMsgList)
	}
}

// OrderStatistics 订单统计
func (o *orders) OrderStatistics(ctx *gin.Context) any {
	if result, getResultError := order.NewOrderServiceClient(app.Nacos.OrderGrpcProvider()).OrderStatistics(context.Background(), &order.OrderBaseInfo{
		UserId: ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.WithData(result.OrderStatistics)
	}
}

// divTotalPrice 计算总价格
func (o *orders) divTotalPrice(prices []*order.ProductPriceCount) (decimal.Decimal, error) {
	if len(prices) == 0 {
		return decimal.NewFromInt(0), errors.New("商品数量不能为0")
	}
	d := decimal.Decimal{}
	for i, v := range prices {
		if i == 0 {
			n1, _ := ksii.ChangeDecimal("0")
			n2, _ := ksii.ChangeDecimal(v.Price)
			num, _ := ksii.ChangeDecimal(v.Quantity)
			d = n1.Add(n2.Mul(num))
		} else {
			n2, _ := ksii.ChangeDecimal(v.Price)
			num, _ := ksii.ChangeDecimal(v.Quantity)
			d = d.Add(n2.Mul(num))
		}
	}
	return d, nil
}
