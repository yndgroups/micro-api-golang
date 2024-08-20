package service

import (
	"context"
	"core/db"
	"core/validate"
	"fmt"
	"protobuf/build/common"
	"protobuf/enum"
)

var RoleService = &roleService{}

type roleService struct {
	common.RoleServiceServer
}

// 创建角色
func (r roleService) Create(ctx context.Context, role *common.Role) (*common.RoleResponse, error) {
	return &common.RoleResponse{Count: db.NewCrud(role, 1)}, nil
}

// 更新角色信息
func (r roleService) Update(ctx context.Context, role *common.Role) (*common.RoleResponse, error) {
	return &common.RoleResponse{Count: db.NewCrud(role, 2)}, nil
}

// 删除角色
func (r roleService) Delete(ctx context.Context, ids *common.RoleIds) (*common.RoleResponse, error) {
	return &common.RoleResponse{Count: db.GrpcBatchDeleteByIds(ids.RoleId, enum.SysRole.TableName, enum.SysRole.PrimaryKey, ids.UserId)}, nil
}

// 根绝角色id查询角色详情
func (r roleService) FindById(ctx context.Context, ids *common.RoleIds) (*common.RoleResponse, error) {
	detail := common.Role{}
	count := db.GetDB().First(&detail, ids.RoleId[0]).RowsAffected
	return &common.RoleResponse{Count: count, Detail: &detail}, nil
}

// 查询角色列表
func (r roleService) FindPageList(ctx context.Context, request *common.RolePageAuthQuery) (*common.RoleResponse, error) {
	// 验证参数合法性
	if _, err := validate.Validated(request); err != nil {
		return &common.RoleResponse{Msg: err.Error()}, nil
	}
	// 执行查询
	countSql := `SELECT COUNT(1) FROM ` + enum.SysRole.TableName + ` A `
	paramSql := ""
	if request.Params != nil {
		paramSql := fmt.Sprintf(` WHERE A.del_status = 0 AND A.app_id = '%v' `, request.Params.AppId)
		if request.Params.Name != "" {
			paramSql += fmt.Sprintf(` AND A.name like CONCAT('%%','%v','%%')`, request.Params.Name)
		}
	}
	var count int64
	var list []*common.Role
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		listSql := `
		SELECT 
			A.role_id,
			A.app_id,
			A.name,
			A.introduce,
			A.unique_identification,
			A.sort_by,
			DATE_FORMAT(A.create_time,'%Y-%m-%d') create_time
		FROM
			` + enum.SysRole.TableName + ` A `
		listSql += paramSql
		listSql += `ORDER BY A.create_time DESC LIMIT ?,?`
		db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&list)
	}
	return &common.RoleResponse{Count: count, List: list}, nil
}

// 查询授权的用户
func (r roleService) FindAuthRoleListByUserId(ctx context.Context, param *common.RoleAuthParam) (*common.RoleResponse, error) {
	var list []*common.Role
	listRoleSql := `
		SELECT 
			A.role_id,
			A.app_id,
			A.name,
			A.introduce,
			A.unique_identification,
			A.sort_by,
			DATE_FORMAT(A.create_time,'%Y-%m-%d') create_time
		FROM
			`
	listRoleSql += fmt.Sprintf(`%v A WHERE app_id = '%v'  ORDER BY sort_by DESC`, enum.SysRole.TableName, param.AppId)
	count := db.DB.Raw(listRoleSql).Scan(&list).RowsAffected
	if count > 0 {
		var authRoleIds []string
		listAuthRoleIdSql := fmt.Sprintf(`SELECT role_id FROM %v WHERE user_id = '%v' AND app_id = '%v'`, enum.SysUserRole.TableName, param.UserId, param.AppId)
		db.DB.Raw(listAuthRoleIdSql).Scan(&authRoleIds)
		return &common.RoleResponse{Count: count, List: list, RoleIds: authRoleIds}, nil
	} else {
		return &common.RoleResponse{Count: 0, Msg: "未查询到角色信息"}, nil
	}
}
