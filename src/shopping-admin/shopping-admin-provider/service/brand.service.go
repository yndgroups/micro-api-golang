package service

import (
	"context"
	"core/config"
	"core/db"
	"core/exception"
	"fmt"
	"protobuf/build/shopping_admin"
)

var BrandService = &brandService{}

type brandService struct {
	shopping_admin.BrandServiceServer
}

// Create 创建品牌
func (b *brandService) Create(ctx context.Context, brand *shopping_admin.Brand) (*shopping_admin.BrandResponse, error) {
	return &shopping_admin.BrandResponse{}, nil
}

// Update 修改品牌
func (b *brandService) Update(ctx context.Context, brand *shopping_admin.Brand) (*shopping_admin.BrandResponse, error) {
	return &shopping_admin.BrandResponse{}, nil
}

// Delete 删除品牌
func (b *brandService) Delete(ctx context.Context, ids *shopping_admin.BrandIds) (*shopping_admin.BrandResponse, error) {
	return &shopping_admin.BrandResponse{}, nil
}

// FindById 根据id查询详情
func (b *brandService) FindById(ctx context.Context, ids *shopping_admin.BrandIds) (*shopping_admin.BrandResponse, error) {
	return &shopping_admin.BrandResponse{}, nil
}

// FindPageList 查询品牌分页
func (b *brandService) FindPageList(ctx context.Context, brandPageAuthQuery *shopping_admin.BrandPageAuthQuery) (*shopping_admin.BrandResponse, error) {
	return &shopping_admin.BrandResponse{}, nil
}

// FindListByProductCategoryId 根据绑定的分类查询品牌
func (b *brandService) FindListByProductCategoryId(ctx context.Context, ids *shopping_admin.BrandIds) (*shopping_admin.BrandResponse, error) {
	sql := `SELECT
	B.brand_id,
	B.name,
	B.introduction,
	B.logo,
	B.display_type,
	B.url
FROM
	ynd_brand B, ynd_product_category_brand PCB WHERE B.brand_id = PCB.brand_id AND PCB.product_category_id = '%v'  ORDER BY B.sort_by`
	var list []*shopping_admin.Brand
	res := db.DB.Raw(fmt.Sprintf(sql, ids.ProductCategoryId)).Scan(list)
	if res.Error != nil {
		return &shopping_admin.BrandResponse{Count: 0, Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().ShoppingAdminProvider)}, nil
	}
	return &shopping_admin.BrandResponse{List: list}, nil
}
