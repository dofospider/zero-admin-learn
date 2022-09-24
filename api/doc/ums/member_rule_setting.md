### 1. N/A

1. route definition

- Url: /api/member/rulesetting/add
- Method: POST
- Request: `addMemberRuleSettingReq`
- Response: `addMemberRuleSettingResp`

2. request definition



```golang
type AddMemberRuleSettingReq struct {
	ContinueSignDay int64 `json:"continueSignDay"` // 连续签到天数
	ContinueSignPoint int64 `json:"continueSignPoint"` // 连续签到赠送数量
	ConsumePerPoint float64 `json:"consumePerPoint"` // 每消费多少元获取1个点
	LowOrderAmount float64 `json:"lowOrderAmount"` // 最低获取点数的订单金额
	MaxPointPerOrder int64 `json:"maxPointPerOrder"` // 每笔订单最高获取点数
	Type int64 `json:"type"` // 类型：0-&gt;积分规则；1-&gt;成长值规则
}
```


3. response definition



```golang
type AddMemberRuleSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/member/rulesetting/list
- Method: POST
- Request: `ListMemberRuleSettingReq`
- Response: `ListMemberRuleSettingResp`

2. request definition



```golang
type ListMemberRuleSettingReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberRuleSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberRuleSettingData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/member/rulesetting/update
- Method: POST
- Request: `UpdateMemberRuleSettingReq`
- Response: `UpdateMemberRuleSettingResp`

2. request definition



```golang
type UpdateMemberRuleSettingReq struct {
	Id int64 `json:"id"`
	ContinueSignDay int64 `json:"continueSignDay"` // 连续签到天数
	ContinueSignPoint int64 `json:"continueSignPoint"` // 连续签到赠送数量
	ConsumePerPoint float64 `json:"consumePerPoint"` // 每消费多少元获取1个点
	LowOrderAmount float64 `json:"lowOrderAmount"` // 最低获取点数的订单金额
	MaxPointPerOrder int64 `json:"maxPointPerOrder"` // 每笔订单最高获取点数
	Type int64 `json:"type"` // 类型：0-&gt;积分规则；1-&gt;成长值规则
}
```


3. response definition



```golang
type UpdateMemberRuleSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/member/rulesetting/delete
- Method: POST
- Request: `DeleteMemberRuleSettingReq`
- Response: `DeleteMemberRuleSettingResp`

2. request definition



```golang
type DeleteMemberRuleSettingReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberRuleSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

