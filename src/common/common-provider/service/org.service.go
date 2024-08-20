package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/common"
	"protobuf/enum"
)

var OrgService = &orgService{}

type orgService struct {
	common.OrgServiceServer
}

func (o orgService) Create(ctx context.Context, org *common.Org) (*common.OrgResponse, error) {
	return &common.OrgResponse{Count: db.NewCrud(org, 1)}, nil
}

func (o orgService) Update(ctx context.Context, org *common.Org) (*common.OrgResponse, error) {
	return &common.OrgResponse{Count: db.NewCrud(org, 2)}, nil
}

func (o orgService) Delete(ctx context.Context, ids *common.OrgIds) (*common.OrgResponse, error) {
	return &common.OrgResponse{Count: db.GrpcBatchDeleteByIds(ids.OrgIds, enum.SysOrg.TableName, enum.SysOrg.PrimaryKey, ids.UserId)}, nil
}

func (o orgService) FindById(ctx context.Context, ids *common.OrgIds) (*common.OrgResponse, error) {
	detail := common.Org{}
	count := db.GetDB().First(&detail, common.Org{OrgId: ids.OrgIds[0]}).RowsAffected
	return &common.OrgResponse{Count: count, Detail: &detail}, nil
}

func (o orgService) FindPageList(ctx context.Context, request *common.OrgPageAuthQuery) (*common.OrgResponse, error) {
	countSql := `SELECT COUNT(1) FROM ` + enum.SysOrg.TableName + ` A`
	paramSql := ` WHERE A.del_status = 0 `
	if request.Params != nil {
		if request.Params.Name != "" {
			paramSql += fmt.Sprintf(` AND A.name like CONCAT('%%','%v','%%')`, request.Params.Name)
		}
		if request.Params.LegalPerson != "" {
			paramSql += fmt.Sprintf(` AND A.name legal_person CONCAT('%%','%v','%%')`, request.Params.LegalPerson)
		}
		if request.Params.ContactPerson != "" {
			paramSql += fmt.Sprintf(` AND A.name contact_person CONCAT('%%','%v','%%')`, request.Params.ContactPerson)
		}
	}
	var count int64
	var list []*common.Org
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		listSql := `
		SELECT 
			A.*
		FROM 
			` + enum.SysOrg.TableName + ` A `
		listSql += paramSql
		listSql += `ORDER BY A.create_time DESC LIMIT ?,?`
		db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&list)
	}
	return &common.OrgResponse{Count: count, List: list}, nil
}
