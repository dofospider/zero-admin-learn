### 1. N/A

1. route definition

- Url: /api/product/comment/add
- Method: POST
- Request: `addProductCommentReq`
- Response: `addProductCommentResp`

2. request definition



```golang
type AddProductCommentReq struct {
	ProductId int64 `json:"productId"`
	MemberNickName string `json:"memberNickName"`
	ProductName string `json:"productName"`
	Star int64 `json:"star"` // 评价星数：0-&gt;5
	MemberIp string `json:"memberIp"` // 评价的ip
	CreateTime string `json:"createTime"`
	ShowStatus int64 `json:"showStatus"`
	ProductAttribute string `json:"productAttribute"` // 购买时的商品属性
	CollectCouont int64 `json:"collectCouont"`
	ReadCount int64 `json:"readCount"`
	Content string `json:"content"`
	Pics string `json:"pics"` // 上传图片地址，以逗号隔开
	MemberIcon string `json:"memberIcon"` // 评论用户头像
	ReplayCount int64 `json:"replayCount"`
}
```


3. response definition



```golang
type AddProductCommentResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/product/comment/list
- Method: POST
- Request: `ListProductCommentReq`
- Response: `ListProductCommentResp`

2. request definition



```golang
type ListProductCommentReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListProductCommentResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtProductCommentData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/product/comment/update
- Method: POST
- Request: `UpdateProductCommentReq`
- Response: `UpdateProductCommentResp`

2. request definition



```golang
type UpdateProductCommentReq struct {
	Id int64 `json:"id"`
	ProductId int64 `json:"productId"`
	MemberNickName string `json:"memberNickName"`
	ProductName string `json:"productName"`
	Star int64 `json:"star"` // 评价星数：0-&gt;5
	MemberIp string `json:"memberIp"` // 评价的ip
	CreateTime string `json:"createTime"`
	ShowStatus int64 `json:"showStatus"`
	ProductAttribute string `json:"productAttribute"` // 购买时的商品属性
	CollectCouont int64 `json:"collectCouont"`
	ReadCount int64 `json:"readCount"`
	Content string `json:"content"`
	Pics string `json:"pics"` // 上传图片地址，以逗号隔开
	MemberIcon string `json:"memberIcon"` // 评论用户头像
	ReplayCount int64 `json:"replayCount"`
}
```


3. response definition



```golang
type UpdateProductCommentResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/product/comment/delete
- Method: POST
- Request: `DeleteProductCommentReq`
- Response: `DeleteProductCommentResp`

2. request definition



```golang
type DeleteProductCommentReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteProductCommentResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

