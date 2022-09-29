### 1. N/A

1. route definition

- Url: /api/member/task/add
- Method: POST
- Request: `addMemberTaskReq`
- Response: `addMemberTaskResp`

2. request definition



```golang
type AddMemberTaskReq struct {
	Name string `json:"name"`
	Growth int64 `json:"growth"` // 赠送成长值
	Intergration int64 `json:"intergration"` // 赠送积分
	Type int64 `json:"type"` // 任务类型：0-&gt;新手任务；1-&gt;日常任务
}
```


3. response definition



```golang
type AddMemberTaskResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/member/task/list
- Method: POST
- Request: `ListMemberTaskReq`
- Response: `ListMemberTaskResp`

2. request definition



```golang
type ListMemberTaskReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberTaskResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberTaskData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/member/task/update
- Method: POST
- Request: `UpdateMemberTaskReq`
- Response: `UpdateMemberTaskResp`

2. request definition



```golang
type UpdateMemberTaskReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Growth int64 `json:"growth"` // 赠送成长值
	Intergration int64 `json:"intergration"` // 赠送积分
	Type int64 `json:"type"` // 任务类型：0-&gt;新手任务；1-&gt;日常任务
}
```


3. response definition



```golang
type UpdateMemberTaskResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/member/task/delete
- Method: POST
- Request: `DeleteMemberTaskReq`
- Response: `DeleteMemberTaskResp`

2. request definition



```golang
type DeleteMemberTaskReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberTaskResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

