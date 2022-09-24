### 1. N/A

1. route definition

- Url: /api/member/statistics/add
- Method: POST
- Request: `addMemberStatisticsInfoReq`
- Response: `addMemberStatisticsInfoResp`

2. request definition



```golang
type AddMemberStatisticsInfoReq struct {
	MemberId int64 `json:"memberId"`
	ConsumeAmount float64 `json:"consumeAmount"` // 累计消费金额
	OrderCount int64 `json:"orderCount"` // 订单数量
	CouponCount int64 `json:"couponCount"` // 优惠券数量
	CommentCount int64 `json:"commentCount"` // 评价数
	ReturnOrderCount int64 `json:"returnOrderCount"` // 退货数量
	LoginCount int64 `json:"loginCount"` // 登录次数
	AttendCount int64 `json:"attendCount"` // 关注数量
	FansCount int64 `json:"fansCount"` // 粉丝数量
	CollectProductCount int64 `json:"collectProductCount"`
	CollectSubjectCount int64 `json:"collectSubjectCount"`
	CollectTopicCount int64 `json:"collectTopicCount"`
	CollectCommentCount int64 `json:"collectCommentCount"`
	InviteFriendCount int64 `json:"inviteFriendCount"`
	RecentOrderTime string `json:"recentOrderTime"` // 最后一次下订单时间
}
```


3. response definition



```golang
type AddMemberStatisticsInfoResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/member/statistics/list
- Method: POST
- Request: `ListMemberStatisticsInfoReq`
- Response: `ListMemberStatisticsInfoResp`

2. request definition



```golang
type ListMemberStatisticsInfoReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberStatisticsInfoResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberStatisticsInfoData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/member/statistics/update
- Method: POST
- Request: `UpdateMemberStatisticsInfoReq`
- Response: `UpdateMemberStatisticsInfoResp`

2. request definition



```golang
type UpdateMemberStatisticsInfoReq struct {
	Id int64 `json:"id"`
	MemberId int64 `json:"memberId"`
	ConsumeAmount float64 `json:"consumeAmount"` // 累计消费金额
	OrderCount int64 `json:"orderCount"` // 订单数量
	CouponCount int64 `json:"couponCount"` // 优惠券数量
	CommentCount int64 `json:"commentCount"` // 评价数
	ReturnOrderCount int64 `json:"returnOrderCount"` // 退货数量
	LoginCount int64 `json:"loginCount"` // 登录次数
	AttendCount int64 `json:"attendCount"` // 关注数量
	FansCount int64 `json:"fansCount"` // 粉丝数量
	CollectProductCount int64 `json:"collectProductCount"`
	CollectSubjectCount int64 `json:"collectSubjectCount"`
	CollectTopicCount int64 `json:"collectTopicCount"`
	CollectCommentCount int64 `json:"collectCommentCount"`
	InviteFriendCount int64 `json:"inviteFriendCount"`
	RecentOrderTime string `json:"recentOrderTime"` // 最后一次下订单时间
}
```


3. response definition



```golang
type UpdateMemberStatisticsInfoResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/member/statistics/delete
- Method: POST
- Request: `DeleteMemberStatisticsInfoReq`
- Response: `DeleteMemberStatisticsInfoResp`

2. request definition



```golang
type DeleteMemberStatisticsInfoReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberStatisticsInfoResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

