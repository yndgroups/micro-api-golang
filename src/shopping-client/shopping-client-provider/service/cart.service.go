package service

import (
	"context"
	"core/config"
	"core/db"
	"core/exception"
	"fmt"
	"protobuf/build/shopping_client"
	"protobuf/enum"

	"gorm.io/gorm"
)

var CartService = &cartService{}

type cartService struct {
	shopping_client.CartServiceServer
}

// Create 添加购物车
func (c *cartService) Create(ctx context.Context, cart *shopping_client.Cart) (*shopping_client.CartResponse, error) {
	var count int64 = 0
	sql := fmt.Sprintf(`SELECT COUNT(1) FROM %s WHERE product_id = %s AND create_by = '%v'`, enum.Cart.TableName, cart.ProductId, cart.CreateBy)
	db.DB.Raw(sql).Scan(&count)
	var result *gorm.DB
	if count == 0 {
		result = db.NewCrudAndMessage(cart, 1)
	} else {
		updateSql := fmt.Sprintf(`UPDATE %s SET quantity = quantity + %v WHERE product_id = '%v' AND create_by = '%v'`, enum.Cart.TableName, cart.Quantity, cart.ProductId, cart.CreateBy)
		result = db.DB.Exec(updateSql)
	}
	if result.Error != nil {
		return &shopping_client.CartResponse{Msg: exception.DbMsg(result.Error.Error(), config.Global.ServerConfig().ShoppingClientProvider)}, nil
	}
	return &shopping_client.CartResponse{Count: result.RowsAffected}, nil
}

// Delete 删除购物车
func (c *cartService) Delete(ctx context.Context, ids *shopping_client.CartIds) (*shopping_client.CartResponse, error) {
	return &shopping_client.CartResponse{Count: db.GrpcBatchDeleteByIds(ids.CartIds, enum.Cart.TableName, enum.Cart.PrimaryKey, ids.UserId)}, nil
}

// FindPageList 查询购物车列表
func (c *cartService) FindPageList(ctx context.Context, request *shopping_client.CartPageAuthQuery) (*shopping_client.CartResponse, error) {
	count := c.count(request.UserId)
	if count > 0 {
		sql := `
SELECT
	C.cart_id,
	C.product_id,
	C.spec_type,
	C.product_attr_unique,
	C.quantity,
	C.status,
	C.sort_by,
	C.create_by,
	C.update_by,
	DATE_FORMAT(C.create_time,'%Y-%m-%d %T') create_time,
	C.update_time,`
		s1 := `
	(IF (C.spec_type = 1,
	(SELECT product_name FROM %v WHERE product_id = C.product_id),
	(SELECT P.product_name FROM %v P, %v PS WHERE PS.product_id = P.product_id AND PS.sku_id = C.product_id))) as product_name,
	(IF (C.spec_type = 1,
	(SELECT image FROM %v WHERE product_id = C.product_id),
	(SELECT image FROM %v WHERE sku_id = C.product_id))) as image,
	(IF (C.spec_type = 1,
	(SELECT price FROM %v WHERE product_id = C.product_id),
	(SELECT price FROM %v WHERE sku_id = C.product_id))) as price
FROM %v C WhERE C.del_status = 0 AND C.create_by = '%v' ORDER BY C.create_time LIMIT %v,%v`
		sql += fmt.Sprintf(s1,
			enum.Product.TableName,
			enum.Product.TableName,
			enum.ProductSku.TableName,
			enum.Product.TableName,
			enum.ProductSku.TableName,
			enum.Product.TableName,
			enum.ProductSku.TableName,
			enum.Cart.TableName,
			request.UserId,
			(request.PageIndex-1)*request.PageSize,
			request.PageSize)
		list := make([]*shopping_client.CartListVo, 0)
		res := db.DB.Raw(sql).Scan(&list)
		if res.Error != nil {
			return &shopping_client.CartResponse{Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().ShoppingClientProvider)}, nil
		}
		return &shopping_client.CartResponse{List: list, Count: count}, nil
	}
	return &shopping_client.CartResponse{Count: count}, nil
}

// FindCount 查询购物车数量
func (c *cartService) FindCount(ctx context.Context, ids *shopping_client.CartIds) (*shopping_client.CartResponse, error) {
	return &shopping_client.CartResponse{Count: c.count(ids.UserId)}, nil
}

// ChangeQuantity 修改购物车商品数量
func (c *cartService) ChangeQuantity(ctx context.Context, param *shopping_client.CartIds) (*shopping_client.CartResponse, error) {
	sql := fmt.Sprintf(`
	SELECT
	if(C.spec_type = 1, 
	(SELECT P.stock FROM %v P, %v C2 WHERE  P.product_id = C2.product_id AND C2.cart_id = '%v'), 
	(SELECT PS.stock FROM %v PS, %v C1 WHERE PS.sku_id = C1.product_id AND C.cart_id = '%v')) as stockNum FROM %v C WHERE C.cart_id = '%v'`,
		enum.Product.TableName,
		enum.Cart.TableName,
		param.CartIds[0],
		enum.ProductSku.TableName,
		enum.Cart.TableName,
		param.CartIds[0],
		enum.Cart.TableName,
		param.CartIds[0],
	)
	var stockNum int64
	db.DB.Raw(sql).Scan(&stockNum)
	if param.Quantity > stockNum {
		return &shopping_client.CartResponse{Count: stockNum, Msg: fmt.Sprintf("当前库存不足，最大数量只能为: %d", stockNum)}, nil
	} else {
		sqlUpdate := fmt.Sprintf(`UPDATE ynd_cart SET quantity = '%v' WHERE cart_id = '%v' AND create_by = '%v'`, param.Quantity, param.CartIds[0], param.UserId)
		return &shopping_client.CartResponse{Count: db.DB.Exec(sqlUpdate).RowsAffected}, nil
	}
}

// count 查询用户的购物车商品数量
func (c *cartService) count(userId string) int64 {
	sql := fmt.Sprintf(`SELECT COUNT(1) FROM %s WHERE del_status = 0 AND create_by = %s`, enum.Cart.TableName, userId)
	var count int64
	db.DB.Raw(sql).Scan(&count)
	return count
}
