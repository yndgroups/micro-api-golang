package shopping_admin

import "protobuf/enum"

func (Product) TableName() string {
	return enum.Product.TableName
}

func (ProductDetail) TableName() string {
	return enum.ProductDetail.TableName
}
