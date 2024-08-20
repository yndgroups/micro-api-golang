package service

import (
	"github.com/gin-gonic/gin"
)

var BusinessFranchisee = &businessFranchisee{}

type businessFranchisee struct{}

// Create 商家加盟新增
func (bf *businessFranchisee) Create(ctx *gin.Context) any {
	return nil
}

// Update 商家加盟修改
func (bf *businessFranchisee) Update(ctx *gin.Context) any {
	return nil
}

// Delete 商家加盟删除
func (bf *businessFranchisee) Delete(ctx *gin.Context) any {
	return nil
}

// FindById 查询商家加盟
func (bf *businessFranchisee) FindById(ctx *gin.Context) any {
	return nil
}

// FindPageList 查询商家加盟分页列表
func (bf *businessFranchisee) FindPageList(ctx *gin.Context) any {
	return nil
}
