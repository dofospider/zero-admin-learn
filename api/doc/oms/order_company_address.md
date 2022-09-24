### 1. N/A

1. route definition

- Url: /api/order/compayaddress/add
- Method: POST
- Request: `addCompayAddressReq`
- Response: `addCompayAddressResp`

2. request definition



```golang
type AddCompayAddressReq struct {
	AddressName string `json:"addressName"` // 地址名称
	SendStatus int64 `json:"sendStatus"` // 默认发货地址：0-&gt;否；1-&gt;是
	ReceiveStatus int64 `json:"receiveStatus"` // 是否默认收货地址：0-&gt;否；1-&gt;是
	Name string `json:"name"` // 收发货人姓名
	Phone string `json:"phone"` // 收货人电话
	Province string `json:"province"` // 省/直辖市
	City string `json:"city"` // 市
	Region string `json:"region"` // 区
	DetailAddress string `json:"detailAddress"` // 详细地址
}
```


3. response definition



```golang
type AddCompayAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/order/compayaddress/list
- Method: POST
- Request: `ListCompayAddressReq`
- Response: `ListCompayAddressResp`

2. request definition



```golang
type ListCompayAddressReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListCompayAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtCompayAddressData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/order/compayaddress/update
- Method: POST
- Request: `UpdateCompayAddressReq`
- Response: `UpdateCompayAddressResp`

2. request definition



```golang
type UpdateCompayAddressReq struct {
	Id int64 `json:"id"`
	AddressName string `json:"addressName"` // 地址名称
	SendStatus int64 `json:"sendStatus"` // 默认发货地址：0-&gt;否；1-&gt;是
	ReceiveStatus int64 `json:"receiveStatus"` // 是否默认收货地址：0-&gt;否；1-&gt;是
	Name string `json:"name"` // 收发货人姓名
	Phone string `json:"phone"` // 收货人电话
	Province string `json:"province"` // 省/直辖市
	City string `json:"city"` // 市
	Region string `json:"region"` // 区
	DetailAddress string `json:"detailAddress"` // 详细地址
}
```


3. response definition



```golang
type UpdateCompayAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/order/compayaddress/delete
- Method: POST
- Request: `DeleteCompayAddressReq`
- Response: `DeleteCompayAddressResp`

2. request definition



```golang
type DeleteCompayAddressReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteCompayAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

