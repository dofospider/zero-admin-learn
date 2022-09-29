### 1. N/A

1. route definition

- Url: /api/product/product/add
- Method: POST
- Request: `addProductReq`
- Response: `addProductResp`

2. request definition



```golang
type AddProductReq struct {
	BrandId int64 `json:"brandId"`
	ProductCategoryId int64 `json:"productCategoryId"`
	FeightTemplateId int64 `json:"feightTemplateId"`
	ProductAttributeCategoryId int64 `json:"productAttributeCategoryId"`
	Name string `json:"name"`
	Pic string `json:"pic"`
	ProductSn string `json:"productSn"` // 货号
	DeleteStatus int64 `json:"deleteStatus"` // 删除状态：0-&gt;未删除；1-&gt;已删除
	PublishStatus int64 `json:"publishStatus"` // 上架状态：0-&gt;下架；1-&gt;上架
	NewStatus int64 `json:"newStatus"` // 新品状态:0-&gt;不是新品；1-&gt;新品
	RecommandStatus int64 `json:"recommandStatus"` // 推荐状态；0-&gt;不推荐；1-&gt;推荐
	VerifyStatus int64 `json:"verifyStatus"` // 审核状态：0-&gt;未审核；1-&gt;审核通过
	Sort int64 `json:"sort"` // 排序
	Sale int64 `json:"sale"` // 销量
	Price float64 `json:"price"`
	PromotionPrice float64 `json:"promotionPrice"` // 促销价格
	GiftGrowth int64 `json:"giftGrowth"` // 赠送的成长值
	GiftPoint int64 `json:"giftPoint"` // 赠送的积分
	UsePointLimit int64 `json:"usePointLimit"` // 限制使用的积分数
	SubTitle string `json:"subTitle"` // 副标题
	Description string `json:"description"` // 商品描述
	OriginalPrice float64 `json:"originalPrice"` // 市场价
	Stock int64 `json:"stock"` // 库存
	LowStock int64 `json:"lowStock"` // 库存预警值
	Unit string `json:"unit"` // 单位
	Weight float64 `json:"weight"` // 商品重量，默认为克
	PreviewStatus int64 `json:"previewStatus"` // 是否为预告商品：0-&gt;不是；1-&gt;是
	ServiceIds string `json:"serviceIds"` // 以逗号分割的产品服务：1-&gt;无忧退货；2-&gt;快速退款；3-&gt;免费包邮
	Keywords string `json:"keywords"`
	Note string `json:"note"`
	AlbumPics string `json:"albumPics"` // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle string `json:"detailTitle"`
	DetailDesc string `json:"detailDesc"`
	DetailHtml string `json:"detailHtml"` // 产品详情网页内容
	DetailMobileHtml string `json:"detailMobileHtml"` // 移动端网页详情
	PromotionStartTime string `json:"promotionStartTime"` // 促销开始时间
	PromotionEndTime string `json:"promotionEndTime"` // 促销结束时间
	PromotionPerLimit int64 `json:"promotionPerLimit"` // 活动限购数量
	PromotionType int64 `json:"promotionType"` // 促销类型：0-&gt;没有促销使用原价;1-&gt;使用促销价；2-&gt;使用会员价；3-&gt;使用阶梯价格；4-&gt;使用满减价格；5-&gt;限时购
	BrandName string `json:"brandName"` // 品牌名称
	ProductCategoryName string `json:"productCategoryName"` // 商品分类名称
}
```


3. response definition



```golang
type AddProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/product/product/list
- Method: POST
- Request: `ListProductReq`
- Response: `ListProductResp`

2. request definition



```golang
type ListProductReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtProductData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/product/product/update
- Method: POST
- Request: `UpdateProductReq`
- Response: `UpdateProductResp`

2. request definition



```golang
type UpdateProductReq struct {
	Id int64 `json:"id"`
	BrandId int64 `json:"brandId"`
	ProductCategoryId int64 `json:"productCategoryId"`
	FeightTemplateId int64 `json:"feightTemplateId"`
	ProductAttributeCategoryId int64 `json:"productAttributeCategoryId"`
	Name string `json:"name"`
	Pic string `json:"pic"`
	ProductSn string `json:"productSn"` // 货号
	DeleteStatus int64 `json:"deleteStatus"` // 删除状态：0-&gt;未删除；1-&gt;已删除
	PublishStatus int64 `json:"publishStatus"` // 上架状态：0-&gt;下架；1-&gt;上架
	NewStatus int64 `json:"newStatus"` // 新品状态:0-&gt;不是新品；1-&gt;新品
	RecommandStatus int64 `json:"recommandStatus"` // 推荐状态；0-&gt;不推荐；1-&gt;推荐
	VerifyStatus int64 `json:"verifyStatus"` // 审核状态：0-&gt;未审核；1-&gt;审核通过
	Sort int64 `json:"sort"` // 排序
	Sale int64 `json:"sale"` // 销量
	Price float64 `json:"price"`
	PromotionPrice float64 `json:"promotionPrice"` // 促销价格
	GiftGrowth int64 `json:"giftGrowth"` // 赠送的成长值
	GiftPoint int64 `json:"giftPoint"` // 赠送的积分
	UsePointLimit int64 `json:"usePointLimit"` // 限制使用的积分数
	SubTitle string `json:"subTitle"` // 副标题
	Description string `json:"description"` // 商品描述
	OriginalPrice float64 `json:"originalPrice"` // 市场价
	Stock int64 `json:"stock"` // 库存
	LowStock int64 `json:"lowStock"` // 库存预警值
	Unit string `json:"unit"` // 单位
	Weight float64 `json:"weight"` // 商品重量，默认为克
	PreviewStatus int64 `json:"previewStatus"` // 是否为预告商品：0-&gt;不是；1-&gt;是
	ServiceIds string `json:"serviceIds"` // 以逗号分割的产品服务：1-&gt;无忧退货；2-&gt;快速退款；3-&gt;免费包邮
	Keywords string `json:"keywords"`
	Note string `json:"note"`
	AlbumPics string `json:"albumPics"` // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle string `json:"detailTitle"`
	DetailDesc string `json:"detailDesc"`
	DetailHtml string `json:"detailHtml"` // 产品详情网页内容
	DetailMobileHtml string `json:"detailMobileHtml"` // 移动端网页详情
	PromotionStartTime string `json:"promotionStartTime"` // 促销开始时间
	PromotionEndTime string `json:"promotionEndTime"` // 促销结束时间
	PromotionPerLimit int64 `json:"promotionPerLimit"` // 活动限购数量
	PromotionType int64 `json:"promotionType"` // 促销类型：0-&gt;没有促销使用原价;1-&gt;使用促销价；2-&gt;使用会员价；3-&gt;使用阶梯价格；4-&gt;使用满减价格；5-&gt;限时购
	BrandName string `json:"brandName"` // 品牌名称
	ProductCategoryName string `json:"productCategoryName"` // 商品分类名称
}
```


3. response definition



```golang
type UpdateProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/product/product/delete
- Method: POST
- Request: `DeleteProductReq`
- Response: `DeleteProductResp`

2. request definition



```golang
type DeleteProductReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

