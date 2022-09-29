### 1. N/A

1. route definition

- Url: /api/sms/flashpromotionsession/add
- Method: POST
- Request: `addFlashPromotionSessionReq`
- Response: `addFlashPromotionSessionResp`

2. request definition



```golang
type AddFlashPromotionSessionReq struct {
	Name string `json:"name"` // 场次名称
	StartTime string `json:"startTime"` // 每日开始时间
	EndTime string `json:"endTime"` // 每日结束时间
	Status int64 `json:"status"` // 启用状态：0-&gt;不启用；1-&gt;启用
	CreateTime string `json:"createTime"` // 创建时间
}
```


3. response definition



```golang
type AddFlashPromotionSessionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sms/flashpromotionsession/list
- Method: POST
- Request: `ListFlashPromotionSessionReq`
- Response: `ListFlashPromotionSessionResp`

2. request definition



```golang
type ListFlashPromotionSessionReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListFlashPromotionSessionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtFlashPromotionSessionData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sms/flashpromotionsession/update
- Method: POST
- Request: `UpdateFlashPromotionSessionReq`
- Response: `UpdateFlashPromotionSessionResp`

2. request definition



```golang
type UpdateFlashPromotionSessionReq struct {
	Id int64 `json:"id"` // 编号
	Name string `json:"name"` // 场次名称
	StartTime string `json:"startTime"` // 每日开始时间
	EndTime string `json:"endTime"` // 每日结束时间
	Status int64 `json:"status"` // 启用状态：0-&gt;不启用；1-&gt;启用
	CreateTime string `json:"createTime"` // 创建时间
}
```


3. response definition



```golang
type UpdateFlashPromotionSessionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sms/flashpromotionsession/delete
- Method: POST
- Request: `DeleteFlashPromotionSessionReq`
- Response: `DeleteFlashPromotionSessionResp`

2. request definition



```golang
type DeleteFlashPromotionSessionReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteFlashPromotionSessionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

