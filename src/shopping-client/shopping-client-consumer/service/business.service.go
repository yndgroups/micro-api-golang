package service

import (
	"github.com/gin-gonic/gin"
)

var Business = &business{}

type business struct{}

// Create 商家新增
func (b *business) Create(ctx *gin.Context) any {
	return nil
}

// UpdateBusiness 商家修改
func (b *business) Update(ctx *gin.Context) any {
	return nil
}

// Delete 商家删除
func (b *business) Delete(ctx *gin.Context) any {
	return nil
}

// FindById 查询商家
func (b *business) FindById(ctx *gin.Context) any {
	return nil
}

// FindPageList 查询商家分页列表
func (b *business) FindPageList(ctx *gin.Context) any {
	return nil
}
