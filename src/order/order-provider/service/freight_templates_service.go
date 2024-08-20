package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/order"
	"protobuf/enum"
)

var FreightTemplatesService = &freightTemplatesService{}

type freightTemplatesService struct {
	order.FreightTemplatesServiceServer
}

// CreateFreightTemplates 创建运费模板
func (f freightTemplatesService) CreateFreightTemplates(ctx context.Context, templates *order.FreightTemplates) (*order.FreightTemplatesResponse, error) {
	return &order.FreightTemplatesResponse{Count: db.NewCrud(templates, 1)}, nil
}

// UpdateFreightTemplates 修改运费模板
func (f freightTemplatesService) UpdateFreightTemplates(ctx context.Context, templates *order.FreightTemplates) (*order.FreightTemplatesResponse, error) {
	return &order.FreightTemplatesResponse{Count: db.NewCrud(templates, 2)}, nil
}

// DeleteFreightTemplates 删除运费模板
func (f freightTemplatesService) DeleteFreightTemplates(ctx context.Context, ids *order.FreightTemplateIds) (*order.FreightTemplatesResponse, error) {
	return &order.FreightTemplatesResponse{Count: db.GrpcBatchDeleteByIds(ids.TempId, enum.FreightTemplates.TableName, enum.FreightTemplates.PrimaryKey, ids.UserId)}, nil
}

func (f freightTemplatesService) FindPageList(ctx context.Context, request *order.FreightTemplatesPageAuthQuery) (*order.FreightTemplatesResponse, error) {
	countSql := `SELECT COUNT(1) FROM ` + enum.FreightTemplates.TableName + ` A`
	paramSql := ` WHERE A.del_status = 0 `
	if request.Params.Name != "" {
		paramSql += fmt.Sprintf(` AND A.name like CONCAT('%%','%v','%%')`, request.Params.Name)
	}
	var count int64
	var list []*order.FreightTemplatesListVo
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		listSql := `
		SELECT 
			A.temp_id TempId,
			A.name Name,
			A.billing_methods BillingMethods,
			(CASE A.billing_methods
				WHEN 1 THEN 'DHL快递收费'
				WHEN 2 THEN 'ups快递收费' 
				WHEN 3 THEN 'FEDEX快递收费' 
				WHEN 4 THEN 'EMS快递收费' END
			) AS BillingMethodsName,
		    A.appoint Appoint,
		    (CASE A.appoint
				WHEN 0 THEN '包邮'
				WHEN 1 THEN '不包邮' END
			) AS AppointName,
		    A.no_delivery NoDelivery,
		    (CASE A.no_delivery
				WHEN 0 THEN '送达'
				WHEN 1 THEN '不送达' END
			) AS NoDeliveryName,
		    A.sort_by SortBy
		FROM 
			` + enum.FreightTemplates.TableName + ` A`
		listSql += paramSql
		listSql += ` ORDER BY A.create_time DESC LIMIT ?,?`
		db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&list)
	}
	return &order.FreightTemplatesResponse{Count: count, List: list}, nil
}

// FindById 根据id查询详情
func (f freightTemplatesService) FindById(ctx context.Context, ids *order.FreightTemplateIds) (*order.FreightTemplatesResponse, error) {
	detail := order.FreightTemplates{}
	count := db.GetDB().First(&detail, order.FreightTemplates{TempId: ids.TempId[0]}).RowsAffected
	return &order.FreightTemplatesResponse{Count: count, Detail: &detail}, nil
}
