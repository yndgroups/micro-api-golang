package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/shopping_client"
	"protobuf/enum"
)

var ProductCategorySerice = &productCategorySerice{}

type productCategorySerice struct {
	shopping_client.ProductCategoryServiceServer
}

// FindCategoryByParentId 查询分了列表
func (pdt *productCategorySerice) FindByParentId(ctx context.Context, params *shopping_client.ProductCategoryAuthParams) (*shopping_client.ProductCategoryResponse, error) {
	listSq := `
SELECT 
	category_id,
	parent_id,
	name,
	type,
	store_id,
	self_rate,
	profit_sharing,
	sort_by,
	del_status,
	create_by,
	DATE_FORMAT(create_time,'%Y-%m-%d') create_time
FROM `
	listSq += fmt.Sprintf(`%v WHERE del_status = 0 AND parent_id = '%v'`, enum.ProductCategory.TableName, params.ParentId)
	var list []*shopping_client.ProductCategory
	count := db.DB.Raw(listSq).Scan(&list).RowsAffected
	return &shopping_client.ProductCategoryResponse{Count: count, List: list}, nil
}

// 递归查询分类列 getCategoryTree是mysql自定义函数
func (pdt *productCategorySerice) FindTreeByParentId(ctx context.Context, params *shopping_client.ProductCategoryAuthParams) (*shopping_client.ProductCategoryResponse, error) {
	listSq := `
	SELECT 
		category_id,
		parent_id,
		name,
		type,
		store_id,
		self_rate,
		profit_sharing,
		sort_by,
		del_status,
		create_by,
		DATE_FORMAT(create_time,'%Y-%m-%d') create_time
	FROM `
	listSq += fmt.Sprintf(` %v WHERE find_in_set(parent_id, (SELECT getCategoryTree('%v'))) `, enum.ProductCategory.TableName, params.ParentId)
	var list []*shopping_client.ProductCategoryTreeList
	count := db.DB.Raw(listSq).Scan(&list).RowsAffected
	return &shopping_client.ProductCategoryResponse{Count: count, TreeList: list}, nil
}
