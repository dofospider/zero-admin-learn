### 1. N/A

1. route definition

- Url: /api/sys/dict/add
- Method: POST
- Request: `addDictReq`
- Response: `addDictResp`

2. request definition



```golang
type AddDictReq struct {
	Value string `json:"value"` // 数据值
	Label string `json:"label"` // 标签名
	Type string `json:"type"` // 类型
	Description string `json:"description"` // 描述
	Sort string `json:"sort"` // 排序（升序）
	Remarks string `json:"remarks"` // 备注信息
}
```


3. response definition



```golang
type AddDictResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sys/dict/list
- Method: POST
- Request: `ListDictReq`
- Response: `ListDictResp`

2. request definition



```golang
type ListDictReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
	Value string `json:"value,optional"`
	Label string `json:"label,optional"`
	DelFlag int64 `json:"delFlag,optional"`
	Type string `json:"type,optional "`
}
```


3. response definition



```golang
type ListDictResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListDictData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sys/dict/update
- Method: POST
- Request: `UpdateDictReq`
- Response: `UpdateDictResp`

2. request definition



```golang
type UpdateDictReq struct {
	Id int64 `json:"id"` // 编号
	Value string `json:"value"` // 数据值
	Label string `json:"label"` // 标签名
	Type string `json:"type"` // 类型
	Description string `json:"description"` // 描述
	Sort float64 `json:"sort"` // 排序（升序）
	Remarks string `json:"remarks"` // 备注信息
}
```


3. response definition



```golang
type UpdateDictResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sys/dict/delete
- Method: POST
- Request: `DeleteDictReq`
- Response: `DeleteDictResp`

2. request definition



```golang
type DeleteDictReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteDictResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

