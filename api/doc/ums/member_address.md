### 1. N/A

1. route definition

- Url: /api/member/address/add
- Method: POST
- Request: `addMemberAddressReq`
- Response: `addMemberAddressResp`

2. request definition



```golang
type AddMemberAddressReq struct {
	MemberId int64 `json:"memberId"`
	Name string `json:"name"` // 收货人名称
	PhoneNumber string `json:"phoneNumber"`
	DefaultStatus int64 `json:"defaultStatus"` // 是否为默认
	PostCode string `json:"postCode"` // 邮政编码
	Province string `json:"province"` // 省份/直辖市
	City string `json:"city"` // 城市
	Region string `json:"region"` // 区
	DetailAddress string `json:"detailAddress"` // 详细地址(街道)
}
```


3. response definition



```golang
type AddMemberAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/member/address/list
- Method: POST
- Request: `ListMemberAddressReq`
- Response: `ListMemberAddressResp`

2. request definition



```golang
type ListMemberAddressReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberAddressResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberAddressData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 3. N/A

1. route definition

- Url: /api/member/address/update
- Method: POST
- Request: `UpdateMemberAddressReq`
- Response: `UpdateMemberAddressResp`

2. request definition



```golang
type UpdateMemberAddressReq struct {
	Id int64 `json:"id"`
	MemberId int64 `json:"memberId"`
	Name string `json:"name"` // 收货人名称
	PhoneNumber string `json:"phoneNumber"`
	DefaultStatus int64 `json:"defaultStatus"` // 是否为默认
	PostCode string `json:"postCode"` // 邮政编码
	Province string `json:"province"` // 省份/直辖市
	City string `json:"city"` // 城市
	Region string `json:"region"` // 区
	DetailAddress string `json:"detailAddress"` // 详细地址(街道)
}
```


3. response definition



```golang
type UpdateMemberAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/member/address/delete
- Method: POST
- Request: `DeleteMemberAddressReq`
- Response: `DeleteMemberAddressResp`

2. request definition



```golang
type DeleteMemberAddressReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

