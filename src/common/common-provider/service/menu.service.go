package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/common"
	"protobuf/enum"
)

var MenuService = &menuService{}

type menuService struct {
	common.MenuServiceServer
}

// 新增菜单
func (m menuService) Create(ctx context.Context, menu *common.Menu) (*common.MenuResponse, error) {
	return &common.MenuResponse{Count: db.NewCrud(menu, 1)}, nil
}

// 删除菜单
func (m menuService) Delete(ctx context.Context, ids *common.MenuIds) (*common.MenuResponse, error) {
	return &common.MenuResponse{Count: db.GrpcBatchDeleteByIds(ids.MenuIds, enum.SysMenu.TableName, enum.SysMenu.PrimaryKey, ids.UserId)}, nil
}

// 修改菜单
func (m menuService) Update(ctx context.Context, menu *common.Menu) (*common.MenuResponse, error) {
	return &common.MenuResponse{Count: db.NewCrud(menu, 2)}, nil
}

// 根据用户id查询菜单及权限值
func (m menuService) FindTreeList(ctx context.Context, param *common.RoleMenuParam) (*common.MenuResponse, error) {
	menuSql := fmt.Sprintf(`
	SELECT 
		SM.menu_id,
		SM.parent_id,
		SM.icon,
		SM.name,
		SM.type,
		SM.url,
		SM.power_sign,
		SM.sort_by 
	FROM
		sys_menu SM,
		sys_role_menu SRM, 
		sys_user_role SUR 
	WHERE 
		SM.del_status = 0 
		AND SM.menu_id = SRM.menu_id 
		AND SM.app_id = SUR.app_id 
		AND SRM.role_id = SUR.role_id 
		AND SUR.user_id = '%s' 
		AND SUR.app_id = '%s' 
		ORDER BY SM.sort_by`, param.UserId, param.AppId)
	menuTreeList := make([]*common.Menu, 0)
	count := db.GetDB().Raw(menuSql).Scan(&menuTreeList).RowsAffected
	if len(menuTreeList) > 0 {
		return &common.MenuResponse{List: menuTreeList, Count: count, PowerSign: UserService.FindPowerSign(param.AppId, param.UserId)}, nil
	} else {
		return &common.MenuResponse{Msg: "未查询菜单数据"}, nil
	}
}

// 根据角色查询当前应用下的菜单及已经授权了的菜单id信息
func (m menuService) FindByRoleIds(ctx context.Context, param *common.RoleMenuParam) (*common.MenuResponse, error) {
	menuSql := fmt.Sprintf(`SELECT M.menu_id,M.parent_id,M.icon,M.name,M.type,M.url,M.power_sign,M.sort_by FROM sys_menu M,sys_role R WHERE M.del_status = 0 AND M.app_id = R.app_id and R.role_id = '%v' ORDER BY sort_by`, param.RoleIds[0])
	menuTreeList := make([]*common.Menu, 0)
	count := db.GetDB().Raw(menuSql).Scan(&menuTreeList).RowsAffected
	if len(menuTreeList) > 0 {
		roleMenuSql := fmt.Sprintf(`SELECT menu_id from sys_role_menu WHERE role_id = '%v'`, param.RoleIds[0])
		var menuIds []string
		db.GetDB().Raw(roleMenuSql).Scan(&menuIds)
		return &common.MenuResponse{List: menuTreeList, Count: count, RoleMenuIds: menuIds}, nil
	} else {
		return &common.MenuResponse{Msg: "该角色还未授权任何菜单数据"}, nil
	}
}

// 根据应用id查询用户权限及菜单数据
func (m menuService) FindAll(ctx context.Context, param *common.RoleMenuParam) (*common.MenuResponse, error) {
	menuSql := `SELECT M.menu_id,M.parent_id,M.icon,M.name,M.type,M.url,M.power_sign,M.sort_by FROM %v M WHERE M.del_status = 0 AND M.app_id = '%v' ORDER BY M.sort_by`
	menuSql = fmt.Sprintf(menuSql, enum.SysMenu.TableName, param.AppId)
	menuTreeList := make([]*common.Menu, 0)
	count := db.GetDB().Raw(menuSql).Scan(&menuTreeList).RowsAffected
	if len(menuTreeList) > 0 {
		return &common.MenuResponse{List: menuTreeList, Count: count}, nil
	} else {
		return &common.MenuResponse{Msg: "没有查询到该应用下的菜单数据"}, nil
	}
}
