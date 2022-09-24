### 1. N/A

1. route definition

- Url: /api/sms/homerecommendsubject/add
- Method: POST
- Request: `addHomeRecommendSubjectReq`
- Response: `addHomeRecommendSubjectResp`

2. request definition



```golang
type AddHomeRecommendSubjectReq struct {
	SubjectId int64 `json:"subjectId"`
	SubjectName string `json:"subjectName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type AddHomeRecommendSubjectResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 2. N/A

1. route definition

- Url: /api/sms/homerecommendsubject/list
- Method: POST
- Request: `ListHomeRecommendSubjectReq`
- Response: `ListHomeRecommendSubjectResp`

2. request definition



```golang
type ListHomeRecommendSubjectReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListHomeRecommendSubjectResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtHomeRecommendSubjectData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 3. N/A

1. route definition

- Url: /api/sms/homerecommendsubject/apisms/homerecommendsubject/update
- Method: POST
- Request: `UpdateHomeRecommendSubjectReq`
- Response: `UpdateHomeRecommendSubjectResp`

2. request definition



```golang
type UpdateHomeRecommendSubjectReq struct {
	Id int64 `json:"id"`
	SubjectId int64 `json:"subjectId"`
	SubjectName string `json:"subjectName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type UpdateHomeRecommendSubjectResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 4. N/A

1. route definition

- Url: /api/sms/homerecommendsubject/delete
- Method: POST
- Request: `DeleteHomeRecommendSubjectReq`
- Response: `DeleteHomeRecommendSubjectResp`

2. request definition



```golang
type DeleteHomeRecommendSubjectReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteHomeRecommendSubjectResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

