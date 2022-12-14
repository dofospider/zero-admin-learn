### 1. N/A

1. route definition

- Url: /api/order/order/add
- Method: POST
- Request: `addOrderReq`
- Response: `addOrderResp`

2. request definition



```golang
type AddOrderReq struct {
	MemberId int64 `json:"memberId"`
	CouponId int64 `json:"couponId"`
	OrderSn string `json:"orderSn"` // 订单编号
	CreateTime string `json:"createTime"` // 提交时间
	MemberUsername string `json:"memberUserName"` // 用户帐号
	TotalAmount float64 `json:"totalAmount"` // 订单总金额
	PayAmount float64 `json:"payAmount"` // 应付金额（实际支付金额）
	FreightAmount float64 `json:"freightAmount"` // 运费金额
	PromotionAmount float64 `json:"promotionAmount"` // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount float64 `json:"integrationAmount"` // 积分抵扣金额
	CouponAmount float64 `json:"couponAmount"` // 优惠券抵扣金额
	DiscountAmount float64 `json:"discountAmount"` // 管理员后台调整订单使用的折扣金额
	PayType int64 `json:"payType"` // 支付方式：0-&gt;未支付；1-&gt;支付宝；2-&gt;微信
	SourceType int64 `json:"sourceType"` // 订单来源：0-&gt;PC订单；1-&gt;app订单
	Status int64 `json:"status"` // 订单状态：0-&gt;待付款；1-&gt;待发货；2-&gt;已发货；3-&gt;已完成；4-&gt;已关闭；5-&gt;无效订单
	OrderType int64 `json:"orderType"` // 订单类型：0-&gt;正常订单；1-&gt;秒杀订单
	DeliveryCompany string `json:"deliveryCompany"` // 物流公司(配送方式)
	DeliverySn string `json:"deliverySn"` // 物流单号
	AutoConfirmDay int64 `json:"autoConfirmDay"` // 自动确认时间（天）
	Integration int64 `json:"integration"` // 可以获得的积分
	Growth int64 `json:"growth"` // 可以活动的成长值
	PromotionInfo string `json:"promotionInfo"` // 活动信息
	BillType int64 `json:"billType"` // 发票类型：0-&gt;不开发票；1-&gt;电子发票；2-&gt;纸质发票
	BillHeader string `json:"billHeader"` // 发票抬头
	BillContent string `json:"billContent"` // 发票内容
	BillReceiverPhone string `json:"billReceiverPhone"` // 收票人电话
	BillReceiverEmail string `json:"billReceiverEmail"` // 收票人邮箱
	ReceiverName string `json:"receiverName"` // 收货人姓名
	ReceiverPhone string `json:"receiverPhone"` // 收货人电话
	ReceiverPostCode string `json:"receiverPostCode"` // 收货人邮编
	ReceiverProvince string `json:"receiverProvince"` // 省份/直辖市
	ReceiverCity string `json:"receiverCity"` // 城市
	ReceiverRegion string `json:"receiverRegion"` // 区
	ReceiverDetailAddress string `json:"receiverDetailAddress"` // 详细地址
	Note string `json:"note"` // 订单备注
	ConfirmStatus int64 `json:"confirmStatus"` // 确认收货状态：0-&gt;未确认；1-&gt;已确认
	DeleteStatus int64 `json:"deleteStatus"` // 删除状态：0-&gt;未删除；1-&gt;已删除
	UseIntegration int64 `json:"useIntegration"` // 下单时使用的积分
}
```


3. response definition



```golang
type AddOrderResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/order/order/list
- Method: POST
- Request: `ListOrderReq`
- Response: `ListOrderResp`

2. request definition



```golang
type ListOrderReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListOrderResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtOrderData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/order/order/update
- Method: POST
- Request: `UpdateOrderReq`
- Response: `UpdateOrderResp`

2. request definition



```golang
type UpdateOrderReq struct {
	Id int64 `json:"id"` // 订单id
	MemberId int64 `json:"memberId"`
	CouponId int64 `json:"couponId"`
	OrderSn string `json:"orderSn"` // 订单编号
	CreateTime string `json:"createTime"` // 提交时间
	MemberUsername string `json:"memberUserName"` // 用户帐号
	TotalAmount float64 `json:"totalAmount"` // 订单总金额
	PayAmount float64 `json:"payAmount"` // 应付金额（实际支付金额）
	FreightAmount float64 `json:"freightAmount"` // 运费金额
	PromotionAmount float64 `json:"promotionAmount"` // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount float64 `json:"integrationAmount"` // 积分抵扣金额
	CouponAmount float64 `json:"couponAmount"` // 优惠券抵扣金额
	DiscountAmount float64 `json:"discountAmount"` // 管理员后台调整订单使用的折扣金额
	PayType int64 `json:"payType"` // 支付方式：0-&gt;未支付；1-&gt;支付宝；2-&gt;微信
	SourceType int64 `json:"sourceType"` // 订单来源：0-&gt;PC订单；1-&gt;app订单
	Status int64 `json:"status"` // 订单状态：0-&gt;待付款；1-&gt;待发货；2-&gt;已发货；3-&gt;已完成；4-&gt;已关闭；5-&gt;无效订单
	OrderType int64 `json:"orderType"` // 订单类型：0-&gt;正常订单；1-&gt;秒杀订单
	DeliveryCompany string `json:"deliveryCompany"` // 物流公司(配送方式)
	DeliverySn string `json:"deliverySn"` // 物流单号
	AutoConfirmDay int64 `json:"autoConfirmDay"` // 自动确认时间（天）
	Integration int64 `json:"integration"` // 可以获得的积分
	Growth int64 `json:"growth"` // 可以活动的成长值
	PromotionInfo string `json:"promotionInfo"` // 活动信息
	BillType int64 `json:"billType"` // 发票类型：0-&gt;不开发票；1-&gt;电子发票；2-&gt;纸质发票
	BillHeader string `json:"billHeader"` // 发票抬头
	BillContent string `json:"billContent"` // 发票内容
	BillReceiverPhone string `json:"billReceiverPhone"` // 收票人电话
	BillReceiverEmail string `json:"billReceiverEmail"` // 收票人邮箱
	ReceiverName string `json:"receiverName"` // 收货人姓名
	ReceiverPhone string `json:"receiverPhone"` // 收货人电话
	ReceiverPostCode string `json:"receiverPostCode"` // 收货人邮编
	ReceiverProvince string `json:"receiverProvince"` // 省份/直辖市
	ReceiverCity string `json:"receiverCity"` // 城市
	ReceiverRegion string `json:"receiverRegion"` // 区
	ReceiverDetailAddress string `json:"receiverDetailAddress"` // 详细地址
	Note string `json:"note"` // 订单备注
	ConfirmStatus int64 `json:"confirmStatus"` // 确认收货状态：0-&gt;未确认；1-&gt;已确认
	DeleteStatus int64 `json:"deleteStatus"` // 删除状态：0-&gt;未删除；1-&gt;已删除
	UseIntegration int64 `json:"useIntegration"` // 下单时使用的积分
}
```


3. response definition



```golang
type UpdateOrderResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/order/order/delete
- Method: POST
- Request: `DeleteOrderReq`
- Response: `DeleteOrderResp`

2. request definition



```golang
type DeleteOrderReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteOrderResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

