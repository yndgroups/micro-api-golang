package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/common"
	"protobuf/build/global"
	"protobuf/enum"
)

var DepartmentService = &departmentService{}

type departmentService struct {
	common.DepartmentServiceServer
}

func (d departmentService) Create(ctx context.Context, department *common.Department) (*common.DepartmentResponse, error) {
	return &common.DepartmentResponse{Count: db.NewCrud(department, 1)}, nil
}

func (d departmentService) Update(ctx context.Context, department *common.Department) (*common.DepartmentResponse, error) {
	return &common.DepartmentResponse{Count: db.NewCrud(department, 2)}, nil
}

func (d departmentService) Delete(ctx context.Context, ids *common.DepartmentIds) (*common.DepartmentResponse, error) {
	return &common.DepartmentResponse{Count: db.GrpcBatchDeleteByIds(ids.DeptIds, enum.SysDepartment.TableName, enum.SysDepartment.PrimaryKey, ids.UserId)}, nil
}

func (d departmentService) FindById(ctx context.Context, ids *common.DepartmentIds) (*common.DepartmentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d departmentService) FindTree(ctx context.Context, auth *global.Auth) (*common.DepartmentResponse, error) {
	// 通过用户查询出其组织机构,在根据组织机构查询其部门信息
	listDeptSql := `SELECT 
			DEPT.dept_id,
			DEPT.parent_id,
			DEPT.tree,
			DEPT.dept_name,
			DEPT.org_id,
			DEPT.leader,
			DEPT.phone,
			DEPT.email,
			DEPT.status,
			DEPT.sort_by,
			DEPT.del_status,
			DEPT.create_by,
			DEPT.update_by,
			DATE_FORMAT(DEPT.create_time,'%Y-%m-%d %T') create_time,
			DATE_FORMAT(DEPT.update_time,'%Y-%m-%d %T') update_time
		FROM `
	listDeptSql += fmt.Sprintf(` %s DEPT,%s ORG WHERE DEPT.org_id = ORG.org_id AND ORG.create_by = '%v'`, enum.SysDepartment.TableName, enum.SysOrg.TableName, auth.UserId)
	var deptList []*common.Department
	count := db.DB.Raw(listDeptSql).Scan(&deptList).RowsAffected
	if count > 0 {
		return &common.DepartmentResponse{Count: count, List: deptList}, nil
	}
	return &common.DepartmentResponse{Count: count, List: deptList, Msg: "未查询到相关部门信息"}, nil
}
