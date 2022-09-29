### 1. N/A

1. route definition

- Url: /api/member/loginlog/add
- Method: POST
- Request: `addMemberLoginLogReq`
- Response: `addMemberLoginLogResp`

2. request definition



```golang
type AddMemberLoginLogReq struct {
	MemberId int64 `json:"memberId"`
	CreateTime string `json:"createTime"`
	Ip string `json:"ip"`
	City string `json:"city"`
	LoginType int64 `json:"loginType"` // 登录类型：0-&gt;PC；1-&gt;android;2-&gt;ios;3-&gt;小程序
	Province string `json:"province"`
}
```


3. response definition



```golang
type AddMemberLoginLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/member/loginlog/list
- Method: POST
- Request: `ListMemberLoginLogReq`
- Response: `ListMemberLoginLogResp`

2. request definition



```golang
type ListMemberLoginLogReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberLoginLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberLoginLogData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/member/loginlog/update
- Method: POST
- Request: `UpdateMemberLoginLogReq`
- Response: `UpdateMemberLoginLogResp`

2. request definition



```golang
type UpdateMemberLoginLogReq struct {
	Id int64 `json:"id"`
	MemberId int64 `json:"memberId"`
	CreateTime string `json:"createTime"`
	Ip string `json:"ip"`
	City string `json:"city"`
	LoginType int64 `json:"loginType"` // 登录类型：0-&gt;PC；1-&gt;android;2-&gt;ios;3-&gt;小程序
	Province string `json:"province"`
}
```


3. response definition



```golang
type UpdateMemberLoginLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/member/loginlog/delete
- Method: POST
- Request: `DeleteMemberLoginLogReq`
- Response: `DeleteMemberLoginLogResp`

2. request definition



```golang
type DeleteMemberLoginLogReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberLoginLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

