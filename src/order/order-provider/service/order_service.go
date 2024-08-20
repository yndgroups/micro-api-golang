package service

import (
	"context"
	"core/config"
	"core/db"
	"core/exception"
	"core/logger"
	"fmt"
	"protobuf/build/order"
	"protobuf/enum"
	"time"
)

var OrderService = &orderService{}

type orderService struct {
	order.OrderServiceServer
}

type CompareMoney struct {
	Balance  float64 `json:"balance"`
	PayPrice float64 `json:"payPrice"`
}

// BalancePayment 余额支付扣款
func (os *orderService) BalancePayment(ctx context.Context, ids *order.OrderIds) (*order.OrderResponse, error) {
	var compareMoney CompareMoney
	sql := fmt.Sprintf(`SELECT  balance Balance,(SELECT pay_price FROM %s WHERE order_no = '%s') AS PayPrice FROM %s WHERE user_id = '%s'`, enum.Order.TableName, ids.OrderNo, enum.SysUser.TableName, ids.UserId)
	if res := db.DB.Raw(sql).Scan(&compareMoney); res.RowsAffected > 0 {
		logger.SugarLogger.Infof("余额日志记录： 用户%s的余额是:%v,订单金额：%v", ids.UserId, compareMoney.Balance, compareMoney.PayPrice)
		if compareMoney.Balance < compareMoney.PayPrice {
			return &order.OrderResponse{Count: 0, Msg: "余额不足"}, nil
		} else {
			logger.SugarLogger.Infof("余额日志记录： 用户%s的余额是:%v,订单金额：%v", ids.UserId, compareMoney.Balance, compareMoney.PayPrice)
			begin := db.DB.Begin()
			payTime := time.Now().Format("2006-01-02 15:04:05")
			updateOrderSql := fmt.Sprintf(`UPDATE %s SET trade_type = 1, pay_status = 2, status = 2, pay_time='%s' WHERE order_no = '%s'`, enum.Order.TableName, payTime, ids.OrderNo)
			affected := begin.Exec(updateOrderSql).RowsAffected
			if affected > 0 { // 订单状态更新完之后扣除余额
				updateUserBalanceSql := fmt.Sprintf("UPDATE %s SET balance = balance - (SELECT pay_price FROM %s WHERE order_no = '%s' ) WHERE user_id = '%v'", enum.SysUser.TableName, enum.Order.TableName, ids.OrderNo, ids.UserId)
				rowsAffected := db.DB.Exec(updateUserBalanceSql).RowsAffected
				if rowsAffected > 0 {
					begin.Commit()
					return &order.OrderResponse{Count: 1}, nil
				} else {
					begin.Rollback()
					return &order.OrderResponse{Count: 0, Msg: "余额扣除失败"}, nil
				}
			} else {
				begin.Rollback()
				return &order.OrderResponse{Count: 0, Msg: "订单状态更新失败"}, nil
			}
		}
	} else {
		return &order.OrderResponse{Count: 0, Msg: "余额及支付金额查询失败"}, nil
	}
}

// Update 更新订单
func (os *orderService) Update(ctx context.Context, param *order.UpdateOrderParam) (*order.OrderResponse, error) {
	// OrderType 订单类型 1: 充值订单 2：商品订单
	if param.OrderType == 1 {
		// 更新用户余额，因为微信充值是安分计算，所以在此要除以 100
		payTime := time.Now().Format("2006-01-02 15:04:05")
		updateSql := `UPDATE ` + enum.PocketMoney.TableName + ` SET pay_type = 1, pay_status = 2, pay_time ='` + payTime + `' WHERE order_no = '` + param.OutTradeNo + `'`
		// 日志存储充值记录sql
		logger.SugarLogger.Info(updateSql)
		res := db.DB.Exec(updateSql)
		if res.Error != nil {
			return &order.OrderResponse{Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().OrderProvider)}, nil
		}
		var userId string
		find := db.DB.Raw(fmt.Sprintf(`SELECT create_by FROM %s WHERE order_no = '%s'`, enum.PocketMoney.TableName, param.OutTradeNo)).Scan(&userId)
		if find.Error != nil {
			return &order.OrderResponse{Msg: exception.DbMsg(find.Error.Error(), config.Global.ServerConfig().OrderProvider)}, nil
		}
		return &order.OrderResponse{OrderBaseInfo: &order.OrderBaseInfo{UserId: userId, OrderNo: param.OutTradeNo}}, nil
		// 更新用户余额
		// updateUserBalanceSql := "UPDATE " + enum.SysUserDetail.TableName + " SET balance = balance + " + param.PayAmount + " WHERE user_id in (SELECT create_by FROM " + enum.PocketMoney.TableName + " WHERE order_no = '" + param.OutTradeNo + "')"
		// if raw.RowsAffected > 0 {
		// 	// 更新状态成功之后需要给用户加上相应的余额
		// 	logger.SugarLogger.Infof("订单回调更新成功，订单号为：%s", param.OutTradeNo)
		// 	// 更新用户余额
		// 	logger.SugarLogger.Infof("更新用户余额,订单号为：%s,余额加 %s", param.OutTradeNo, param.PayAmount)
		// 	db.DB.Exec(updateUserBalanceSql)
		// } else {
		// 	// 假如更新失败 尝试再调用一次更新
		// 	logger.SugarLogger.Infof("更新失败，尝试再调用一次更新，订单号为：%s", param.OutTradeNo)
		// 	db.DB.Exec(updateSql)
		// 	// 更新用户余额
		// 	logger.SugarLogger.Infof("更新用户余额,订单号为：%s,余额加 %s", param.OutTradeNo, param.PayAmount)
		// 	db.DB.Exec(updateUserBalanceSql)
		// }
	} else if param.OrderType == 2 {
		// 更新用户余额，因为微信充值是安分计算，所以在此要除以 100
		payTime := time.Now().Format("2006-01-02 15:04:05")
		updateOrderSql := `UPDATE ` + enum.Order.TableName + ` SET trade_type = 2,pay_status = 2, status = 2, pay_time='` + payTime + `' WHERE order_no = '` + param.OutTradeNo + `'`
		var payType int
		db.DB.Raw("SELECT pay_type FROM " + enum.Order.TableName + " WHERE order_no = '" + param.OutTradeNo + "'").Scan(&payType)
		// 余额减去下单金额
		updateRowsAffected := int64(0)
		// 更新订单状态
		updateRowsAffected = db.DB.Exec(updateOrderSql).RowsAffected
		if updateRowsAffected > 0 {
			// 更新状态成功之后需要给用户加上相应的余额
			logger.SugarLogger.Infof("订单回调更新成功，订单号为：%v", param.OutTradeNo)
			// 更新用户余额
			logger.SugarLogger.Infof("更新用户余额,订单号为：%v,余额扣除 %s", param.OutTradeNo, param.PayAmount)
			if payType == 2 {
				updateUserBalanceSql := "UPDATE " + enum.SysUserDetail.TableName + " SET balance = balance - " + param.PayAmount + " WHERE user_id in (SELECT create_by FROM " + enum.Order.TableName + " WHERE order_no = '" + param.OutTradeNo + "')"
				db.DB.Exec(updateUserBalanceSql)
			}
		} else {
			// 假如更新失败 尝试再调用一次更新
			logger.SugarLogger.Infof("更新失败，尝试再调用一次更新，订单号为：%s", param.OutTradeNo)
			db.DB.Exec(updateOrderSql)
			// 更新用户余额
			logger.SugarLogger.Infof("更新用户余额,订单号为：%s,余额扣除 %s", param.OutTradeNo, param.PayAmount)
			if payType == 2 {
				updateUserBalanceSql := "UPDATE " + enum.SysUserDetail.TableName + " SET balance = balance - " + param.PayAmount + " WHERE user_id in (SELECT create_by FROM " + enum.Order.TableName + " WHERE order_no = '" + param.OutTradeNo + "')"
				db.DB.Exec(updateUserBalanceSql)
			}
		}
	} else {
		logger.SugarLogger.Errorf("未识别到的订单号：info = %v", param)
	}
	return &order.OrderResponse{Count: 0}, nil
}

// Save 保存订单信息
func (os *orderService) Save(ctx context.Context, orderInfo *order.SaveOrder) (*order.OrderResponse, error) {
	tx := db.DB.Begin()
	var count int64
	// 保存订单信息
	if err := tx.Create(&orderInfo.Order).Error; err != nil {
		tx.Rollback()
		return &order.OrderResponse{Msg: "下单信息保持失败，请重试！"}, nil
	} else {
		count += 1
		if err := tx.Create(&orderInfo.OrderDetails).Error; err != nil {
			tx.Rollback()
			return &order.OrderResponse{Count: 0, Msg: "下单信息保持失败，请重试！"}, nil
		} else {
			tx.Commit()
			return &order.OrderResponse{Count: count}, nil
		}
	}
}

// Delete 删除订单
func (os *orderService) Delete(ctx context.Context, ids *order.OrderIds) (*order.OrderResponse, error) {
	count := db.GrpcBatchDeleteByIds(ids.OrderId, enum.FreightTemplates.TableName, enum.FreightTemplates.PrimaryKey, ids.UserId)
	if count == 0 {
		return &order.OrderResponse{Msg: "删除失败"}, nil
	}
	return &order.OrderResponse{Count: count}, nil
}

// FindPageList 查询订单列表
func (os *orderService) FindPageList(ctx context.Context, request *order.OrderPageAuthQuery) (*order.OrderResponse, error) {
	countSql := `SELECT COUNT(1) FROM ` + enum.Order.TableName + ` A`
	paramSql := ` WHERE A.del_status = 0 `
	if request.Params != nil {
		if request.Params.OrderNo != "" {
			paramSql += fmt.Sprintf(` AND A.order_no = %s`, request.Params.OrderNo)
		}
		if request.Params.PayStatus != 0 {
			paramSql += fmt.Sprintf(` AND A.pay_status = %v`, request.Params.PayStatus)
		}
		if request.Params.UserPhone != "" {
			paramSql += fmt.Sprintf(` AND A.user_phone = %s`, request.Params.UserPhone)
		}
	}
	var count int64
	var list []*order.OrderListVo
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		listSql := `
		SELECT 
			A.order_id OrderId,
			A.parent_id ParentId,
			A.order_no OrderNo,
			A.trade_type TradeType,
			(CASE A.trade_type
					WHEN 1 THEN '微信'
					WHEN 2 THEN '支付宝' END
			) AS TradeTypeName,
			A.trade_no TradeNo,
			A.total_price TotalPrice,
			A.pay_price PayPrice,
			A.total_quantity TotalQuantity,
			A.status Status,
			(CASE A.status
					WHEN 1 THEN '待付款'
					WHEN 2 THEN '待发货'
					WHEN 3 THEN '已发货'
					WHEN 4 THEN '待收货'
					WHEN 5 THEN '已收货'
					WHEN 6 THEN '待评价'
					WHEN 7 THEN '订单关闭' END
			) AS StatusName,
			A.pay_type PayType,
			(CASE A.pay_type
					WHEN 1 THEN '直接支付'
					WHEN 2 THEN '余额支付' END
			) AS PayTypeName,
			A.pay_status PayStatus,
			(CASE A.pay_status
					WHEN 1 THEN '待支付'
					WHEN 2 THEN '已支付'
					WHEN 3 THEN '支付超时关闭'
					WHEN 4 THEN '支付失败' END
			) AS PayStatusName,
			DATE_FORMAT(A.pay_time,'%Y-%m-%d %T') PayTime,
			A.user_phone UserPhone,
			A.user_address UserAddress,
			A.refund_status RefundStatus,
			(CASE A.refund_status
					WHEN 1 THEN '未退款'
					WHEN 2 THEN '申请中'
					WHEN 3 THEN '已退款'
					WHEN 4 THEN '拒绝退款' END
			) AS RefundStatusName,
			A.create_by CreateBy,
			DATE_FORMAT(A.create_time,'%Y-%m-%d %T') CreateTime
		FROM 
			` + enum.Order.TableName + ` A `
		listSql += paramSql
		listSql += ` LIMIT ?,?`
		db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&list)
	}
	return &order.OrderResponse{Count: count, List: list}, nil
}

// FindById 查询订单详情
func (os *orderService) FindById(ctx context.Context, ids *order.OrderIds) (*order.OrderResponse, error) {
	sql := `
		SELECT 
			A.order_id OrderId,
			A.parent_id ParentId,
			A.order_no OrderNo,
			A.trade_type TradeType,
			A.conf_id confId,
			(CASE A.trade_type
					WHEN 1 THEN '微信'
					WHEN 2 THEN '支付宝' END
			) AS TradeTypeName,
			A.trade_no TradeNo,
			A.total_price TotalPrice,
			A.pay_price PayPrice,
			A.total_quantity TotalQuantity,
			A.status Status,
			(CASE A.status
					WHEN 1 THEN '待付款'
					WHEN 2 THEN '待发货'
					WHEN 3 THEN '已发货'
					WHEN 4 THEN '待收货'
					WHEN 5 THEN '已收货'
					WHEN 6 THEN '待评价'
					WHEN 7 THEN '订单关闭' END
			) AS StatusName,
			A.pay_type PayType,
			(CASE A.pay_type
					WHEN 1 THEN '直接支付'
					WHEN 2 THEN '余额支付' END
			) AS PayTypeName,
			A.pay_status PayStatus,
			(CASE A.pay_status
					WHEN 1 THEN '待支付'
					WHEN 2 THEN '已支付'
					WHEN 3 THEN '支付超时关闭'
					WHEN 4 THEN '支付失败' END
			) AS PayStatusName,
			A.pay_time PayTime,
			A.user_phone UserPhone,
			A.user_address UserAddress,
			A.refund_status RefundStatus,
			(CASE A.refund_status
					WHEN 1 THEN '未退款'
					WHEN 2 THEN '申请中'
					WHEN 3 THEN '已退款'
					WHEN 4 THEN '拒绝退款' END
			) AS RefundStatusName,
			A.create_by CreateBy,
			A.create_time CreateTime
		FROM  ` + enum.Order.TableName
	if ids.OrderNo != "" {
		sql += fmt.Sprintf(` A WHERE order_no = '%v'`, ids.OrderNo)
	} else {
		sql += fmt.Sprintf(` A WHERE order_id = '%v'`, ids.OrderId[0])
	}
	detail := order.OrderDetailVo{}
	result := db.DB.Raw(sql).Scan(&detail)
	if ids.OrderNo == "" && result.RowsAffected > 0 {
		var productList []*order.OrderDetail
		db.DB.Find(&productList)
		detail.ProductList = productList
	}
	return &order.OrderResponse{Count: result.RowsAffected, Detail: &detail}, nil
}

// OrderStatistics 订单统计
func (os *orderService) OrderStatistics(ctx context.Context, param *order.OrderBaseInfo) (*order.OrderResponse, error) {
	sql := `
		select count(1) total,
		sum(case when status='1' then 1 else 0 end) daifukuan,
		sum(case when status='2' then 1 else 0 end) daifahuo,
		sum(case when status='3' then 1 else 0 end) yifahuo,
		sum(case when status='4' then 1 else 0 end) daishouhuo,
		sum(case when status='5' then 1 else 0 end) yishouhuo,
		sum(case when status='6' then 1 else 0 end) daipinjia,
		sum(case when status='7' then 1 else 0 end) dingdanguanbi
	FROM 
	 ` + enum.Order.TableName
	//sql += fmt.Sprintf(` A WHERE store_id = '%v'`, param.UserId)
	orderStatistics := order.OrderStatistics{}
	result := db.DB.Raw(sql).Scan(&orderStatistics)
	if result.Error != nil {
		return &order.OrderResponse{Msg: exception.DbMsg(result.Error.Error(), config.Global.ServerConfig().OrderProvider)}, nil
	}
	return &order.OrderResponse{Count: result.RowsAffected, OrderStatistics: &orderStatistics}, nil
}

// FindByParams 查询订单参数
func (os *orderService) FindByParams(ctx context.Context, params *order.OrderParams) (*order.OrderResponse, error) {
	if params != nil && (params.OrderNo == "" && params.TradeNo == "") {
		return &order.OrderResponse{Msg: "核心参数不能为空"}, nil
	}
	sql := `
	SELECT
	    O.real_name,
	    O.user_phone,
	    O.user_address,
	    OD.product_name,
	    OD.price,
	    OD.quantity
	FROM
	    ynd_order_detail OD LEFT JOIN ynd_order O ON OD.order_id = O.order_id
	WHERE `
	if params.OrderNo != "" && params.TradeNo != "" {
		sql += fmt.Sprintf(" O.order_no = '%v' AND trade_no = '%v'", params.OrderNo, params.TradeNo)
	} else if params.OrderNo != "" && params.TradeNo == "" {
		sql += fmt.Sprintf(" O.order_no = '%v'", params.OrderNo)
	} else if params.OrderNo == "" && params.TradeNo != "" {
		sql += fmt.Sprintf(" O.trade_no = '%v'", params.TradeNo)
	} else {
		return &order.OrderResponse{Msg: "参数处理异常"}, nil
	}
	var list []*order.OrderMsg
	res := db.DB.Raw(sql).Scan(&list)
	if res.Error != nil {
		return &order.OrderResponse{Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().OrderProvider)}, nil
	}
	return &order.OrderResponse{OrderMsgList: list, Count: res.RowsAffected}, nil
}
