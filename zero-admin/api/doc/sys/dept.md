### 1. N/A

1. route definition

- Url: /api/sys/dept/add
- Method: POST
- Request: `addDeptReq`
- Response: `addDeptResp`

2. request definition



```golang
type AddDeptReq struct {
	Name string `json:"name"` // 机构名称
	ParentId int64 `json:"parentId,optional"` // 上级机构ID，一级机构为0
	OrderNum int64 `json:"orderNum"` // 排序
}
```


3. response definition



```golang
type AddDeptResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sys/dept/list
- Method: POST
- Request: `ListDeptReq`
- Response: `ListDeptResp`

2. request definition



```golang
type ListDeptReq struct {
	Name string `json:"name,optional"`
	CreateBy string `json:"createBy,optional"`
}
```


3. response definition



```golang
type ListDeptResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data []*ListDeptData `json:"data"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sys/dept/update
- Method: POST
- Request: `UpdateDeptReq`
- Response: `UpdateDeptResp`

2. request definition



```golang
type UpdateDeptReq struct {
	Id int64 `json:"id"` // 编号
	Name string `json:"name"` // 机构名称
	ParentId int64 `json:"parentId"` // 上级机构ID，一级机构为0
	OrderNum int64 `json:"orderNum"` // 排序
}
```


3. response definition



```golang
type UpdateDeptResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sys/dept/delete
- Method: POST
- Request: `DeleteDeptReq`
- Response: `DeleteDeptResp`

2. request definition



```golang
type DeleteDeptReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteDeptResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

