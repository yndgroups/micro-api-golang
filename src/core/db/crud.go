package db

import (
	"core/model"
	"core/resp"
	"fmt"

	"gorm.io/gorm"
)

func BatchDeleteByIds(ids model.Ids, tableName string, filedIdName string, userId string) any {
	if rowsAffected := GetDB().Table(tableName).Where(filedIdName+" IN ?", ids).Updates(map[string]interface{}{"del_status": "1", "update_by": userId}).RowsAffected; rowsAffected > 0 {
		return resp.DeleteSuccess.WithMsg(fmt.Sprintf("成功删除%v条数据", rowsAffected))
	} else {
		return resp.DeleteFail
	}
}

func GrpcBatchDeleteByIds(ids model.Ids, tableName string, filedIdName string, userId string) int64 {
	return GetDB().Table(tableName).Where(filedIdName+" IN ?", ids).Updates(map[string]interface{}{"del_status": "1", "update_by": userId}).RowsAffected
}

// Crud 公共新增雨修改
func Crud[T any](data T, cuType int) any {
	if cuType == 1 {
		if GetDB().Create(data).RowsAffected > 0 {
			return resp.InsertSuccess
		} else {
			return resp.InsertFail
		}
	} else if cuType == 2 {
		if GetDB().Updates(data).RowsAffected > 0 {
			return resp.UpdateSuccess
		} else {
			return resp.InsertFail
		}
	}
	return resp.Fail.WithMsg("程序异常，请联系管理员处理")
}

// NewCrud 公共新增与修改
func NewCrud[T any](data T, cuType int) int64 {
	if cuType == 1 {
		return GetDB().Create(data).RowsAffected
	} else if cuType == 2 {
		return GetDB().Updates(data).RowsAffected
	} else {
		return 0
	}
}

// NewCrud 公共新增与修改
func NewCrudAndMessage[T any](data T, cuType int) *gorm.DB {
	if cuType == 1 {
		return GetDB().Create(data)
	} else if cuType == 2 {
		return GetDB().Updates(data)
	} else {
		return nil
	}
}
