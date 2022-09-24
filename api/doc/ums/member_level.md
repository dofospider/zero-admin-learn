### 1. N/A

1. route definition

- Url: /api/member/level/add
- Method: POST
- Request: `addMemberLevelReq`
- Response: `addMemberLevelResp`

2. request definition



```golang
type AddMemberLevelReq struct {
	Name string `json:"name"`
	GrowthPoint int64 `json:"growthPoint"`
	DefaultStatus int64 `json:"defaultStatus"` // 是否为默认等级：0-&gt;不是；1-&gt;是
	FreeFreightPoint float64 `json:"freeFreightPoint"` // 免运费标准
	CommentGrowthPoint int64 `json:"commentGrowthPoint"` // 每次评价获取的成长值
	PriviledgeFreeFreight int64 `json:"priviledgeFreeFreight"` // 是否有免邮特权
	PriviledgeSignIn int64 `json:"priviledgeSignIn"` // 是否有签到特权
	PriviledgeComment int64 `json:"priviledgeComment"` // 是否有评论获奖励特权
	PriviledgePromotion int64 `json:"priviledgePromotion"` // 是否有专享活动特权
	PriviledgeMemberPrice int64 `json:"priviledgeMemberPrice"` // 是否有会员价格特权
	PriviledgeBirthday int64 `json:"priviledgeBirthday"` // 是否有生日特权
	Note string `json:"note"`
}
```


3. response definition



```golang
type AddMemberLevelResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/member/level/list
- Method: POST
- Request: `ListMemberLevelReq`
- Response: `ListMemberLevelResp`

2. request definition



```golang
type ListMemberLevelReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberLevelResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberLevelData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/member/level/update
- Method: POST
- Request: `UpdateMemberLevelReq`
- Response: `UpdateMemberLevelResp`

2. request definition



```golang
type UpdateMemberLevelReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	GrowthPoint int64 `json:"growthPoint"`
	DefaultStatus int64 `json:"defaultStatus"` // 是否为默认等级：0-&gt;不是；1-&gt;是
	FreeFreightPoint float64 `json:"freeFreightPoint"` // 免运费标准
	CommentGrowthPoint int64 `json:"commentGrowthPoint"` // 每次评价获取的成长值
	PriviledgeFreeFreight int64 `json:"priviledgeFreeFreight"` // 是否有免邮特权
	PriviledgeSignIn int64 `json:"priviledgeSignIn"` // 是否有签到特权
	PriviledgeComment int64 `json:"priviledgeComment"` // 是否有评论获奖励特权
	PriviledgePromotion int64 `json:"priviledgePromotion"` // 是否有专享活动特权
	PriviledgeMemberPrice int64 `json:"priviledgeMemberPrice"` // 是否有会员价格特权
	PriviledgeBirthday int64 `json:"priviledgeBirthday"` // 是否有生日特权
	Note string `json:"note"`
}
```


3. response definition



```golang
type UpdateMemberLevelResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/member/level/delete
- Method: POST
- Request: `DeleteMemberLevelReq`
- Response: `DeleteMemberLevelResp`

2. request definition



```golang
type DeleteMemberLevelReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberLevelResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

