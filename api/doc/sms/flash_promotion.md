### 1. N/A

1. route definition

- Url: /api/sms/flashpromotion/add
- Method: POST
- Request: `addFlashPromotionReq`
- Response: `addFlashPromotionResp`

2. request definition



```golang
type AddFlashPromotionReq struct {
	Title string `json:"title"`
	StartDate string `json:"startDate"` // 开始日期
	EndDate string `json:"endDate"` // 结束日期
	Status int64 `json:"status"` // 上下线状态
	CreateTime string `json:"createTime"` // 秒杀时间段名称
}
```


3. response definition



```golang
type AddFlashPromotionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sms/flashpromotion/list
- Method: POST
- Request: `ListFlashPromotionReq`
- Response: `ListFlashPromotionResp`

2. request definition



```golang
type ListFlashPromotionReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListFlashPromotionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtFlashPromotionData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sms/flashpromotion/update
- Method: POST
- Request: `UpdateFlashPromotionReq`
- Response: `UpdateFlashPromotionResp`

2. request definition



```golang
type UpdateFlashPromotionReq struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	StartDate string `json:"startDate"` // 开始日期
	EndDate string `json:"endDate"` // 结束日期
	Status int64 `json:"status"` // 上下线状态
	CreateTime string `json:"createTime"` // 秒杀时间段名称
}
```


3. response definition



```golang
type UpdateFlashPromotionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sms/flashpromotion/delete
- Method: POST
- Request: `DeleteFlashPromotionReq`
- Response: `DeleteFlashPromotionResp`

2. request definition



```golang
type DeleteFlashPromotionReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteFlashPromotionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

