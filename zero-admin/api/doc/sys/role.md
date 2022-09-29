### 1. N/A

1. route definition

- Url: /api/sys/role/add
- Method: POST
- Request: `addRoleReq`
- Response: `addRoleResp`

2. request definition



```golang
type AddRoleReq struct {
	Name string `json:"name"` // 角色名称
	Remark string `json:"remark"` // 备注
	Status int64 `json:"status,optional"` // 状态
}
```


3. response definition



```golang
type AddRoleResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sys/role/list
- Method: POST
- Request: `ListRoleReq`
- Response: `ListRoleResp`

2. request definition



```golang
type ListRoleReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
	Name string `json:"name,optional "`
	Status int64 `json:"delFlag,optional "`
}
```


3. response definition



```golang
type ListRoleResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListRoleData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sys/role/update
- Method: POST
- Request: `UpdateRoleReq`
- Response: `UpdateRoleResp`

2. request definition



```golang
type UpdateRoleReq struct {
	Id int64 `json:"id"` // 编号
	Name string `json:"name"` // 角色名称
	Remark string `json:"remark"` // 备注
	Status int64 `json:"status"` // 状态
}
```


3. response definition



```golang
type UpdateRoleResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sys/role/delete
- Method: POST
- Request: `DeleteRoleReq`
- Response: `DeleteRoleResp`

2. request definition



```golang
type DeleteRoleReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteRoleResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 5. N/A

1. route definition

- Url: /api/sys/role/roleMenuIds
- Method: POST
- Request: `RoleMenuIdsReq`
- Response: `RoleMenuIdsResp`

2. request definition



```golang
type RoleMenuIdsReq struct {
}
```


3. response definition



```golang
type RoleMenuIdsResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 6. N/A

1. route definition

- Url: /api/sys/role/queryMenuByRoleId
- Method: POST
- Request: `RoleMenuReq`
- Response: `RoleMenuResp`

2. request definition



```golang
type RoleMenuReq struct {
	Id int64 `json:"id,optional"`
}
```


3. response definition



```golang
type RoleMenuResp struct {
	AllData []*ListMenuData `json:"allData"`
	RoleData []int64 `json:"userData"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 7. N/A

1. route definition

- Url: /api/sys/role/updateRoleMenu
- Method: POST
- Request: `UpdateRoleMenuReq`
- Response: `UpdateRoleMenuResp`

2. request definition



```golang
type UpdateRoleMenuReq struct {
	RoleId int64 `json:"roleId"`
	MenuIds []int64 `json:"menuIds"`
}
```


3. response definition



```golang
type UpdateRoleMenuResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

