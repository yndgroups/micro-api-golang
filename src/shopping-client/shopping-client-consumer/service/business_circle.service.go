package service

import (
	"github.com/gin-gonic/gin"
)

var BusinessCircle = businessCircle{}

type businessCircle struct{}

// Delete 商圈删除
func (bc *businessCircle) Delete(ctx *gin.Context) any {
	return nil
}

// FindPageList 查询商圈分页列表
func (bc *businessCircle) FindPageList(ctx *gin.Context) any {
	return nil
}

// FindById 查询商圈
func (bc *businessCircle) FindById(ctx *gin.Context) any {
	return nil
}
