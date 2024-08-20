package service

import (
	"context"
	"core/db"
	"protobuf/build/common"
	"protobuf/enum"
)

var AppModuleService = &appModuleService{}

type appModuleService struct {
	common.AppModuleServiceServer
}

func (a *appModuleService) Create(ctx context.Context, module *common.AppModule) (*common.AppModuleResponse, error) {
	return &common.AppModuleResponse{Count: db.NewCrud(module, 1)}, nil
}

func (a *appModuleService) Delete(ctx context.Context, ids *common.AppModuleIds) (*common.AppModuleResponse, error) {
	return &common.AppModuleResponse{Count: db.GrpcBatchDeleteByIds(ids.MdIds, enum.SysAppModule.TableName, enum.SysAppModule.PrimaryKey, ids.UserId)}, nil
}

func (a *appModuleService) Update(ctx context.Context, module *common.AppModule) (*common.AppModuleResponse, error) {
	return &common.AppModuleResponse{Count: db.NewCrud(module, 2)}, nil
}

func (a *appModuleService) FindById(ctx context.Context, ids *common.AppModuleIds) (*common.AppModuleResponse, error) {
	return &common.AppModuleResponse{}, nil
}

func (a *appModuleService) FindPageList(ctx context.Context, query *common.AppModulePageAuthQuery) (*common.AppModuleResponse, error) {
	return &common.AppModuleResponse{}, nil
	/*listModuleVo := make([]common.AppModule, 0)
	sql := `SELECT A.md_id, A.name, A.host, A.match_path, A.js_path, A.css_path, A.del_status, A.create_by, A.update_by,A.create_time,
			A.update_time FROM sys_app_module A WHERE  A.md_id = (SELECT FIND_IN_SET(A.md_id,(SELECT md_id from sys_app B WHERE B.app_id = ?)))`
	if db.DB.Raw(sql, appId).Scan(&listModuleVo).RowsAffected > 0 {
		return resp.Success.WithData(listModuleVo)
	}
	return resp.NotData*/
}
