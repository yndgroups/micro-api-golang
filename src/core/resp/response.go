package resp

import (
	"encoding/json"
	"fmt"
)

var (
	Back = Result(1, "统一返回结果")

	Success       = Result(1, "请求成功！")
	Fail          = Result(0, "请求失败！")
	InsertSuccess = Result(1, "新增成功！")
	UpdateSuccess = Result(1, "修改成功！")
	DeleteSuccess = Result(1, "删除成功！")

	InsertFail = Result(0, "新增失败！")
	UpdateFail = Result(0, "修改失败，请检查参数是否已经更改，与上一次更新的致完全一致，则不做更新！")
	DeleteFail = Result(0, "删除失败！")
	NotData    = Result(0, "未查询到相关数据")

	Err = Result(500, "服务器处理错误！")

	// NoAuthorized 403 服务器拒绝请求。可以理解为没有权限访问此网站，服务器能够收到请求但拒绝提供服务。
	AuthorizedErr = Result(400, "非法请求！")
	NoAuthorized  = Result(403001, "你没有权限调用此接口！")

	ErrValidate    = Result(412000, "参数验证不通过！")
	ValidateErr    = Result(412000, "参数验证不通过！")
	ErrParam       = Result(412001, "参数有误！")
	ErrSignParam   = Result(412002, "签名参数有误！")
	NotFindParamId = Result(412003, "请正确传入必须的id参数！")
	ErrUserService = Result(412004, "用户服务异常！")
	ErrUserPhone   = Result(412005, "手机校验不通过！")
	ErrUserCaptcha = Result(412006, "验证码错误！")

	//100：继续
	//101：切换协议
	//200：确定，表示客户端请求已成功
	//201：已创建
	//202：已接收
	//203：非权威性信息
	//204：无内容
	//205：充值内容
	//206：部分内容。表明已部分下载了一个文件，可以续传损坏的下载或者将下载拆分为多个并发的流。
	//207：多状态（WebDAV）。此消息之前应该还有一条XML消息，其中可能包含几个单独的响应代码，具体取决于发出了多少个子请求。
	//301：已永久移动。此请求和之后所有的请求都应该转到制定的URI。
	//302：对象已移动。对于基于表单的身份验证，此消息通常表示为“对象已移动”。请求的资源临时驻留在不同的URI。由于重定向有时可能会改变，客户端将来在请求时应该继续使用 RequestURI。只有在 CacheControl 或 Expires 标题字段中指示，此响应才能够缓存。
	//304：未修改。客户端请求的文档已在其缓存中，文档自缓存以来尚未被修改过。客户端使用文档的缓存副本，而不从服务器下载文档。
	//307：临时重定向。
	//406：客户端浏览器不接受所请求页面的 MIME 类型
	//407：要求进行代理身份验证
	//412：前提条件失败
	//413：请求实体太大
	//414：请求URI太长
	//415：不支持的媒体类型
	//416：无法满足请求的范围
	//417：执行失败
	//423：锁定的错误
	//500：内部服务器错误。很多服务器端错误都可能导致此错误消息。事件查看器日志包含更详细的错误原因。此外，您可以禁用友好 HTTP 错误消息以便收到详细的错误说明。IIS 定义了几个不同的500错误，用于指示更为具体的错误原因。
	//501：页眉指定了未实现的配置
	//502：Web 服务器作为网关或代理服务器时，从上游服务器收到了无效响应。此类错误一般与服务器本身有关（与请求无关）。IIS 定义了几个不同的502错误，用于指示更为具体的错误原因。
	//503：目前服务器无法使用，一般是因为服务器超载或停止维护
	//504：网关超时
	//505：HTTP版本不受支持
)

/*
Response
@description 返回参数统一封装
@author yangDaqiong
@date  2022-01-02
*/
type Response struct {
	// 结果状态码
	Code uint32 `json:"code" default:"0"`
	// 状态信息
	Msg string `json:"msg" default:"请求失败"`
	// 结果
	Data any `json:"data"`
} // @name Response

// 构造函数
func Result(code uint32, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// 返回删除结果函数
func (r *Response) Result(count int64, data any) Response {
	resp := Response{
		Code: 0,
		Msg:  "未查询到相关数据",
		Data: nil,
	}
	if count > 0 {
		resp.Data = data
		resp.Msg = fmt.Sprintf("成功查询%v条数据", count)
		resp.Code = 1
	}
	return resp
}

// 返回带提示信息
func (r *Response) WithMsg(msg string) Response {
	return Response{
		Code: r.Code,
		Msg:  msg,
		Data: nil,
	}
}

// 返回复杂数据结果
func (r *Response) WithData(data any) Response {
	return Response{
		Code: r.Code,
		Msg:  r.Msg,
		Data: data,
	}
}

// 返回复杂数据结果
func (r *Response) WithDataAndMsg(data any, msg string) Response {
	return Response{
		Code: r.Code,
		Msg:  msg,
		Data: data,
	}
}

// Insert 返回新增结果函数
func (r *Response) Insert(count int64, msg ...string) Response {
	resp := Response{
		Code: 0,
		Msg:  "新增失败",
		Data: nil,
	}
	if count > 0 {
		resp.Code = 1
		resp.Msg = fmt.Sprintf("成功新增%d条数据", count)
	} else {
		if len(msg) > 0 {
			resp.Msg = msg[0]
		}
	}
	return resp
}

// Update 返回修改结果函数
func (r *Response) Update(count int64, msg ...string) Response {
	resp := Response{
		Code: r.Code,
		Msg:  "修改失败",
		Data: nil,
	}
	if count > 0 {
		resp.Code = 1
		resp.Msg = fmt.Sprintf("成功修改%d条数据", count)
	} else {
		if len(msg) > 0 {
			resp.Msg = msg[0]
		}
	}
	return resp
}

// 返回删除结果函数
func (r *Response) Delete(count int64, msg ...string) Response {
	resp := Response{
		Code: r.Code,
		Msg:  "删除失败",
		Data: nil,
	}
	if count > 0 {
		resp.Code = 1
		resp.Msg = fmt.Sprintf("成功删除%d条数据", count)
	} else {
		if len(msg) > 0 {
			resp.Msg = msg[0]
		}
	}
	return resp
}

// 结构体转字符串
func (resp *Response) ToString() string {
	err := &struct {
		Code uint32      `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: resp.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}
