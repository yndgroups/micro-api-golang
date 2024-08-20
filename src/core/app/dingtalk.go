package app

import (
	"core/config"
	"fmt"
	"protobuf/build/order"

	"github.com/blinkbean/dingtalk"
)

var DD *dingtalk.DingTalk

// InitDingtalk 初始化
func InitDingtalk(config config.Dingtalk) {
	DD = dingtalk.InitDingTalkWithSecret(config.Token, config.Secret)
}

// SendText 文本发送
func SendText(msg string) {
	DD.SendTextMessage(msg, dingtalk.WithAtAll())
}

// SendMap 发送消息到钉钉 按行写入数组，增强markdown的可读性
func SendMap(maps map[string]string) {
	msg := []string{
		"### 新订单提醒",
		"---",
		"- 商品名称：<font color=#ff0000 size=6>" + maps["productName"] + "</font>",
		"- 商品单价：<font color=#ff0000 size=6>" + maps["price"] + "元</font>",
		"- 商品单号：<font color=#ff0000 size=6>" + maps["orderNo"] + "</font>",
		"- 客户姓名：<font color=#ff0000 size=6>" + maps["customerName"] + "</font>",
		"- 联系电话：<font color=#ff0000 size=6>" + maps["phpne"] + "</font>",
		"- 收货地址：<font color=#ff0000 size=6>" + maps["address"] + "</font>",
	}
	DD.SendMarkDownMessageBySlice("新订单提醒", msg, dingtalk.WithAtAll())
}

// ProductItem 商品项
type ProductItem struct {
	// 商品名称
	ProductName string `json:"productName"`
	// 商品数量
	Quantity int64 `json:"quantity"`
	// 价格
	Price string `json:"price"`
}

// OrderMessage 订单消息推送模版
type OrderMessage struct {
	// 收货姓名
	CustomerName string `json:"customerName"`
	// 收货电话
	Phone string `json:"phone"`
	// 收货地址
	Address string `json:"address"`
	// 商品信息
	ProductList []*order.OrderMsg `json:"productList"`
}

// SendOrderMessage 发送订单MarkDown消息
func SendOrderMessage(list []*order.OrderMsg) {
	if len(list) == 0 {
		SendText("商品查询失败，请去后台查看是爱原因")
		return
	}
	dm := dingtalk.DingMap()
	dm.Set("新订单提醒", dingtalk.H2)
	dm.Set("---", dingtalk.N)
	for i, v := range list {
		dm.Set(fmt.Sprintf("商品%d: $$ %v $$", i+1, v.ProductName), dingtalk.RED)
		dm.Set(fmt.Sprintf("数  量: $$ %v $$", i+1), dingtalk.RED)
		//dm.Set(fmt.Sprintf("价  格: $$ %v $$", i+1), dingtalk.RED)
		dm.Set("---", dingtalk.N)
	}
	dm.Set(fmt.Sprintf("客户姓名: %v", list[0].RealName), dingtalk.N)
	dm.Set(fmt.Sprintf("联系电话: %v", list[0].UserPhone), dingtalk.N)
	dm.Set(fmt.Sprintf("收货地址: $$ %v $$", list[0].UserAddress), dingtalk.RED)
	DD.SendMarkDownMessageBySlice("新订单提醒", dm.Slice())
}
