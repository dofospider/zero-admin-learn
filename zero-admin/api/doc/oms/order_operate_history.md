### 1. N/A

1. route definition

- Url: /api/order/operatehistory/add
- Method: POST
- Request: `addOperateHistoryReq`
- Response: `addOperateHistoryResp`

2. request definition



```golang
type AddOperateHistoryReq struct {
	OrderId int64 `json:"orderId"` // 订单id
	OperateMan string `json:"operateMan"` // 操作人：用户；系统；后台管理员
	CreateTime string `json:"createTime"` // 操作时间
	OrderStatus int64 `json:"orderStatus"` // 订单状态：0-&gt;待付款；1-&gt;待发货；2-&gt;已发货；3-&gt;已完成；4-&gt;已关闭；5-&gt;无效订单
	Note string `json:"note"` // 备注
}
```


3. response definition



```golang
type AddOperateHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/order/operatehistory/list
- Method: POST
- Request: `ListOperateHistoryReq`
- Response: `ListOperateHistoryResp`

2. request definition



```golang
type ListOperateHistoryReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListOperateHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtOperateHistoryData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/order/operatehistory/update
- Method: POST
- Request: `UpdateOperateHistoryReq`
- Response: `UpdateOperateHistoryResp`

2. request definition



```golang
type UpdateOperateHistoryReq struct {
	Id int64 `json:"id"`
	OrderId int64 `json:"orderId"` // 订单id
	OperateMan string `json:"operateMan"` // 操作人：用户；系统；后台管理员
	CreateTime string `json:"createTime"` // 操作时间
	OrderStatus int64 `json:"orderStatus"` // 订单状态：0-&gt;待付款；1-&gt;待发货；2-&gt;已发货；3-&gt;已完成；4-&gt;已关闭；5-&gt;无效订单
	Note string `json:"note"` // 备注
}
```


3. response definition



```golang
type UpdateOperateHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/order/operatehistory/delete
- Method: POST
- Request: `DeleteOperateHistoryReq`
- Response: `DeleteOperateHistoryResp`

2. request definition



```golang
type DeleteOperateHistoryReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteOperateHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

