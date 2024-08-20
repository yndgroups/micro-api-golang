package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/shopping_client"
	"protobuf/enum"
)

var AdService = &adService{}

type adService struct {
	shopping_client.AdServiceServer
}

// FindById 根据广告id查询广告详情
func (ad *adService) FindById(ctx context.Context, ids *shopping_client.AdIds) (*shopping_client.AdResponse, error) {
	detailSql := `
		SELECT
			ad_id,
			ad_type,
			position,
			store_id,
			view_type,
			ad_info,
			del_status,
			create_by,
			update_by,
			create_time,
			DATE_FORMAT(create_time,'%Y-%m-%d') create_time
		FROM`
	detailSql += fmt.Sprintf(` %v WHERE del_status = 0 AND ad_id = '%v'`, enum.Ad.TableName, ids.AdIds[0])
	var detail *shopping_client.AdDetailVo
	count := db.DB.Raw(detailSql).Scan(&detail).RowsAffected
	return &shopping_client.AdResponse{Count: count, Detail: detail}, nil
}

// FindAdList 根据广告id查询广告详情
func (ad *adService) FindAdList(ctx context.Context, params *shopping_client.AddListParma) (*shopping_client.AdResponse, error) {
	listSql := `
		SELECT
			ad_id,
			ad_type,
			position,
			store_id,
			view_type,
			ad_info,
			del_status,
			create_by,
			update_by,
			create_time,
			DATE_FORMAT(create_time,'%Y-%m-%d') create_time
		FROM`
	listSql += fmt.Sprintf(` %v WHERE del_status = 0 AND ad_type = '%v' AND position = '%v'`, enum.Ad.TableName, params.Type, params.Position)
	var list []*shopping_client.AdListVo
	count := db.DB.Raw(listSql).Scan(&list).RowsAffected
	return &shopping_client.AdResponse{Count: count, List: list}, nil
}
