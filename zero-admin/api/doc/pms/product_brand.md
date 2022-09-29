### 1. N/A

1. route definition

- Url: /api/product/brand/add
- Method: POST
- Request: `addProductBrandReq`
- Response: `addProductBrandResp`

2. request definition



```golang
type AddProductBrandReq struct {
	Name string `json:"name"`
	FirstLetter string `json:"firstLetter"` // 首字母
	Sort int64 `json:"sort"`
	FactoryStatus int64 `json:"factoryStatus"` // 是否为品牌制造商：0-&gt;不是；1-&gt;是
	ShowStatus int64 `json:"showStatus"`
	ProductCount int64 `json:"productCount"` // 产品数量
	ProductCommentCount int64 `json:"productCommentCount"` // 产品评论数量
	Logo string `json:"logo"` // 品牌logo
	BigPic string `json:"bigPic"` // 专区大图
	BrandStory string `json:"brandStory"` // 品牌故事
}
```


3. response definition



```golang
type AddProductBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/product/brand/list
- Method: POST
- Request: `ListProductBrandReq`
- Response: `ListProductBrandResp`

2. request definition



```golang
type ListProductBrandReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListProductBrandResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtProductBrandData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 3. N/A

1. route definition

- Url: /api/product/brand/update
- Method: POST
- Request: `UpdateProductBrandReq`
- Response: `UpdateProductBrandResp`

2. request definition



```golang
type UpdateProductBrandReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	FirstLetter string `json:"firstLetter"` // 首字母
	Sort int64 `json:"sort"`
	FactoryStatus int64 `json:"factoryStatus"` // 是否为品牌制造商：0-&gt;不是；1-&gt;是
	ShowStatus int64 `json:"showStatus"`
	ProductCount int64 `json:"productCount"` // 产品数量
	ProductCommentCount int64 `json:"productCommentCount"` // 产品评论数量
	Logo string `json:"logo"` // 品牌logo
	BigPic string `json:"bigPic"` // 专区大图
	BrandStory string `json:"brandStory"` // 品牌故事
}
```


3. response definition



```golang
type UpdateProductBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/product/brand/delete
- Method: POST
- Request: `DeleteProductBrandReq`
- Response: `DeleteProductBrandResp`

2. request definition



```golang
type DeleteProductBrandReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteProductBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

