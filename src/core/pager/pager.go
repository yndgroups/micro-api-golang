package pager

/**
 * Pager 分页
 * @description Pager 请求参数统一封装
 * @author yangDaqiong
 * @date  2022-01-02
 */
type Pager struct {
	// 分页索引
	PageIndex int `json:"pageIndex"`
	// 分页长度
	PageSize int `json:"pageSize"`
	// 列表
	List any `json:"list"`
	// 数据长度
	TotalCount int64 `json:"totalCount"`
} // @Name Pager 分页参数封装

type NewPager struct {
	// 分页索引
	PageIndex int64 `json:"pageIndex"`
	// 分页长度
	PageSize int64 `json:"pageSize"`
	// 列表
	List any `json:"list"`
	// 数据长度
	TotalCount int64 `json:"totalCount"`
} // @Name Pager 分页参数封装

/*
@description PageInfo 列表分页请求参数
@author xaringan
@date  2022-06-19
*/
type PageInfo[T any] struct {
	// 分页起始页
	PageIndex int `json:"pageIndex" validate:"required,min=1" example:"1"`
	// 分页长度
	PageSize int `json:"pageSize" validate:"required,min=10,max=100" example:"10"`
	// 请求参数
	Params T `json:"params"`
}
