package test

import (
	"testing"

	"github.com/blinkbean/dingtalk"
)

func TestDinGtalkSendMessage(t *testing.T) {
	var tokens = []string{"cd6af33e4ffd4756ec1aaf9e0e71e0e349db978cedb0e20e425a18641d47eb08"}
	cli := dingtalk.InitDingTalk(tokens, "Order")
	cli.SendTextMessage("您有一个新订单，测试", dingtalk.WithAtAll())
	print(cli)
}
