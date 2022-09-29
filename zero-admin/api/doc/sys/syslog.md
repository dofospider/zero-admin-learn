### 1. N/A

1. route definition

- Url: /api/sys/sysLog/list
- Method: POST
- Request: `ListSysLogReq`
- Response: `ListSysLogResp`

2. request definition



```golang
type ListSysLogReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListSysLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListSysLogData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 2. N/A

1. route definition

- Url: /api/sys/sysLog/delete
- Method: POST
- Request: `DeleteSysLogReq`
- Response: `DeleteSysLogResp`

2. request definition



```golang
type DeleteSysLogReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteSysLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

