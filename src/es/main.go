package main

import (
	"es/config"
	"es/mapping"
	"fmt"
)

func main() {

	client, err := config.EsService.CreateClient()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("连接成功: ", client)
	}

	// // 创建索引 /{indexName}
	// if index, err := config.EsService.CreateIndex(mapping.UserInfo.GetIndexName(), &mapping.UserInfo); err != nil {
	// 	fmt.Println(err.Error())
	// 	fmt.Println("===>", index)
	// }

	// mapping.UserInfo.Age = "12"
	// mapping.UserInfo.UserId = 1
	// mapping.UserInfo.UserName = "sb1"

	// 保存数据 /{indexName}/_doc/1
	// if res, err := config.EsService.Save(mapping.UserInfo.GetIndexName(), "1", mapping.UserInfo); err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("%v", res)
	// }

	if res, err := config.EsService.SearchByDocId(mapping.UserInfo.GetIndexName(), "1", mapping.UserInfo); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}

}
