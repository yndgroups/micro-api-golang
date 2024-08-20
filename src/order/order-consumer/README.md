# GO语言微服务实

# 代理设置
- File > Settings > Go > Go Modules
  在 Environment 输入如下
``` 
GOPROXY=https://goproxy.cn,direct
```

## 使用 fresh 实现热部署

- 安装 fresh
``` 
go get github.com/pilu/fresh
```

- 跳转到项目目录,例如项目名为`myapp`
```
cd projectDir
```

- 启动
``` 
fresh
```

## swagger 文档生产命令
```
swag init -d ./,../../core,../../protos --exclude ./param/,./mapper/,../../core/pager,../../core/validate
```

``` 
// getProductList 获取商品数据
func getProductList(products []model.ProductParam, accessToken string) ([]model.ProductExpand, error) {
	if len(products) < 0 {
		return nil, errors.New("参数不能为空")
	}
	list, err := req.MicroServiceHttpRequest[[]model.ProductExpand](model.MicroServiceHttpParam[any]{
		ServiceName: viper.GetString("api.shopping.server-name"),
		HttpMethod:  viper.GetString("api.http-method"),
		ApiPath:     viper.GetString("api.shopping.findProductsByProductIds"),
		Method:      "POST",
		AccessToken: accessToken,
		ReqParam:    products,
	})
	return list, err
}

// 保存商品信息
func saveOrder(productList []core.ProductExpand, detail model.CreateOrderInfo) error {
	payPrice, _ := decimal.NewFromString(detail.PayPrice) //  变成分 乘100
	rmb, _ := decimal.NewFromString("100")                //  变成分 乘100
	order := model.Order{
		OrderId:       redis.GET.GetPrimaryKey("SYS_ORDER_REQUEST_COUNT"),
		ParentId:      "1",
		ConfId:        detail.ConfId,
		OrderNo:       detail.TradeNo, // 订单号
		TradeType:     1,
		TradeNo:       detail.PayNo,
		TotalPrice:    detail.TotalPrice,
		PayPrice:      ksii.ChangeString(payPrice.Div(rmb)),
		TotalQuantity: detail.ProductTotalQuantity,
		PayStatus:     1,
		PayType:       1,
		Status:        1,
	}
	order.CreateBy = detail.UserId
	tx := db.DB.Begin()
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return errors.New("下单失败，请重试！")
	}
	tx.Commit()
	return nil
}
```
