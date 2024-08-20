package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/order"
	"protobuf/enum"
)

var FreightTemplatesFreeService = &freightTemplatesFreeService{}

type freightTemplatesFreeService struct {
	order.FreightTemplatesFreeServiceServer
}

func (f freightTemplatesFreeService) CreateFreightTemplatesFree(ctx context.Context, free *order.FreightTemplatesFree) (*order.FreightTemplatesFreeResponse, error) {
	return &order.FreightTemplatesFreeResponse{Count: db.NewCrud(free, 1)}, nil
}

func (f freightTemplatesFreeService) UpdateFreightTemplatesFree(ctx context.Context, free *order.FreightTemplatesFree) (*order.FreightTemplatesFreeResponse, error) {
	return &order.FreightTemplatesFreeResponse{Count: db.NewCrud(free, 2)}, nil
}

func (f freightTemplatesFreeService) DeleteFreightTemplatesFree(ctx context.Context, ids *order.FreightTemplatesFreeIds) (*order.FreightTemplatesFreeResponse, error) {
	return &order.FreightTemplatesFreeResponse{Count: db.GrpcBatchDeleteByIds(ids.FreeId, enum.FreightTemplates.TableName, enum.FreightTemplates.PrimaryKey, ids.UserId)}, nil
}

func (f freightTemplatesFreeService) FindPageList(ctx context.Context, request *order.FreightTemplatesFreePageAuthQuery) (*order.FreightTemplatesFreeResponse, error) {
	countSql := `SELECT COUNT(1) FROM ` + enum.FreightTemplatesFree.TableName + ` A`
	paramSql := ` WHERE A.del_status = 0 `
	if request.Params.AreaCode != "" {
		paramSql += fmt.Sprintf(` AND A.area_code like CONCAT('%%','%v','%%')`, request.Params.AreaCode)
	}
	var count int64
	var list []*order.FreightTemplatesFree
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		listSql := `
		SELECT 
			A.free_id FreeId,
			A.billing_methods BillingMethods,
			A.area_code AreaCode,
			A.quantity Quantity,
			A.price Price,
			T.name TempName
		FROM 
			` + enum.FreightTemplatesFree.TableName + ` A LEFT JOIN ` + enum.FreightTemplates.TableName + ` T ON A.temp_id=T.temp_id`
		listSql += paramSql
		listSql += `ORDER BY A.create_time DESC LIMIT ?,?`
		db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&list)
	}
	return &order.FreightTemplatesFreeResponse{Count: count, List: list}, nil
}

func (f freightTemplatesFreeService) FindById(ctx context.Context, ids *order.FreightTemplatesFreeIds) (*order.FreightTemplatesFreeResponse, error) {
	detail := order.FreightTemplatesFree{}
	count := db.GetDB().First(&detail, order.FreightTemplatesFree{FreeId: ids.FreeId[0]}).RowsAffected
	return &order.FreightTemplatesFreeResponse{Count: count, Detail: &detail}, nil
}
