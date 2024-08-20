package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/common"
	"protobuf/enum"
	"strings"
)

var ConfigService = &configService{}

type configService struct {
	common.ConfigServiceServer
}

// Create 创建配置
func (c configService) Create(ctx context.Context, config *common.Config) (*common.ConfigResponse, error) {
	return &common.ConfigResponse{Count: db.NewCrud(config, 1)}, nil
}

// Update 修改配置
func (c configService) Update(ctx context.Context, config *common.Config) (*common.ConfigResponse, error) {
	return &common.ConfigResponse{Count: db.NewCrud(config, 2)}, nil
}

// Delete 删除配置
func (c configService) Delete(ctx context.Context, ids *common.ConfigIds) (*common.ConfigResponse, error) {
	return &common.ConfigResponse{Count: db.GrpcBatchDeleteByIds(ids.ConfIds, enum.SysConfig.TableName, enum.SysConfig.PrimaryKey, ids.UserId)}, nil
}

// FindById 查询配置详情
func (c configService) FindById(ctx context.Context, ids *common.ConfigIds) (*common.ConfigResponse, error) {
	var config common.Config
	sql := `
	SELECT
		conf_id,
		conf_type,
		app_id,
		outside_app_id,
		outside_app_name,
		outside_app_secret,
		mch_id,
		pater_ner_key,
		phone,
		email,
		return_url,
		notify_url,
		js_api_ticket,
		token,
		amount,
		msg_data_format,
		serial_no,
		api_v2_key,
		api_v3_key,
		private_key,
		aes_key,
		timestamp,
		del_status,
		create_by,
		update_by,
		DATE_FORMAT(create_time,'%Y-%m-%d') create_time,
		DATE_FORMAT(update_time,'%Y-%m-%d') update_time
	FROM ` + enum.SysConfig.TableName
	sql += fmt.Sprintf(` WHERE del_status = 0 AND conf_id = '%v'`, ids.ConfIds[0])
	affected := db.DB.Raw(sql).Scan(&config).RowsAffected
	return &common.ConfigResponse{Detail: &config, Count: affected}, nil
}

// FindByConfigExpandParam 根据查询参数获取配置
func (c configService) FindByConfigExpandParam(ctx context.Context, param *common.ConfigIds) (*common.ConfigResponse, error) {
	if param == nil {
		return &common.ConfigResponse{Msg: "必须传参数不能为空"}, nil
	}
	if param.AppId == "" {
		return &common.ConfigResponse{Msg: "必须传参数AppId不能为空"}, nil
	}
	sql := `
	SELECT
		conf_id,
		conf_type,
		app_id,
		outside_app_id,
		outside_app_name,
		outside_app_secret,
		mch_id,
		pater_ner_key,
		phone,
		email,
		return_url,
		notify_url,
		js_api_ticket,
		token,
		amount,
		msg_data_format,
		serial_no,
		api_v2_key,
		api_v3_key,
		private_key,
		aes_key,
		timestamp,
		del_status,
		create_by,
		update_by,
		DATE_FORMAT(create_time,'%Y-%m-%d') create_time,
		DATE_FORMAT(update_time,'%Y-%m-%d') update_time
	FROM ` + enum.SysConfig.TableName
	sql += fmt.Sprintf(` WHERE del_status = 0 AND app_id = '%v'`, param.AppId)
	if len(param.ConfIds) > 0 {
		sql += fmt.Sprintf(` AND conf_id in (%v)`, strings.Join(param.ConfIds, ","))
	}
	if len(param.StoreIds) > 0 {
		sql += fmt.Sprintf(` AND create_by in (%v)`, strings.Join(param.StoreIds, ","))
	}
	list := make([]*common.Config, 0)
	rowsAffected := db.DB.Raw(sql).Scan(&list).RowsAffected
	if rowsAffected > 0 {
		if (len(param.StoreIds) == 1 || len(param.ConfIds) == 1) && len(list) > 0 {
			return &common.ConfigResponse{Detail: list[0], Count: rowsAffected}, nil
		}
		return &common.ConfigResponse{List: list, Count: rowsAffected}, nil
	}
	return &common.ConfigResponse{Msg: "未查询到配置"}, nil
}

// FindPageList 查询配置列表
func (c configService) FindPageList(ctx context.Context, query *common.ConfigPageAuthQuery) (*common.ConfigResponse, error) {
	return &common.ConfigResponse{}, nil
}
