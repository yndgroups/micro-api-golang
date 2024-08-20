package enum

import "core/enum"

var (

	// 用户
	SysUser = enum.TableEnum("sys_user", "user_id")

	// 用户基础信息
	SysUserDetail = enum.TableEnum("sys_user_detail", "user_id")

	// 角色
	SysRole = enum.TableEnum("sys_role", "role_id")

	// 协议
	SysAgreement = enum.TableEnum("sys_agreement", "agree_id")

	// 应用
	SysApp = enum.TableEnum("sys_app", "app_id")

	// 应用模块
	SysAppModule = enum.TableEnum("sys_app_module", "md_id")

	// 地区信息表
	SysArea = enum.TableEnum("sys_area", "area_id")

	// 系统配置表
	SysConfig = enum.TableEnum("sys_conf", "conf_id")

	// 客户表
	SysCustomer = enum.TableEnum("sys_customer", "customer_id")

	// 字典表
	SysDict = enum.TableEnum("sys_dict", "dict_id")

	// 枚举表
	SysEnum = enum.TableEnum("sys_enum", "enum_id")

	// 菜单表
	SysMenu = enum.TableEnum("sys_menu", "menu_id")

	// 功能
	SysFunc = enum.TableEnum("sys_func", "func_id")

	// 小程序表
	SysMiniApp = enum.TableEnum("sys_mini_app", "mini_app_id")

	// 收货地址
	SysUserAddress = enum.TableEnum("sys_user_address", "address_id")

	// 用户应用表
	UserApp = enum.TableEnum("sys_user_app", "user_id")

	// 站点声明表
	SysWebsiteStated = enum.TableEnum("sys_website_stated", "state_id")

	// 角色菜单表
	SysRoleMenu = enum.TableEnum("sys_role_menu", "role_id")

	// 角色功能表
	SysRoleFunc = enum.TableEnum("sys_role_func", "func_id")

	// 模板表
	Template = enum.TableEnum("sys_template", "template_id")

	// 用户组
	SysUserGroup = enum.TableEnum("sys_user_group", "user_group_id")

	// 用户角色表
	SysUserRole = enum.TableEnum("sys_user_role", "user_id")

	// 部门表
	SysDepartment = enum.TableEnum("sys_department", "dept_id")

	// 岗位
	SysPost = enum.TableEnum("sys_post", "post_id")

	// 组织机构表
	SysOrg = enum.TableEnum("sys_org", "org_id")

	// 账本
	AccountBook = enum.TableEnum("ynd_account_book", "book_id")

	// 账本记录
	AccountUser = enum.TableEnum("ynd_account_user", "record_id")

	// 在线充值 ========================== 订单服务 =====================================
	PocketMoney = enum.TableEnum("ynd_pocket_money", "pm_id")

	// 订单
	Order = enum.TableEnum("ynd_order", "order_id")

	// 订单详情
	OrderDetail = enum.TableEnum("ynd_order_detail", "detail_id")

	// 模板
	FreightTemplates = enum.TableEnum("ynd_freight_templates", "temp_id")

	// 免邮费
	FreightTemplatesFree = enum.TableEnum("ynd_freight_templates_free", "free_id")

	// 区域运费
	FreightTemplatesRegion = enum.TableEnum("ynd_freight_templates_region", "region_id")

	// 不送达区域
	FreightTemplatesUndelivered = enum.TableEnum("ynd_freight_templates_undelivered", "undelivered_id")

	// 商家表
	Business = enum.TableEnum("ynd_business", "business_id")

	// 广告表
	Ad = enum.TableEnum("ynd_ad", "ad_id")

	// 商家加盟表
	BusinessFranchisee = enum.TableEnum("ynd_business_franchisee", "franchisee_id")

	// 商圈
	BusinessCircle = enum.TableEnum("ynd_business_circle", "circle_id")

	// 商家提现申请
	BusinessApply = enum.TableEnum("ynd_business_apply", "apply_id")

	// 商家余额
	BusinessBalance = enum.TableEnum("ynd_business_balance", "business_id")

	// 商家余额提现申请
	BusinessBalanceApply = enum.TableEnum("t_business_balance_apply", "balance_apply_id")

	// 商家门店
	BusinessStore = enum.TableEnum("ynd_business_store", "store_id")

	// 商品
	Product = enum.TableEnum("ynd_product", "product_id")

	// 商品详情
	ProductDetail = enum.TableEnum("ynd_product_detail", "product_id")

	// 商品
	ProductSku = enum.TableEnum("ynd_product_sku", "sku_id")

	// 商品分类
	ProductCategory = enum.TableEnum("ynd_product_category", "product_category_id")

	// 品牌
	Brand = enum.TableEnum("ynd_brand", "brand_id")

	// 商品分类品牌
	ProductCategoryBrand = enum.TableEnum("ynd_product_category_brand", "brand_id")

	// 商品规格
	ProductSpec = enum.TableEnum("ynd_product_spec", "spec_id")

	// 应用版本
	AppVersion = enum.TableEnum("sys_agreement", "agree_id")

	// 快递
	Express = enum.TableEnum("ynd_express", "express_id")

	// 店铺
	Store = enum.TableEnum("ynd_store", "store_id")

	// 优惠券
	Coupon = enum.TableEnum("ynd_coupon", "coupon_id")

	// 店铺员工
	StoreStaff = enum.TableEnum("ynd_store_staff", "staff_id")

	// 购物车
	Cart = enum.TableEnum("ynd_cart", "cart_id")
)
