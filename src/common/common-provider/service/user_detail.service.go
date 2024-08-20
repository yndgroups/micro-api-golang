package service

import (
	"context"
	"core/config"
	"core/db"
	"core/exception"
	"fmt"
	"protobuf/build/common"
	"protobuf/enum"
)

var UserDetailService = &userDetailService{}

type userDetailService struct {
	common.UserDetailServiceServer
}

// Create 创建用户详细信息
func (ud *userDetailService) Create(context context.Context, userInfo *common.UserDetail) (*common.UserDetailResponse, error) {
	var count int64
	sql := fmt.Sprintf(`SELECT COUNT(1) FROM %s WHERE user_id = '%s'`, enum.SysUserDetail.TableName, userInfo.UserId)
	db.DB.Raw(sql).Scan(&count)
	if count > 0 {
		return &common.UserDetailResponse{Count: db.NewCrud(userInfo, 2)}, nil
	}
	return &common.UserDetailResponse{Count: db.NewCrud(userInfo, 1)}, nil
}

// Update 更新用户详细信息
func (ud *userDetailService) Update(context context.Context, userInfo *common.UserDetail) (*common.UserDetailResponse, error) {
	return &common.UserDetailResponse{Count: db.NewCrud(userInfo, 2)}, nil
}

// FindById 查询用户详细信息
func (ud *userDetailService) FindById(context context.Context, ids *common.UserDetailIds) (*common.UserDetailResponse, error) {
	sql := `
	SELECT user_id,
		balance,
		gender,
		nick_name,
		real_name,
		id_card,
		province,
		city,
		mark,
		avatar,
		phone,
		email,
		status,
		address,
		user_type,
		login_type,
		is_money_level,
		is_ever_level,
		lon,
		lat,
		birth_day
	FROM  `
	sql += fmt.Sprintf("%s WHERE user_id = '%v'", enum.SysUserDetail.TableName, ids.UserId[0])
	var userDetail common.UserDetail
	res := db.DB.Raw(sql).Scan(&userDetail)
	if res.Error != nil {
		return &common.UserDetailResponse{Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().CommonProvider)}, nil
	}
	return &common.UserDetailResponse{Detail: &userDetail, Count: res.RowsAffected}, nil
}

// FindBalanceByUserInfoId 查询用户余额
func (ud *userDetailService) FindBalanceByUserInfoId(context context.Context, ids *common.UserDetailIds) (*common.UserDetailResponse, error) {
	var balance common.UserBalance
	sql := fmt.Sprintf(`SELECT balance from sys_user_info WHERE user_id = '%v'`, ids.UserId[0])
	count := db.DB.Raw(sql).Scan(&balance).RowsAffected
	return &common.UserDetailResponse{Count: count, UserBalance: &balance}, nil
}

// UpdatePayPwd 更新支付密码
func (ud *userDetailService) UpdatePayPwd(context context.Context, param *common.PaymentPassword) (*common.UserDetailResponse, error) {
	return &common.UserDetailResponse{}, nil
}

// UpdateBalanceByUserId 更新用户余额
func (ud *userDetailService) UpdateBalanceByUserId(context context.Context, param *common.UserBalance) (*common.UserDetailResponse, error) {
	updateUserBalanceSql := "UPDATE " + enum.SysUserDetail.TableName + " SET balance = balance + " + param.Balance + " WHERE user_id = " + param.UserId
	res := db.DB.Exec(updateUserBalanceSql)
	if res.Error != nil {
		return &common.UserDetailResponse{Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().CommonProvider), Count: res.RowsAffected}, nil
	}
	return &common.UserDetailResponse{Count: res.RowsAffected}, nil
}
