package order

import "protobuf/enum"

func (Order) TableName() string {
	return enum.Order.TableName
}

func (OrderDetail) TableName() string {
	return enum.OrderDetail.TableName
}

func (FreightTemplates) TableName() string {
	return enum.FreightTemplates.TableName
}

func (FreightTemplatesFree) TableName() string {
	return enum.FreightTemplatesFree.TableName
}

func (FreightTemplatesRegion) TableName() string {
	return enum.FreightTemplatesRegion.TableName
}

func (FreightTemplatesUndelivered) TableName() string {
	return enum.FreightTemplatesUndelivered.TableName
}

func (PocketMoney) TableName() string {
	return enum.PocketMoney.TableName
}
