package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/common"
	"protobuf/enum"
)

var AreaService = &areaService{}

type areaService struct {
	common.AreaServiceServer
}

func (a *areaService) Create(ctx context.Context, area *common.Area) (*common.AreaResponse, error) {
	return &common.AreaResponse{Count: db.NewCrud(area, 1)}, nil
}

func (a *areaService) Update(ctx context.Context, area *common.Area) (*common.AreaResponse, error) {
	return &common.AreaResponse{Count: db.NewCrud(area, 2)}, nil
}

func (a *areaService) Delete(ctx context.Context, ids *common.AreaIds) (*common.AreaResponse, error) {
	return &common.AreaResponse{Count: db.GrpcBatchDeleteByIds(ids.AreaIds, enum.SysArea.TableName, enum.SysArea.PrimaryKey, ids.UserId)}, nil
}

// FindById 地区详情查询
func (a *areaService) FindById(ctx context.Context, ids *common.AreaIds) (*common.AreaResponse, error) {
	areaDetail := common.Area{}
	affected := db.DB.First(&areaDetail, &common.Area{AreaId: ids.AreaIds[0]}).RowsAffected
	return &common.AreaResponse{
		Count:  affected,
		Detail: &areaDetail,
	}, nil
}

// FindPageList 地区分页查询
func (a *areaService) FindPageList(ctx context.Context, request *common.AreaPageAuthQuery) (*common.AreaResponse, error) {
	countSql := `SELECT COUNT(area_id) FROM ` + enum.SysArea.TableName + ` A`
	paramSql := ` WHERE A.del_status = 0`
	if request.Params != nil {
		if request.Params.AreaName != "" {
			paramSql += fmt.Sprintf(` AND A.area_name like CONCAT('%%','%v','%%')`, request.Params.AreaName)
		}
		if request.Params.AreaCode != "" {
			paramSql += fmt.Sprintf(` AND A.area_code like CONCAT('%%','%v','%%')`, request.Params.AreaCode)
		}
	}
	var count int64
	var list []*common.Area
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		listSql := `
		SELECT 
			A.area_id,
			A.area_code,
			A.area_name,
			A.full_area_name,
			A.area_tag,
			A.parent_code,
			A.is_standard,
			A.year,
			A.del_status,
			A.create_by,
			DATE_FORMAT(A.create_time,'%Y-%m-%d') create_time
		FROM 
			` + enum.SysArea.TableName + ` A`
		listSql += paramSql
		listSql += ` LIMIT ?,?`
		db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&list)
	}
	return &common.AreaResponse{Count: count, List: list}, nil
}

// FindAreaTree 根据AreaTags查询地区信息
func (a *areaService) FindTree(ctx context.Context, tags *common.AreaTags) (*common.AreaResponse, error) {
	list := make([]*common.AreaTree, 0)
	sql := fmt.Sprintf(`SELECT area_name,area_code,parent_code,area_tag FROM sys_area WHERE area_tag <= '%v'`, tags.AreaTag)
	affected := db.DB.Raw(sql).Scan(&list).RowsAffected
	return &common.AreaResponse{TreeList: list, Count: affected}, nil
}
