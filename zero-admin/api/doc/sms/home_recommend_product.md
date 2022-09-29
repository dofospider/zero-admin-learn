### 1. N/A

1. route definition

- Url: /api/sms/homerecommendproduct/add
- Method: POST
- Request: `addHomeRecommendProductReq`
- Response: `addHomeRecommendProductResp`

2. request definition



```golang
type AddHomeRecommendProductReq struct {
	ProductId int64 `json:"productId"`
	ProductName string `json:"productName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type AddHomeRecommendProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sms/homerecommendproduct/list
- Method: POST
- Request: `ListHomeRecommendProductReq`
- Response: `ListHomeRecommendProductResp`

2. request definition



```golang
type ListHomeRecommendProductReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListHomeRecommendProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtHomeRecommendProductData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sms/homerecommendproduct/update
- Method: POST
- Request: `UpdateHomeRecommendProductReq`
- Response: `UpdateHomeRecommendProductResp`

2. request definition



```golang
type UpdateHomeRecommendProductReq struct {
	Id int64 `json:"id"`
	ProductId int64 `json:"productId"`
	ProductName string `json:"productName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type UpdateHomeRecommendProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sms/homerecommendproduct/delete
- Method: POST
- Request: `DeleteHomeRecommendProductReq`
- Response: `DeleteHomeRecommendProductResp`

2. request definition



```golang
type DeleteHomeRecommendProductReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteHomeRecommendProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

