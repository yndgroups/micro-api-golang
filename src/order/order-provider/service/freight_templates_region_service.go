package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/order"
	"protobuf/enum"
)

var FreightTemplatesRegionService = &freightTemplatesRegionService{}

type freightTemplatesRegionService struct {
	order.FreightTemplatesRegionServiceServer
}

func (f freightTemplatesRegionService) CreateFreightTemplatesRegion(ctx context.Context, region *order.FreightTemplatesRegion) (*order.FreightTemplatesRegionResponse, error) {
	return &order.FreightTemplatesRegionResponse{Count: db.NewCrud(region, 1)}, nil
}

func (f freightTemplatesRegionService) UpdateFreightTemplatesRegion(ctx context.Context, region *order.FreightTemplatesRegion) (*order.FreightTemplatesRegionResponse, error) {
	return &order.FreightTemplatesRegionResponse{Count: db.NewCrud(region, 2)}, nil
}

func (f freightTemplatesRegionService) DeleteFreightTemplatesRegion(ctx context.Context, ids *order.FreightTemplatesRegionIds) (*order.FreightTemplatesRegionResponse, error) {
	return &order.FreightTemplatesRegionResponse{Count: db.GrpcBatchDeleteByIds(ids.RegionId, enum.FreightTemplates.TableName, enum.FreightTemplates.PrimaryKey, ids.UserId)}, nil
}

func (f freightTemplatesRegionService) FindPageList(ctx context.Context, request *order.FreightTemplatesRegionPageAuthQuery) (*order.FreightTemplatesRegionResponse, error) {
	countSql := `SELECT COUNT(1) FROM ` + enum.FreightTemplatesRegion.TableName + ` A`
	paramSql := ` WHERE A.del_status = 0 `
	if request.Params.Name != "" {
		paramSql += fmt.Sprintf(` AND A.name like CONCAT('%%','%v','%%')`, request.Params.Name)
	}
	var count int64
	var list []*order.FreightTemplatesRegion
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		listSql := `
		SELECT 
			A.region_id RegionId,
			A.area_code AreaCode,
			A.del_status DelStatus,
			A.create_time CreateTime,
			A.temp_id TempId,
			A.billing_methods BillingMethods,
			A.first_price FirstPrice,
			A.continued Continued,
			A.continued_price ContinuedPrice,
			T.name TempName
		FROM 
			` + enum.FreightTemplatesRegion.TableName + ` A LEFT JOIN ` + enum.FreightTemplates.TableName + ` T ON A.temp_id=T.temp_id`
		listSql += paramSql
		listSql += `ORDER BY A.create_time DESC LIMIT ?,?`
		db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&list)
	}
	return &order.FreightTemplatesRegionResponse{Count: count, List: list}, nil
}

func (f freightTemplatesRegionService) FindById(ctx context.Context, ids *order.FreightTemplatesRegionIds) (*order.FreightTemplatesRegionResponse, error) {
	detail := order.FreightTemplatesRegion{}
	count := db.GetDB().First(&detail, order.FreightTemplatesRegion{RegionId: ids.RegionId[0]}).RowsAffected
	return &order.FreightTemplatesRegionResponse{Count: count, Detail: &detail}, nil
}
