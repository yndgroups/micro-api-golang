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

var AgreementService = &agreementService{}

type agreementService struct {
	common.AgreementServiceServer
}

func (a *agreementService) Create(ctx context.Context, agreement *common.Agreement) (*common.AgreementResponse, error) {
	result := db.NewCrudAndMessage(agreement, 1)
	if result.Error != nil {
		return &common.AgreementResponse{Count: result.RowsAffected, Msg: exception.DbMsg(result.Error.Error(), config.Global.ServerConfig().CommonProvider)}, nil
	}
	return &common.AgreementResponse{Count: result.RowsAffected}, nil
}

func (a *agreementService) Update(ctx context.Context, agreement *common.Agreement) (*common.AgreementResponse, error) {
	return &common.AgreementResponse{Count: db.NewCrud(agreement, 2)}, nil
}

func (a *agreementService) Delete(ctx context.Context, ids *common.AgreementIds) (*common.AgreementResponse, error) {
	return &common.AgreementResponse{Count: db.GrpcBatchDeleteByIds(ids.AgreeIds, enum.SysAgreement.TableName, enum.SysAgreement.PrimaryKey, ids.UserId)}, nil
}

func (a *agreementService) FindById(ctx context.Context, ids *common.AgreementIds) (*common.AgreementResponse, error) {
	detail := common.AgreementDetailVo{}
	sql := `
SELECT 
	agree_id AgreeId,
	type Type,
	(CASE type
		WHEN 0 THEN '会员协议'
		WHEN 1 THEN '代理商协议' END
	) AS TypeName,
	title Title,
	content Content,
	sort_by SortBy,
	display Display,
	(CASE display
		WHEN 0 THEN '不显示'
		WHEN 1 THEN '显示' END
	) AS DisplayName,
	del_status DelStatus,
	DATE_FORMAT(create_time,'%Y-%m-%d') CreateTime,
	DATE_FORMAT(update_time,'%Y-%m-%d') UpdateTime
FROM ` + enum.SysAgreement.TableName + ` WHERE agree_id = ` + ids.AgreeIds[0]
	rowsAffected := db.GetDB().Raw(sql).Scan(&detail).RowsAffected
	return &common.AgreementResponse{Count: rowsAffected, Detail: &detail}, nil
}

// FindPageList 查询协议列表
func (a *agreementService) FindPageList(ctx context.Context, request *common.AgreementPageAuthQuery) (*common.AgreementResponse, error) {
	countSql := fmt.Sprintf(`SELECT COUNT(1) FROM %s A `, enum.SysAgreement.TableName)
	paramSql := ` WHERE A.del_status = 0 `
	if request.Params != nil {
		if request.Params.Title != "" {
			paramSql += fmt.Sprintf(` AND A.title like CONCAT('%%','%v','%%')`, request.Params.Title)
		}
		if request.Params.Content != "" {
			paramSql += fmt.Sprintf(` AND A.content like CONCAT('%%','%v','%%')`, request.Params.Content)
		}
	}
	var count int64
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		sql := `
	SELECT 
		A.agree_id,
		A.type,
		A.title,
		A.content,
		A.sort_by,
		A.display,
		A.del_status,
		A.create_by,
		A.update_by,
		DATE_FORMAT(A.create_time,'%Y-%m-%d %T') create_time,
		DATE_FORMAT(A.update_time,'%Y-%m-%d %T') update_time `
		sql += fmt.Sprintf(` FROM %s A`, enum.SysAgreement.TableName)
		sql += paramSql
		sql += ` LIMIT ?,?`
		var list []*common.Agreement
		db.DB.Raw(sql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&list)
		return &common.AgreementResponse{List: list, Count: count}, nil
	} else {
		return &common.AgreementResponse{Msg: "未查询到相关数据"}, nil
	}
}
