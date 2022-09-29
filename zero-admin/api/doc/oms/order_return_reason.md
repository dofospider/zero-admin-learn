### 1. N/A

1. route definition

- Url: /api/order/returnreason/add
- Method: POST
- Request: `addReturnResonReq`
- Response: `addReturnResonResp`

2. request definition



```golang
type AddReturnResonReq struct {
	Name string `json:"name"` // 退货类型
	Sort int64 `json:"sort"`
	Status int64 `json:"status"` // 状态：0-&gt;不启用；1-&gt;启用
}
```


3. response definition



```golang
type AddReturnResonResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/order/returnreason/list
- Method: POST
- Request: `ListReturnResonReq`
- Response: `ListReturnResonResp`

2. request definition



```golang
type ListReturnResonReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListReturnResonResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtReturnResonData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/order/returnreason/update
- Method: POST
- Request: `UpdateReturnResonReq`
- Response: `UpdateReturnResonResp`

2. request definition



```golang
type UpdateReturnResonReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"` // 退货类型
	Sort int64 `json:"sort"`
	Status int64 `json:"status"` // 状态：0-&gt;不启用；1-&gt;启用
}
```


3. response definition



```golang
type UpdateReturnResonResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/order/returnreason/delete
- Method: POST
- Request: `DeleteReturnResonReq`
- Response: `DeleteReturnResonResp`

2. request definition



```golang
type DeleteReturnResonReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteReturnResonResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

