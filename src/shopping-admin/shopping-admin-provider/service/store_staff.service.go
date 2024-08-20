package service

import (
	"context"
	"core/db"
	"errors"
	"fmt"
	"protobuf/build/shopping_admin"
	"protobuf/enum"
)

var StoreStaffService = &storeStaffService{}

type storeStaffService struct {
	shopping_admin.StoreStaffServiceServer
}

// CreateStoreStaff 添加员工信息
func (s *storeStaffService) CreateStoreStaff(ctx context.Context, staff *shopping_admin.StoreStaff) (*shopping_admin.StoreStaffResponse, error) {
	return &shopping_admin.StoreStaffResponse{Count: db.NewCrud(staff, 1)}, nil
}

// UpdateStoreStaff 更新员工信息
func (s *storeStaffService) UpdateStoreStaff(ctx context.Context, staff *shopping_admin.StoreStaff) (*shopping_admin.StoreStaffResponse, error) {
	return &shopping_admin.StoreStaffResponse{Count: db.NewCrud(staff, 2)}, nil
}

// DeleteStoreStaff 删除员工信息
func (s *storeStaffService) DeleteStoreStaff(ctx context.Context, ids *shopping_admin.StoreStaffIds) (*shopping_admin.StoreStaffResponse, error) {
	return &shopping_admin.StoreStaffResponse{Count: db.GrpcBatchDeleteByIds(ids.StaffId, enum.FreightTemplates.TableName, enum.FreightTemplates.PrimaryKey, ids.UserId)}, nil
}

// FindPageList  查询员工列表
func (s *storeStaffService) FindPageList(ctx context.Context, request *shopping_admin.StoreStaffPageAuthQuery) (*shopping_admin.StoreStaffResponse, error) {
	countSql := `SELECT COUNT(1) FROM ` + enum.StoreStaff.TableName + ` A`
	paramSql := ` WHERE A.del_status = 0 `
	if request.Params.JobName != "" {
		paramSql += fmt.Sprintf(` AND A.job_name like CONCAT('%%','%v','%%')`, request.Params.JobName)
	}
	if request.Params.RealName != "" {
		paramSql += fmt.Sprintf(` AND A.real_name like CONCAT('%%','%v','%%')`, request.Params.RealName)
	}
	if request.Params.Phone != "" {
		paramSql += fmt.Sprintf(` AND A.phone like CONCAT('%%','%v','%%')`, request.Params.Phone)
	}
	var count int64
	var list []*shopping_admin.StoreStaff
	err := db.DB.Raw(countSql + paramSql).Count(&count).Error
	if count > 0 {
		listSql := `
		SELECT 
			A.staff_id,
			A.user_id,
			A.store_id,
			A.real_name,
			A.avatar,
			A.job_name,
			A.phone,
			A.email,
			A.push_status,
			A.verify_status,
			A.del_status,
			A.create_by,
			A.update_by,
			DATE_FORMAT(A.create_time,'%Y-%m-%d') AS create_time,
			A.update_time
		FROM 
			` + enum.StoreStaff.TableName + ` A`
		listSql += paramSql
		listSql += ` LIMIT ?,?`
		err = db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&list).Error
	}
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &shopping_admin.StoreStaffResponse{Count: count, List: list}, nil
}

// FindById 查询员工详情
func (s *storeStaffService) FindById(ctx context.Context, ids *shopping_admin.StoreStaffIds) (*shopping_admin.StoreStaffResponse, error) {
	detail := shopping_admin.StoreStaff{}
	count := db.GetDB().First(&detail, shopping_admin.StoreStaff{StaffId: ids.StaffId[0]}).RowsAffected
	return &shopping_admin.StoreStaffResponse{Count: count, Detail: &detail}, nil
}