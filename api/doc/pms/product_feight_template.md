### 1. N/A

1. route definition

- Url: /api/product/feighttemplate/add
- Method: POST
- Request: `addFeightTemplateReq`
- Response: `addFeightTemplateResp`

2. request definition



```golang
type AddFeightTemplateReq struct {
	Name string `json:"name"`
	ChargeType int64 `json:"chargeType"` // 计费类型:0-&gt;按重量；1-&gt;按件数
	FirstWeight float64 `json:"first_weight"` // 首重kg
	FirstFee float64 `json:"firstFee"` // 首费（元）
	ContinueWeight float64 `json:"continue_weight"`
	ContinmeFee float64 `json:"continmeFee"`
	Dest string `json:"dest"` // 目的地（省、市）
}
```


3. response definition



```golang
type AddFeightTemplateResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/product/feighttemplate/list
- Method: POST
- Request: `ListFeightTemplateReq`
- Response: `ListFeightTemplateResp`

2. request definition



```golang
type ListFeightTemplateReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListFeightTemplateResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtFeightTemplateData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/product/feighttemplate/update
- Method: POST
- Request: `UpdateFeightTemplateReq`
- Response: `UpdateFeightTemplateResp`

2. request definition



```golang
type UpdateFeightTemplateReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	ChargeType int64 `json:"chargeType"` // 计费类型:0-&gt;按重量；1-&gt;按件数
	FirstWeight float64 `json:"first_weight"` // 首重kg
	FirstFee float64 `json:"firstFee"` // 首费（元）
	ContinueWeight float64 `json:"continue_weight"`
	ContinmeFee float64 `json:"continmeFee"`
	Dest string `json:"dest"` // 目的地（省、市）
}
```


3. response definition



```golang
type UpdateFeightTemplateResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/product/feighttemplate/delete
- Method: POST
- Request: `DeleteFeightTemplateReq`
- Response: `DeleteFeightTemplateResp`

2. request definition



```golang
type DeleteFeightTemplateReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteFeightTemplateResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

