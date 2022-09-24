### 1. N/A

1. route definition

- Url: /api/order/setting/add
- Method: POST
- Request: `addOrderSettingReq`
- Response: `addOrderSettingResp`

2. request definition



```golang
type AddOrderSettingReq struct {
	FlashOrderOvertime int64 `json:"flashOrderOvertime"` // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int64 `json:"normalOrderOvertime"` // 正常订单超时时间(分)
	ConfirmOvertime int64 `json:"confirmOvertime"` // 发货后自动确认收货时间（天）
	FinishOvertime int64 `json:"finishOvertime"` // 自动完成交易时间，不能申请售后（天）
	CommentOvertime int64 `json:"commentOvertime"` // 订单完成后自动好评时间（天）
}
```


3. response definition



```golang
type AddOrderSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/order/setting/list
- Method: POST
- Request: `ListOrderSettingReq`
- Response: `ListOrderSettingResp`

2. request definition



```golang
type ListOrderSettingReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListOrderSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtOrderSettingData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/order/setting/update
- Method: POST
- Request: `UpdateOrderSettingReq`
- Response: `UpdateOrderSettingResp`

2. request definition



```golang
type UpdateOrderSettingReq struct {
	Id int64 `json:"id"`
	FlashOrderOvertime int64 `json:"flashOrderOvertime"` // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int64 `json:"normalOrderOvertime"` // 正常订单超时时间(分)
	ConfirmOvertime int64 `json:"confirmOvertime"` // 发货后自动确认收货时间（天）
	FinishOvertime int64 `json:"finishOvertime"` // 自动完成交易时间，不能申请售后（天）
	CommentOvertime int64 `json:"commentOvertime"` // 订单完成后自动好评时间（天）
}
```


3. response definition



```golang
type UpdateOrderSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/order/setting/delete
- Method: POST
- Request: `DeleteOrderSettingReq`
- Response: `DeleteOrderSettingResp`

2. request definition



```golang
type DeleteOrderSettingReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteOrderSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

