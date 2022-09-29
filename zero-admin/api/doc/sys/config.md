### 1. N/A

1. route definition

- Url: /api/sys/config/add
- Method: POST
- Request: `addConfigReq`
- Response: `addConfigResp`

2. request definition



```golang
type AddConfigReq struct {
	Value string `json:"value"` // 数据值
	Label string `json:"label"` // 标签名
	Type string `json:"type"` // 类型
	Description string `json:"description"` // 描述
	Sort int64 `json:"sort"` // 排序（升序）
	Remarks string `json:"remarks"` // 备注信息
}
```


3. response definition



```golang
type AddConfigResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sys/config/list
- Method: POST
- Request: `ListConfigReq`
- Response: `ListConfigResp`

2. request definition



```golang
type ListConfigReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListConfigResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListConfigData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sys/config/update
- Method: POST
- Request: `UpdateConfigReq`
- Response: `UpdateConfigResp`

2. request definition



```golang
type UpdateConfigReq struct {
	Id int64 `json:"id"` // 编号
	Value string `json:"value"` // 数据值
	Label string `json:"label"` // 标签名
	Type string `json:"type"` // 类型
	Description string `json:"description"` // 描述
	Sort int64 `json:"sort"` // 排序（升序）
	Remarks string `json:"remarks"` // 备注信息
}
```


3. response definition



```golang
type UpdateConfigResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sys/config/delete
- Method: POST
- Request: `DeleteConfigReq`
- Response: `DeleteConfigResp`

2. request definition



```golang
type DeleteConfigReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteConfigResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

