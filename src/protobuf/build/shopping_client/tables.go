package shopping_client

import "protobuf/enum"

// 购物车表映射
func (c *Cart) TableName() string {
	return enum.Cart.TableName
}
