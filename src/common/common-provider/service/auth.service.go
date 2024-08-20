package service

import (
	"context"
	"core/config"
	"core/db"
	"core/exception"
	"fmt"
	"protobuf/build/common"
	"protobuf/build/global"
	"protobuf/enum"
)

var AuthService = &authService{}

type authService struct {
	common.AuthServer
}

// 新增角色关联菜单数据
func (r authService) CreateRoleMenu(ctx context.Context, params *common.RoleMenuList) (*common.AuthResponse, error) {
	insertSql := fmt.Sprintf("INSERT INTO `%v` (`role_id`,`menu_id`,`create_by`) VALUES ", enum.SysRoleMenu.TableName)
	for i, v := range params.List {
		if i == 0 { // 删除历史授权数据重新插入
			sql := fmt.Sprintf(`DELETE FROM %v WHERE role_id='%v'`, enum.SysRoleMenu.TableName, v.RoleId)
			db.GetDB().Exec(sql)
			insertSql += fmt.Sprintf(`('%v','%v','%v')`, v.RoleId, v.MenuId, v.CreateBy)
		} else {
			insertSql += fmt.Sprintf(`,('%v','%v','%v')`, v.RoleId, v.MenuId, v.CreateBy)
		}
	}
	res := db.GetDB().Exec(insertSql)
	if res.Error != nil {
		if res.Error != nil {
			return &common.AuthResponse{Count: 0, Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().CommonProvider)}, nil
		}
	}
	var userIds []string
	if res.RowsAffected > 0 {
		// 查询该角色影响的用户，通知其进行权限变更
		findUserRoleSql := fmt.Sprintf(`SELECT user_id FROM %v WHERE role_id = '%v'`, enum.SysUserRole.TableName, params.List[0].RoleId)
		db.GetDB().Raw(findUserRoleSql).Scan(&userIds)
	}
	return &common.AuthResponse{Count: res.RowsAffected, UserIds: userIds}, nil
}

// 新增角色关联功能数据
func (r authService) CreateRoleFunc(ctx context.Context, params *common.RoleFuncList) (*common.AuthResponse, error) {
	insertSql := fmt.Sprintf("INSERT INTO `%v` (`role_id`,`func_id`,`create_by`) VALUES ", enum.SysRoleFunc.TableName)
	for i, v := range params.List {
		if i == 0 { // 删除历史授权数据重新插入
			sql := fmt.Sprintf(`DELETE FROM %v WHERE role_id='%v'`, enum.SysRoleFunc.TableName, v.RoleId)
			db.GetDB().Exec(sql)
			insertSql += fmt.Sprintf(`('%v','%v','%v')`, v.RoleId, v.FuncId, v.CreateBy)
		} else {
			insertSql += fmt.Sprintf(`,('%v','%v','%v')`, v.RoleId, v.FuncId, v.CreateBy)
		}
	}
	res := db.GetDB().Exec(insertSql)
	if res.Error != nil {
		if res.Error != nil {
			return &common.AuthResponse{Count: 0, Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().CommonProvider)}, nil
		}
	}
	var userIds []string
	if res.RowsAffected > 0 {
		// 查询该角色影响的用户，通知其进行权限变更
		findUserRoleSql := fmt.Sprintf(`SELECT user_id FROM %v WHERE role_id = '%v'`, enum.SysUserRole.TableName, params.List[0].RoleId)
		db.GetDB().Raw(findUserRoleSql).Scan(&userIds)
	}
	return &common.AuthResponse{Count: res.RowsAffected, UserIds: userIds}, nil
}

// CreateUserRole 创建角色与用户关联
func (r authService) CreateUserRole(ctx context.Context, params *common.UserRoleList) (*common.AuthResponse, error) {
	insertSql := fmt.Sprintf("INSERT INTO `%v` (`role_id`,`user_id`,`app_id`,`create_by`) VALUES ", enum.SysUserRole.TableName)
	// 批量插入sql INSERT INTO `sys_user_role` (`role_id`,`user_id`,`app_id`,`create_by`) VALUES ('1','1','1','1'),VALUES ('2','2','2','2')
	for i, v := range params.List {
		if i == 0 { // 删除历史授权数据重新插入
			sql := fmt.Sprintf(`DELETE FROM %v WHERE user_id='%v' AND app_id='%v'`, enum.SysUserRole.TableName, v.UserId, v.AppId)
			db.GetDB().Exec(sql)
			insertSql += fmt.Sprintf(`('%v','%v','%v','%v')`, v.RoleId, v.UserId, v.AppId, v.CreateBy)
		} else {
			insertSql += fmt.Sprintf(`,('%v','%v','%v','%v')`, v.RoleId, v.UserId, v.AppId, v.CreateBy)
		}
	}
	res := db.GetDB().Exec(insertSql)
	if res.Error != nil {
		return &common.AuthResponse{Count: 0, Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().CommonProvider)}, nil
	}
	return &common.AuthResponse{Count: res.RowsAffected}, nil
}

func (r authService) GetToken(ctx context.Context, auth *global.Auth) (*common.AuthResponse, error) {
	listSql := fmt.Sprintf(`SELECT COUNT(app_id) FROM %v WHERE del_status = 0 AND app_id = '%v'`, enum.SysApp.TableName, auth.AppId)
	var count int64
	db.DB.Raw(listSql).Scan(&count)
	if count > 0 {
		return &common.AuthResponse{Count: 0, Msg: "test"}, nil
	}
	return &common.AuthResponse{Count: 0, Msg: "应用不存在"}, nil
}
