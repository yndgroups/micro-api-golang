package service

import (
	"context"
	"core/config"
	"core/db"
	"core/exception"
	"fmt"
	"protobuf/build/shopping_admin"
	"protobuf/enum"
)

var ProductSpecService = &productSpecService{}

type productSpecService struct {
	shopping_admin.ProductSpecServiceServer
}

// Create 新增
func (ps *productSpecService) Create(cxt context.Context, params *shopping_admin.ProductSpec) (*shopping_admin.ProductSpecResponse, error) {
	return &shopping_admin.ProductSpecResponse{}, nil
}

// Update 修改
func (ps *productSpecService) Update(cxt context.Context, params *shopping_admin.ProductSpec) (*shopping_admin.ProductSpecResponse, error) {
	return &shopping_admin.ProductSpecResponse{}, nil
}

// Delete 删除
func (ps *productSpecService) Delete(cxt context.Context, params *shopping_admin.ProductSpecIds) (*shopping_admin.ProductSpecResponse, error) {
	return &shopping_admin.ProductSpecResponse{}, nil
}

// FindById 查询 详情
func (ps *productSpecService) FindById(cxt context.Context, params *shopping_admin.ProductSpecIds) (*shopping_admin.ProductSpecResponse, error) {
	return &shopping_admin.ProductSpecResponse{}, nil
}

// FindPageList 查询 分页
func (ps *productSpecService) FindPageList(cxt context.Context, params *shopping_admin.ProductSpecPageAuthQuery) (*shopping_admin.ProductSpecResponse, error) {
	return &shopping_admin.ProductSpecResponse{}, nil
}

// FindListByProductCategoryId 根据分类id查询规格
func (ps *productSpecService) FindListByProductCategoryId(cxt context.Context, params *shopping_admin.ProductSpecIds) (*shopping_admin.ProductSpecResponse, error) {
	sql := fmt.Sprintf(`SELECT spec_id,name,options from %s WHERE product_category_id = '%v'`, enum.ProductSpec.TableName, params.ProductCategoryId)
	list := make([]*shopping_admin.ProductSpec, 0)
	res := db.DB.Raw(sql).Scan(&list)
	if res.Error != nil {
		return &shopping_admin.ProductSpecResponse{Count: 0, Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().ShoppingAdminProvider)}, nil
	}
	return &shopping_admin.ProductSpecResponse{List: list}, nil
}
