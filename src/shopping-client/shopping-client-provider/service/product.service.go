package service

import (
	"context"
	"core/config"
	"core/db"
	"core/exception"
	"fmt"
	"protobuf/build/global"
	"protobuf/build/shopping_client"
	"protobuf/enum"
)

var ProductService = &productService{}

type productService struct {
	shopping_client.ProductServiceServer
}

// FindListByProductIds 根据集合查询商品数据
func (p *productService) FindListByProductIds(ctx context.Context, params *shopping_client.CreateOrderParams) (*shopping_client.ProductResponse, error) {
	spuIds := make([]string, 0)
	skuIds := make([]string, 0)
	if len(params.ProductList) > 0 {
		for _, v := range params.ProductList {
			if v.SpecType == 1 { // 单规格
				spuIds = append(spuIds, v.ProductId)
			} else if v.SpecType == 2 { // 多规格
				skuIds = append(skuIds, v.ProductId)
			}
		}
	} else {
		return &shopping_client.ProductResponse{Msg: "请至少传入一个商品参数"}, nil
	}
	var list []*global.ProductExpand
	spuSql := `
		SELECT
		   SPU.product_id,
		   SPU.product_name,
		   SPU.category_id,
		   SPU.temp_id,
		   SPU.slider_image,
		   SPU.recommend_image,
		   SPU.spec_type,
		   SPU.price,
		   SPU.cost_price,
		   SPU.vip_price,
		   SPU.market_price,
		   SPU.image,
		   SPU.stock,
		   SPU.stock_warning
		FROM ` + enum.Product.TableName + ` SPU WHERE SPU.product_id in ?`
	skuSql := `SELECT
			SKU.sku_id product_id,
		   	SPU.product_name,
		   	SPU.category_id,
		   	SPU.temp_id,
		   	SPU.slider_image,
			SPU.recommend_image,
			SPU.spec_type,
			SKU.price,
			SKU.cost_price,
			SKU.vip_price,
			SKU.market_price,
			SKU.image,
			SKU.stock,
			SKU.stock_warning
		FROM ` + enum.Product.TableName + ` SPU LEFT JOIN ` + enum.ProductSku.TableName + ` SKU ON SPU.product_id = SKU.product_id WHERE SKU.sku_id = ?`
	var count int64
	if len(spuIds) > 0 && len(skuIds) > 0 {
		db.DB.Raw(spuSql+` UNION ALL `+skuSql, spuIds, skuIds).Scan(&list)
	} else if len(spuIds) > 0 && len(skuIds) == 0 {
		db.DB.Raw(spuSql, spuIds).Scan(&list)
	} else if len(spuIds) == 0 && len(skuIds) > 0 {
		db.DB.Raw(skuSql, skuIds).Scan(&list)
	}
	return &shopping_client.ProductResponse{ExpandList: list, Count: count}, nil
}

// FindById 商品详情
func (p *productService) FindById(ctx context.Context, ids *shopping_client.ProductIds) (*shopping_client.ProductResponse, error) {
	sql := `
		SELECT A.product_id,
		A.category_id,
		A.product_name,
		A.keyword,
		A.price,
		A.cost_price,
		A.vip_price,
		A.market_price,
		A.image,
		A.slider_image,
		A.recommend_image,
		A.stock,
		A.stock_warning,
		A.unit_name,
		A.temp_id,
		A.spec_type,
		A.is_display,
		A.introduction,
		A.is_virtual,
		A.brand_id,
		A.virtual_type,
		A.virtual_sales,
		A.is_limit,
		A.limit_type,
		A.limit_num,
		A.views,
		A.mer_id,
		A.sn,
		A.mer_use,
		A.bar_code,
		A.sales,
		A.is_hot,
		A.is_benefit,
		A.is_best,
		A.is_new,
		A.give_integral,
		A.is_seckill,
		A.is_bargain,
		A.is_good,
		A.is_sub,
		A.is_vip,
		A.code_path,
		A.activity,
		A.video_link,
		A.spu,
		A.label_id,
		A.command_word,
		A.recommend_list,
		A.vip_product,
		A.presale,
		A.presale_start_time,
		A.presale_end_time,
		A.presale_day,
		A.logistics,
		A.freight,
		A.custom_form,
		A.sort_by,
		A.del_status,
		A.create_by,
		A.update_by,
		DATE_FORMAT(A.create_time,'%Y-%m-%d') create_time,
		DATE_FORMAT(A.update_time,'%Y-%m-%d') update_time,
		C.content
	`
	sql += fmt.Sprintf(` FROM %v A, %v C WHERE A.product_id = '%v'`, enum.Product.TableName, enum.ProductDetail.TableName, ids.ProductIds[0])
	detail := shopping_client.ProductDetaiVo{}
	res := db.DB.Raw(sql).Scan(&detail)
	if res.Error != nil {
		return &shopping_client.ProductResponse{Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().ShoppingClientProvider)}, nil
	} else {
		if detail.SpecType == 2 {
			skuSql := fmt.Sprintf(`
			SELECT 
			sku_id,
			product_id,
			suk_name,
			sales,
			price,
			cost_price,
			vip_price,
			market_price,
			image,
			sn,
			stock,
			stock_warning,
			bar_code,
			weight,
			volume,
			type,
			quota,
			quota_show,
			coupon_id,
			brokerage,
			brokerage_two FROM %v WHERE product_id = '%v'`, enum.ProductSku.TableName, detail.ProductId)
			skuList := shopping_client.ProductDetaiVo{}.SkuList
			db.DB.Raw(skuSql).Scan(&skuList)
			if len(skuList) > 0 {
				detail.SkuList = skuList
			}
		}
		return &shopping_client.ProductResponse{
			Count:  1,
			Detail: &detail,
		}, nil
	}
}

// FindPageList 查询商品列表
func (p *productService) FindPageList(ctx context.Context, request *shopping_client.ProductPageAuthQuery) (*shopping_client.ProductResponse, error) {
	countSql := `SELECT COUNT(1) FROM ` + enum.Product.TableName + ` SPU`
	paramSql := ` WHERE SPU.del_status = 0 `
	if request.Params != nil {
		if request.Params.ProductName != "" {
			paramSql += fmt.Sprintf(` AND SPU.product_name like CONCAT('%%','%v','%%')`, request.Params.ProductName)
		}
		if request.Params.Keyword != "" {
			paramSql += fmt.Sprintf(`AND SPU.keyword like CONCAT('%%','%v','%%')`, request.Params.Keyword)
		}
	}
	var count int64
	db.DB.Raw(countSql + paramSql).Count(&count)
	if count > 0 {
		listSql := `
			SELECT 
				SPU.product_id,
				SPU.product_name,
				SPU.category_id,
				SPU.freight_temp_id,
				SPU.slider_image,
				SPU.recommend_image,
				SPU.price,
				SPU.cost_price,
				SPU.vip_price,
				SPU.market_price,
				SPU.image,
				SPU.unit_name,
				SPU.stock Stock,
				SPU.stock_warning,
				SPU.del_status,
				SPU.create_by,
				DATE_FORMAT(SPU.create_time,'%Y-%m-%d %T') create_time
			FROM ` + enum.Product.TableName + ` SPU`
		listSql += paramSql
		listSql += ` LIMIT ?,?`
		var productList []*shopping_client.ProductListVo
		db.DB.Raw(listSql, (request.PageIndex-1)*request.PageSize, request.PageSize).Scan(&productList)
		return &shopping_client.ProductResponse{Count: count, List: productList}, nil
	}
	return &shopping_client.ProductResponse{Count: count, Msg: "未查询到相关数据"}, nil
}
