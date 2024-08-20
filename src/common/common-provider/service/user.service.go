package service

import (
	"context"
	"core/config"
	"core/db"
	"core/exception"
	"fmt"
	"protobuf/build/common"
	"protobuf/enum"
)

// 定义helloService并实现约定的接口
// type helloService struct {
// 	pb.UnimplementedHelloServer
// }
// HelloService Hello服务
// var HelloService = helloService{}

type userService struct {
	common.UserServiceServer
}

var UserService = &userService{}

func (u *userService) Delete(ctx context.Context, ids *common.UserIds) (*common.UserResponse, error) {
	return &common.UserResponse{Count: db.GrpcBatchDeleteByIds(ids.UserId, enum.SysUser.TableName, enum.SysUser.PrimaryKey, ids.OperatorUserid)}, nil
}

// 用户新增
func (u *userService) Create(ctx context.Context, user *common.CreateUserDTO) (*common.UserResponse, error) {
	// 保存用户信息
	// 创建事务
	tx := db.DB.Begin()
	res := tx.Create(user.User)
	if res.Error != nil {
		// 回滚事务
		tx.Rollback()
		return &common.UserResponse{Count: 0, Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().CommonProvider)}, nil
	}
	if res.RowsAffected > 0 {
		// 查询角色信息
		var roleId string
		roleCount := db.DB.Raw(fmt.Sprintf(`SELECT role_id from sys_role WHERE app_id = '%s' ORDER BY unique_identification DESC LIMIT 1`, user.PowerParma.AppId)).Scan(&roleId).RowsAffected
		if roleCount > 0 {
			// 保存角色信息
			insertRoleSql := fmt.Sprintf("INSERT INTO `%v` (`role_id`,`user_id`,`app_id`,`create_by`) VALUES('%v','%v','%v', '%v')", enum.SysUserRole.TableName, roleId, user.User.UserId, user.PowerParma.AppId, user.User.UserId)
			insertRoleResult := tx.Exec(insertRoleSql)
			if insertRoleResult.Error != nil {
				// 回滚事务
				tx.Rollback()
				return &common.UserResponse{Count: 0, Msg: exception.DbMsg(insertRoleResult.Error.Error(), config.Global.ServerConfig().CommonProvider)}, nil
			}
			// 提交事务
			tx.Commit()
			// 角色存在 查询角色相关的权限信息存储在redis,便于权限验证
			return &common.UserResponse{Count: res.RowsAffected, PowerSign: u.FindPowerSign(user.PowerParma.AppId, user.User.UserId)}, nil
		} else {
			return &common.UserResponse{Count: 0, Msg: "该应用下未查询到任何角色信息"}, nil
		}
	}
	return &common.UserResponse{Count: 0, Msg: "用户创建是失败"}, nil
}

// 查询用户权限值
func (u *userService) FindPermissions(ctx context.Context, params *common.UserParam) (*common.UserResponse, error) {
	return &common.UserResponse{PowerSign: u.FindPowerSign(params.AppId, params.UserId)}, nil
}

// 用户修改
func (p *userService) Update(ctx context.Context, user *common.User) (*common.UserResponse, error) {
	return &common.UserResponse{Count: db.NewCrud(user, 2)}, nil
}

// 查询用户列表
func (p *userService) FindPageList(ctx context.Context, request *common.UserPageAuthQuery) (*common.UserResponse, error) {
	countSql := `SELECT COUNT(1) FROM ` + enum.SysUser.TableName + ` A`
	paramSql := ` WHERE A.del_status = 0 `
	if request.Params.UserName != "" {
		paramSql += fmt.Sprintf(` AND A.user_name like CONCAT('%%','%v','%%')`, request.Params.UserName)
	}
	if request.Params.Phone != "" {
		paramSql += fmt.Sprintf(` AND A.phone like CONCAT('%%','%v','%%')`, request.Params.Phone)
	}
	if request.Params.Email != "" {
		paramSql += fmt.Sprintf(` AND A.email like CONCAT('%%','%v','%%')`, request.Params.Email)
	}
	if request.Params.OpenId != "" {
		paramSql += fmt.Sprintf(` AND A.open_id like CONCAT('%%','%v','%%')`, request.Params.OpenId)
	}
	if request.Params.RegType != 0 {
		paramSql += fmt.Sprintf(` AND A.reg_type = %v`, request.Params.RegType)
	}
	var count int64
	var list []*common.UserListVo
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		listSql := `
		SELECT 
			A.user_id,
			A.user_name,
			A.user_password,
			A.salt,
			A.phone,
			A.email,
			A.reasons_prohibition,
			A.account_type,
			A.status,
			A.reg_type,
			A.source_id,
			A.del_status,
			A.create_by,
			DATE_FORMAT(A.create_time,'%Y-%m-%d') create_time
		FROM
			` + enum.SysUser.TableName + ` A`
		listSql += paramSql
		listSql += ` LIMIT ?,?`
		db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&list)
	}
	return &common.UserResponse{Count: count, List: list}, nil
}

// 根据用户id查询详情
func (p *userService) FindById(ctx context.Context, ids *common.UserIds) (*common.UserResponse, error) {
	user := common.UserBaseDetail{}
	count := db.DB.First(&user, ids.UserId[0]).RowsAffected
	return &common.UserResponse{Detail: &user, Count: count}, nil
}

// 获取本人基本信息
func (p *userService) FindBalance(ctx context.Context, ids *common.UserIds) (*common.UserResponse, error) {
	balance := "0"
	count := db.DB.Raw("SELECT balance FROM " + enum.SysUserDetail.TableName + " WHERE user_id = " + ids.UserId[0]).Scan(&balance).RowsAffected
	return &common.UserResponse{Count: count, Balance: balance}, nil
}

// 根据用户名或者手机号查询用户信息
func (p *userService) FindByUserParam(ctx context.Context, params *common.UserParam) (*common.UserResponse, error) {
	if params != nil {
		if params.UserName == "" && params.Phone == "" && params.OpenId == "" && params.Email == "" {
			return &common.UserResponse{Detail: nil, Count: 0, Msg: "未传入核心参数！"}, nil
		}
		if params.AppId == "" {
			return &common.UserResponse{Detail: nil, Count: 0, Msg: "未传入AppId参数！"}, nil
		}
		// 查询应用是否存在
		var appCount int
		findAppSql := fmt.Sprintf(`SELECT count(app_id) from %v WHERE app_id = '%v'`, enum.SysApp.TableName, params.AppId)
		db.DB.Raw(findAppSql).Scan(&appCount)
		if appCount == 0 {
			return &common.UserResponse{Detail: nil, Count: 0, Msg: "应用不存在,请联系管理员处理！"}, nil
		}

		// 查询角色信息
		roles := make([]*common.Role, 0)
		findRoleSql := fmt.Sprintf(`SELECT role_id from %v WHERE app_id = '%v'`, enum.SysRole.TableName, params.AppId)
		rolesCount := db.DB.Raw(findRoleSql).Scan(&roles).RowsAffected
		if rolesCount == 0 {
			return &common.UserResponse{Detail: nil, Count: 0, Msg: "应用下不存在角色信息,请联系管理员处理！", Roles: roles}, nil
		}

		// 查询用户信息
		findUserSql := `
			SELECT 
				A.user_id,
				A.user_name,
				A.user_password,
				A.salt,
				A.open_id,
				A.phone,
				A.email,
				A.reasons_prohibition,
				A.account_type,
				A.status,
				A.reg_type,
				A.source_id,
				A.del_status,
				A.create_by,
				DATE_FORMAT(A.create_time,'%Y-%m-%d') create_time,`
		findUserSql += fmt.Sprintf(`(SELECT GROUP_CONCAT(unique_identification) FROM %v CR,%v CUR WHERE CUR.user_id = A.user_id AND CUR.role_id = CR.role_id ) as unique_identification `, enum.SysRole.TableName, enum.SysUserRole.TableName)
		findUserSql += fmt.Sprintf(` FROM  %v A WHERE A.del_status = 0 `, enum.SysUser.TableName)
		// 两个参数
		if params.OpenId != "" {
			findUserSql += fmt.Sprintf(` AND A.open_id ='%v'`, params.OpenId)
		}
		if params.Email != "" {
			findUserSql += fmt.Sprintf(` AND A.email ='%v'`, params.Email)
		}
		if params.UserName != "" {
			findUserSql += fmt.Sprintf(` AND A.user_name ='%v'`, params.UserName)
		}
		if params.Phone != "" {
			findUserSql += fmt.Sprintf(` AND A.phone ='%v'`, params.Phone)
		}
		user := common.UserBaseDetail{}
		res := db.DB.Raw(findUserSql).Scan(&user)
		// 判断查询是否异常
		if res.Error != nil {
			return &common.UserResponse{Detail: nil, Count: 0, Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().CommonProvider), Roles: roles}, nil
		}
		// 判断用户信息
		if res.RowsAffected <= 0 {
			return &common.UserResponse{Detail: nil, Count: 0, Msg: "未查询到您的账号信息！请检查是否正确，或者换其他登陆方式", Roles: roles}, nil
		}
		// 判断角色信息
		if user.UniqueIdentification == "" {
			return &common.UserResponse{Detail: nil, Count: 0, Msg: "您的账号不存在角色信息,请联系管理员处理！", Roles: roles}, nil
		}
		// 查询权限sql
		sqlPowerSign := `
		SELECT CONCAT_WS(",",
			(SELECT GROUP_CONCAT(FN.power_sign) as power_sign  from sys_func FN,sys_role_func RF WHERE FN.func_id = RF.func_id AND RF.role_id IN (
			(SELECT UR.role_id FROM sys_role R,sys_user_role UR  WHERE R.role_id = UR.role_id AND UR.app_id = '%v' AND UR.user_id = '%v')
			)),
			(SELECT ifnull(c.power_sign, '') from (SELECT GROUP_CONCAT(M.power_sign)  as power_sign  from sys_menu M,sys_role_menu RM WHERE RM.menu_id = M.menu_id AND RM.role_id IN 
			(SELECT UR.role_id FROM sys_role R,sys_user_role UR  WHERE R.role_id = UR.role_id AND UR.app_id = '%v' AND UR.user_id = '%v')) as c
			)
		) as power_sign`
		sqlPowerSign = fmt.Sprintf(sqlPowerSign, params.AppId, user.UserId, params.AppId, user.UserId)
		var powerSign string
		db.DB.Raw(sqlPowerSign).Scan(&powerSign)
		return &common.UserResponse{Detail: &user, Count: res.RowsAffected, Roles: roles, PowerSign: powerSign}, nil
	}
	return &common.UserResponse{Msg: "必传参数为空"}, nil
}

// FindPowerSign 根据应用i和用户id查询权限值
func (u *userService) FindPowerSign(appId string, userId string) string {
	sqlPowerSign := `
	SELECT CONCAT_WS(",",
	(SELECT GROUP_CONCAT(FN.power_sign) as power_sign  from sys_func FN,sys_role_func RF WHERE FN.func_id = RF.func_id AND RF.role_id IN (
	(SELECT UR.role_id FROM sys_role R,sys_user_role UR  WHERE R.role_id = UR.role_id AND UR.app_id = '%v' AND UR.user_id = '%v'))),
	(SELECT GROUP_CONCAT(M.power_sign)  as power_sign  from sys_menu M,sys_role_menu RM WHERE RM.menu_id = M.menu_id AND RM.role_id IN 
	(SELECT UR.role_id FROM sys_role R,sys_user_role UR  WHERE R.role_id = UR.role_id AND UR.app_id = '%v' AND UR.user_id = '%v'))) as power_sign`
	sqlPowerSign = fmt.Sprintf(sqlPowerSign, appId, userId, appId, userId)
	var powerSign string
	db.DB.Raw(sqlPowerSign).Scan(&powerSign)
	return powerSign
}
