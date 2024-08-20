package common

import "protobuf/enum"

func (User) TableName() string {
	return enum.SysUser.TableName
}

func (UserDetail) TableName() string {
	return enum.SysUserDetail.TableName
}

func (Agreement) TableName() string {
	return enum.SysAgreement.TableName
}

func (App) TableName() string {
	return enum.SysApp.TableName
}

func (AppModule) TableName() string {
	return enum.SysAppModule.TableName
}

func (Area) TableName() string {
	return enum.SysArea.TableName
}

func (Config) TableName() string {
	return enum.SysConfig.TableName
}

func (Customer) TableName() string {
	return enum.SysCustomer.TableName
}

func (Dict) TableName() string {
	return enum.SysDict.TableName
}

func (Enum) TableName() string {
	return enum.SysEnum.TableName
}

func (Menu) TableName() string {
	return enum.SysMenu.TableName
}

func (MiniApp) TableName() string {
	return enum.SysMiniApp.TableName
}

func (Role) TableName() string {
	return enum.SysRole.TableName
}

func (UserAddress) TableName() string {
	return enum.SysUserAddress.TableName
}

func (WebsiteStated) TableName() string {
	return enum.SysWebsiteStated.TableName
}

func (Org) TableName() string {
	return enum.SysOrg.TableName
}

func (UserRole) TableName() string {
	return enum.SysUserRole.TableName
}

func (Department) TableName() string {
	return enum.SysDepartment.TableName
}

func (Post) TableName() string {
	return enum.SysPost.TableName
}

func (SysFunc) TableName() string {
	return enum.SysFunc.TableName
}
