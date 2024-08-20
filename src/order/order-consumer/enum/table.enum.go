package enum

import "core/enum"

var (
	// OnlineRecharge 在线充值
	OnlineRecharge = enum.TableEnum("ynd_online_recharge", "or_id")
	// User 订单用户信息
	User = enum.TableEnum("ynd_user", "user_id")
	// Order 订单
	Order = enum.TableEnum("ynd_order", "order_id")
	// FreightTemplates 模板
	FreightTemplates = enum.TableEnum("ynd_freight_templates", "temp_id")
	// FreightTemplatesFree 免邮费
	FreightTemplatesFree = enum.TableEnum("ynd_freight_templates_free", "free_id")
	// FreightTemplatesRegion 区域运费
	FreightTemplatesRegion = enum.TableEnum("ynd_freight_templates_region", "region_id")
	// FreightTemplatesUndelivered 不送达区域
	FreightTemplatesUndelivered = enum.TableEnum("ynd_freight_templates_undelivered", "undelivered_id")
)
