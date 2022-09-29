### 1. N/A

1. route definition

- Url: /api/sms/coupon/add
- Method: POST
- Request: `addCouponReq`
- Response: `addCouponResp`

2. request definition



```golang
type AddCouponReq struct {
	Type int64 `json:"type"` // 优惠券类型；0-&gt;全场赠券；1-&gt;会员赠券；2-&gt;购物赠券；3-&gt;注册赠券
	Name string `json:"name"`
	Platform int64 `json:"platform"` // 使用平台：0-&gt;全部；1-&gt;移动；2-&gt;PC
	Count int64 `json:"count"` // 数量
	Amount float64 `json:"amount"` // 金额
	PerLimit int64 `json:"perLimit"` // 每人限领张数
	MinPoint float64 `json:"minPoint"` // 使用门槛；0表示无门槛
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	UseType int64 `json:"useType"` // 使用类型：0-&gt;全场通用；1-&gt;指定分类；2-&gt;指定商品
	Note string `json:"note"` // 备注
	PublishCount int64 `json:"publishCount"` // 发行数量
	UseCount int64 `json:"useCount"` // 已使用数量
	ReceiveCount int64 `json:"receiveCount"` // 领取数量
	EnableTime string `json:"enableTime"` // 可以领取的日期
	Code string `json:"code"` // 优惠码
	MemberLevel int64 `json:"memberLevel"` // 可领取的会员类型：0-&gt;无限时
}
```


3. response definition



```golang
type AddCouponResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sms/coupon/list
- Method: POST
- Request: `ListCouponReq`
- Response: `ListCouponResp`

2. request definition



```golang
type ListCouponReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListCouponResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtCouponData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sms/coupon/update
- Method: POST
- Request: `UpdateCouponReq`
- Response: `UpdateCouponResp`

2. request definition



```golang
type UpdateCouponReq struct {
	Id int64 `json:"id"`
	Type int64 `json:"type"` // 优惠券类型；0-&gt;全场赠券；1-&gt;会员赠券；2-&gt;购物赠券；3-&gt;注册赠券
	Name string `json:"name"`
	Platform int64 `json:"platform"` // 使用平台：0-&gt;全部；1-&gt;移动；2-&gt;PC
	Count int64 `json:"count"` // 数量
	Amount float64 `json:"amount"` // 金额
	PerLimit int64 `json:"perLimit"` // 每人限领张数
	MinPoint float64 `json:"minPoint"` // 使用门槛；0表示无门槛
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	UseType int64 `json:"useType"` // 使用类型：0-&gt;全场通用；1-&gt;指定分类；2-&gt;指定商品
	Note string `json:"note"` // 备注
	PublishCount int64 `json:"publishCount"` // 发行数量
	UseCount int64 `json:"useCount"` // 已使用数量
	ReceiveCount int64 `json:"receiveCount"` // 领取数量
	EnableTime string `json:"enableTime"` // 可以领取的日期
	Code string `json:"code"` // 优惠码
	MemberLevel int64 `json:"memberLevel"` // 可领取的会员类型：0-&gt;无限时
}
```


3. response definition



```golang
type UpdateCouponResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sms/coupon/delete
- Method: POST
- Request: `DeleteCouponReq`
- Response: `DeleteCouponResp`

2. request definition



```golang
type DeleteCouponReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteCouponResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

