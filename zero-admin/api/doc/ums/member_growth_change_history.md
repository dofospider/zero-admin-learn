### 1. N/A

1. route definition

- Url: /api/member/growthchangehistory/add
- Method: POST
- Request: `addGrowthChangeHistoryReq`
- Response: `addGrowthChangeHistoryResp`

2. request definition



```golang
type AddGrowthChangeHistoryReq struct {
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
type AddGrowthChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/member/growthchangehistory/list
- Method: POST
- Request: `ListGrowthChangeHistoryReq`
- Response: `ListGrowthChangeHistoryResp`

2. request definition



```golang
type ListGrowthChangeHistoryReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListGrowthChangeHistoryResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtGrowthChangeHistoryData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 3. N/A

1. route definition

- Url: /api/member/growthchangehistory/update
- Method: POST
- Request: `UpdateGrowthChangeHistoryReq`
- Response: `UpdateGrowthChangeHistoryResp`

2. request definition



```golang
type UpdateGrowthChangeHistoryReq struct {
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
type UpdateGrowthChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/member/growthchangehistory/delete
- Method: POST
- Request: `DeleteGrowthChangeHistoryReq`
- Response: `DeleteGrowthChangeHistoryResp`

2. request definition



```golang
type DeleteGrowthChangeHistoryReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteGrowthChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

