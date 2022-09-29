### 1. N/A

1. route definition

- Url: /api/product/skustock/add
- Method: POST
- Request: `addSkuStockReq`
- Response: `addSkuStockResp`

2. request definition



```golang
type AddSkuStockReq struct {
	ProductId int64 `json:"productId"`
	SkuCode string `json:"skuCode"` // sku编码
	Price float64 `json:"price"`
	Stock int64 `json:"stock"` // 库存
	LowStock int64 `json:"lowStock"` // 预警库存
	Pic string `json:"pic"` // 展示图片
	Sale int64 `json:"sale"` // 销量
	PromotionPrice float64 `json:"promotionPrice"` // 单品促销价格
	LockStock int64 `json:"lockStock"` // 锁定库存
	SpData string `json:"spData"` // 商品销售属性，json格式
}
```


3. response definition



```golang
type AddSkuStockResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/product/skustock/list
- Method: POST
- Request: `ListSkuStockReq`
- Response: `ListSkuStockResp`

2. request definition



```golang
type ListSkuStockReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListSkuStockResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtSkuStockData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 3. N/A

1. route definition

- Url: /api/product/skustock/update
- Method: POST
- Request: `UpdateSkuStockReq`
- Response: `UpdateSkuStockResp`

2. request definition



```golang
type UpdateSkuStockReq struct {
	Id int64 `json:"id"`
	ProductId int64 `json:"productId"`
	SkuCode string `json:"skuCode"` // sku编码
	Price float64 `json:"price"`
	Stock int64 `json:"stock"` // 库存
	LowStock int64 `json:"lowStock"` // 预警库存
	Pic string `json:"pic"` // 展示图片
	Sale int64 `json:"sale"` // 销量
	PromotionPrice float64 `json:"promotionPrice"` // 单品促销价格
	LockStock int64 `json:"lockStock"` // 锁定库存
	SpData string `json:"spData"` // 商品销售属性，json格式
}
```


3. response definition



```golang
type UpdateSkuStockResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/product/skustock/delete
- Method: POST
- Request: `DeleteSkuStockReq`
- Response: `DeleteSkuStockResp`

2. request definition



```golang
type DeleteSkuStockReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteSkuStockResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

