### 1. N/A

1. route definition

- Url: /api/sms/homebrand/add
- Method: POST
- Request: `addHomeBrandReq`
- Response: `addHomeBrandResp`

2. request definition



```golang
type AddHomeBrandReq struct {
	BrandId int64 `json:"brandId"`
	BrandName string `json:"brandName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type AddHomeBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sms/homebrand/list
- Method: POST
- Request: `ListHomeBrandReq`
- Response: `ListHomeBrandResp`

2. request definition



```golang
type ListHomeBrandReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListHomeBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtHomeBrandData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sms/homebrand/update
- Method: POST
- Request: `UpdateHomeBrandReq`
- Response: `UpdateHomeBrandResp`

2. request definition



```golang
type UpdateHomeBrandReq struct {
	Id int64 `json:"id"`
	BrandId int64 `json:"brandId"`
	BrandName string `json:"brandName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type UpdateHomeBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sms/homebrand/delete
- Method: POST
- Request: `DeleteHomeBrandReq`
- Response: `DeleteHomeBrandResp`

2. request definition



```golang
type DeleteHomeBrandReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteHomeBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

