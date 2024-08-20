package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/order"
	"protobuf/enum"
)

var FreightTemplatesUndeliveredService = &freightTemplatesUndeliveredService{}

type freightTemplatesUndeliveredService struct {
	order.FreightTemplatesUndeliveredServiceServer
}

func (f freightTemplatesUndeliveredService) CreateFreightTemplatesUndelivered(ctx context.Context, undelivered *order.FreightTemplatesUndelivered) (*order.FreightTemplatesUndeliveredResponse, error) {
	return &order.FreightTemplatesUndeliveredResponse{Count: db.NewCrud(undelivered, 1)}, nil
}

func (f freightTemplatesUndeliveredService) UpdateFreightTemplatesUndelivered(ctx context.Context, undelivered *order.FreightTemplatesUndelivered) (*order.FreightTemplatesUndeliveredResponse, error) {
	return &order.FreightTemplatesUndeliveredResponse{Count: db.NewCrud(undelivered, 2)}, nil
}

func (f freightTemplatesUndeliveredService) DeleteFreightTemplatesUndelivered(ctx context.Context, ids *order.FreightTemplatesUndeliveredIds) (*order.FreightTemplatesUndeliveredResponse, error) {
	return &order.FreightTemplatesUndeliveredResponse{Count: db.GrpcBatchDeleteByIds(ids.UndeliveredId, enum.FreightTemplates.TableName, enum.FreightTemplates.PrimaryKey, ids.UserId)}, nil

}

func (f freightTemplatesUndeliveredService) FindPageList(ctx context.Context, request *order.FreightTemplatesUndeliveredPageAuthQuery) (*order.FreightTemplatesUndeliveredResponse, error) {
	countSql := `SELECT COUNT(1) FROM ` + enum.FreightTemplatesUndelivered.TableName + ` A`
	paramSql := ` WHERE A.del_status = 0 `
	if request.Params.AreaCode != "" {
		paramSql += fmt.Sprintf(` AND A.area_code like CONCAT('%%','%v','%%')`, request.Params.AreaCode)
	}
	var count int64
	var list []*order.FreightTemplatesUndelivered
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		listSql := `
		SELECT 
			A.undelivered_id UndeliveredId,
			A.area_code AreaCode,
			A.del_status DelStatus,
			A.create_time CreateTime,
			A.temp_id TempId,
			T.name TempName
		FROM 
			` + enum.FreightTemplatesUndelivered.TableName + ` A LEFT JOIN ` + enum.FreightTemplates.TableName + ` T ON A.temp_id=T.temp_id`
		listSql += paramSql
		listSql += `ORDER BY A.create_time DESC LIMIT ?,?`
		db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&list)
	}
	return &order.FreightTemplatesUndeliveredResponse{Count: count, List: list}, nil
}

func (f freightTemplatesUndeliveredService) FindById(ctx context.Context, ids *order.FreightTemplatesUndeliveredIds) (*order.FreightTemplatesUndeliveredResponse, error) {
	detail := order.FreightTemplatesUndelivered{}
	count := db.GetDB().First(&detail, order.FreightTemplatesUndelivered{UndeliveredId: ids.UndeliveredId[0]}).RowsAffected
	return &order.FreightTemplatesUndeliveredResponse{Count: count, Detail: &detail}, nil
}
