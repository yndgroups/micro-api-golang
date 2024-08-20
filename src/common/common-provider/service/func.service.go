package service

import (
	"context"
	"core/config"
	"core/db"
	"fmt"
	"protobuf/build/common"
	"protobuf/enum"
	"strings"
)

var SysFuncService = &sysFuncService{}

type sysFuncService struct {
	common.SysFuncServiceServer
}

// 创建功能
func (s sysFuncService) Create(cxt context.Context, param *common.SysFunc) (*common.SysFuncResponse, error) {
	return &common.SysFuncResponse{Count: db.NewCrud(param, 1)}, nil
}

// 修改功能
func (s sysFuncService) Update(cxt context.Context, param *common.SysFunc) (*common.SysFuncResponse, error) {
	return &common.SysFuncResponse{Count: db.NewCrud(param, 2)}, nil
}

// 删除功能
func (s sysFuncService) Delete(cxt context.Context, ids *common.SysFuncIds) (*common.SysFuncResponse, error) {
	return &common.SysFuncResponse{Count: db.GrpcBatchDeleteByIds(ids.FuncIds, enum.SysFunc.TableName, enum.SysFunc.PrimaryKey, ids.UserId)}, nil
}

// 删除功能
func (s sysFuncService) FindById(cxt context.Context, ids *common.SysFuncIds) (*common.SysFuncResponse, error) {
	return &common.SysFuncResponse{Count: db.GrpcBatchDeleteByIds(ids.FuncIds, enum.SysFunc.TableName, enum.SysFunc.PrimaryKey, ids.UserId)}, nil
}

func (s sysFuncService) FindList(cxt context.Context, request *common.SysFuncAuthQuery) (*common.SysFuncResponse, error) {
	if !strings.Contains(config.Global.ServerConfig().Administrator, request.Auth.UserId) { // 非超级用户
		roleIds := make([]string, 0)
		db.GetDB().Raw("SELECT role_id FROM sys_user_role WHERE app_id = ? AND user_id = ?", request.Auth.AppId, request.Auth.UserId).Scan(&roleIds)
		if len(roleIds) == 0 {
			return &common.SysFuncResponse{Msg: "不存在您的的角色所关联的功能信息"}, nil
		}
		// 查询角色关联的功能id集合
		funcIds := make([]string, 0)
		db.GetDB().Raw(`SELECT DISTINCT func_id from sys_role_menu WHERE role_id in ?`, roleIds).Scan(&funcIds)
		if len(funcIds) == 0 {
			return &common.SysFuncResponse{Msg: "该账号对应下的角色未关联任何功能信息"}, nil
		}
		// 查询功能
		menuSql := `SELECT M.func_id,M.parent_id,M.name,M.power_sign,M.sort_by FROM ` + enum.SysFunc.TableName + ` M WHERE M.del_status = 0 AND M.func_id in ? ORDER BY M.sort_by`
		menuTreeList := make([]*common.SysFuncListVo, 0)
		count := db.GetDB().Raw(menuSql, funcIds).Scan(&menuTreeList).RowsAffected
		if len(menuTreeList) > 0 {
			return &common.SysFuncResponse{List: menuTreeList, Count: count}, nil
		} else {
			return &common.SysFuncResponse{Msg: "没有查询到该应用下的功能数据！"}, nil
		}
	} else {
		// 超级用户查询功能
		menuSql := `SELECT M.func_id,M.parent_id,M.name,M.power_sign,M.sort_by FROM  ` + enum.SysFunc.TableName + `  M WHERE M.del_status = 0 AND M.app_id = ? ORDER BY M.sort_by`
		funcTreeList := make([]*common.SysFuncListVo, 0)
		count := db.GetDB().Raw(menuSql, request.Auth.AppId).Scan(&funcTreeList).RowsAffected
		if len(funcTreeList) > 0 {
			return &common.SysFuncResponse{List: funcTreeList, Count: count}, nil
		} else {
			return &common.SysFuncResponse{Msg: "没有查询到该应用下的功能数据！"}, nil
		}
	}
}

// 根据角色查询当前应用下的功能及已经授权了的功能id信息
func (m sysFuncService) FindFuncInfoByRoleIds(ctx context.Context, param *common.RoleFuncParam) (*common.SysFuncResponse, error) {
	funcSql := fmt.Sprintf(`SELECT M.func_id,M.parent_id,M.name,M.power_sign,M.sort_by FROM sys_func M,sys_role R  WHERE M.del_status = 0 AND M.app_id = R.app_id and R.role_id = '%v' ORDER BY sort_by`, param.RoleId)
	var list []*common.SysFuncListVo
	count := db.GetDB().Raw(funcSql).Scan(&list).RowsAffected
	if len(list) > 0 {
		roleMenuSql := fmt.Sprintf(`SELECT func_id from sys_role_func WHERE role_id = '%v'`, param.RoleId)
		var funcIds []string
		db.GetDB().Raw(roleMenuSql).Scan(&funcIds)
		return &common.SysFuncResponse{List: list, Count: count, RoleFuncIds: funcIds}, nil
	} else {
		return &common.SysFuncResponse{Msg: "该角色还未授权任何功能数据！"}, nil
	}
}
