### 1. N/A

1. route definition

- Url: /api/order/cart/add
- Method: POST
- Request: `addCartItemReq`
- Response: `addCartItemResp`

2. request definition



```golang
type AddCartItemReq struct {
	ProductId int64 `json:"productId"`
	ProductSkuId int64 `json:"productSkuId"`
	MemberId int64 `json:"memberId"`
	Quantity int64 `json:"quantity"` // 购买数量
	Price float64 `json:"price"` // 添加到购物车的价格
	ProductPic string `json:"productPic"` // 商品主图
	ProductName string `json:"productName"` // 商品名称
	ProductSubTitle string `json:"productSubTitle"` // 商品副标题（卖点）
	ProductSkuCode string `json:"productSkuCode"` // 商品sku条码
	MemberNickname string `json:"memberNickName"` // 会员昵称
	DeleteStatus int64 `json:"deleteStatus"` // 是否删除
	ProductCategoryId int64 `json:"productCategoryId"` // 商品分类
	ProductBrand string `json:"productBrand"`
	ProductSn string `json:"productSn"`
	ProductAttr string `json:"productAttr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}
```


3. response definition



```golang
type AddCartItemResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/order/cart/list
- Method: POST
- Request: `ListCartItemReq`
- Response: `ListCartItemResp`

2. request definition



```golang
type ListCartItemReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListCartItemResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtCartItemData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/order/cart/update
- Method: POST
- Request: `UpdateCartItemReq`
- Response: `UpdateCartItemResp`

2. request definition



```golang
type UpdateCartItemReq struct {
	Id int64 `json:"id"`
	ProductId int64 `json:"productId"`
	ProductSkuId int64 `json:"productSkuId"`
	MemberId int64 `json:"memberId"`
	Quantity int64 `json:"quantity"` // 购买数量
	Price float64 `json:"price"` // 添加到购物车的价格
	ProductPic string `json:"productPic"` // 商品主图
	ProductName string `json:"productName"` // 商品名称
	ProductSubTitle string `json:"productSubTitle"` // 商品副标题（卖点）
	ProductSkuCode string `json:"productSkuCode"` // 商品sku条码
	MemberNickname string `json:"memberNickName"` // 会员昵称
	DeleteStatus int64 `json:"deleteStatus"` // 是否删除
	ProductCategoryId int64 `json:"productCategoryId"` // 商品分类
	ProductBrand string `json:"productBrand"`
	ProductSn string `json:"productSn"`
	ProductAttr string `json:"productAttr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}
```


3. response definition



```golang
type UpdateCartItemResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/order/cart/delete
- Method: POST
- Request: `DeleteCartItemReq`
- Response: `DeleteCartItemResp`

2. request definition



```golang
type DeleteCartItemReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteCartItemResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

