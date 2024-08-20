package enum

import "core/enum"

var (
	// NoAuthorized 不需要验证，即开放权限接口
	NoAuthorized = enum.ApiEnum("", "")
	// CreateAd 广告添加
	CreateAd = enum.ApiEnum("CreateAd", "admin:shopping:createAd")
	// UpdateAd 广告修改
	UpdateAd = enum.ApiEnum("UpdateAd", "admin:shopping:updateAd")
	// DeleteAd 广告删除
	DeleteAd = enum.ApiEnum("DeleteAd", "admin:shopping:deleteAd")
)
