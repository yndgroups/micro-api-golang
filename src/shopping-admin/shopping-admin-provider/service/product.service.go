package service

import (
	"context"
	"core/config"
	"core/db"
	"core/exception"
	"core/redis"
	"fmt"
	"protobuf/build/global"
	"protobuf/build/shopping_admin"
	"protobuf/enum"
)

var ProductService = &productService{}

type productService struct {
	shopping_admin.ProductServiceServer
}

func (p *productService) FindProductListByProductIds(ctx context.Context, request *shopping_admin.ProductRequestParamMicroList) (*shopping_admin.ProductParamsResponse, error) {
	spuIds := make([]string, 0)
	skuIds := make([]string, 0)
	if len(request.ParamList) > 0 {
		for _, v := range request.ParamList {
			if v.SpecType == 0 {
				spuIds = append(spuIds, v.ProductId)
			} else {
				skuIds = append(skuIds, v.ProductId)
			}
		}
	} else {
		return &shopping_admin.ProductParamsResponse{List: nil, Count: 0}, nil
	}
	var list []*global.ProductExpand
	spuSql := `
		SELECT
		   SPU.product_id ProductId,
		   SPU.product_name ProductName,
		   SPU.category_id CategoryId,
		   SPU.temp_id TempId,
		   SPU.slider_image SliderImage,
		   SPU.recommend_image RecommendImage,
		   SPU.price Price,
		   SPU.cost_price CostPrice,
		   SPU.vip_price VipPrice,
		   SPU.market_price MarketPrice,
		   SPU.image Image,
		   SPU.stock Stock,
		   SPU.stock_warning StockWarning
		FROM
		    ` + enum.Product.TableName + ` SPU
		WHERE SPU.product_id in ?`
	skuSql := `SELECT
		   	SPU.product_name ProductName,
		   	SPU.category_id CategoryId,
		   	SPU.temp_id TempId,
		   	SPU.slider_image SliderImage,
			SPU.recommend_image RecommendImage,
			SKU.price Price,
			SKU.cost_price CostPrice,
			SKU.vip_price VipPrice,
			SKU.market_price MarketPrice,
			SKU.image Image,
			SKU.stock Stock,
			SKU.sku_id ProductId,
			SKU.stock_warning StockWarning
		FROM
		    ` + enum.Product.TableName + ` SPU LEFT JOIN ` + enum.ProductSku.TableName + ` SKU ON SPU.product_id = SKU.product_id
		WHERE SKU.sku_id = ?`
	var count int64
	if len(spuIds) > 0 && len(skuIds) > 0 {
		db.DB.Raw(spuSql+` UNION ALL `+skuSql, spuIds, skuIds).Scan(&list)
	} else if len(spuIds) > 0 && len(skuIds) == 0 {
		db.DB.Raw(spuSql, spuIds).Scan(&list)
	} else if len(spuIds) == 0 && len(skuIds) > 0 {
		db.DB.Raw(skuSql, skuIds).Scan(&list)
	}
	return &shopping_admin.ProductParamsResponse{List: list, Count: count}, nil
}

func (p *productService) CreateProduct(ctx context.Context, product *shopping_admin.Product) (*shopping_admin.ProductResponse, error) {
	tx := db.DB.Begin()
	// 保存详情数据
	if err := tx.Create(shopping_admin.ProductDetail{
		ProductId: product.ProductId,
		Content:   product.Content,
	}).Error; err != nil {
		tx.Rollback()
		return &shopping_admin.ProductResponse{Count: 0, Msg: "保存商品详情数据失败," + exception.DbMsg(err.Error(), config.Global.ServerConfig().CommonProvider)}, nil
	}

	// 保存sku数据
	if len(product.SkuList) > 0 && product.SpecType == 1 {
		for i, v := range product.SkuList {
			v.SkuId = redis.GET.GetPrimaryKey("SYS_SHOPP_ADMIN_REQUEST_COUNT")
			if err := tx.Create(v).Error; err != nil {
				tx.Rollback()
				return &shopping_admin.ProductResponse{Count: 0, Msg: fmt.Sprintf("保存商品SKU第%v数据失败,", i+1) + exception.DbMsg(err.Error(), config.Global.ServerConfig().CommonProvider)}, nil
			}
		}
	}

	// 保存商品基本数据
	if res := tx.Create(product); res.Error != nil {
		tx.Rollback()
		return &shopping_admin.ProductResponse{Count: 0, Msg: "保存商品基本数据失败," + exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().CommonProvider)}, nil
	}
	tx.Commit()
	return &shopping_admin.ProductResponse{Msg: "数据添加成功！", Count: 1}, nil
}

func (p *productService) UpdateProduct(ctx context.Context, product *shopping_admin.Product) (*shopping_admin.ProductResponse, error) {
	return &shopping_admin.ProductResponse{}, nil
}

func (p *productService) DeleteProduct(ctx context.Context, ids *shopping_admin.ProductIds) (*shopping_admin.ProductResponse, error) {
	return &shopping_admin.ProductResponse{}, nil
}

func (p *productService) FindById(ctx context.Context, ids *shopping_admin.ProductIds) (*shopping_admin.ProductResponse, error) {
	product := shopping_admin.ProductDetailVO{}
	sql := `
		SELECT 
			PT.product_id ProductId,
			PT.category_id CategoryId,
			PT.product_name ProductName,
			PT.keyword Keyword,
			PT.introduction Introduction,
			PT.price Price,
			PT.cost_price CostPrice,
			PT.vip_price VipPrice,
			PT.market_price MarketPrice,
			PT.image Image,
			PT.slider_image SliderImage,
			PT.recommend_image RecommendImage,
			PT.stock Stock,
			PT.stock_warning StockWarning,
			PT.unit_name UnitName,
			PT.temp_id TempId,
			PT.spec_type SpecType,
			(CASE PT.spec_type
					WHEN 0 THEN '单规格'
					WHEN 1 THEN '多规格' END
			) AS SpecTypeName,
			PT.is_display IsDisplay,
			(CASE PT.is_display
					WHEN 0 THEN '未上架'
					WHEN 1 THEN '上架' END
			) AS IsDisplayName,
			PT.is_virtual IsVirtual,
			(CASE PT.is_virtual
					WHEN 0 THEN '实物商品'
					WHEN 1 THEN '虚拟商品' END
			) AS IsVirtualName,
			PT.virtual_type VirtualType,
			PT.virtual_sales VirtualSales,
			PT.is_limit IsLimit,
			(CASE PT.is_limit
					WHEN 0 THEN '否'
					WHEN 1 THEN '是' END
			) AS IsLimitName,
			PT.limit_type LimitType,
			(CASE PT.limit_type
					WHEN 0 THEN '单次限购'
					WHEN 1 THEN '永久限购' END
			) AS LimitTypeName,
			PT.limit_num LimitNum,
			PT.sn Sn,
			PT.brand_id BrandId,
			PT.sort_by SortBy,
			DATE_FORMAT(PT.create_time,'%Y-%m-%d %T') CreateTime,
			PD.content Content
		FROM 
		    ` + enum.Product.TableName + ` PT left join ` + enum.ProductDetail.TableName + ` PD on PT.product_id = PD.product_id 
		WHERE PT.product_id = ?`
	if db.DB.Raw(sql, ids.ProductId[0]).Scan(&product).RowsAffected > 0 {
		if product.SpecType == 1 {
			product.SkuList = p.FindProductSkuList(ids.ProductId[0])
		}
	}
	return &shopping_admin.ProductResponse{Detail: &product}, nil
}

func (p *productService) FindProductSkuList(product string) []*shopping_admin.ProductSkuListVo {
	list := make([]*shopping_admin.ProductSkuListVo, 0)
	listSql := `
			SELECT
				SKU.*
			FROM
				ynd_product_sku SKU
			WHERE
				SKU.product_id = ?
	`
	db.DB.Raw(listSql, product).Scan(&list)
	return list
}

func (p *productService) FindPageList(ctx context.Context, request *shopping_admin.ProductPageAuthQuery) (*shopping_admin.ProductResponse, error) {
	countSql := `SELECT COUNT(1) FROM ` + enum.Product.TableName + ` SPU WHERE del_status = 0`
	paramSql := ""
	if request.Params != nil {
		if request.Params.ProductName != "" {
			paramSql += fmt.Sprintf(" AND  SPU.`product_name` like CONCAT('%%','%v','%%')", request.Params.ProductName)
		}
		if request.Params.Keyword != "" {
			paramSql += fmt.Sprintf(" AND SPU.`keyword` like CONCAT('%%','%v','%%')", request.Params.Keyword)
		}
	}
	var count int64
	println(countSql + paramSql)
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		listSql := `SELECT 
			SPU.product_id ProductId,
			SPU.category_id CategoryId,
			SPU.product_name ProductName,
			SPU.keyword Keyword,
			SPU.introduction Introduction,
			SPU.price Price,
			SPU.cost_price CostPrice,
			SPU.vip_price VipPrice,
			SPU.market_price MarketPrice,
			SPU.image Image,
			SPU.slider_image SliderImage,
			SPU.recommend_image RecommendImage,
			SPU.stock Stock,
			SPU.stock_warning StockWarning,
			SPU.unit_name UnitName,
			SPU.temp_id TempId,
			SPU.spec_type SpecType,
			(CASE SPU.spec_type
					WHEN 0 THEN '单规格'
					WHEN 1 THEN '多规格' END
			) AS SpecTypeName,
			SPU.is_display IsDisplay,
			(CASE SPU.is_display
					WHEN 0 THEN '未上架'
					WHEN 1 THEN '上架' END
			) AS IsDisplayName,
			SPU.is_virtual IsVirtual,
			(CASE SPU.is_virtual
					WHEN 0 THEN '实物商品'
					WHEN 1 THEN '虚拟商品' END
			) AS IsVirtualName,
			SPU.virtual_type VirtualType,
			SPU.virtual_sales VirtualSales,
			SPU.is_limit IsLimit,
			(CASE SPU.is_limit
					WHEN 0 THEN '否'
					WHEN 1 THEN '是' END
			) AS IsLimitName,
			SPU.limit_type LimitType,
			(CASE SPU.limit_type
					WHEN 0 THEN '单次限购'
					WHEN 1 THEN '永久限购' END
			) AS LimitTypeName,
			SPU.limit_num LimitNum,
			SPU.sn Sn,
			SPU.brand_id BrandId,
			SPU.sort_by SortBy,
			DATE_FORMAT(SPU.create_time,'%Y-%m-%d %T') CreateTime,
		    SPUD.content Content
		FROM 
		    ` + enum.Product.TableName + ` SPU LEFT JOIN ` + enum.ProductDetail.TableName + ` SPUD ON SPU.product_id = SPUD.product_id
		WHERE SPU.del_status = 0` + paramSql
		listSql += fmt.Sprintf(" ORDER BY SPU.create_time DESC LIMIT %d,%d", request.PageIndex, request.PageSize)
		list := make([]*shopping_admin.Product, 0)
		db.DB.Raw(listSql).Scan(&list) // 执行数据查询
		if db.DB.Raw(listSql).Scan(&list).RowsAffected > 0 {
			// for i, v := range list {
			// 	if v.SpecType == 1 {
			// 		skuList := FindProductSkuList(v.ProductId)
			// 		if len(skuList) > 0 {
			// 			list[i].SkuList = skuList
			// 		}
			// 	}
			// }
			return &shopping_admin.ProductResponse{List: list, Count: count}, nil
		} else {
			return &shopping_admin.ProductResponse{Count: 0, Msg: "为查询到相关数据"}, nil
		}
	} else {
		return &shopping_admin.ProductResponse{Count: 0}, nil
	}
}
