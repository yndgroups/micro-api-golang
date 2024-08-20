package resp

import (
	"core/req"
)

// @description Pager 请求参数统一封装
// @author yangDaqiong
// @date  2022-01-02
type Pager struct {
	// 分页索引
	PageIndex int64 `json:"pageIndex"`
	// 分页长度
	PageSize int64 `json:"pageSize"`
	// 列表
	List any `json:"list"`
	// 数据长度
	TotalCount int64 `json:"totalCount"`
	// 分页数
	PageNum int64 `json:"pageNum"`
} // @Name Pager 分页参数封装

// 构建分野参数
func CreatePager[T any](req req.Request[T], totalCount int64, list any) *Pager {
	var pageNum int64 = totalCount / req.PageSize
	if totalCount%req.PageSize != 0 {
		pageNum++
	}
	return &Pager{
		PageIndex:  req.PageIndex,
		PageSize:   req.PageSize,
		PageNum:    pageNum,
		List:       list,
		TotalCount: totalCount,
	}
}
