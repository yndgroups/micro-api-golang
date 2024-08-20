package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
)

type esService struct {
	EsClient *elastic.Client
}

var EsService = &esService{}

// CreateClient 创建一个客户端
func (es *esService) CreateClient() (esc *esService, err error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://10.211.55.10:9200"),
		elastic.SetBasicAuth("elastic", "123456"), // 用户名和密码
		elastic.SetGzip(true),
		elastic.SetSniff(false), // 处理elastic ip地址和本机地址不一致的问题
		elastic.SetHealthcheck(true),
		elastic.SetHealthcheckTimeout(100000*time.Second),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC", log.LstdFlags)), // 设置日志输出的名字
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),         // 输出日志级别
	)
	if err != nil {
		return nil, errors.New("连接es失败")
	}
	es.EsClient = client
	return es, nil
}

// CreateIndex 创建索引库
func (esc *esService) CreateIndex(indexName string, mapping any) (bool, error) {
	exists, err := esc.EsClient.IndexExists(indexName).Do(context.Background())
	if err != nil {
		return false, err
	}

	if !exists {
		if res, ok := mapping.(string); ok {
			if do, err := esc.EsClient.CreateIndex(indexName).BodyString(res).Do(context.Background()); err != nil {
				return false, err
			} else {
				return do.Acknowledged, nil
			}

		} else {
			if do, err := esc.EsClient.CreateIndex(indexName).BodyJson(res).Do(context.Background()); err != nil {
				return false, err
			} else {
				return do.Acknowledged, nil
			}
		}
	}
	return false, errors.New("索引库已存在")
}

func (esc *esService) Save(IndexName string, id string, model any) (bool, error) {
	put, err := esc.EsClient.Index().
		Index(IndexName).        // 设置索引名称
		Id(id).                  // 设置文档ID
		BodyJson(model).         // 指定前面声明的对象信息
		Do(context.Background()) // 执行请求，需要传入一个上下文对象

	if err != nil {
		panic(err)
	}
	fmt.Printf("文档Id %s, 索引名 %s\n", put.Id, put.Index)
	return true, nil
}

// 根据id更新数据
func (esc *esService) UpdateById(indexName string, id string, age int) (bool, error) {
	put, err := esc.EsClient.Update().
		Index(indexName).
		Id(id).
		Doc(map[string]interface{}{"userAge": age}).Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Printf("文档Id %s, 索引名 %s\n", put.Id, put.Index)
	return true, nil
}

// UpdateQuery 根据查询条件更新文档
func (esc *esService) UpdateQuery(indexName string, uid int, age int) (bool, error) {
	do, err := esc.EsClient.UpdateByQuery(indexName).
		Query(elastic.NewTermQuery("userId", uid)).
		Script(elastic.NewScript("ctx._source['userAge'] = " + strconv.Itoa(age))).
		ProceedOnVersionConflict().Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Printf("文档一共更新了 %d\n", do.Total)
	return true, nil
}

// 根据ID删除文档
func (esc *esService) DeleteById(indexName string, id string) (bool, error) {
	put, err := esc.EsClient.Delete().
		Index(indexName).
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Printf("文档Id %s, 索引名 %s\n", put.Id, put.Index)
	return true, nil
}

// 根据查询条件删除文档
func (esc *esService) DeleteQuery(indexName string, uid int) (bool, error) {
	do, err := esc.EsClient.DeleteByQuery(indexName).
		Query(elastic.NewTermQuery("userId", uid)).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Printf("文档一共删除了 %d\n", do.Total)
	return true, nil
}

// 根据ID查询文档
func (esc *esService) SearchByDocId(indexName string, docId string, u any) (any, error) {
	getResult, err := esc.EsClient.Get().
		Index(indexName).
		Id(docId).
		Do(context.Background())
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if getResult.Found {
		fmt.Printf("文档id=%s 版本号=%d 索引名=%s\n", getResult.Id, getResult.Version, getResult.Index)
	}
	data, _ := getResult.Source.MarshalJSON()
	jsonParseErr := json.Unmarshal(data, &u)
	if jsonParseErr != nil {
		return nil, errors.New(jsonParseErr.Error())
	}
	return u, nil
}

// 精确匹配单个字段
func (esc *esService) SearchByTermFild(indexName string, field string, val any, resp any) ([]any, error) {
	result, err := esc.EsClient.Search(indexName).
		Query(elastic.NewTermQuery(field, val)).
		Sort("userAge", false). // 设置排序字段 false表示升序
		From(0).                // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(10).               // 设置分页参数 - 每页大小
		Pretty(true).           // 查询结果返回可读性较好的JSON格式
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	users := make([]any, 0)
	if result.TotalHits() > 0 {
		// 查询结果不为空，遍历结果
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range result.Each(reflect.TypeOf(resp)) {
			users = append(users, item)
		}
	}

	return users, nil
}
