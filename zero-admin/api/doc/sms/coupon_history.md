### 1. N/A

1. route definition

- Url: /api/sms/couponhistory/add
- Method: POST
- Request: `addCouponHistoryReq`
- Response: `addCouponHistoryResp`

2. request definition



```golang
type AddCouponHistoryReq struct {
	CouponId int64 `json:"couponId"`
	MemberId int64 `json:"memberId"`
	CouponCode string `json:"couponCode"`
	MemberNickname string `json:"memberNickName"` // 领取人昵称
	GetType int64 `json:"getType"` // 获取类型：0-&gt;后台赠送；1-&gt;主动获取
	CreateTime string `json:"createTime"`
	UseStatus int64 `json:"useStatus"` // 使用状态：0-&gt;未使用；1-&gt;已使用；2-&gt;已过期
	UseTime string `json:"useTime"` // 使用时间
	OrderId int64 `json:"orderId"` // 订单编号
	OrderSn string `json:"orderSn"` // 订单号码
}
```


3. response definition



```golang
type AddCouponHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sms/couponhistory/list
- Method: POST
- Request: `ListCouponHistoryReq`
- Response: `ListCouponHistoryResp`

2. request definition



```golang
type ListCouponHistoryReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListCouponHistoryResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtCouponHistoryData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 3. N/A

1. route definition

- Url: /api/sms/couponhistory/update
- Method: POST
- Request: `UpdateCouponHistoryReq`
- Response: `UpdateCouponHistoryResp`

2. request definition



```golang
type UpdateCouponHistoryReq struct {
	Id int64 `json:"id"`
	CouponId int64 `json:"couponId"`
	MemberId int64 `json:"memberId"`
	CouponCode string `json:"couponCode"`
	MemberNickname string `json:"memberNickName"` // 领取人昵称
	GetType int64 `json:"getType"` // 获取类型：0-&gt;后台赠送；1-&gt;主动获取
	CreateTime string `json:"createTime"`
	UseStatus int64 `json:"useStatus"` // 使用状态：0-&gt;未使用；1-&gt;已使用；2-&gt;已过期
	UseTime string `json:"useTime"` // 使用时间
	OrderId int64 `json:"orderId"` // 订单编号
	OrderSn string `json:"orderSn"` // 订单号码
}
```


3. response definition



```golang
type UpdateCouponHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sms/couponhistory/delete
- Method: POST
- Request: `DeleteCouponHistoryReq`
- Response: `DeleteCouponHistoryResp`

2. request definition



```golang
type DeleteCouponHistoryReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteCouponHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

