### 1. N/A

1. route definition

- Url: /api/sms/flashpromotionlog/add
- Method: POST
- Request: `addFlashPromotionLogReq`
- Response: `addFlashPromotionLogResp`

2. request definition



```golang
type AddFlashPromotionLogReq struct {
	MemberId int64 `json:"memberId"`
	ProductId int64 `json:"productId"`
	MemberPhone string `json:"memberPhone"`
	ProductName string `json:"productName"`
	SubscribeTime string `json:"subscribeTime"` // 会员订阅时间
	SendTime string `json:"sendTime"`
}
```


3. response definition



```golang
type AddFlashPromotionLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sms/flashpromotionlog/list
- Method: POST
- Request: `ListFlashPromotionLogReq`
- Response: `ListFlashPromotionLogResp`

2. request definition



```golang
type ListFlashPromotionLogReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListFlashPromotionLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtFlashPromotionLogData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sms/flashpromotionlog/update
- Method: POST
- Request: `UpdateFlashPromotionLogReq`
- Response: `UpdateFlashPromotionLogResp`

2. request definition



```golang
type UpdateFlashPromotionLogReq struct {
	Id int64 `json:"id"`
	MemberId int64 `json:"memberId"`
	ProductId int64 `json:"productId"`
	MemberPhone string `json:"memberPhone"`
	ProductName string `json:"productName"`
	SubscribeTime string `json:"subscribeTime"` // 会员订阅时间
	SendTime string `json:"sendTime"`
}
```


3. response definition



```golang
type UpdateFlashPromotionLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sms/flashpromotionlog/delete
- Method: POST
- Request: `DeleteFlashPromotionLogReq`
- Response: `DeleteFlashPromotionLogResp`

2. request definition



```golang
type DeleteFlashPromotionLogReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteFlashPromotionLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

