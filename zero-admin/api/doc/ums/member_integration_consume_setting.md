### 1. N/A

1. route definition

- Url: /api/member/integrationconsumesetting/add
- Method: POST
- Request: `addIntegrationConsumeSettingReq`
- Response: `addIntegrationConsumeSettingResp`

2. request definition



```golang
type AddIntegrationConsumeSettingReq struct {
	DeductionPerAmount int64 `json:"deductionPerAmount"` // 每一元需要抵扣的积分数量
	MaxPercentPerOrder int64 `json:"maxPercentPerOrder"` // 每笔订单最高抵用百分比
	UseUnit int64 `json:"useUnit"` // 每次使用积分最小单位100
	CouponStatus int64 `json:"couponStatus"` // 是否可以和优惠券同用；0-&gt;不可以；1-&gt;可以
}
```


3. response definition



```golang
type AddIntegrationConsumeSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/member/integrationconsumesetting/list
- Method: POST
- Request: `ListIntegrationConsumeSettingReq`
- Response: `ListIntegrationConsumeSettingResp`

2. request definition



```golang
type ListIntegrationConsumeSettingReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListIntegrationConsumeSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtIntegrationConsumeSettingData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/member/integrationconsumesetting/update
- Method: POST
- Request: `UpdateIntegrationConsumeSettingReq`
- Response: `UpdateIntegrationConsumeSettingResp`

2. request definition



```golang
type UpdateIntegrationConsumeSettingReq struct {
	Id int64 `json:"id"`
	DeductionPerAmount int64 `json:"deductionPerAmount"` // 每一元需要抵扣的积分数量
	MaxPercentPerOrder int64 `json:"maxPercentPerOrder"` // 每笔订单最高抵用百分比
	UseUnit int64 `json:"useUnit"` // 每次使用积分最小单位100
	CouponStatus int64 `json:"couponStatus"` // 是否可以和优惠券同用；0-&gt;不可以；1-&gt;可以
}
```


3. response definition



```golang
type UpdateIntegrationConsumeSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/member/integrationconsumesetting/delete
- Method: POST
- Request: `DeleteIntegrationConsumeSettingReq`
- Response: `DeleteIntegrationConsumeSettingResp`

2. request definition



```golang
type DeleteIntegrationConsumeSettingReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteIntegrationConsumeSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

