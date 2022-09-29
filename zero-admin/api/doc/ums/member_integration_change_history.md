### 1. N/A

1. route definition

- Url: /api/member/integrationchangehistory/add
- Method: POST
- Request: `addIntegrationChangeHistoryReq`
- Response: `addIntegrationChangeHistoryResp`

2. request definition



```golang
type AddIntegrationChangeHistoryReq struct {
	MemberId int64 `json:"memberId"`
	CreateTime string `json:"createTime"`
	ChangeType int64 `json:"changeType"` // 改变类型：0-&gt;增加；1-&gt;减少
	ChangeCount int64 `json:"changeCount"` // 积分改变数量
	OperateMan string `json:"operateMan"` // 操作人员
	OperateNote string `json:"operateNote"` // 操作备注
	SourceType int64 `json:"sourceType"` // 积分来源：0-&gt;购物；1-&gt;管理员修改
}
```


3. response definition



```golang
type AddIntegrationChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/member/integrationchangehistory/list
- Method: POST
- Request: `ListIntegrationChangeHistoryReq`
- Response: `ListIntegrationChangeHistoryResp`

2. request definition



```golang
type ListIntegrationChangeHistoryReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListIntegrationChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtIntegrationChangeHistoryData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/member/integrationchangehistory/update
- Method: POST
- Request: `UpdateIntegrationChangeHistoryReq`
- Response: `UpdateIntegrationChangeHistoryResp`

2. request definition



```golang
type UpdateIntegrationChangeHistoryReq struct {
	Id int64 `json:"id"`
	MemberId int64 `json:"memberId"`
	CreateTime string `json:"createTime"`
	ChangeType int64 `json:"changeType"` // 改变类型：0-&gt;增加；1-&gt;减少
	ChangeCount int64 `json:"changeCount"` // 积分改变数量
	OperateMan string `json:"operateMan"` // 操作人员
	OperateNote string `json:"operateNote"` // 操作备注
	SourceType int64 `json:"sourceType"` // 积分来源：0-&gt;购物；1-&gt;管理员修改
}
```


3. response definition



```golang
type UpdateIntegrationChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/member/integrationchangehistory/delete
- Method: POST
- Request: `DeleteIntegrationChangeHistoryReq`
- Response: `DeleteIntegrationChangeHistoryResp`

2. request definition



```golang
type DeleteIntegrationChangeHistoryReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteIntegrationChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

