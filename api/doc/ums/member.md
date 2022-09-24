### 1. N/A

1. route definition

- Url: /api/member/member/list
- Method: POST
- Request: `ListMemberReq`
- Response: `ListMemberResp`

2. request definition



```golang
type ListMemberReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 2. N/A

1. route definition

- Url: /api/member/member/update
- Method: POST
- Request: `UpdateMemberReq`
- Response: `UpdateMemberResp`

2. request definition



```golang
type UpdateMemberReq struct {
	Id int64 `json:"id"`
	MemberLevelId int64 `json:"memberLevelId"`
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Nickname string `json:"nickname"` // 昵称
	Phone string `json:"phone"` // 手机号码
	Status int64 `json:"status"` // 帐号启用状态:0-&gt;禁用；1-&gt;启用
	CreateTime string `json:"createTime"` // 注册时间
	Icon string `json:"icon"` // 头像
	Gender int64 `json:"gender"` // 性别：0-&gt;未知；1-&gt;男；2-&gt;女
	Birthday string `json:"birthday"` // 生日
	City string `json:"city"` // 所做城市
	Job string `json:"job"` // 职业
	PersonalizedSignature string `json:"personalizedSignature"` // 个性签名
	SourceType int64 `json:"sourceType"` // 用户来源
	Integration int64 `json:"integration"` // 积分
	Growth int64 `json:"growth"` // 成长值
	LuckeyCount int64 `json:"luckeyCount"` // 剩余抽奖次数
	HistoryIntegration int64 `json:"historyIntegration"` // 历史积分数量
}
```


3. response definition



```golang
type UpdateMemberResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 3. N/A

1. route definition

- Url: /api/member/member/delete
- Method: POST
- Request: `DeleteMemberReq`
- Response: `DeleteMemberResp`

2. request definition



```golang
type DeleteMemberReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

