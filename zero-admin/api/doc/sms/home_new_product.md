### 1. N/A

1. route definition

- Url: /api/sms/homenewproduct/add
- Method: POST
- Request: `addHomeNewProductReq`
- Response: `addHomeNewProductResp`

2. request definition



```golang
type AddHomeNewProductReq struct {
	ProductId int64 `json:"productId"`
	ProductName string `json:"productName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type AddHomeNewProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sms/homenewproduct/list
- Method: POST
- Request: `ListHomeNewProductReq`
- Response: `ListHomeNewProductResp`

2. request definition



```golang
type ListHomeNewProductReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListHomeNewProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtHomeNewProductData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sms/homenewproduct/update
- Method: POST
- Request: `UpdateHomeNewProductReq`
- Response: `UpdateHomeNewProductResp`

2. request definition



```golang
type UpdateHomeNewProductReq struct {
	Id int64 `json:"id"`
	ProductId int64 `json:"productId"`
	ProductName string `json:"productName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type UpdateHomeNewProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sms/homenewproduct/delete
- Method: POST
- Request: `DeleteHomeNewProductReq`
- Response: `DeleteHomeNewProductResp`

2. request definition



```golang
type DeleteHomeNewProductReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteHomeNewProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

