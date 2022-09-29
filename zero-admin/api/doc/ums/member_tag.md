### 1. N/A

1. route definition

- Url: /api/member/tag/add
- Method: POST
- Request: `addMemberTagReq`
- Response: `addMemberTagResp`

2. request definition



```golang
type AddMemberTagReq struct {
	Name string `json:"name"`
	FinishOrderCount int64 `json:"finishOrderCount"` // 自动打标签完成订单数量
	FinishOrderAmount float64 `json:"finishOrderAmount"` // 自动打标签完成订单金额
}
```


3. response definition



```golang
type AddMemberTagResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/member/tag/list
- Method: POST
- Request: `ListMemberTagReq`
- Response: `ListMemberTagResp`

2. request definition



```golang
type ListMemberTagReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberTagResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberTagData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/member/tag/update
- Method: POST
- Request: `UpdateMemberTagReq`
- Response: `UpdateMemberTagResp`

2. request definition



```golang
type UpdateMemberTagReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	FinishOrderCount int64 `json:"finishOrderCount"` // 自动打标签完成订单数量
	FinishOrderAmount float64 `json:"finishOrderAmount"` // 自动打标签完成订单金额
}
```


3. response definition



```golang
type UpdateMemberTagResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/member/tag/delete
- Method: POST
- Request: `DeleteMemberTagReq`
- Response: `DeleteMemberTagResp`

2. request definition



```golang
type DeleteMemberTagReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberTagResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

