package mshop

import (
	"errors"
	"fmt"
	"sports_service/models"
	tc "sports_service/tools/tencentCloud"
)

type ProductSimpleInfo struct {
	Id           int          `json:"id" xorm:"not null pk autoincr comment('商品id') INT(11)"`
	ProductName  string       `json:"product_name" xorm:"not null default '' comment('商品名称') VARCHAR(255)"`
	ProductImage tc.BucketURI `json:"product_image" xorm:"not null default '' comment('商品主图路径') VARCHAR(512)"`
	Status       int          `json:"status" xorm:"not null default 0 comment('商品状态（0. 正常 1. 下架）') TINYINT(2)"`
	IsFreeShip   int          `json:"is_free_ship" xorm:"not null default 0 comment('是否免邮 0 免邮') TINYINT(2)"`
	Introduction []string     `json:"introduction" xorm:"not null default '' comment('促销语') VARCHAR(255)"`
	Keywords     string       `json:"keywords" xorm:"not null default '' comment('关键词') VARCHAR(255)"`
	Sortorder    int          `json:"sortorder" xorm:"not null default 0 comment('排序') index INT(11)"`
	VideoUrl     tc.BucketURI `json:"video_url" xorm:"not null default '' comment('视频地址') VARCHAR(512)"`
	SaleNum      int          `json:"sale_num" xorm:"not null default 0 comment('销量') INT(11)"`
	CurPrice     int          `json:"cur_price" xorm:"not null default 0 comment('商品价格（分）') INT(10)"`
	MarketPrice  int          `json:"market_price" xorm:"not null default 0 comment('划线价格（分）') INT(10)"`
	SkuNum       int          `json:"sku_num"`
	CategoryId   int          `json:"category_id"`
	CategoryName string       `json:"category_name"`
}

type ProductDetailInfo struct {
	Id             int               `json:"sku_id"`
	ProductId      int64             `json:"product_id"`
	Title          string            `json:"title"`
	SkuImage       tc.BucketURI      `json:"sku_image"`
	SkuNo          string            `json:"sku_no"`
	Images         []tc.BucketURI    `json:"images"`
	CurPrice       int               `json:"cur_price"`
	MarketPrice    int               `json:"market_price"`
	VideoUrl       tc.BucketURI      `json:"video_url"`
	IsFreeShip     int               `json:"is_free_ship"`
	DiscountPrice  int               `json:"discount_price"`
	StartTime      int64             `json:"start_time"`
	EndTime        int64             `json:"end_time"`
	RemainDuration int64             `json:"remain_duration"`  // 活动剩余时长
	HasActivities  int32             `json:"has_activities"`   // 1 有活动
	ProductDetail  []tc.BucketURI    `json:"product_detail"`   // 商品详情 长图/描述 多张
	OwnSpec        []OwnSpec         `json:"own_spec"`         // 商品实体的特有规格参数
	AfterService   []*AfterService   `json:"after_service"`    // 服务
	Specifications []SpecInfo        `json:"specifications"`   // 全部规格参数
	SpecTemplate   []SpecTemplate    `json:"spec_template"`    // 特有规格参数
	Indexes        string            `json:"indexes"`          // 特有规格属性在商品属性模板中的对应下标组合
	Stock          int               `json:"stock"`            // 库存
	MaxBuy         int               `json:"max_buy"`          // 限购 0 表示无限制
	MinBuy         int               `json:"min_buy"`          // 起购数
	SaleNum        int               `json:"sale_num"`         // 销量
	ProductCartNum int64             `json:"product_cart_num"` // 购物车数量
	Introduction   []string          `json:"introduction"`     // 促销语
	Postage        int               `json:"postage"`          // 邮费
	SkuList        []*ProductSkuInfo `json:"sku_list"`
}

type AfterService struct {
	Id       int64        `json:"id"`
	Service  string       `json:"service"`  // 服务名称
	Icon     tc.BucketURI `json:"icon"`     // 图标
	Describe string       `json:"describe"` // 描述
}

// 商品实体的特有规格参数
type OwnSpec struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

// 规格信息
type SpecInfo struct {
	Group  string `json:"group"` // 规格组名称
	Params []Spec `json:"params"`
}

type Spec struct {
	Key        string   `json:"key"`                 // 属性名称
	Val        string   `json:"val"`                 // 属性值
	Searchable int32    `json:"searchable"`          // 是否作为搜索条件  1 是
	Global     int32    `json:"global"`              // 是否为全局属性 1 是
	Unit       string   `json:"unit,omitempty"`      // 单位
	Numerical  int32    `json:"numerical,omitempty"` // 参数是否为数值类型 1是
	Icon       string   `json:"icon,omitempty"`      // icon图标
	Options    []string `json:"options,omitempty"`   // 参数选项
}

// 特有规格参数选项
type SpecTemplate struct {
	Key     string   `json:"key"`
	Options []string `json:"options"`
}

const (
	GET_SPU_LIST_BY_CATEGORY = "SELECT p.* FROM products AS p LEFT JOIN product_category_related AS pc " +
		"ON p.id = pc.product_id WHERE pc.category_id = ? AND p.status = 0 AND p.is_delete = 0 " +
		"ORDER BY p.sale_num DESC, p.is_top DESC, p.is_recommend DESC, p.is_cream DESC, p.sortorder DESC, p.id DESC LIMIT ?, ?"
)

// 通过分类获取spu列表
// sortType 0 销量倒序 1 销量正序 2 价格倒序 3 价格正序
func (m *ShopModel) GetSpuListByCategory(categoryId, sortType string, offset, size int) ([]*ProductSimpleInfo, error) {
	var list []*ProductSimpleInfo
	sql := "SELECT p.* FROM products AS p LEFT JOIN product_category_related AS pc " +
		"ON p.id = pc.product_id WHERE pc.category_id = ? AND p.status = 0 AND p.is_delete = 0 ORDER BY "

	switch sortType {
	case "0":
		sql += "p.sale_num DESC, "
	case "1":
		sql += "p.sale_num ASC, "
	case "2":
		sql += "p.cur_price DESC, "
	case "3":
		sql += "p.cur_price ASC, "
	default:
		sql += "p.sale_num DESC, "
	}

	sql += "p.is_top DESC, p.is_recommend DESC, p.is_cream DESC, p.sortorder DESC, p.id DESC LIMIT ?, ?"
	if err := m.Engine.SQL(sql, categoryId, offset, size).Find(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 获取所有商品spu
// sortType 0 销量倒序 1 销量正序 2 价格倒序 3 价格正序
func (m *ShopModel) GetAllSpu(sortType, keyword string, offset, size int) ([]*ProductSimpleInfo, error) {
	var list []*ProductSimpleInfo
	engine := m.Engine.Table(&models.Products{})
	engine.Where("status=0 AND is_delete=0")
	if keyword != "" {
		engine.Where("product_name like '%" + keyword + "%'")
	}

	switch sortType {
	case "0":
		engine.OrderBy("sale_num DESC, is_top DESC, is_recommend DESC, is_cream DESC, sortorder DESC")
	case "1":
		engine.OrderBy("sale_num ASC, is_top DESC, is_recommend DESC, is_cream DESC, sortorder DESC")
	case "2":
		engine.OrderBy("cur_price DESC, is_top DESC, is_recommend DESC, is_cream DESC, sortorder DESC")
	case "3":
		engine.OrderBy("cur_price ASC, is_top DESC, is_recommend DESC, is_cream DESC, sortorder DESC")
	default:
		engine.OrderBy("sale_num DESC, is_top DESC, is_recommend DESC, is_cream DESC, sortorder DESC")
	}

	if err := engine.Limit(size, offset).Find(&list); err != nil {
		return nil, err
	}

	return list, nil
}

const (
	GET_RECOMMEND_PRODUCTS = "SELECT DISTINCT(p.id), p.* FROM `products` AS p JOIN (SELECT RAND() * " +
		"((SELECT MAX(id) FROM `products`)-(SELECT MIN(id) FROM `products`)) + (SELECT MIN(id) FROM `products`) AS id) AS p2 " +
		"WHERE p.id >= p2.id-1 AND p.id != ? AND p.status=0 AND p.is_delete=0 ORDER BY p.id LIMIT ?"
)

func (m *ShopModel) GetRecommendProducts(productId string, limit int) ([]ProductSimpleInfo, error) {
	var list []ProductSimpleInfo
	if err := m.Engine.SQL(GET_RECOMMEND_PRODUCTS, productId, limit).Find(&list); err != nil {
		return list, err
	}

	return list, nil
}

// 获取sku列表
func (m *ShopModel) GetProductSkuList(productId string) ([]*ProductSkuInfo, error) {
	var list []*ProductSkuInfo
	if err := m.Engine.Table(&models.ProductSku{}).Where("product_id=? AND is_delete = 0", productId).Desc("sortorder").Find(&list); err != nil {
		return list, err
	}

	return list, nil
}

func (m *ShopModel) GetProductSku(condition string) (*models.ProductSku, error) {
	sku := &models.ProductSku{}
	ok, err := m.Engine.Where(condition).Get(sku)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("sku not found")
	}

	return sku, nil
}

// 获取权重最高的sku
func (m *ShopModel) GetProductSkuBySort(productId string) (*models.ProductSku, error) {
	sku := &models.ProductSku{}
	ok, err := m.Engine.Where("product_id=? AND status=0", productId).Desc("sortorder").Get(sku)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("sku not found")
	}

	return sku, nil
}

// 获取商品spu
func (m *ShopModel) GetProductSpu(productId string) (*models.Products, error) {
	spu := &models.Products{}
	ok, err := m.Engine.Where("id=?", productId).Get(spu)
	if !ok || err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("spu not found")
	}

	return spu, nil
}

// indexes 特有规格属性在商品属性模板中的对应下标组合
func (m *ShopModel) GetProductDetail(productId, indexes string) (*ProductDetailInfo, error) {
	sql := "SELECT ps.*, p.introduction, p.specifications, p.spec_template, p.video_url, p.sale_num, p.product_detail FROM product_sku AS ps " +
		"LEFT JOIN products AS p ON ps.product_id = p.id WHERE ps.status=0 AND ps.is_delete=0 AND ps.product_id=? "

	if indexes != "" {
		sql += fmt.Sprintf(" AND ps.indexes='%s'", indexes)
	} else {
		sql += "ORDER BY ps.sortorder DESC, id DESC LIMIT 1"
	}

	detail := &ProductDetailInfo{}
	ok, err := m.Engine.SQL(sql, productId).Get(detail)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("get detail fail")
	}

	return detail, nil
}

const (
	GET_PRODUCT_SERVICE_INFO = "SELECT svc.service, svc.id, svc.icon, svc.describe FROM shop_service_conf AS svc " +
		"LEFT JOIN product_service AS ps ON svc.id=ps.service_id WHERE product_id=?"
)

// 获取商品服务信息
func (m *ShopModel) GetProductServiceInfo(productId string) ([]*AfterService, error) {
	var list []*AfterService
	if err := m.Engine.SQL(GET_PRODUCT_SERVICE_INFO, productId).Find(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 获取商品sku库存
func (m *ShopModel) GetProductSkuStock(skuId string) (*models.ProductSkuStock, error) {
	stock := &models.ProductSkuStock{}
	ok, err := m.Engine.Where("sku_id=?", skuId).Get(stock)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("get detail fail")
	}

	return stock, nil
}

//const (
//	UPDATE_PRODUCT_SKU_STOCK = "UPDATE `product_sku_stock` SET `stock`= `stock`-? WHERE sku_id=? AND stock >= ? AND stock > 0"
//)
// 更新商品sku库存
//func (m *ShopModel) UpdateProductSkuStock(skuId string, count int) (int64, error) {
//	res, err := m.Engine.Exec(UPDATE_PRODUCT_SKU_STOCK, count, skuId, count)
//	if err != nil {
//		return 0, err
//	}
//
//	affected, err := res.RowsAffected()
//	if err != nil {
//		return 0, err
//	}
//
//	return affected, nil
//}

const (
	UPDATE_PRODUCT_SKU_STOCK = "UPDATE `product_sku_stock` SET `purchased_num`= `purchased_num`+ ? WHERE sku_id=? AND " +
		"`stock` >= `purchased_num`+ ? AND `purchased_num` + ? >= 0 LIMIT 1"
)

func (m *ShopModel) UpdateProductSkuStock(skuId string, count int) (int64, error) {
	res, err := m.Engine.Exec(UPDATE_PRODUCT_SKU_STOCK, count, skuId, count, count)
	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}

// 更新商品信息
func (m *ShopModel) UpdateProductInfo(condition, cols string, product *models.Products) (int64, error) {
	engine := m.Engine
	if condition != "" {
		engine.Where(condition)
	}

	if cols != "" {
		engine.Cols(cols)
	}

	return engine.Update(product)
}
