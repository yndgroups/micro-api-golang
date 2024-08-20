package service

import (
	"context"
	"core/config"
	"core/db"
	"fmt"
	"protobuf/build/common"
	"protobuf/build/global"
	"protobuf/enum"
	"strings"
)

var AppService = &appService{}

type appService struct {
	common.AppServiceServer
}

// CreateApp 创建应用
func (a *appService) Create(ctx context.Context, app *common.App) (*common.AppResponse, error) {
	return &common.AppResponse{Count: db.NewCrud(app, 1)}, nil
}

// DeleteApp 删除应用
func (a *appService) Delete(ctx context.Context, ids *common.AppIds) (*common.AppResponse, error) {
	return &common.AppResponse{Count: db.GrpcBatchDeleteByIds(ids.AppIds, enum.SysApp.TableName, enum.SysApp.PrimaryKey, ids.UserId)}, nil
}

// UpdateApp 更新应用
func (a *appService) Update(ctx context.Context, app *common.App) (*common.AppResponse, error) {
	return &common.AppResponse{Count: db.NewCrud(app, 2)}, nil
}

// FindById 查询详情
func (a *appService) FindById(ctx context.Context, ids *common.AppIds) (*common.AppResponse, error) {
	return &common.AppResponse{
		Count: 0,
	}, nil
}

// FindPageList 查询应用分页数据
func (a *appService) FindPageList(ctx context.Context, request *common.AppPageAuthQuery) (*common.AppResponse, error) {
	listSql := `SELECT app_id,name,icon,introduce FROM ` + enum.SysApp.TableName
	paramSql := ` WHERE del_status = 0`
	if request.Params != nil {
		if request.Params.Name != "" {
			paramSql += fmt.Sprintf(` AND name like CONCAT('%%','%v','%%')`, request.Params.Name)
		}
		if request.Params.Introduce != "" {
			paramSql += fmt.Sprintf(` AND introduce like CONCAT('%%','%v','%%')`, request.Params.Introduce)
		}
	}
	listSql += paramSql
	listSql += ` LIMIT ?,?`
	var appList []*common.App
	var count int64
	db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&appList).Count(&count)
	return &common.AppResponse{Count: count, List: appList}, nil
}

// FindAppList 查询权限下的应用数据
func (a *appService) FindList(ctx context.Context, auth *global.Auth) (*common.AppResponse, error) {
	var appList []*common.App
	var count int64
	if strings.Contains(config.Global.ServerConfig().Administrator, auth.UserId) {
		sqlStr := `
		SELECT 
			A.app_id,
			A.name,
			A.icon,
			A.introduce 
		FROM 
		`
		sqlStr += enum.SysApp.TableName + " A WHERE A.del_status = 0"
		db.DB.Raw(sqlStr).Scan(&appList).Count(&count)
	} else {
		sqlStr := `
     SELECT 
         A.app_id,
         A.name,
         A.icon,
         A.introduce 
     FROM 
	`
		sqlStr += enum.SysApp.TableName + " A, " + enum.UserApp.TableName + " U  WHERE U.del_status = 0 AND U.app_id = A.app_id AND U.user_id = '" + auth.UserId + "'"
		db.DB.Raw(sqlStr).Scan(&appList).Count(&count)
	}
	return &common.AppResponse{Count: count, List: appList}, nil
}
