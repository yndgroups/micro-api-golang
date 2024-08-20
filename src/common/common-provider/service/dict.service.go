package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/common"
	"protobuf/enum"
)

var DictService = &dictService{}

type dictService struct {
	common.DictServiceServer
}

// FindDictListByParentId 根据父id查询字典
func (d dictService) FindListByParentId(ctx context.Context, ids *common.DictIds) (*common.DictResponse, error) {
	var dictList []*common.Dict
	sql := `
SELECT
	dict_id,
	parent_id,
	dict_value,
	description,
	sort_by
FROM ` + enum.SysDict.TableName + fmt.Sprintf(` WHERE del_status = 0 AND parent_id ='%v' ORDER BY sort_by ASC`, ids.ParentId)
	affected := db.DB.Raw(sql).Scan(&dictList).RowsAffected
	return &common.DictResponse{Count: affected, List: dictList}, nil
}

func (d dictService) Create(ctx context.Context, dict *common.Dict) (*common.DictResponse, error) {
	return &common.DictResponse{Count: db.NewCrud(dict, 1)}, nil
}

func (d dictService) Update(ctx context.Context, dict *common.Dict) (*common.DictResponse, error) {
	return &common.DictResponse{Count: db.NewCrud(dict, 2)}, nil
}

func (d dictService) Delete(ctx context.Context, ids *common.DictIds) (*common.DictResponse, error) {
	return &common.DictResponse{Count: db.GrpcBatchDeleteByIds(ids.DictIds, enum.SysDict.TableName, enum.SysDict.PrimaryKey, ids.UserId)}, nil
}

// FindById 根据字典id获取字典
func (d dictService) FindById(ctx context.Context, ids *common.DictIds) (*common.DictResponse, error) {
	var detail *common.Dict
	sql := `
SELECT
	dict_id,
	parent_id,
	dict_value,
	description,
	sort_by,
	del_status,
	create_by,
	update_by,
	DATE_FORMAT(create_time,'%Y-%m-%d') create_time,
	DATE_FORMAT(update_time,'%Y-%m-%d') update_time
FROM ` + enum.SysDict.TableName + fmt.Sprintf(` WHERE del_status = 0 AND dict_id ='%v'`, ids.DictIds[0])
	// 不用&符号会报错
	affected := db.DB.Raw(sql).Scan(&detail).RowsAffected
	return &common.DictResponse{Count: affected, Detail: detail}, nil
}

func (d dictService) FindPageList(ctx context.Context, query *common.DictPageAuthQuery) (*common.DictResponse, error) {
	return &common.DictResponse{}, nil
}
