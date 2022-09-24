### 1. N/A

1. route definition

- Url: /api/sys/job/add
- Method: POST
- Request: `addJobReq`
- Response: `addJobResp`

2. request definition



```golang
type AddJobReq struct {
	JobName string `json:"jobName"` // 职位名称
	Remarks string `json:"remarks"` // 备注
	OrderNum int64 `json:"orderNum"` // 排序
}
```


3. response definition



```golang
type AddJobResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sys/job/list
- Method: POST
- Request: `ListJobReq`
- Response: `ListJobResp`

2. request definition



```golang
type ListJobReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
	JobName string `json:"jobName,optional"`
}
```


3. response definition



```golang
type ListJobResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListJobData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sys/job/update
- Method: POST
- Request: `UpdateJobReq`
- Response: `UpdateJobResp`

2. request definition



```golang
type UpdateJobReq struct {
	Id int64 `json:"id"` // 编号
	JobName string `json:"jobName"` // 职位名称
	OrderNum int64 `json:"orderNum"` // 排序
	Remarks string `json:"remarks"` // 备注
}
```


3. response definition



```golang
type UpdateJobResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sys/job/delete
- Method: POST
- Request: `DeleteJobReq`
- Response: `DeleteJobResp`

2. request definition



```golang
type DeleteJobReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteJobResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

