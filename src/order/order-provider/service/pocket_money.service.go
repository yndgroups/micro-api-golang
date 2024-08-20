package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/order"
	"protobuf/enum"
)

var PocketMoneyService = &pocketMoneyService{}

type pocketMoneyService struct {
	order.PocketMoneyServiceServer
}

func (o *pocketMoneyService) Create(ctx context.Context, recharge *order.PocketMoney) (*order.PocketMoneyResponse, error) {
	return &order.PocketMoneyResponse{Count: db.NewCrud(recharge, 1)}, nil
}

func (o *pocketMoneyService) Update(ctx context.Context, recharge *order.PocketMoney) (*order.PocketMoneyResponse, error) {
	// 更新就只能做支付状态更新
	return &order.PocketMoneyResponse{Count: db.NewCrud(order.PocketMoney{PmId: recharge.PmId, PayStatus: 2}, 2)}, nil
}

// FindPageList 充值记录列表
func (o *pocketMoneyService) FindPageList(ctx context.Context, request *order.PocketMoneyPageAuthQuery) (*order.PocketMoneyResponse, error) {
	countSql := `SELECT COUNT(1) FROM ` + enum.PocketMoney.TableName + ` A WHERE del_status = 0 `
	paramSql := ""
	if request.Params != nil {
		if request.Params.PayType != "" {
			paramSql += fmt.Sprintf(" AND A.pay_type = '%v'", request.Params.PayType)
		}
		if request.Auth.UserId != "" {
			paramSql += fmt.Sprintf(" AND A.create_by = '%v' ", request.Auth.UserId)
		}
	}
	var count int64
	db.DB.Raw(countSql + paramSql).Count(&count)
	var list []*order.PocketMoneyListVo
	if count > 0 {
		listSql := `
SELECT
	A.pm_id,
	A.order_no,
	A.trade_no,
	A.amount,
	A.pay_type,
	A.payment_name,
	A.pay_status,
	DATE_FORMAT(A.create_time,'%Y-%m-%d %T') create_time,
	DATE_FORMAT(A.pay_time,'%Y-%m-%d %T') pay_time,
	(CASE A.pay_status
		WHEN 1 THEN '待支付'
		WHEN 2 THEN '已支付' 
		WHEN 3 THEN '支付超时关闭' 
		WHEN 4 THEN '支付失败' END
	) AS sb
FROM ` + enum.PocketMoney.TableName + ` A WHERE A.del_status = 0 ` + paramSql
		listSql += " ORDER BY A.create_time DESC "
		if request.PageIndex > 0 && request.PageSize > 0 {
			listSql += fmt.Sprintf(" LIMIT %v,%v", (request.PageIndex-1)*request.PageSize, request.PageSize)
		}
		db.DB.Raw(listSql).Scan(&list)
	}
	return &order.PocketMoneyResponse{Count: count, List: list}, nil
}

// FindById 查询订单详情
func (o *pocketMoneyService) FindById(ctx context.Context, param *order.PocketMoneyIds) (*order.PocketMoneyResponse, error) {
	var rechargeDetail order.PocketMoney
	if param.OrderNo != "" {
		sql := fmt.Sprintf(`SELECT * FROM %v WHERE order_no = '%v'`, enum.PocketMoney.TableName, param.OrderNo)
		affected := db.DB.Raw(sql).Scan(&rechargeDetail).RowsAffected
		return &order.PocketMoneyResponse{Count: affected, Detail: &rechargeDetail}, nil
	} else {
		affected := db.DB.First(&rechargeDetail, order.PocketMoney{PmId: param.PmId[0]}).RowsAffected
		return &order.PocketMoneyResponse{Count: affected, Detail: &rechargeDetail}, nil
	}
}
