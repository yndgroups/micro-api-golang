package service

import (
	"github.com/gin-gonic/gin"
)

var BusinessStore = &businessStore{}

type businessStore struct{}

// Delete 商家门店删除
func (bs *businessStore) Delete(ctx *gin.Context) any {
	return nil
}

// FindPageList 查询商家门店分页列表
func (bs *businessStore) FindPageList(ctx *gin.Context) any {
	return nil
}

// FindById 查询商家门店
func (bs *businessStore) FindById(ctx *gin.Context) any {
	return nil
}
