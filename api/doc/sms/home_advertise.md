### 1. N/A

1. route definition

- Url: /api/sms/homeadvertise/add
- Method: POST
- Request: `addHomeAdvertiseReq`
- Response: `addHomeAdvertiseResp`

2. request definition



```golang
type AddHomeAdvertiseReq struct {
	Name string `json:"name"`
	Type int64 `json:"type"` // 轮播位置：0-&gt;PC首页轮播；1-&gt;app首页轮播
	Pic string `json:"pic"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	Status int64 `json:"status"` // 上下线状态：0-&gt;下线；1-&gt;上线
	ClickCount int64 `json:"clickCount"` // 点击数
	OrderCount int64 `json:"orderCount"` // 下单数
	Url string `json:"url"` // 链接地址
	Note string `json:"note"` // 备注
	Sort int64 `json:"sort"` // 排序
}
```


3. response definition



```golang
type AddHomeAdvertiseResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sms/homeadvertise/list
- Method: POST
- Request: `ListHomeAdvertiseReq`
- Response: `ListHomeAdvertiseResp`

2. request definition



```golang
type ListHomeAdvertiseReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListHomeAdvertiseResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtHomeAdvertiseData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sms/homeadvertise/update
- Method: POST
- Request: `UpdateHomeAdvertiseReq`
- Response: `UpdateHomeAdvertiseResp`

2. request definition



```golang
type UpdateHomeAdvertiseReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Type int64 `json:"type"` // 轮播位置：0-&gt;PC首页轮播；1-&gt;app首页轮播
	Pic string `json:"pic"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	Status int64 `json:"status"` // 上下线状态：0-&gt;下线；1-&gt;上线
	ClickCount int64 `json:"clickCount"` // 点击数
	OrderCount int64 `json:"orderCount"` // 下单数
	Url string `json:"url"` // 链接地址
	Note string `json:"note"` // 备注
	Sort int64 `json:"sort"` // 排序
}
```


3. response definition



```golang
type UpdateHomeAdvertiseResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sms/homeadvertise/delete
- Method: POST
- Request: `DeleteHomeAdvertiseReq`
- Response: `DeleteHomeAdvertiseResp`

2. request definition



```golang
type DeleteHomeAdvertiseReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteHomeAdvertiseResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

