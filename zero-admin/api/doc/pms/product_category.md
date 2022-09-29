### 1. N/A

1. route definition

- Url: /api/product/category/add
- Method: POST
- Request: `addProductCategoryReq`
- Response: `addProductCategoryResp`

2. request definition



```golang
type AddProductCategoryReq struct {
	ParentId int64 `json:"parentId"` // 上机分类的编号：0表示一级分类
	Name string `json:"name"`
	Level int64 `json:"level"` // 分类级别：0-&gt;1级；1-&gt;2级
	ProductCount int64 `json:"productCount"`
	ProductUnit string `json:"productUnit"`
	NavStatus int64 `json:"navStatus"` // 是否显示在导航栏：0-&gt;不显示；1-&gt;显示
	ShowStatus int64 `json:"showStatus"` // 显示状态：0-&gt;不显示；1-&gt;显示
	Sort int64 `json:"sort"`
	Icon string `json:"icon"` // 图标
	Keywords string `json:"keywords"`
	Description string `json:"description"` // 描述
}
```


3. response definition



```golang
type AddProductCategoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/product/category/list
- Method: POST
- Request: `ListProductCategoryReq`
- Response: `ListProductCategoryResp`

2. request definition



```golang
type ListProductCategoryReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListProductCategoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtProductCategoryData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/product/category/update
- Method: POST
- Request: `UpdateProductCategoryReq`
- Response: `UpdateProductCategoryResp`

2. request definition



```golang
type UpdateProductCategoryReq struct {
	Id int64 `json:"id"`
	ParentId int64 `json:"parentId"` // 上机分类的编号：0表示一级分类
	Name string `json:"name"`
	Level int64 `json:"level"` // 分类级别：0-&gt;1级；1-&gt;2级
	ProductCount int64 `json:"productCount"`
	ProductUnit string `json:"productUnit"`
	NavStatus int64 `json:"navStatus"` // 是否显示在导航栏：0-&gt;不显示；1-&gt;显示
	ShowStatus int64 `json:"showStatus"` // 显示状态：0-&gt;不显示；1-&gt;显示
	Sort int64 `json:"sort"`
	Icon string `json:"icon"` // 图标
	Keywords string `json:"keywords"`
	Description string `json:"description"` // 描述
}
```


3. response definition



```golang
type UpdateProductCategoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/product/category/delete
- Method: POST
- Request: `DeleteProductCategoryReq`
- Response: `DeleteProductCategoryResp`

2. request definition



```golang
type DeleteProductCategoryReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteProductCategoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

