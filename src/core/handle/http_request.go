package handle

import (
	"bytes"
	"core/model"
	"core/resp"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/kirinlabs/HttpRequest"
	"github.com/mitchellh/mapstructure"
)

type GetParam struct {
	Url         string            `json:"url"`     // 请求地址
	Params      string            `json:"params"`  // 请求地址栏后需要拼接参数操作
	Headers     map[string]string `json:"headers"` // 请求header头设置
	ContentType string            `json:"contentType"`
}

type PostParam struct {
	Url         string                 `json:"url"`     // 请求地址
	Params      map[string]string      `json:"params"`  // 请求地址栏后需要拼接参数操作
	Headers     map[string]string      `json:"headers"` // 请求header头设置
	Data        map[string]interface{} `json:"data"`    // 请求报文
	ContentType string                 `json:"contentType"`
}

/*
GET 请求
*/
func GET(param GetParam) (interface{}, error) {
	httpClient := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", param.Url, bytes.NewBuffer(nil))
	// 设置请求头类型
	if param.ContentType != "" {
		req.Header.Set("Content-type", param.ContentType)
	} else {
		req.Header.Set("Content-type", "application/json")
	}
	if err != nil {
		return nil, errors.New("请求失败: %v \n")
	}
	if param.Headers != nil {
		for key, val := range param.Headers {
			req.Header.Add(key, val)
		}
	}
	// 发起请求
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result := make(map[string]interface{}, 0)
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func POST(param PostParam) (interface{}, error) {
	httpClient := &http.Client{Timeout: 30 * time.Second}
	body, _ := json.Marshal(param.Data)
	req, err := http.NewRequest("POST", param.Url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New("new request is fail: %v \n")
	}
	// 设置请求头类型
	if param.ContentType != "" {
		req.Header.Set("Content-type", param.ContentType)
	} else {
		req.Header.Set("Content-type", "application/json")
	}
	// 设置请求头信息
	if param.Headers != nil {
		for key, val := range param.Headers {
			req.Header.Add(key, val)
		}
	}
	// 请求参数编码
	q := req.URL.Query()
	if param.Params != nil {
		for key, val := range param.Params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	// 发起请求
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	result := make(map[string]interface{}, 0)
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

/*
@Description HttpRequest  支持任意方式的http请求 超时时间30秒
@Param url 请求地址
@Param method 请求方式
@Param headers 请求header头设置
@Param params 请求地址栏后需要拼接参数操作
@Param data 请求报文
@Return interface{}, error
*/
func HttpRequests(url, method string, params, headers map[string]string, data []byte) (interface{}, error) {
	httpClient := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.New("new request is fail: %v \n")
	}
	// 设置请求头类型
	req.Header.Set("Content-type", "application/json")
	// 设置请求头信息
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	// 请求参数编码
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	// 发起请求
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result := make(map[string]interface{}, 0)
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func MicroServiceHttpRequest[T any](httpParam model.MicroServiceHttpParam[any]) (T, error) {
	var url string
	var target T
	if httpParam.Url != "" {
		if strings.Contains(httpParam.Url, "http") || strings.Contains(httpParam.Url, "http") {
			return target, errors.New("请传入完整的带http或者https请求地址")
		}
		url = httpParam.Url
	} else {
		// if getUrl, err := nacos.GetURLByServiceInfo(httpParam.ServiceName, httpParam.HttpMethod, httpParam.ApiPath); err != nil {
		// 	return target, errors.New(err.Error())
		// } else {
		// 	url = getUrl
		// }
	}
	request := HttpRequest.NewRequest()
	request.SetHeaders(map[string]string{
		"accessToken": httpParam.AccessToken,
	})
	var post *HttpRequest.Response
	switch httpParam.Method {
	case "GET":
		post, _ = request.JSON().Get(url)
		break
	case "POST":
		post, _ = request.JSON().Post(url, httpParam.ReqParam)
		break
	}
	if body, err := post.Body(); err != nil {
		return target, errors.New("服务调用失败")
	} else {
		response := resp.Response{}
		if err := json.Unmarshal(body, &response); err != nil {
			return target, errors.New("解析Response返回参数失败！")
		} else {
			if response.Code != 1 {
				return target, errors.New(fmt.Sprintf("查询失败！code = %v", response.Code))
			}
			if err := mapstructure.Decode(response.Data, &target); err != nil {
				return target, errors.New("数据列表结构体转换失败")
			}
			return target, nil
		}
	}
}
