### 1. N/A

1. route definition

- Url: /api/product/memberprice/add
- Method: POST
- Request: `addMemberPriceReq`
- Response: `addMemberPriceResp`

2. request definition



```golang
type AddMemberPriceReq struct {
	ProductId int64 `json:"productId"`
	MemberLevelId int64 `json:"memberLevelId"`
	MemberPrice float64 `json:"memberPrice"` // 会员价格
	MemberLevelName string `json:"memberLevelName"`
}
```


3. response definition



```golang
type AddMemberPriceResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/product/memberprice/list
- Method: POST
- Request: `ListMemberPriceReq`
- Response: `ListMemberPriceResp`

2. request definition



```golang
type ListMemberPriceReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberPriceResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberPriceData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 3. N/A

1. route definition

- Url: /api/product/memberprice/update
- Method: POST
- Request: `UpdateMemberPriceReq`
- Response: `UpdateMemberPriceResp`

2. request definition



```golang
type UpdateMemberPriceReq struct {
	Id int64 `json:"id"`
	ProductId int64 `json:"productId"`
	MemberLevelId int64 `json:"memberLevelId"`
	MemberPrice float64 `json:"memberPrice"` // 会员价格
	MemberLevelName string `json:"memberLevelName"`
}
```


3. response definition



```golang
type UpdateMemberPriceResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/product/memberprice/delete
- Method: POST
- Request: `DeleteMemberPriceReq`
- Response: `DeleteMemberPriceResp`

2. request definition



```golang
type DeleteMemberPriceReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberPriceResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

