package service

import (
	"context"
	"core/db"
	"protobuf/build/shopping_admin"
	"shopping-admin/shopping-admin-consumer/enum"
)

var ProductCategoryService = &productCategoryService{}

type productCategoryService struct {
	shopping_admin.ProductCategoryServiceServer
}

func (p *productCategoryService) CreateProductCategory(ctx context.Context, product *shopping_admin.ProductCategory) (*shopping_admin.ProductCategoryResponse, error) {
	return &shopping_admin.ProductCategoryResponse{}, nil
}

func (p *productCategoryService) UpdateProductCategory(ctx context.Context, product *shopping_admin.ProductCategory) (*shopping_admin.ProductCategoryResponse, error) {
	return &shopping_admin.ProductCategoryResponse{}, nil
}

func (p *productCategoryService) DeleteProductCategory(ctx context.Context, product *shopping_admin.ProductCategoryIds) (*shopping_admin.ProductCategoryResponse, error) {
	return &shopping_admin.ProductCategoryResponse{}, nil
}

func (p *productCategoryService) FindById(ctx context.Context, product *shopping_admin.ProductCategoryIds) (*shopping_admin.ProductCategoryResponse, error) {
	return &shopping_admin.ProductCategoryResponse{}, nil
}

func (p *productCategoryService) FindPageList(ctx context.Context, product *shopping_admin.ProductCategoryPageAuthQuery) (*shopping_admin.ProductCategoryResponse, error) {
	return &shopping_admin.ProductCategoryResponse{}, nil
}

func (p *productCategoryService) FindTreeByParentId(ctx context.Context, request *shopping_admin.ProductCategoryIds) (*shopping_admin.ProductCategoryResponse, error) {
	sql := `
			SELECT 
				A.product_category_id,
				A.parent_id,
				A.name,
				A.self_rate,
				A.profit_sharing,
				A.sort_by
			FROM
				` + enum.ProductCategory.TableName + ` A
			WHERE A.del_status = 0 AND FIND_IN_SET(product_category_id,getProductCategoryChildren(?))
			ORDER BY A.sort_by ASC`
	list := make([]*shopping_admin.ProductCategoryListVO, 0)
	db.DB.Raw(sql, request.ParentId).Scan(&list)
	return &shopping_admin.ProductCategoryResponse{
		List: list,
	}, nil
}

func (p *productCategoryService) FindListByParentId(ctx context.Context, request *shopping_admin.ProductCategoryIds) (*shopping_admin.ProductCategoryResponse, error) {
	sql := `
			SELECT 
				A.product_category_id,
				A.parent_id,
				A.name name,
				A.self_rate,
				A.profit_sharing,
				A.sort_by,
				(SELECT COUNT(1) from ` + enum.ProductCategory.TableName + ` B WHERE B.parent_id = A.product_category_id) as childCount
			FROM
				` + enum.ProductCategory.TableName + ` A
			WHERE A.del_status = 0 AND A.parent_id = ?
			ORDER BY A.sort_by ASC`
	list := make([]*shopping_admin.ProductCategoryListVO, 0)
	db.DB.Raw(sql, request.ParentId).Scan(&list)
	return &shopping_admin.ProductCategoryResponse{List: list}, nil

}
