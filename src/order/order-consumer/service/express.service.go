package service

import (
	"bytes"
	"core/resp"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

var Express = &express{}

type express struct{}

// 获取快递公司类型
func AutoComNum(ctx *gin.Context) any {
	//fmt.Println("AutoComNum...")
	text := ctx.Query("num") //单号
	url := "https://www.kuaidi100.com/autonumber/autoComNum?resultv2=1&text=" + text
	paramData := gin.H{}
	headerData := make(map[string]string, 1)
	headerData["User-Agent"] = "Mozilla/5.0 (X11; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0"
	response := SendPostWithHeard(url, paramData, headerData)
	type resAuto struct {
		ComCode   string `json:"comCode"`
		LengthPre int    `json:"lengthPre"`
		NoCount   int    `json:"noCount"`
		NoPre     string `json:"noPre"`
		Name      string `json:"name"`
	}

	type respData struct {
		Num     string     `json:"num"`
		ComCode string     `json:"comCode"`
		Auto    []*resAuto `json:"auto"`
	}
	var resdata *respData
	if err := json.Unmarshal([]byte(response), &resdata); err != nil {
		fmt.Println("Unmarshal resdata fail,", err)
		// httpext.Error(ctx, e.ERROR)
		return resp.Fail.WithData(err.Error())
	}
	for _, v := range resdata.Auto {
		kdName, ok := KD100Flags[v.ComCode]
		if ok {
			v.Name = kdName
		} else {
			v.Name = "未知快递"
		}
	}
	return resp.Success.WithData(resdata)
}

// 获取快递跟踪信息
func QueryNumDetail(ctx *gin.Context) any {
	//fmt.Println("AutoComNum...")
	num := ctx.Query("num")        //单号
	kdtype := ctx.Query("comCode") //快递类型

	key := "XXXXXXXX"                     //客户授权key
	customer := "XXXXXXXXXXXXXXXXXXXXXXX" //查询公司编号

	posturl := "http://poll.kuaidi100.com/poll/query.do" //实时查询请求地址

	paramData := make(map[string]string)
	paramData["com"] = kdtype //快递公司编码
	paramData["num"] = num    //快递单号

	paramDataSlice, _ := json.Marshal(paramData)
	paramjson := string(paramDataSlice)
	str := paramjson + key + customer
	h := md5.New()
	io.WriteString(h, str)
	sign := strings.ToUpper(string(h.Sum(nil)))
	fmt.Println("p=====>", customer)

	//POST请求需要三个参数，分别为customer(CustomerId)和sign(签名)和param(参数)
	postRes, postErr := http.PostForm(posturl, url.Values{"customer": {customer}, "sign": {sign}, "param": {paramjson}})
	if postErr != nil {
		fmt.Println("查询失败" + postErr.Error())
		return resp.Fail.WithData("查询失败," + postErr.Error())
	}
	postBody, err := ioutil.ReadAll(postRes.Body)
	if err != nil {
		fmt.Println("查询失败,请至快递公司官网自行查询" + err.Error())
		return resp.Fail.WithData("查询失败,请至快递公司官网自行查询" + err.Error())
	}
	fmt.Println(string(postBody))
	response := make(map[string]interface{})
	err = json.Unmarshal(postBody, &response)
	if err != nil {
		fmt.Println("json.Unmarshal error", err.Error())
		// httpext.ErrorExt(ctx, e.ERROR, "查询失败,请至快递公司官网自行查询"+err.Error())
		return resp.Fail.WithData(err.Error())
	}

	return resp.Success.WithData(response)
}

// SendPostWithHeard ...   发送POST请求 ...
// url：         请求地址
// data：        POST请求提交的数据
// hearderParam  设置header头信息
// contentType： 请求体格式，如：json
// content：     请求放回的内容
// func SendPostWithHeard(url string, data interface{}, hearderParam map[string]string, contentType string) string {
func SendPostWithHeard(url string, data interface{}, hearderParam map[string]string) string {
	bytesData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	//增加header选项 添加请求头
	if len(hearderParam) > 0 {
		for k, v := range hearderParam {
			request.Header.Set(k, v)
		}
	}
	//request.Header.Set("Content-Type", contentTypeParam[contentType])

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(1111111)
	fmt.Println(respBytes)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	return string(respBytes)
}

var KD100Flags = map[string]string{
	"ane66":          "安能快递",
	"debangwuliu":    "德邦物流",
	"debangkuaidi":   "德邦快递",
	"ems":            "EMS",
	"guotongkuaidi":  "国通快递",
	"huitongkuaidi":  "百世快递",
	"jd":             "京东物流",
	"kuayue":         "跨越速运",
	"pjbest":         "品骏快递",
	"shentong":       "申通快递",
	"shunfeng":       "顺丰速运",
	"suer":           "速尔快递",
	"xinfengwuliu":   "信丰物流",
	"youshuwuliu":    "优速物流",
	"youzhengguonei": "邮政快递包裹",
	"yuantong":       "圆通速递",
	"yuantongguoji":  "圆通国际",
	"yunda":          "韵达快递",
	"zhaijisong":     "宅急送",
	"zhongtong":      "中通快递",
	"ewe":            "EWE全球快递",
	"quanyikuaidi":   "全一快递",
	"tiantian":       "天天快递",
	"sxjdfreight":    "顺心捷达",
	"dhl":            "DHL-中国件",
	"tnt":            "TNT",
	"other":          "其它快递",
}
