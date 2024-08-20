package enum

import "core/enum"

var (
	// NoAuthorized 不需要验证，即开放权限接口
	NoAuthorized = enum.ApiEnum("", "")
	// CreateApp 应用添加
	CreateApp = enum.ApiEnum("CreateApp", "system:common:CreateApp")
	// UpdateApp 应用修改
	UpdateApp = enum.ApiEnum("UpdateApp", "system:common:updateApp")
	// DeleteApp 应用删除
	DeleteApp = enum.ApiEnum("DeleteApp", "system:common:deleteApp")
)
