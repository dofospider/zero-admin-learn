### 1. N/A

1. route definition

- Url: /api/sys/loginLog/list
- Method: POST
- Request: `ListLoginLogReq`
- Response: `ListLoginLogResp`

2. request definition



```golang
type ListLoginLogReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListLoginLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListLoginLogData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 2. N/A

1. route definition

- Url: /api/sys/loginLog/delete
- Method: POST
- Request: `DeleteLoginLogReq`
- Response: `DeleteLoginLogResp`

2. request definition



```golang
type DeleteLoginLogReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteLoginLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

