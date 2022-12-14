### 1. N/A

1. route definition

- Url: /api/order/returnapply/add
- Method: POST
- Request: `addReturnApplyReq`
- Response: `addReturnApplyResp`

2. request definition



```golang
type AddReturnApplyReq struct {
	OrderId int64 `json:"orderId"` // 订单id
	CompanyAddressId int64 `json:"companyAddressId"` // 收货地址表id
	ProductId int64 `json:"productId"` // 退货商品id
	OrderSn string `json:"orderSn"` // 订单编号
	CreateTime string `json:"createTime"` // 申请时间
	MemberUsername string `json:"memberUserName"` // 会员用户名
	ReturnAmount float64 `json:"returnAmount"` // 退款金额
	ReturnName string `json:"returnName"` // 退货人姓名
	ReturnPhone string `json:"returnPhone"` // 退货人电话
	Status int64 `json:"status"` // 申请状态：0-&gt;待处理；1-&gt;退货中；2-&gt;已完成；3-&gt;已拒绝
	HandleTime string `json:"handleTime"` // 处理时间
	ProductPic string `json:"productPic"` // 商品图片
	ProductName string `json:"productName"` // 商品名称
	ProductBrand string `json:"productBrand"` // 商品品牌
	ProductAttr string `json:"productAttr"` // 商品销售属性：颜色：红色；尺码：xl;
	ProductCount int64 `json:"productCount"` // 退货数量
	ProductPrice float64 `json:"productPrice"` // 商品单价
	ProductRealPrice float64 `json:"productRealPrice"` // 商品实际支付单价
	Reason string `json:"reason"` // 原因
	Description string `json:"description"` // 描述
	ProofPics string `json:"proofPics"` // 凭证图片，以逗号隔开
	HandleNote string `json:"handleNote"` // 处理备注
	HandleMan string `json:"handleMan"` // 处理人员
	ReceiveMan string `json:"receiveMan"` // 收货人
	ReceiveTime string `json:"receiveTime"` // 收货时间
	ReceiveNote string `json:"receiveNote"` // 收货备注
}
```


3. response definition



```golang
type AddReturnApplyResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/order/returnapply/list
- Method: POST
- Request: `ListReturnApplyReq`
- Response: `ListReturnApplyResp`

2. request definition



```golang
type ListReturnApplyReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListReturnApplyResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtReturnApplyData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/order/returnapply/update
- Method: POST
- Request: `UpdateReturnApplyReq`
- Response: `UpdateReturnApplyResp`

2. request definition



```golang
type UpdateReturnApplyReq struct {
	Id int64 `json:"id"`
	OrderId int64 `json:"orderId"` // 订单id
	CompanyAddressId int64 `json:"companyAddressId"` // 收货地址表id
	ProductId int64 `json:"productId"` // 退货商品id
	OrderSn string `json:"orderSn"` // 订单编号
	CreateTime string `json:"createTime"` // 申请时间
	MemberUsername string `json:"memberUserName"` // 会员用户名
	ReturnAmount float64 `json:"returnAmount"` // 退款金额
	ReturnName string `json:"returnName"` // 退货人姓名
	ReturnPhone string `json:"returnPhone"` // 退货人电话
	Status int64 `json:"status"` // 申请状态：0-&gt;待处理；1-&gt;退货中；2-&gt;已完成；3-&gt;已拒绝
	HandleTime string `json:"handleTime"` // 处理时间
	ProductPic string `json:"productPic"` // 商品图片
	ProductName string `json:"productName"` // 商品名称
	ProductBrand string `json:"productBrand"` // 商品品牌
	ProductAttr string `json:"productAttr"` // 商品销售属性：颜色：红色；尺码：xl;
	ProductCount int64 `json:"productCount"` // 退货数量
	ProductPrice float64 `json:"productPrice"` // 商品单价
	ProductRealPrice float64 `json:"productRealPrice"` // 商品实际支付单价
	Reason string `json:"reason"` // 原因
	Description string `json:"description"` // 描述
	ProofPics string `json:"proofPics"` // 凭证图片，以逗号隔开
	HandleNote string `json:"handleNote"` // 处理备注
	HandleMan string `json:"handleMan"` // 处理人员
	ReceiveMan string `json:"receiveMan"` // 收货人
	ReceiveTime string `json:"receiveTime"` // 收货时间
	ReceiveNote string `json:"receiveNote"` // 收货备注
}
```


3. response definition



```golang
type UpdateReturnApplyResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/order/returnapply/delete
- Method: POST
- Request: `DeleteReturnApplyReq`
- Response: `DeleteReturnApplyResp`

2. request definition



```golang
type DeleteReturnApplyReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteReturnApplyResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

