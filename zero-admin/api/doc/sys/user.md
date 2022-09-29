### 1. N/A

1. route definition

- Url: /api/sys/user/currentUser
- Method: GET
- Request: `-`
- Response: `userInfoResp`

2. request definition



3. response definition



```golang
type UserInfoResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Avatar string `json:"avatar"`
	Name string `json:"name"`
	MenuTree []*ListMenuTree `json:"menuTree"`
	MenuTreeVue []*ListMenuTreeVue `json:"menuTreeVue"`
}
```

### 2. N/A

1. route definition

- Url: /api/sys/user/add
- Method: POST
- Request: `addUserReq`
- Response: `addUserResp`

2. request definition



```golang
type AddUserReq struct {
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	Name string `json:"name"`
	NickName string `json:"nickName"`
	DeptId int64 `json:"deptId"`
	RoleId int64 `json:"roleId"`
	JobId int64 `json:"jobId"`
}
```


3. response definition



```golang
type AddUserResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 3. N/A

1. route definition

- Url: /api/sys/user/list
- Method: POST
- Request: `ListUserReq`
- Response: `ListUserResp`

2. request definition



```golang
type ListUserReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
	Name string `json:"name,optional"`
	NickName string `json:"nickName,optional"`
	Mobile string `json:"mobile,optional"`
	Email string `json:"email,optional"`
	Status int64 `json:"status,optional"`
	DeptId int64 `json:"deptId,optional"`
	JobId int64 `json:"deptId,optional"`
}
```


3. response definition



```golang
type ListUserResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListUserData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 4. N/A

1. route definition

- Url: /api/sys/user/update
- Method: POST
- Request: `UpdateUserReq`
- Response: `UpdateUserResp`

2. request definition



```golang
type UpdateUserReq struct {
	Id int64 `json:"id"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	Name string `json:"name"`
	NickName string `json:"nickName"`
	DeptId int64 `json:"deptId"`
	RoleId int64 `json:"roleId"`
	Status int64 `json:"status"`
	JobId int64 `json:"jobId"`
}
```


3. response definition



```golang
type UpdateUserResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 5. N/A

1. route definition

- Url: /api/sys/user/delete
- Method: POST
- Request: `DeleteUserReq`
- Response: `DeleteUserResp`

2. request definition



```golang
type DeleteUserReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteUserResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 6. N/A

1. route definition

- Url: /api/sys/user/reSetPassword
- Method: POST
- Request: `ReSetPasswordReq`
- Response: `ReSetPasswordResp`

2. request definition



```golang
type ReSetPasswordReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type ReSetPasswordResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 7. N/A

1. route definition

- Url: /api/sys/user/UpdateUserStatus
- Method: POST
- Request: `UserStatusReq`
- Response: `UserStatusResp`

2. request definition



```golang
type UserStatusReq struct {
	Id int64 `json:"id"`
	Status int64 `json:"status"` // 状态  0：禁用   1：正常
}
```


3. response definition



```golang
type UserStatusResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 8. N/A

1. route definition

- Url: /api/sys/user/selectAllData
- Method: POST
- Request: `SelectDataReq`
- Response: `SelectDataResp`

2. request definition



```golang
type SelectDataReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type SelectDataResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	RoleAll []*RoleAllResp `json:"roleAll"`
	DeptAll []*DeptAllResp `json:"deptAll"`
	JobAll []*JobAllResp `json:"jobAll"`
}
```

### 9. N/A

1. route definition

- Url: /api/sys/user/login
- Method: POST
- Request: `loginReq`
- Response: `loginResp`

2. request definition



```golang
type LoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
	CurrentAuthority string `json:"currentAuthority"`
	Id int64 `json:"id"`
	UserName string `json:"userName"`
	AccessToken string `json:"token"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

