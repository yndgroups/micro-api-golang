package enum

import "core/enum"

var (
	// Business 商家表
	Business = enum.TableEnum("ynd_business", "business_id")
	// Ad 广告表
	Ad = enum.TableEnum("ynd_ad", "ad_id")
	// BusinessFranchisee 商家加盟表
	BusinessFranchisee = enum.TableEnum("ynd_business_franchisee", "franchisee_id")
	// BusinessCircle 商圈
	BusinessCircle = enum.TableEnum("ynd_business_circle", "circle_id")
	// BusinessApply 商家提现申请
	BusinessApply = enum.TableEnum("ynd_business_apply", "apply_id")
	// BusinessBalance 商家余额
	BusinessBalance = enum.TableEnum("ynd_business_balance", "business_id")
	// BusinessBalanceApply 商家余额提现申请
	BusinessBalanceApply = enum.TableEnum("t_business_balance_apply", "balance_apply_id")
	// BusinessStore 商家门店
	BusinessStore = enum.TableEnum("ynd_business_store", "store_id")
	// Product 商品
	Product = enum.TableEnum("ynd_product", "product_id")
	// ProductDetail 商品详情
	ProductDetail = enum.TableEnum("ynd_product_detail", "product_id")
	// ProductSku 商品
	ProductSku = enum.TableEnum("ynd_product_sku", "sku_id")
	// ProductCategory 商品分类
	ProductCategory = enum.TableEnum("ynd_product_category", "product_category_id")
	// Brand 品牌
	Brand = enum.TableEnum("ynd_brand", "brand_id")
	// ProductCategoryBrand 商品分类品牌
	ProductCategoryBrand = enum.TableEnum("ynd_product_category_brand", "brand_id")
	// ProductSpec 商品规格
	ProductSpec = enum.TableEnum("ynd_product_spec", "spec_id")
	// AppVersion 应用版本
	AppVersion = enum.TableEnum("sys_agreement", "agree_id")
	// Express 快递
	Express = enum.TableEnum("ynd_express", "express_id")
	// Store 店铺
	Store = enum.TableEnum("ynd_store", "store_id")
	// Coupon 优惠券
	Coupon = enum.TableEnum("ynd_coupon", "coupon_id")
)
