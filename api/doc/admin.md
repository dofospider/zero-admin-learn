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

### 10. N/A

1. route definition

- Url: /api/sys/role/add
- Method: POST
- Request: `addRoleReq`
- Response: `addRoleResp`

2. request definition



```golang
type AddRoleReq struct {
	Name string `json:"name"` // 角色名称
	Remark string `json:"remark"` // 备注
	Status int64 `json:"status,optional"` // 状态
}
```


3. response definition



```golang
type AddRoleResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 11. N/A

1. route definition

- Url: /api/sys/role/list
- Method: POST
- Request: `ListRoleReq`
- Response: `ListRoleResp`

2. request definition



```golang
type ListRoleReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
	Name string `json:"name,optional "`
	Status int64 `json:"delFlag,optional "`
}
```


3. response definition



```golang
type ListRoleResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListRoleData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 12. N/A

1. route definition

- Url: /api/sys/role/update
- Method: POST
- Request: `UpdateRoleReq`
- Response: `UpdateRoleResp`

2. request definition



```golang
type UpdateRoleReq struct {
	Id int64 `json:"id"` // 编号
	Name string `json:"name"` // 角色名称
	Remark string `json:"remark"` // 备注
	Status int64 `json:"status"` // 状态
}
```


3. response definition



```golang
type UpdateRoleResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 13. N/A

1. route definition

- Url: /api/sys/role/delete
- Method: POST
- Request: `DeleteRoleReq`
- Response: `DeleteRoleResp`

2. request definition



```golang
type DeleteRoleReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteRoleResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 14. N/A

1. route definition

- Url: /api/sys/role/roleMenuIds
- Method: POST
- Request: `RoleMenuIdsReq`
- Response: `RoleMenuIdsResp`

2. request definition



```golang
type RoleMenuIdsReq struct {
}
```


3. response definition



```golang
type RoleMenuIdsResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 15. N/A

1. route definition

- Url: /api/sys/role/queryMenuByRoleId
- Method: POST
- Request: `RoleMenuReq`
- Response: `RoleMenuResp`

2. request definition



```golang
type RoleMenuReq struct {
	Id int64 `json:"id,optional"`
}
```


3. response definition



```golang
type RoleMenuResp struct {
	AllData []*ListMenuData `json:"allData"`
	RoleData []int64 `json:"userData"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 16. N/A

1. route definition

- Url: /api/sys/role/updateRoleMenu
- Method: POST
- Request: `UpdateRoleMenuReq`
- Response: `UpdateRoleMenuResp`

2. request definition



```golang
type UpdateRoleMenuReq struct {
	RoleId int64 `json:"roleId"`
	MenuIds []int64 `json:"menuIds"`
}
```


3. response definition



```golang
type UpdateRoleMenuResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 17. N/A

1. route definition

- Url: /api/sys/menu/add
- Method: POST
- Request: `addMenuReq`
- Response: `addMenuResp`

2. request definition



```golang
type AddMenuReq struct {
	Name string `json:"name"` // 菜单名称
	ParentId int64 `json:"parentId,optional"` // 父菜单ID，一级菜单为0
	Url string `json:"url,optional"` // 菜单URL,类型：1.普通页面（如用户管理， /sysmodel/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀&#43;目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)
	Perms string `json:"perms,optional"` // 授权(多个用逗号分隔，如：sysmodel:user:add,sysmodel:user:edit)
	Type int64 `json:"type,optional"` // 类型   0：目录   1：菜单   2：按钮
	Icon string `json:"icon,optional"` // 菜单图标
	OrderNum int64 `json:"orderNum,optional"` // 排序
	VuePath string `json:"vuePath,optional"` // vue系统的path
	VueComponent string `json:"vueComponent,optional"` // vue的页面
	VueIcon string `json:"vueIcon,optional"` // vue的图标
	VueRedirect string `json:"vueRedirect,optional"` // vue的路由重定向
}
```


3. response definition



```golang
type AddMenuResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 18. N/A

1. route definition

- Url: /api/sys/menu/list
- Method: POST
- Request: `ListMenuReq`
- Response: `ListMenuResp`

2. request definition



```golang
type ListMenuReq struct {
	Name string `json:"name,optional"`
	Url string `json:"url,optional "`
}
```


3. response definition



```golang
type ListMenuResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data []*ListtMenuData `json:"data"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 19. N/A

1. route definition

- Url: /api/sys/menu/update
- Method: POST
- Request: `UpdateMenuReq`
- Response: `UpdateMenuResp`

2. request definition



```golang
type UpdateMenuReq struct {
	Id int64 `json:"id"` // 编号
	Name string `json:"name"` // 菜单名称
	ParentId int64 `json:"parentId"` // 父菜单ID，一级菜单为0
	Url string `json:"url,optional"` // 菜单URL,类型：1.普通页面（如用户管理， /sysmodel/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀&#43;目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)
	Perms string `json:"perms,optional"` // 授权(多个用逗号分隔，如：sysmodel:user:add,sysmodel:user:edit)
	Type int64 `json:"type,optional"` // 类型   0：目录   1：菜单   2：按钮
	Icon string `json:"icon,optional"` // 菜单图标
	OrderNum int64 `json:"orderNum,optional"` // 排序
	VuePath string `json:"vuePath,optional"` // vue系统的path
	VueComponent string `json:"vueComponent,optional"` // vue的页面
	VueIcon string `json:"vueIcon,optional"` // vue的图标
	VueRedirect string `json:"vueRedirect,optional"` // vue的路由重定向
}
```


3. response definition



```golang
type UpdateMenuResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 20. N/A

1. route definition

- Url: /api/sys/menu/delete
- Method: POST
- Request: `DeleteMenuReq`
- Response: `DeleteMenuResp`

2. request definition



```golang
type DeleteMenuReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMenuResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 21. N/A

1. route definition

- Url: /api/sys/dict/add
- Method: POST
- Request: `addDictReq`
- Response: `addDictResp`

2. request definition



```golang
type AddDictReq struct {
	Value string `json:"value"` // 数据值
	Label string `json:"label"` // 标签名
	Type string `json:"type"` // 类型
	Description string `json:"description"` // 描述
	Sort string `json:"sort"` // 排序（升序）
	Remarks string `json:"remarks"` // 备注信息
}
```


3. response definition



```golang
type AddDictResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 22. N/A

1. route definition

- Url: /api/sys/dict/list
- Method: POST
- Request: `ListDictReq`
- Response: `ListDictResp`

2. request definition



```golang
type ListDictReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
	Value string `json:"value,optional"`
	Label string `json:"label,optional"`
	DelFlag int64 `json:"delFlag,optional"`
	Type string `json:"type,optional "`
}
```


3. response definition



```golang
type ListDictResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListDictData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 23. N/A

1. route definition

- Url: /api/sys/dict/update
- Method: POST
- Request: `UpdateDictReq`
- Response: `UpdateDictResp`

2. request definition



```golang
type UpdateDictReq struct {
	Id int64 `json:"id"` // 编号
	Value string `json:"value"` // 数据值
	Label string `json:"label"` // 标签名
	Type string `json:"type"` // 类型
	Description string `json:"description"` // 描述
	Sort float64 `json:"sort"` // 排序（升序）
	Remarks string `json:"remarks"` // 备注信息
}
```


3. response definition



```golang
type UpdateDictResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 24. N/A

1. route definition

- Url: /api/sys/dict/delete
- Method: POST
- Request: `DeleteDictReq`
- Response: `DeleteDictResp`

2. request definition



```golang
type DeleteDictReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteDictResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 25. N/A

1. route definition

- Url: /api/sys/dept/add
- Method: POST
- Request: `addDeptReq`
- Response: `addDeptResp`

2. request definition



```golang
type AddDeptReq struct {
	Name string `json:"name"` // 机构名称
	ParentId int64 `json:"parentId,optional"` // 上级机构ID，一级机构为0
	OrderNum int64 `json:"orderNum"` // 排序
}
```


3. response definition



```golang
type AddDeptResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 26. N/A

1. route definition

- Url: /api/sys/dept/list
- Method: POST
- Request: `ListDeptReq`
- Response: `ListDeptResp`

2. request definition



```golang
type ListDeptReq struct {
	Name string `json:"name,optional"`
	CreateBy string `json:"createBy,optional"`
}
```


3. response definition



```golang
type ListDeptResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data []*ListDeptData `json:"data"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 27. N/A

1. route definition

- Url: /api/sys/dept/update
- Method: POST
- Request: `UpdateDeptReq`
- Response: `UpdateDeptResp`

2. request definition



```golang
type UpdateDeptReq struct {
	Id int64 `json:"id"` // 编号
	Name string `json:"name"` // 机构名称
	ParentId int64 `json:"parentId"` // 上级机构ID，一级机构为0
	OrderNum int64 `json:"orderNum"` // 排序
}
```


3. response definition



```golang
type UpdateDeptResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 28. N/A

1. route definition

- Url: /api/sys/dept/delete
- Method: POST
- Request: `DeleteDeptReq`
- Response: `DeleteDeptResp`

2. request definition



```golang
type DeleteDeptReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteDeptResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 29. N/A

1. route definition

- Url: /api/sys/loginLog/list
- Method: POST
- Request: `ListLoginLogReq`
- Response: `ListLoginLogResp`

2. request definition



```golang
type ListLoginLogReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListLoginLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListLoginLogData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 30. N/A

1. route definition

- Url: /api/sys/loginLog/delete
- Method: POST
- Request: `DeleteLoginLogReq`
- Response: `DeleteLoginLogResp`

2. request definition



```golang
type DeleteLoginLogReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteLoginLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 31. N/A

1. route definition

- Url: /api/sys/sysLog/list
- Method: POST
- Request: `ListSysLogReq`
- Response: `ListSysLogResp`

2. request definition



```golang
type ListSysLogReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListSysLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListSysLogData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 32. N/A

1. route definition

- Url: /api/sys/sysLog/delete
- Method: POST
- Request: `DeleteSysLogReq`
- Response: `DeleteSysLogResp`

2. request definition



```golang
type DeleteSysLogReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteSysLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 33. N/A

1. route definition

- Url: /api/sys/config/add
- Method: POST
- Request: `addConfigReq`
- Response: `addConfigResp`

2. request definition



```golang
type AddConfigReq struct {
	Value string `json:"value"` // 数据值
	Label string `json:"label"` // 标签名
	Type string `json:"type"` // 类型
	Description string `json:"description"` // 描述
	Sort int64 `json:"sort"` // 排序（升序）
	Remarks string `json:"remarks"` // 备注信息
}
```


3. response definition



```golang
type AddConfigResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 34. N/A

1. route definition

- Url: /api/sys/config/list
- Method: POST
- Request: `ListConfigReq`
- Response: `ListConfigResp`

2. request definition



```golang
type ListConfigReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListConfigResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListConfigData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 35. N/A

1. route definition

- Url: /api/sys/config/update
- Method: POST
- Request: `UpdateConfigReq`
- Response: `UpdateConfigResp`

2. request definition



```golang
type UpdateConfigReq struct {
	Id int64 `json:"id"` // 编号
	Value string `json:"value"` // 数据值
	Label string `json:"label"` // 标签名
	Type string `json:"type"` // 类型
	Description string `json:"description"` // 描述
	Sort int64 `json:"sort"` // 排序（升序）
	Remarks string `json:"remarks"` // 备注信息
}
```


3. response definition



```golang
type UpdateConfigResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 36. N/A

1. route definition

- Url: /api/sys/config/delete
- Method: POST
- Request: `DeleteConfigReq`
- Response: `DeleteConfigResp`

2. request definition



```golang
type DeleteConfigReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteConfigResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 37. N/A

1. route definition

- Url: /api/sys/job/add
- Method: POST
- Request: `addJobReq`
- Response: `addJobResp`

2. request definition



```golang
type AddJobReq struct {
	JobName string `json:"jobName"` // 职位名称
	Remarks string `json:"remarks"` // 备注
	OrderNum int64 `json:"orderNum"` // 排序
}
```


3. response definition



```golang
type AddJobResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 38. N/A

1. route definition

- Url: /api/sys/job/list
- Method: POST
- Request: `ListJobReq`
- Response: `ListJobResp`

2. request definition



```golang
type ListJobReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
	JobName string `json:"jobName,optional"`
}
```


3. response definition



```golang
type ListJobResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListJobData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 39. N/A

1. route definition

- Url: /api/sys/job/update
- Method: POST
- Request: `UpdateJobReq`
- Response: `UpdateJobResp`

2. request definition



```golang
type UpdateJobReq struct {
	Id int64 `json:"id"` // 编号
	JobName string `json:"jobName"` // 职位名称
	OrderNum int64 `json:"orderNum"` // 排序
	Remarks string `json:"remarks"` // 备注
}
```


3. response definition



```golang
type UpdateJobResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 40. N/A

1. route definition

- Url: /api/sys/job/delete
- Method: POST
- Request: `DeleteJobReq`
- Response: `DeleteJobResp`

2. request definition



```golang
type DeleteJobReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteJobResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 41. N/A

1. route definition

- Url: /api/order/cart/add
- Method: POST
- Request: `addCartItemReq`
- Response: `addCartItemResp`

2. request definition



```golang
type AddCartItemReq struct {
	ProductId int64 `json:"productId"`
	ProductSkuId int64 `json:"productSkuId"`
	MemberId int64 `json:"memberId"`
	Quantity int64 `json:"quantity"` // 购买数量
	Price float64 `json:"price"` // 添加到购物车的价格
	ProductPic string `json:"productPic"` // 商品主图
	ProductName string `json:"productName"` // 商品名称
	ProductSubTitle string `json:"productSubTitle"` // 商品副标题（卖点）
	ProductSkuCode string `json:"productSkuCode"` // 商品sku条码
	MemberNickname string `json:"memberNickName"` // 会员昵称
	DeleteStatus int64 `json:"deleteStatus"` // 是否删除
	ProductCategoryId int64 `json:"productCategoryId"` // 商品分类
	ProductBrand string `json:"productBrand"`
	ProductSn string `json:"productSn"`
	ProductAttr string `json:"productAttr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}
```


3. response definition



```golang
type AddCartItemResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 42. N/A

1. route definition

- Url: /api/order/cart/list
- Method: POST
- Request: `ListCartItemReq`
- Response: `ListCartItemResp`

2. request definition



```golang
type ListCartItemReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListCartItemResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtCartItemData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 43. N/A

1. route definition

- Url: /api/order/cart/update
- Method: POST
- Request: `UpdateCartItemReq`
- Response: `UpdateCartItemResp`

2. request definition



```golang
type UpdateCartItemReq struct {
	Id int64 `json:"id"`
	ProductId int64 `json:"productId"`
	ProductSkuId int64 `json:"productSkuId"`
	MemberId int64 `json:"memberId"`
	Quantity int64 `json:"quantity"` // 购买数量
	Price float64 `json:"price"` // 添加到购物车的价格
	ProductPic string `json:"productPic"` // 商品主图
	ProductName string `json:"productName"` // 商品名称
	ProductSubTitle string `json:"productSubTitle"` // 商品副标题（卖点）
	ProductSkuCode string `json:"productSkuCode"` // 商品sku条码
	MemberNickname string `json:"memberNickName"` // 会员昵称
	DeleteStatus int64 `json:"deleteStatus"` // 是否删除
	ProductCategoryId int64 `json:"productCategoryId"` // 商品分类
	ProductBrand string `json:"productBrand"`
	ProductSn string `json:"productSn"`
	ProductAttr string `json:"productAttr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}
```


3. response definition



```golang
type UpdateCartItemResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 44. N/A

1. route definition

- Url: /api/order/cart/delete
- Method: POST
- Request: `DeleteCartItemReq`
- Response: `DeleteCartItemResp`

2. request definition



```golang
type DeleteCartItemReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteCartItemResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 45. N/A

1. route definition

- Url: /api/order/order/add
- Method: POST
- Request: `addOrderReq`
- Response: `addOrderResp`

2. request definition



```golang
type AddOrderReq struct {
	MemberId int64 `json:"memberId"`
	CouponId int64 `json:"couponId"`
	OrderSn string `json:"orderSn"` // 订单编号
	CreateTime string `json:"createTime"` // 提交时间
	MemberUsername string `json:"memberUserName"` // 用户帐号
	TotalAmount float64 `json:"totalAmount"` // 订单总金额
	PayAmount float64 `json:"payAmount"` // 应付金额（实际支付金额）
	FreightAmount float64 `json:"freightAmount"` // 运费金额
	PromotionAmount float64 `json:"promotionAmount"` // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount float64 `json:"integrationAmount"` // 积分抵扣金额
	CouponAmount float64 `json:"couponAmount"` // 优惠券抵扣金额
	DiscountAmount float64 `json:"discountAmount"` // 管理员后台调整订单使用的折扣金额
	PayType int64 `json:"payType"` // 支付方式：0-&gt;未支付；1-&gt;支付宝；2-&gt;微信
	SourceType int64 `json:"sourceType"` // 订单来源：0-&gt;PC订单；1-&gt;app订单
	Status int64 `json:"status"` // 订单状态：0-&gt;待付款；1-&gt;待发货；2-&gt;已发货；3-&gt;已完成；4-&gt;已关闭；5-&gt;无效订单
	OrderType int64 `json:"orderType"` // 订单类型：0-&gt;正常订单；1-&gt;秒杀订单
	DeliveryCompany string `json:"deliveryCompany"` // 物流公司(配送方式)
	DeliverySn string `json:"deliverySn"` // 物流单号
	AutoConfirmDay int64 `json:"autoConfirmDay"` // 自动确认时间（天）
	Integration int64 `json:"integration"` // 可以获得的积分
	Growth int64 `json:"growth"` // 可以活动的成长值
	PromotionInfo string `json:"promotionInfo"` // 活动信息
	BillType int64 `json:"billType"` // 发票类型：0-&gt;不开发票；1-&gt;电子发票；2-&gt;纸质发票
	BillHeader string `json:"billHeader"` // 发票抬头
	BillContent string `json:"billContent"` // 发票内容
	BillReceiverPhone string `json:"billReceiverPhone"` // 收票人电话
	BillReceiverEmail string `json:"billReceiverEmail"` // 收票人邮箱
	ReceiverName string `json:"receiverName"` // 收货人姓名
	ReceiverPhone string `json:"receiverPhone"` // 收货人电话
	ReceiverPostCode string `json:"receiverPostCode"` // 收货人邮编
	ReceiverProvince string `json:"receiverProvince"` // 省份/直辖市
	ReceiverCity string `json:"receiverCity"` // 城市
	ReceiverRegion string `json:"receiverRegion"` // 区
	ReceiverDetailAddress string `json:"receiverDetailAddress"` // 详细地址
	Note string `json:"note"` // 订单备注
	ConfirmStatus int64 `json:"confirmStatus"` // 确认收货状态：0-&gt;未确认；1-&gt;已确认
	DeleteStatus int64 `json:"deleteStatus"` // 删除状态：0-&gt;未删除；1-&gt;已删除
	UseIntegration int64 `json:"useIntegration"` // 下单时使用的积分
}
```


3. response definition



```golang
type AddOrderResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 46. N/A

1. route definition

- Url: /api/order/order/list
- Method: POST
- Request: `ListOrderReq`
- Response: `ListOrderResp`

2. request definition



```golang
type ListOrderReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListOrderResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtOrderData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 47. N/A

1. route definition

- Url: /api/order/order/update
- Method: POST
- Request: `UpdateOrderReq`
- Response: `UpdateOrderResp`

2. request definition



```golang
type UpdateOrderReq struct {
	Id int64 `json:"id"` // 订单id
	MemberId int64 `json:"memberId"`
	CouponId int64 `json:"couponId"`
	OrderSn string `json:"orderSn"` // 订单编号
	CreateTime string `json:"createTime"` // 提交时间
	MemberUsername string `json:"memberUserName"` // 用户帐号
	TotalAmount float64 `json:"totalAmount"` // 订单总金额
	PayAmount float64 `json:"payAmount"` // 应付金额（实际支付金额）
	FreightAmount float64 `json:"freightAmount"` // 运费金额
	PromotionAmount float64 `json:"promotionAmount"` // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount float64 `json:"integrationAmount"` // 积分抵扣金额
	CouponAmount float64 `json:"couponAmount"` // 优惠券抵扣金额
	DiscountAmount float64 `json:"discountAmount"` // 管理员后台调整订单使用的折扣金额
	PayType int64 `json:"payType"` // 支付方式：0-&gt;未支付；1-&gt;支付宝；2-&gt;微信
	SourceType int64 `json:"sourceType"` // 订单来源：0-&gt;PC订单；1-&gt;app订单
	Status int64 `json:"status"` // 订单状态：0-&gt;待付款；1-&gt;待发货；2-&gt;已发货；3-&gt;已完成；4-&gt;已关闭；5-&gt;无效订单
	OrderType int64 `json:"orderType"` // 订单类型：0-&gt;正常订单；1-&gt;秒杀订单
	DeliveryCompany string `json:"deliveryCompany"` // 物流公司(配送方式)
	DeliverySn string `json:"deliverySn"` // 物流单号
	AutoConfirmDay int64 `json:"autoConfirmDay"` // 自动确认时间（天）
	Integration int64 `json:"integration"` // 可以获得的积分
	Growth int64 `json:"growth"` // 可以活动的成长值
	PromotionInfo string `json:"promotionInfo"` // 活动信息
	BillType int64 `json:"billType"` // 发票类型：0-&gt;不开发票；1-&gt;电子发票；2-&gt;纸质发票
	BillHeader string `json:"billHeader"` // 发票抬头
	BillContent string `json:"billContent"` // 发票内容
	BillReceiverPhone string `json:"billReceiverPhone"` // 收票人电话
	BillReceiverEmail string `json:"billReceiverEmail"` // 收票人邮箱
	ReceiverName string `json:"receiverName"` // 收货人姓名
	ReceiverPhone string `json:"receiverPhone"` // 收货人电话
	ReceiverPostCode string `json:"receiverPostCode"` // 收货人邮编
	ReceiverProvince string `json:"receiverProvince"` // 省份/直辖市
	ReceiverCity string `json:"receiverCity"` // 城市
	ReceiverRegion string `json:"receiverRegion"` // 区
	ReceiverDetailAddress string `json:"receiverDetailAddress"` // 详细地址
	Note string `json:"note"` // 订单备注
	ConfirmStatus int64 `json:"confirmStatus"` // 确认收货状态：0-&gt;未确认；1-&gt;已确认
	DeleteStatus int64 `json:"deleteStatus"` // 删除状态：0-&gt;未删除；1-&gt;已删除
	UseIntegration int64 `json:"useIntegration"` // 下单时使用的积分
}
```


3. response definition



```golang
type UpdateOrderResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 48. N/A

1. route definition

- Url: /api/order/order/delete
- Method: POST
- Request: `DeleteOrderReq`
- Response: `DeleteOrderResp`

2. request definition



```golang
type DeleteOrderReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteOrderResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 49. N/A

1. route definition

- Url: /api/order/compayaddress/add
- Method: POST
- Request: `addCompayAddressReq`
- Response: `addCompayAddressResp`

2. request definition



```golang
type AddCompayAddressReq struct {
	AddressName string `json:"addressName"` // 地址名称
	SendStatus int64 `json:"sendStatus"` // 默认发货地址：0-&gt;否；1-&gt;是
	ReceiveStatus int64 `json:"receiveStatus"` // 是否默认收货地址：0-&gt;否；1-&gt;是
	Name string `json:"name"` // 收发货人姓名
	Phone string `json:"phone"` // 收货人电话
	Province string `json:"province"` // 省/直辖市
	City string `json:"city"` // 市
	Region string `json:"region"` // 区
	DetailAddress string `json:"detailAddress"` // 详细地址
}
```


3. response definition



```golang
type AddCompayAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 50. N/A

1. route definition

- Url: /api/order/compayaddress/list
- Method: POST
- Request: `ListCompayAddressReq`
- Response: `ListCompayAddressResp`

2. request definition



```golang
type ListCompayAddressReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListCompayAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtCompayAddressData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 51. N/A

1. route definition

- Url: /api/order/compayaddress/update
- Method: POST
- Request: `UpdateCompayAddressReq`
- Response: `UpdateCompayAddressResp`

2. request definition



```golang
type UpdateCompayAddressReq struct {
	Id int64 `json:"id"`
	AddressName string `json:"addressName"` // 地址名称
	SendStatus int64 `json:"sendStatus"` // 默认发货地址：0-&gt;否；1-&gt;是
	ReceiveStatus int64 `json:"receiveStatus"` // 是否默认收货地址：0-&gt;否；1-&gt;是
	Name string `json:"name"` // 收发货人姓名
	Phone string `json:"phone"` // 收货人电话
	Province string `json:"province"` // 省/直辖市
	City string `json:"city"` // 市
	Region string `json:"region"` // 区
	DetailAddress string `json:"detailAddress"` // 详细地址
}
```


3. response definition



```golang
type UpdateCompayAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 52. N/A

1. route definition

- Url: /api/order/compayaddress/delete
- Method: POST
- Request: `DeleteCompayAddressReq`
- Response: `DeleteCompayAddressResp`

2. request definition



```golang
type DeleteCompayAddressReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteCompayAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 53. N/A

1. route definition

- Url: /api/order/operatehistory/add
- Method: POST
- Request: `addOperateHistoryReq`
- Response: `addOperateHistoryResp`

2. request definition



```golang
type AddOperateHistoryReq struct {
	OrderId int64 `json:"orderId"` // 订单id
	OperateMan string `json:"operateMan"` // 操作人：用户；系统；后台管理员
	CreateTime string `json:"createTime"` // 操作时间
	OrderStatus int64 `json:"orderStatus"` // 订单状态：0-&gt;待付款；1-&gt;待发货；2-&gt;已发货；3-&gt;已完成；4-&gt;已关闭；5-&gt;无效订单
	Note string `json:"note"` // 备注
}
```


3. response definition



```golang
type AddOperateHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 54. N/A

1. route definition

- Url: /api/order/operatehistory/list
- Method: POST
- Request: `ListOperateHistoryReq`
- Response: `ListOperateHistoryResp`

2. request definition



```golang
type ListOperateHistoryReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListOperateHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtOperateHistoryData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 55. N/A

1. route definition

- Url: /api/order/operatehistory/update
- Method: POST
- Request: `UpdateOperateHistoryReq`
- Response: `UpdateOperateHistoryResp`

2. request definition



```golang
type UpdateOperateHistoryReq struct {
	Id int64 `json:"id"`
	OrderId int64 `json:"orderId"` // 订单id
	OperateMan string `json:"operateMan"` // 操作人：用户；系统；后台管理员
	CreateTime string `json:"createTime"` // 操作时间
	OrderStatus int64 `json:"orderStatus"` // 订单状态：0-&gt;待付款；1-&gt;待发货；2-&gt;已发货；3-&gt;已完成；4-&gt;已关闭；5-&gt;无效订单
	Note string `json:"note"` // 备注
}
```


3. response definition



```golang
type UpdateOperateHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 56. N/A

1. route definition

- Url: /api/order/operatehistory/delete
- Method: POST
- Request: `DeleteOperateHistoryReq`
- Response: `DeleteOperateHistoryResp`

2. request definition



```golang
type DeleteOperateHistoryReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteOperateHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 57. N/A

1. route definition

- Url: /api/order/returnapply/add
- Method: POST
- Request: `addReturnApplyReq`
- Response: `addReturnApplyResp`

2. request definition



```golang
type AddReturnApplyReq struct {
	OrderId int64 `json:"orderId"` // 订单id
	CompanyAddressId int64 `json:"companyAddressId"` // 收货地址表id
	ProductId int64 `json:"productId"` // 退货商品id
	OrderSn string `json:"orderSn"` // 订单编号
	CreateTime string `json:"createTime"` // 申请时间
	MemberUsername string `json:"memberUserName"` // 会员用户名
	ReturnAmount float64 `json:"returnAmount"` // 退款金额
	ReturnName string `json:"returnName"` // 退货人姓名
	ReturnPhone string `json:"returnPhone"` // 退货人电话
	Status int64 `json:"status"` // 申请状态：0-&gt;待处理；1-&gt;退货中；2-&gt;已完成；3-&gt;已拒绝
	HandleTime string `json:"handleTime"` // 处理时间
	ProductPic string `json:"productPic"` // 商品图片
	ProductName string `json:"productName"` // 商品名称
	ProductBrand string `json:"productBrand"` // 商品品牌
	ProductAttr string `json:"productAttr"` // 商品销售属性：颜色：红色；尺码：xl;
	ProductCount int64 `json:"productCount"` // 退货数量
	ProductPrice float64 `json:"productPrice"` // 商品单价
	ProductRealPrice float64 `json:"productRealPrice"` // 商品实际支付单价
	Reason string `json:"reason"` // 原因
	Description string `json:"description"` // 描述
	ProofPics string `json:"proofPics"` // 凭证图片，以逗号隔开
	HandleNote string `json:"handleNote"` // 处理备注
	HandleMan string `json:"handleMan"` // 处理人员
	ReceiveMan string `json:"receiveMan"` // 收货人
	ReceiveTime string `json:"receiveTime"` // 收货时间
	ReceiveNote string `json:"receiveNote"` // 收货备注
}
```


3. response definition



```golang
type AddReturnApplyResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 58. N/A

1. route definition

- Url: /api/order/returnapply/list
- Method: POST
- Request: `ListReturnApplyReq`
- Response: `ListReturnApplyResp`

2. request definition



```golang
type ListReturnApplyReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListReturnApplyResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtReturnApplyData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 59. N/A

1. route definition

- Url: /api/order/returnapply/update
- Method: POST
- Request: `UpdateReturnApplyReq`
- Response: `UpdateReturnApplyResp`

2. request definition



```golang
type UpdateReturnApplyReq struct {
	Id int64 `json:"id"`
	OrderId int64 `json:"orderId"` // 订单id
	CompanyAddressId int64 `json:"companyAddressId"` // 收货地址表id
	ProductId int64 `json:"productId"` // 退货商品id
	OrderSn string `json:"orderSn"` // 订单编号
	CreateTime string `json:"createTime"` // 申请时间
	MemberUsername string `json:"memberUserName"` // 会员用户名
	ReturnAmount float64 `json:"returnAmount"` // 退款金额
	ReturnName string `json:"returnName"` // 退货人姓名
	ReturnPhone string `json:"returnPhone"` // 退货人电话
	Status int64 `json:"status"` // 申请状态：0-&gt;待处理；1-&gt;退货中；2-&gt;已完成；3-&gt;已拒绝
	HandleTime string `json:"handleTime"` // 处理时间
	ProductPic string `json:"productPic"` // 商品图片
	ProductName string `json:"productName"` // 商品名称
	ProductBrand string `json:"productBrand"` // 商品品牌
	ProductAttr string `json:"productAttr"` // 商品销售属性：颜色：红色；尺码：xl;
	ProductCount int64 `json:"productCount"` // 退货数量
	ProductPrice float64 `json:"productPrice"` // 商品单价
	ProductRealPrice float64 `json:"productRealPrice"` // 商品实际支付单价
	Reason string `json:"reason"` // 原因
	Description string `json:"description"` // 描述
	ProofPics string `json:"proofPics"` // 凭证图片，以逗号隔开
	HandleNote string `json:"handleNote"` // 处理备注
	HandleMan string `json:"handleMan"` // 处理人员
	ReceiveMan string `json:"receiveMan"` // 收货人
	ReceiveTime string `json:"receiveTime"` // 收货时间
	ReceiveNote string `json:"receiveNote"` // 收货备注
}
```


3. response definition



```golang
type UpdateReturnApplyResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 60. N/A

1. route definition

- Url: /api/order/returnapply/delete
- Method: POST
- Request: `DeleteReturnApplyReq`
- Response: `DeleteReturnApplyResp`

2. request definition



```golang
type DeleteReturnApplyReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteReturnApplyResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 61. N/A

1. route definition

- Url: /api/order/returnreason/add
- Method: POST
- Request: `addReturnResonReq`
- Response: `addReturnResonResp`

2. request definition



```golang
type AddReturnResonReq struct {
	Name string `json:"name"` // 退货类型
	Sort int64 `json:"sort"`
	Status int64 `json:"status"` // 状态：0-&gt;不启用；1-&gt;启用
}
```


3. response definition



```golang
type AddReturnResonResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 62. N/A

1. route definition

- Url: /api/order/returnreason/list
- Method: POST
- Request: `ListReturnResonReq`
- Response: `ListReturnResonResp`

2. request definition



```golang
type ListReturnResonReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListReturnResonResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtReturnResonData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 63. N/A

1. route definition

- Url: /api/order/returnreason/update
- Method: POST
- Request: `UpdateReturnResonReq`
- Response: `UpdateReturnResonResp`

2. request definition



```golang
type UpdateReturnResonReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"` // 退货类型
	Sort int64 `json:"sort"`
	Status int64 `json:"status"` // 状态：0-&gt;不启用；1-&gt;启用
}
```


3. response definition



```golang
type UpdateReturnResonResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 64. N/A

1. route definition

- Url: /api/order/returnreason/delete
- Method: POST
- Request: `DeleteReturnResonReq`
- Response: `DeleteReturnResonResp`

2. request definition



```golang
type DeleteReturnResonReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteReturnResonResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 65. N/A

1. route definition

- Url: /api/order/setting/add
- Method: POST
- Request: `addOrderSettingReq`
- Response: `addOrderSettingResp`

2. request definition



```golang
type AddOrderSettingReq struct {
	FlashOrderOvertime int64 `json:"flashOrderOvertime"` // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int64 `json:"normalOrderOvertime"` // 正常订单超时时间(分)
	ConfirmOvertime int64 `json:"confirmOvertime"` // 发货后自动确认收货时间（天）
	FinishOvertime int64 `json:"finishOvertime"` // 自动完成交易时间，不能申请售后（天）
	CommentOvertime int64 `json:"commentOvertime"` // 订单完成后自动好评时间（天）
}
```


3. response definition



```golang
type AddOrderSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 66. N/A

1. route definition

- Url: /api/order/setting/list
- Method: POST
- Request: `ListOrderSettingReq`
- Response: `ListOrderSettingResp`

2. request definition



```golang
type ListOrderSettingReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListOrderSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtOrderSettingData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 67. N/A

1. route definition

- Url: /api/order/setting/update
- Method: POST
- Request: `UpdateOrderSettingReq`
- Response: `UpdateOrderSettingResp`

2. request definition



```golang
type UpdateOrderSettingReq struct {
	Id int64 `json:"id"`
	FlashOrderOvertime int64 `json:"flashOrderOvertime"` // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int64 `json:"normalOrderOvertime"` // 正常订单超时时间(分)
	ConfirmOvertime int64 `json:"confirmOvertime"` // 发货后自动确认收货时间（天）
	FinishOvertime int64 `json:"finishOvertime"` // 自动完成交易时间，不能申请售后（天）
	CommentOvertime int64 `json:"commentOvertime"` // 订单完成后自动好评时间（天）
}
```


3. response definition



```golang
type UpdateOrderSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 68. N/A

1. route definition

- Url: /api/order/setting/delete
- Method: POST
- Request: `DeleteOrderSettingReq`
- Response: `DeleteOrderSettingResp`

2. request definition



```golang
type DeleteOrderSettingReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteOrderSettingResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 69. N/A

1. route definition

- Url: /api/product/product/add
- Method: POST
- Request: `addProductReq`
- Response: `addProductResp`

2. request definition



```golang
type AddProductReq struct {
	BrandId int64 `json:"brandId"`
	ProductCategoryId int64 `json:"productCategoryId"`
	FeightTemplateId int64 `json:"feightTemplateId"`
	ProductAttributeCategoryId int64 `json:"productAttributeCategoryId"`
	Name string `json:"name"`
	Pic string `json:"pic"`
	ProductSn string `json:"productSn"` // 货号
	DeleteStatus int64 `json:"deleteStatus"` // 删除状态：0-&gt;未删除；1-&gt;已删除
	PublishStatus int64 `json:"publishStatus"` // 上架状态：0-&gt;下架；1-&gt;上架
	NewStatus int64 `json:"newStatus"` // 新品状态:0-&gt;不是新品；1-&gt;新品
	RecommandStatus int64 `json:"recommandStatus"` // 推荐状态；0-&gt;不推荐；1-&gt;推荐
	VerifyStatus int64 `json:"verifyStatus"` // 审核状态：0-&gt;未审核；1-&gt;审核通过
	Sort int64 `json:"sort"` // 排序
	Sale int64 `json:"sale"` // 销量
	Price float64 `json:"price"`
	PromotionPrice float64 `json:"promotionPrice"` // 促销价格
	GiftGrowth int64 `json:"giftGrowth"` // 赠送的成长值
	GiftPoint int64 `json:"giftPoint"` // 赠送的积分
	UsePointLimit int64 `json:"usePointLimit"` // 限制使用的积分数
	SubTitle string `json:"subTitle"` // 副标题
	Description string `json:"description"` // 商品描述
	OriginalPrice float64 `json:"originalPrice"` // 市场价
	Stock int64 `json:"stock"` // 库存
	LowStock int64 `json:"lowStock"` // 库存预警值
	Unit string `json:"unit"` // 单位
	Weight float64 `json:"weight"` // 商品重量，默认为克
	PreviewStatus int64 `json:"previewStatus"` // 是否为预告商品：0-&gt;不是；1-&gt;是
	ServiceIds string `json:"serviceIds"` // 以逗号分割的产品服务：1-&gt;无忧退货；2-&gt;快速退款；3-&gt;免费包邮
	Keywords string `json:"keywords"`
	Note string `json:"note"`
	AlbumPics string `json:"albumPics"` // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle string `json:"detailTitle"`
	DetailDesc string `json:"detailDesc"`
	DetailHtml string `json:"detailHtml"` // 产品详情网页内容
	DetailMobileHtml string `json:"detailMobileHtml"` // 移动端网页详情
	PromotionStartTime string `json:"promotionStartTime"` // 促销开始时间
	PromotionEndTime string `json:"promotionEndTime"` // 促销结束时间
	PromotionPerLimit int64 `json:"promotionPerLimit"` // 活动限购数量
	PromotionType int64 `json:"promotionType"` // 促销类型：0-&gt;没有促销使用原价;1-&gt;使用促销价；2-&gt;使用会员价；3-&gt;使用阶梯价格；4-&gt;使用满减价格；5-&gt;限时购
	BrandName string `json:"brandName"` // 品牌名称
	ProductCategoryName string `json:"productCategoryName"` // 商品分类名称
}
```


3. response definition



```golang
type AddProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 70. N/A

1. route definition

- Url: /api/product/product/list
- Method: POST
- Request: `ListProductReq`
- Response: `ListProductResp`

2. request definition



```golang
type ListProductReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtProductData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 71. N/A

1. route definition

- Url: /api/product/product/update
- Method: POST
- Request: `UpdateProductReq`
- Response: `UpdateProductResp`

2. request definition



```golang
type UpdateProductReq struct {
	Id int64 `json:"id"`
	BrandId int64 `json:"brandId"`
	ProductCategoryId int64 `json:"productCategoryId"`
	FeightTemplateId int64 `json:"feightTemplateId"`
	ProductAttributeCategoryId int64 `json:"productAttributeCategoryId"`
	Name string `json:"name"`
	Pic string `json:"pic"`
	ProductSn string `json:"productSn"` // 货号
	DeleteStatus int64 `json:"deleteStatus"` // 删除状态：0-&gt;未删除；1-&gt;已删除
	PublishStatus int64 `json:"publishStatus"` // 上架状态：0-&gt;下架；1-&gt;上架
	NewStatus int64 `json:"newStatus"` // 新品状态:0-&gt;不是新品；1-&gt;新品
	RecommandStatus int64 `json:"recommandStatus"` // 推荐状态；0-&gt;不推荐；1-&gt;推荐
	VerifyStatus int64 `json:"verifyStatus"` // 审核状态：0-&gt;未审核；1-&gt;审核通过
	Sort int64 `json:"sort"` // 排序
	Sale int64 `json:"sale"` // 销量
	Price float64 `json:"price"`
	PromotionPrice float64 `json:"promotionPrice"` // 促销价格
	GiftGrowth int64 `json:"giftGrowth"` // 赠送的成长值
	GiftPoint int64 `json:"giftPoint"` // 赠送的积分
	UsePointLimit int64 `json:"usePointLimit"` // 限制使用的积分数
	SubTitle string `json:"subTitle"` // 副标题
	Description string `json:"description"` // 商品描述
	OriginalPrice float64 `json:"originalPrice"` // 市场价
	Stock int64 `json:"stock"` // 库存
	LowStock int64 `json:"lowStock"` // 库存预警值
	Unit string `json:"unit"` // 单位
	Weight float64 `json:"weight"` // 商品重量，默认为克
	PreviewStatus int64 `json:"previewStatus"` // 是否为预告商品：0-&gt;不是；1-&gt;是
	ServiceIds string `json:"serviceIds"` // 以逗号分割的产品服务：1-&gt;无忧退货；2-&gt;快速退款；3-&gt;免费包邮
	Keywords string `json:"keywords"`
	Note string `json:"note"`
	AlbumPics string `json:"albumPics"` // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle string `json:"detailTitle"`
	DetailDesc string `json:"detailDesc"`
	DetailHtml string `json:"detailHtml"` // 产品详情网页内容
	DetailMobileHtml string `json:"detailMobileHtml"` // 移动端网页详情
	PromotionStartTime string `json:"promotionStartTime"` // 促销开始时间
	PromotionEndTime string `json:"promotionEndTime"` // 促销结束时间
	PromotionPerLimit int64 `json:"promotionPerLimit"` // 活动限购数量
	PromotionType int64 `json:"promotionType"` // 促销类型：0-&gt;没有促销使用原价;1-&gt;使用促销价；2-&gt;使用会员价；3-&gt;使用阶梯价格；4-&gt;使用满减价格；5-&gt;限时购
	BrandName string `json:"brandName"` // 品牌名称
	ProductCategoryName string `json:"productCategoryName"` // 商品分类名称
}
```


3. response definition



```golang
type UpdateProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 72. N/A

1. route definition

- Url: /api/product/product/delete
- Method: POST
- Request: `DeleteProductReq`
- Response: `DeleteProductResp`

2. request definition



```golang
type DeleteProductReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 73. N/A

1. route definition

- Url: /api/product/brand/add
- Method: POST
- Request: `addProductBrandReq`
- Response: `addProductBrandResp`

2. request definition



```golang
type AddProductBrandReq struct {
	Name string `json:"name"`
	FirstLetter string `json:"firstLetter"` // 首字母
	Sort int64 `json:"sort"`
	FactoryStatus int64 `json:"factoryStatus"` // 是否为品牌制造商：0-&gt;不是；1-&gt;是
	ShowStatus int64 `json:"showStatus"`
	ProductCount int64 `json:"productCount"` // 产品数量
	ProductCommentCount int64 `json:"productCommentCount"` // 产品评论数量
	Logo string `json:"logo"` // 品牌logo
	BigPic string `json:"bigPic"` // 专区大图
	BrandStory string `json:"brandStory"` // 品牌故事
}
```


3. response definition



```golang
type AddProductBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 74. N/A

1. route definition

- Url: /api/product/brand/list
- Method: POST
- Request: `ListProductBrandReq`
- Response: `ListProductBrandResp`

2. request definition



```golang
type ListProductBrandReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListProductBrandResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtProductBrandData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 75. N/A

1. route definition

- Url: /api/product/brand/update
- Method: POST
- Request: `UpdateProductBrandReq`
- Response: `UpdateProductBrandResp`

2. request definition



```golang
type UpdateProductBrandReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	FirstLetter string `json:"firstLetter"` // 首字母
	Sort int64 `json:"sort"`
	FactoryStatus int64 `json:"factoryStatus"` // 是否为品牌制造商：0-&gt;不是；1-&gt;是
	ShowStatus int64 `json:"showStatus"`
	ProductCount int64 `json:"productCount"` // 产品数量
	ProductCommentCount int64 `json:"productCommentCount"` // 产品评论数量
	Logo string `json:"logo"` // 品牌logo
	BigPic string `json:"bigPic"` // 专区大图
	BrandStory string `json:"brandStory"` // 品牌故事
}
```


3. response definition



```golang
type UpdateProductBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 76. N/A

1. route definition

- Url: /api/product/brand/delete
- Method: POST
- Request: `DeleteProductBrandReq`
- Response: `DeleteProductBrandResp`

2. request definition



```golang
type DeleteProductBrandReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteProductBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 77. N/A

1. route definition

- Url: /api/product/category/add
- Method: POST
- Request: `addProductCategoryReq`
- Response: `addProductCategoryResp`

2. request definition



```golang
type AddProductCategoryReq struct {
	ParentId int64 `json:"parentId"` // 上机分类的编号：0表示一级分类
	Name string `json:"name"`
	Level int64 `json:"level"` // 分类级别：0-&gt;1级；1-&gt;2级
	ProductCount int64 `json:"productCount"`
	ProductUnit string `json:"productUnit"`
	NavStatus int64 `json:"navStatus"` // 是否显示在导航栏：0-&gt;不显示；1-&gt;显示
	ShowStatus int64 `json:"showStatus"` // 显示状态：0-&gt;不显示；1-&gt;显示
	Sort int64 `json:"sort"`
	Icon string `json:"icon"` // 图标
	Keywords string `json:"keywords"`
	Description string `json:"description"` // 描述
}
```


3. response definition



```golang
type AddProductCategoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 78. N/A

1. route definition

- Url: /api/product/category/list
- Method: POST
- Request: `ListProductCategoryReq`
- Response: `ListProductCategoryResp`

2. request definition



```golang
type ListProductCategoryReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListProductCategoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtProductCategoryData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 79. N/A

1. route definition

- Url: /api/product/category/update
- Method: POST
- Request: `UpdateProductCategoryReq`
- Response: `UpdateProductCategoryResp`

2. request definition



```golang
type UpdateProductCategoryReq struct {
	Id int64 `json:"id"`
	ParentId int64 `json:"parentId"` // 上机分类的编号：0表示一级分类
	Name string `json:"name"`
	Level int64 `json:"level"` // 分类级别：0-&gt;1级；1-&gt;2级
	ProductCount int64 `json:"productCount"`
	ProductUnit string `json:"productUnit"`
	NavStatus int64 `json:"navStatus"` // 是否显示在导航栏：0-&gt;不显示；1-&gt;显示
	ShowStatus int64 `json:"showStatus"` // 显示状态：0-&gt;不显示；1-&gt;显示
	Sort int64 `json:"sort"`
	Icon string `json:"icon"` // 图标
	Keywords string `json:"keywords"`
	Description string `json:"description"` // 描述
}
```


3. response definition



```golang
type UpdateProductCategoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 80. N/A

1. route definition

- Url: /api/product/category/delete
- Method: POST
- Request: `DeleteProductCategoryReq`
- Response: `DeleteProductCategoryResp`

2. request definition



```golang
type DeleteProductCategoryReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteProductCategoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 81. N/A

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

### 82. N/A

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

### 83. N/A

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

### 84. N/A

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

### 85. N/A

1. route definition

- Url: /api/product/feighttemplate/add
- Method: POST
- Request: `addFeightTemplateReq`
- Response: `addFeightTemplateResp`

2. request definition



```golang
type AddFeightTemplateReq struct {
	Name string `json:"name"`
	ChargeType int64 `json:"chargeType"` // 计费类型:0-&gt;按重量；1-&gt;按件数
	FirstWeight float64 `json:"first_weight"` // 首重kg
	FirstFee float64 `json:"firstFee"` // 首费（元）
	ContinueWeight float64 `json:"continue_weight"`
	ContinmeFee float64 `json:"continmeFee"`
	Dest string `json:"dest"` // 目的地（省、市）
}
```


3. response definition



```golang
type AddFeightTemplateResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 86. N/A

1. route definition

- Url: /api/product/feighttemplate/list
- Method: POST
- Request: `ListFeightTemplateReq`
- Response: `ListFeightTemplateResp`

2. request definition



```golang
type ListFeightTemplateReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListFeightTemplateResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtFeightTemplateData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 87. N/A

1. route definition

- Url: /api/product/feighttemplate/update
- Method: POST
- Request: `UpdateFeightTemplateReq`
- Response: `UpdateFeightTemplateResp`

2. request definition



```golang
type UpdateFeightTemplateReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	ChargeType int64 `json:"chargeType"` // 计费类型:0-&gt;按重量；1-&gt;按件数
	FirstWeight float64 `json:"first_weight"` // 首重kg
	FirstFee float64 `json:"firstFee"` // 首费（元）
	ContinueWeight float64 `json:"continue_weight"`
	ContinmeFee float64 `json:"continmeFee"`
	Dest string `json:"dest"` // 目的地（省、市）
}
```


3. response definition



```golang
type UpdateFeightTemplateResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 88. N/A

1. route definition

- Url: /api/product/feighttemplate/delete
- Method: POST
- Request: `DeleteFeightTemplateReq`
- Response: `DeleteFeightTemplateResp`

2. request definition



```golang
type DeleteFeightTemplateReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteFeightTemplateResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 89. N/A

1. route definition

- Url: /api/product/memberprice/add
- Method: POST
- Request: `addMemberPriceReq`
- Response: `addMemberPriceResp`

2. request definition



```golang
type AddMemberPriceReq struct {
	ProductId int64 `json:"productId"`
	MemberLevelId int64 `json:"memberLevelId"`
	MemberPrice float64 `json:"memberPrice"` // 会员价格
	MemberLevelName string `json:"memberLevelName"`
}
```


3. response definition



```golang
type AddMemberPriceResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 90. N/A

1. route definition

- Url: /api/product/memberprice/list
- Method: POST
- Request: `ListMemberPriceReq`
- Response: `ListMemberPriceResp`

2. request definition



```golang
type ListMemberPriceReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberPriceResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberPriceData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 91. N/A

1. route definition

- Url: /api/product/memberprice/update
- Method: POST
- Request: `UpdateMemberPriceReq`
- Response: `UpdateMemberPriceResp`

2. request definition



```golang
type UpdateMemberPriceReq struct {
	Id int64 `json:"id"`
	ProductId int64 `json:"productId"`
	MemberLevelId int64 `json:"memberLevelId"`
	MemberPrice float64 `json:"memberPrice"` // 会员价格
	MemberLevelName string `json:"memberLevelName"`
}
```


3. response definition



```golang
type UpdateMemberPriceResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 92. N/A

1. route definition

- Url: /api/product/memberprice/delete
- Method: POST
- Request: `DeleteMemberPriceReq`
- Response: `DeleteMemberPriceResp`

2. request definition



```golang
type DeleteMemberPriceReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberPriceResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 93. N/A

1. route definition

- Url: /api/product/skustock/add
- Method: POST
- Request: `addSkuStockReq`
- Response: `addSkuStockResp`

2. request definition



```golang
type AddSkuStockReq struct {
	ProductId int64 `json:"productId"`
	SkuCode string `json:"skuCode"` // sku编码
	Price float64 `json:"price"`
	Stock int64 `json:"stock"` // 库存
	LowStock int64 `json:"lowStock"` // 预警库存
	Pic string `json:"pic"` // 展示图片
	Sale int64 `json:"sale"` // 销量
	PromotionPrice float64 `json:"promotionPrice"` // 单品促销价格
	LockStock int64 `json:"lockStock"` // 锁定库存
	SpData string `json:"spData"` // 商品销售属性，json格式
}
```


3. response definition



```golang
type AddSkuStockResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 94. N/A

1. route definition

- Url: /api/product/skustock/list
- Method: POST
- Request: `ListSkuStockReq`
- Response: `ListSkuStockResp`

2. request definition



```golang
type ListSkuStockReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListSkuStockResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtSkuStockData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 95. N/A

1. route definition

- Url: /api/product/skustock/update
- Method: POST
- Request: `UpdateSkuStockReq`
- Response: `UpdateSkuStockResp`

2. request definition



```golang
type UpdateSkuStockReq struct {
	Id int64 `json:"id"`
	ProductId int64 `json:"productId"`
	SkuCode string `json:"skuCode"` // sku编码
	Price float64 `json:"price"`
	Stock int64 `json:"stock"` // 库存
	LowStock int64 `json:"lowStock"` // 预警库存
	Pic string `json:"pic"` // 展示图片
	Sale int64 `json:"sale"` // 销量
	PromotionPrice float64 `json:"promotionPrice"` // 单品促销价格
	LockStock int64 `json:"lockStock"` // 锁定库存
	SpData string `json:"spData"` // 商品销售属性，json格式
}
```


3. response definition



```golang
type UpdateSkuStockResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 96. N/A

1. route definition

- Url: /api/product/skustock/delete
- Method: POST
- Request: `DeleteSkuStockReq`
- Response: `DeleteSkuStockResp`

2. request definition



```golang
type DeleteSkuStockReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteSkuStockResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 97. N/A

1. route definition

- Url: /api/sms/coupon/add
- Method: POST
- Request: `addCouponReq`
- Response: `addCouponResp`

2. request definition



```golang
type AddCouponReq struct {
	Type int64 `json:"type"` // 优惠券类型；0-&gt;全场赠券；1-&gt;会员赠券；2-&gt;购物赠券；3-&gt;注册赠券
	Name string `json:"name"`
	Platform int64 `json:"platform"` // 使用平台：0-&gt;全部；1-&gt;移动；2-&gt;PC
	Count int64 `json:"count"` // 数量
	Amount float64 `json:"amount"` // 金额
	PerLimit int64 `json:"perLimit"` // 每人限领张数
	MinPoint float64 `json:"minPoint"` // 使用门槛；0表示无门槛
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	UseType int64 `json:"useType"` // 使用类型：0-&gt;全场通用；1-&gt;指定分类；2-&gt;指定商品
	Note string `json:"note"` // 备注
	PublishCount int64 `json:"publishCount"` // 发行数量
	UseCount int64 `json:"useCount"` // 已使用数量
	ReceiveCount int64 `json:"receiveCount"` // 领取数量
	EnableTime string `json:"enableTime"` // 可以领取的日期
	Code string `json:"code"` // 优惠码
	MemberLevel int64 `json:"memberLevel"` // 可领取的会员类型：0-&gt;无限时
}
```


3. response definition



```golang
type AddCouponResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 98. N/A

1. route definition

- Url: /api/sms/coupon/list
- Method: POST
- Request: `ListCouponReq`
- Response: `ListCouponResp`

2. request definition



```golang
type ListCouponReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListCouponResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtCouponData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 99. N/A

1. route definition

- Url: /api/sms/coupon/update
- Method: POST
- Request: `UpdateCouponReq`
- Response: `UpdateCouponResp`

2. request definition



```golang
type UpdateCouponReq struct {
	Id int64 `json:"id"`
	Type int64 `json:"type"` // 优惠券类型；0-&gt;全场赠券；1-&gt;会员赠券；2-&gt;购物赠券；3-&gt;注册赠券
	Name string `json:"name"`
	Platform int64 `json:"platform"` // 使用平台：0-&gt;全部；1-&gt;移动；2-&gt;PC
	Count int64 `json:"count"` // 数量
	Amount float64 `json:"amount"` // 金额
	PerLimit int64 `json:"perLimit"` // 每人限领张数
	MinPoint float64 `json:"minPoint"` // 使用门槛；0表示无门槛
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	UseType int64 `json:"useType"` // 使用类型：0-&gt;全场通用；1-&gt;指定分类；2-&gt;指定商品
	Note string `json:"note"` // 备注
	PublishCount int64 `json:"publishCount"` // 发行数量
	UseCount int64 `json:"useCount"` // 已使用数量
	ReceiveCount int64 `json:"receiveCount"` // 领取数量
	EnableTime string `json:"enableTime"` // 可以领取的日期
	Code string `json:"code"` // 优惠码
	MemberLevel int64 `json:"memberLevel"` // 可领取的会员类型：0-&gt;无限时
}
```


3. response definition



```golang
type UpdateCouponResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 100. N/A

1. route definition

- Url: /api/sms/coupon/delete
- Method: POST
- Request: `DeleteCouponReq`
- Response: `DeleteCouponResp`

2. request definition



```golang
type DeleteCouponReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteCouponResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 101. N/A

1. route definition

- Url: /api/sms/couponhistory/add
- Method: POST
- Request: `addCouponHistoryReq`
- Response: `addCouponHistoryResp`

2. request definition



```golang
type AddCouponHistoryReq struct {
	CouponId int64 `json:"couponId"`
	MemberId int64 `json:"memberId"`
	CouponCode string `json:"couponCode"`
	MemberNickname string `json:"memberNickName"` // 领取人昵称
	GetType int64 `json:"getType"` // 获取类型：0-&gt;后台赠送；1-&gt;主动获取
	CreateTime string `json:"createTime"`
	UseStatus int64 `json:"useStatus"` // 使用状态：0-&gt;未使用；1-&gt;已使用；2-&gt;已过期
	UseTime string `json:"useTime"` // 使用时间
	OrderId int64 `json:"orderId"` // 订单编号
	OrderSn string `json:"orderSn"` // 订单号码
}
```


3. response definition



```golang
type AddCouponHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 102. N/A

1. route definition

- Url: /api/sms/couponhistory/list
- Method: POST
- Request: `ListCouponHistoryReq`
- Response: `ListCouponHistoryResp`

2. request definition



```golang
type ListCouponHistoryReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListCouponHistoryResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtCouponHistoryData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 103. N/A

1. route definition

- Url: /api/sms/couponhistory/update
- Method: POST
- Request: `UpdateCouponHistoryReq`
- Response: `UpdateCouponHistoryResp`

2. request definition



```golang
type UpdateCouponHistoryReq struct {
	Id int64 `json:"id"`
	CouponId int64 `json:"couponId"`
	MemberId int64 `json:"memberId"`
	CouponCode string `json:"couponCode"`
	MemberNickname string `json:"memberNickName"` // 领取人昵称
	GetType int64 `json:"getType"` // 获取类型：0-&gt;后台赠送；1-&gt;主动获取
	CreateTime string `json:"createTime"`
	UseStatus int64 `json:"useStatus"` // 使用状态：0-&gt;未使用；1-&gt;已使用；2-&gt;已过期
	UseTime string `json:"useTime"` // 使用时间
	OrderId int64 `json:"orderId"` // 订单编号
	OrderSn string `json:"orderSn"` // 订单号码
}
```


3. response definition



```golang
type UpdateCouponHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 104. N/A

1. route definition

- Url: /api/sms/couponhistory/delete
- Method: POST
- Request: `DeleteCouponHistoryReq`
- Response: `DeleteCouponHistoryResp`

2. request definition



```golang
type DeleteCouponHistoryReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteCouponHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 105. N/A

1. route definition

- Url: /api/sms/flashpromotion/add
- Method: POST
- Request: `addFlashPromotionReq`
- Response: `addFlashPromotionResp`

2. request definition



```golang
type AddFlashPromotionReq struct {
	Title string `json:"title"`
	StartDate string `json:"startDate"` // 开始日期
	EndDate string `json:"endDate"` // 结束日期
	Status int64 `json:"status"` // 上下线状态
	CreateTime string `json:"createTime"` // 秒杀时间段名称
}
```


3. response definition



```golang
type AddFlashPromotionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 106. N/A

1. route definition

- Url: /api/sms/flashpromotion/list
- Method: POST
- Request: `ListFlashPromotionReq`
- Response: `ListFlashPromotionResp`

2. request definition



```golang
type ListFlashPromotionReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListFlashPromotionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtFlashPromotionData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 107. N/A

1. route definition

- Url: /api/sms/flashpromotion/update
- Method: POST
- Request: `UpdateFlashPromotionReq`
- Response: `UpdateFlashPromotionResp`

2. request definition



```golang
type UpdateFlashPromotionReq struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	StartDate string `json:"startDate"` // 开始日期
	EndDate string `json:"endDate"` // 结束日期
	Status int64 `json:"status"` // 上下线状态
	CreateTime string `json:"createTime"` // 秒杀时间段名称
}
```


3. response definition



```golang
type UpdateFlashPromotionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 108. N/A

1. route definition

- Url: /api/sms/flashpromotion/delete
- Method: POST
- Request: `DeleteFlashPromotionReq`
- Response: `DeleteFlashPromotionResp`

2. request definition



```golang
type DeleteFlashPromotionReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteFlashPromotionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 109. N/A

1. route definition

- Url: /api/sms/flashpromotionlog/add
- Method: POST
- Request: `addFlashPromotionLogReq`
- Response: `addFlashPromotionLogResp`

2. request definition



```golang
type AddFlashPromotionLogReq struct {
	MemberId int64 `json:"memberId"`
	ProductId int64 `json:"productId"`
	MemberPhone string `json:"memberPhone"`
	ProductName string `json:"productName"`
	SubscribeTime string `json:"subscribeTime"` // 会员订阅时间
	SendTime string `json:"sendTime"`
}
```


3. response definition



```golang
type AddFlashPromotionLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 110. N/A

1. route definition

- Url: /api/sms/flashpromotionlog/list
- Method: POST
- Request: `ListFlashPromotionLogReq`
- Response: `ListFlashPromotionLogResp`

2. request definition



```golang
type ListFlashPromotionLogReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListFlashPromotionLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtFlashPromotionLogData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 111. N/A

1. route definition

- Url: /api/sms/flashpromotionlog/update
- Method: POST
- Request: `UpdateFlashPromotionLogReq`
- Response: `UpdateFlashPromotionLogResp`

2. request definition



```golang
type UpdateFlashPromotionLogReq struct {
	Id int64 `json:"id"`
	MemberId int64 `json:"memberId"`
	ProductId int64 `json:"productId"`
	MemberPhone string `json:"memberPhone"`
	ProductName string `json:"productName"`
	SubscribeTime string `json:"subscribeTime"` // 会员订阅时间
	SendTime string `json:"sendTime"`
}
```


3. response definition



```golang
type UpdateFlashPromotionLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 112. N/A

1. route definition

- Url: /api/sms/flashpromotionlog/delete
- Method: POST
- Request: `DeleteFlashPromotionLogReq`
- Response: `DeleteFlashPromotionLogResp`

2. request definition



```golang
type DeleteFlashPromotionLogReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteFlashPromotionLogResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 113. N/A

1. route definition

- Url: /api/sms/flashpromotionsession/add
- Method: POST
- Request: `addFlashPromotionSessionReq`
- Response: `addFlashPromotionSessionResp`

2. request definition



```golang
type AddFlashPromotionSessionReq struct {
	Name string `json:"name"` // 场次名称
	StartTime string `json:"startTime"` // 每日开始时间
	EndTime string `json:"endTime"` // 每日结束时间
	Status int64 `json:"status"` // 启用状态：0-&gt;不启用；1-&gt;启用
	CreateTime string `json:"createTime"` // 创建时间
}
```


3. response definition



```golang
type AddFlashPromotionSessionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 114. N/A

1. route definition

- Url: /api/sms/flashpromotionsession/list
- Method: POST
- Request: `ListFlashPromotionSessionReq`
- Response: `ListFlashPromotionSessionResp`

2. request definition



```golang
type ListFlashPromotionSessionReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListFlashPromotionSessionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtFlashPromotionSessionData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 115. N/A

1. route definition

- Url: /api/sms/flashpromotionsession/update
- Method: POST
- Request: `UpdateFlashPromotionSessionReq`
- Response: `UpdateFlashPromotionSessionResp`

2. request definition



```golang
type UpdateFlashPromotionSessionReq struct {
	Id int64 `json:"id"` // 编号
	Name string `json:"name"` // 场次名称
	StartTime string `json:"startTime"` // 每日开始时间
	EndTime string `json:"endTime"` // 每日结束时间
	Status int64 `json:"status"` // 启用状态：0-&gt;不启用；1-&gt;启用
	CreateTime string `json:"createTime"` // 创建时间
}
```


3. response definition



```golang
type UpdateFlashPromotionSessionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 116. N/A

1. route definition

- Url: /api/sms/flashpromotionsession/delete
- Method: POST
- Request: `DeleteFlashPromotionSessionReq`
- Response: `DeleteFlashPromotionSessionResp`

2. request definition



```golang
type DeleteFlashPromotionSessionReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteFlashPromotionSessionResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 117. N/A

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

### 118. N/A

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

### 119. N/A

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

### 120. N/A

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

### 121. N/A

1. route definition

- Url: /api/sms/homebrand/add
- Method: POST
- Request: `addHomeBrandReq`
- Response: `addHomeBrandResp`

2. request definition



```golang
type AddHomeBrandReq struct {
	BrandId int64 `json:"brandId"`
	BrandName string `json:"brandName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type AddHomeBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 122. N/A

1. route definition

- Url: /api/sms/homebrand/list
- Method: POST
- Request: `ListHomeBrandReq`
- Response: `ListHomeBrandResp`

2. request definition



```golang
type ListHomeBrandReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListHomeBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtHomeBrandData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 123. N/A

1. route definition

- Url: /api/sms/homebrand/update
- Method: POST
- Request: `UpdateHomeBrandReq`
- Response: `UpdateHomeBrandResp`

2. request definition



```golang
type UpdateHomeBrandReq struct {
	Id int64 `json:"id"`
	BrandId int64 `json:"brandId"`
	BrandName string `json:"brandName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type UpdateHomeBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 124. N/A

1. route definition

- Url: /api/sms/homebrand/delete
- Method: POST
- Request: `DeleteHomeBrandReq`
- Response: `DeleteHomeBrandResp`

2. request definition



```golang
type DeleteHomeBrandReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteHomeBrandResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 125. N/A

1. route definition

- Url: /api/sms/homenewproduct/add
- Method: POST
- Request: `addHomeNewProductReq`
- Response: `addHomeNewProductResp`

2. request definition



```golang
type AddHomeNewProductReq struct {
	ProductId int64 `json:"productId"`
	ProductName string `json:"productName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type AddHomeNewProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 126. N/A

1. route definition

- Url: /api/sms/homenewproduct/list
- Method: POST
- Request: `ListHomeNewProductReq`
- Response: `ListHomeNewProductResp`

2. request definition



```golang
type ListHomeNewProductReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListHomeNewProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtHomeNewProductData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 127. N/A

1. route definition

- Url: /api/sms/homenewproduct/update
- Method: POST
- Request: `UpdateHomeNewProductReq`
- Response: `UpdateHomeNewProductResp`

2. request definition



```golang
type UpdateHomeNewProductReq struct {
	Id int64 `json:"id"`
	ProductId int64 `json:"productId"`
	ProductName string `json:"productName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type UpdateHomeNewProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 128. N/A

1. route definition

- Url: /api/sms/homenewproduct/delete
- Method: POST
- Request: `DeleteHomeNewProductReq`
- Response: `DeleteHomeNewProductResp`

2. request definition



```golang
type DeleteHomeNewProductReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteHomeNewProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 129. N/A

1. route definition

- Url: /api/sms/homerecommendproduct/add
- Method: POST
- Request: `addHomeRecommendProductReq`
- Response: `addHomeRecommendProductResp`

2. request definition



```golang
type AddHomeRecommendProductReq struct {
	ProductId int64 `json:"productId"`
	ProductName string `json:"productName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type AddHomeRecommendProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 130. N/A

1. route definition

- Url: /api/sms/homerecommendproduct/list
- Method: POST
- Request: `ListHomeRecommendProductReq`
- Response: `ListHomeRecommendProductResp`

2. request definition



```golang
type ListHomeRecommendProductReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListHomeRecommendProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtHomeRecommendProductData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 131. N/A

1. route definition

- Url: /api/sms/homerecommendproduct/update
- Method: POST
- Request: `UpdateHomeRecommendProductReq`
- Response: `UpdateHomeRecommendProductResp`

2. request definition



```golang
type UpdateHomeRecommendProductReq struct {
	Id int64 `json:"id"`
	ProductId int64 `json:"productId"`
	ProductName string `json:"productName"`
	RecommendStatus int64 `json:"recommendStatus"`
	Sort int64 `json:"sort"`
}
```


3. response definition



```golang
type UpdateHomeRecommendProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 132. N/A

1. route definition

- Url: /api/sms/homerecommendproduct/delete
- Method: POST
- Request: `DeleteHomeRecommendProductReq`
- Response: `DeleteHomeRecommendProductResp`

2. request definition



```golang
type DeleteHomeRecommendProductReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteHomeRecommendProductResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 133. N/A

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

### 134. N/A

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

### 135. N/A

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

### 136. N/A

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

### 137. N/A

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

### 138. N/A

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

### 139. N/A

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

### 140. N/A

1. route definition

- Url: /api/member/address/add
- Method: POST
- Request: `addMemberAddressReq`
- Response: `addMemberAddressResp`

2. request definition



```golang
type AddMemberAddressReq struct {
	MemberId int64 `json:"memberId"`
	Name string `json:"name"` // 收货人名称
	PhoneNumber string `json:"phoneNumber"`
	DefaultStatus int64 `json:"defaultStatus"` // 是否为默认
	PostCode string `json:"postCode"` // 邮政编码
	Province string `json:"province"` // 省份/直辖市
	City string `json:"city"` // 城市
	Region string `json:"region"` // 区
	DetailAddress string `json:"detailAddress"` // 详细地址(街道)
}
```


3. response definition



```golang
type AddMemberAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 141. N/A

1. route definition

- Url: /api/member/address/list
- Method: POST
- Request: `ListMemberAddressReq`
- Response: `ListMemberAddressResp`

2. request definition



```golang
type ListMemberAddressReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberAddressResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberAddressData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 142. N/A

1. route definition

- Url: /api/member/address/update
- Method: POST
- Request: `UpdateMemberAddressReq`
- Response: `UpdateMemberAddressResp`

2. request definition



```golang
type UpdateMemberAddressReq struct {
	Id int64 `json:"id"`
	MemberId int64 `json:"memberId"`
	Name string `json:"name"` // 收货人名称
	PhoneNumber string `json:"phoneNumber"`
	DefaultStatus int64 `json:"defaultStatus"` // 是否为默认
	PostCode string `json:"postCode"` // 邮政编码
	Province string `json:"province"` // 省份/直辖市
	City string `json:"city"` // 城市
	Region string `json:"region"` // 区
	DetailAddress string `json:"detailAddress"` // 详细地址(街道)
}
```


3. response definition



```golang
type UpdateMemberAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 143. N/A

1. route definition

- Url: /api/member/address/delete
- Method: POST
- Request: `DeleteMemberAddressReq`
- Response: `DeleteMemberAddressResp`

2. request definition



```golang
type DeleteMemberAddressReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberAddressResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 144. N/A

1. route definition

- Url: /api/member/growthchangehistory/add
- Method: POST
- Request: `addGrowthChangeHistoryReq`
- Response: `addGrowthChangeHistoryResp`

2. request definition



```golang
type AddGrowthChangeHistoryReq struct {
	MemberId int64 `json:"memberId"`
	CreateTime string `json:"createTime"`
	ChangeType int64 `json:"changeType"` // 改变类型：0-&gt;增加；1-&gt;减少
	ChangeCount int64 `json:"changeCount"` // 积分改变数量
	OperateMan string `json:"operateMan"` // 操作人员
	OperateNote string `json:"operateNote"` // 操作备注
	SourceType int64 `json:"sourceType"` // 积分来源：0-&gt;购物；1-&gt;管理员修改
}
```


3. response definition



```golang
type AddGrowthChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 145. N/A

1. route definition

- Url: /api/member/growthchangehistory/list
- Method: POST
- Request: `ListGrowthChangeHistoryReq`
- Response: `ListGrowthChangeHistoryResp`

2. request definition



```golang
type ListGrowthChangeHistoryReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListGrowthChangeHistoryResp struct {
	Current int64 `json:"current,default=1"`
	Data []*ListtGrowthChangeHistoryData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 146. N/A

1. route definition

- Url: /api/member/growthchangehistory/update
- Method: POST
- Request: `UpdateGrowthChangeHistoryReq`
- Response: `UpdateGrowthChangeHistoryResp`

2. request definition



```golang
type UpdateGrowthChangeHistoryReq struct {
	Id int64 `json:"id"`
	MemberId int64 `json:"memberId"`
	CreateTime string `json:"createTime"`
	ChangeType int64 `json:"changeType"` // 改变类型：0-&gt;增加；1-&gt;减少
	ChangeCount int64 `json:"changeCount"` // 积分改变数量
	OperateMan string `json:"operateMan"` // 操作人员
	OperateNote string `json:"operateNote"` // 操作备注
	SourceType int64 `json:"sourceType"` // 积分来源：0-&gt;购物；1-&gt;管理员修改
}
```


3. response definition



```golang
type UpdateGrowthChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 147. N/A

1. route definition

- Url: /api/member/growthchangehistory/delete
- Method: POST
- Request: `DeleteGrowthChangeHistoryReq`
- Response: `DeleteGrowthChangeHistoryResp`

2. request definition



```golang
type DeleteGrowthChangeHistoryReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteGrowthChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 148. N/A

1. route definition

- Url: /api/member/integrationchangehistory/add
- Method: POST
- Request: `addIntegrationChangeHistoryReq`
- Response: `addIntegrationChangeHistoryResp`

2. request definition



```golang
type AddIntegrationChangeHistoryReq struct {
	MemberId int64 `json:"memberId"`
	CreateTime string `json:"createTime"`
	ChangeType int64 `json:"changeType"` // 改变类型：0-&gt;增加；1-&gt;减少
	ChangeCount int64 `json:"changeCount"` // 积分改变数量
	OperateMan string `json:"operateMan"` // 操作人员
	OperateNote string `json:"operateNote"` // 操作备注
	SourceType int64 `json:"sourceType"` // 积分来源：0-&gt;购物；1-&gt;管理员修改
}
```


3. response definition



```golang
type AddIntegrationChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 149. N/A

1. route definition

- Url: /api/member/integrationchangehistory/list
- Method: POST
- Request: `ListIntegrationChangeHistoryReq`
- Response: `ListIntegrationChangeHistoryResp`

2. request definition



```golang
type ListIntegrationChangeHistoryReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListIntegrationChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtIntegrationChangeHistoryData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 150. N/A

1. route definition

- Url: /api/member/integrationchangehistory/update
- Method: POST
- Request: `UpdateIntegrationChangeHistoryReq`
- Response: `UpdateIntegrationChangeHistoryResp`

2. request definition



```golang
type UpdateIntegrationChangeHistoryReq struct {
	Id int64 `json:"id"`
	MemberId int64 `json:"memberId"`
	CreateTime string `json:"createTime"`
	ChangeType int64 `json:"changeType"` // 改变类型：0-&gt;增加；1-&gt;减少
	ChangeCount int64 `json:"changeCount"` // 积分改变数量
	OperateMan string `json:"operateMan"` // 操作人员
	OperateNote string `json:"operateNote"` // 操作备注
	SourceType int64 `json:"sourceType"` // 积分来源：0-&gt;购物；1-&gt;管理员修改
}
```


3. response definition



```golang
type UpdateIntegrationChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 151. N/A

1. route definition

- Url: /api/member/integrationchangehistory/delete
- Method: POST
- Request: `DeleteIntegrationChangeHistoryReq`
- Response: `DeleteIntegrationChangeHistoryResp`

2. request definition



```golang
type DeleteIntegrationChangeHistoryReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteIntegrationChangeHistoryResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 152. N/A

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

### 153. N/A

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

### 154. N/A

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

### 155. N/A

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

### 156. N/A

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

### 157. N/A

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

### 158. N/A

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

### 159. N/A

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

### 160. N/A

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

### 161. N/A

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

### 162. N/A

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

### 163. N/A

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

### 164. N/A

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

### 165. N/A

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

### 166. N/A

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

### 167. N/A

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

### 168. N/A

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

### 169. N/A

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

### 170. N/A

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

### 171. N/A

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

### 172. N/A

1. route definition

- Url: /api/member/tag/add
- Method: POST
- Request: `addMemberTagReq`
- Response: `addMemberTagResp`

2. request definition



```golang
type AddMemberTagReq struct {
	Name string `json:"name"`
	FinishOrderCount int64 `json:"finishOrderCount"` // 自动打标签完成订单数量
	FinishOrderAmount float64 `json:"finishOrderAmount"` // 自动打标签完成订单金额
}
```


3. response definition



```golang
type AddMemberTagResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 173. N/A

1. route definition

- Url: /api/member/tag/list
- Method: POST
- Request: `ListMemberTagReq`
- Response: `ListMemberTagResp`

2. request definition



```golang
type ListMemberTagReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberTagResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberTagData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 174. N/A

1. route definition

- Url: /api/member/tag/update
- Method: POST
- Request: `UpdateMemberTagReq`
- Response: `UpdateMemberTagResp`

2. request definition



```golang
type UpdateMemberTagReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	FinishOrderCount int64 `json:"finishOrderCount"` // 自动打标签完成订单数量
	FinishOrderAmount float64 `json:"finishOrderAmount"` // 自动打标签完成订单金额
}
```


3. response definition



```golang
type UpdateMemberTagResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 175. N/A

1. route definition

- Url: /api/member/tag/delete
- Method: POST
- Request: `DeleteMemberTagReq`
- Response: `DeleteMemberTagResp`

2. request definition



```golang
type DeleteMemberTagReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberTagResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 176. N/A

1. route definition

- Url: /api/member/task/add
- Method: POST
- Request: `addMemberTaskReq`
- Response: `addMemberTaskResp`

2. request definition



```golang
type AddMemberTaskReq struct {
	Name string `json:"name"`
	Growth int64 `json:"growth"` // 赠送成长值
	Intergration int64 `json:"intergration"` // 赠送积分
	Type int64 `json:"type"` // 任务类型：0-&gt;新手任务；1-&gt;日常任务
}
```


3. response definition



```golang
type AddMemberTaskResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 177. N/A

1. route definition

- Url: /api/member/task/list
- Method: POST
- Request: `ListMemberTaskReq`
- Response: `ListMemberTaskResp`

2. request definition



```golang
type ListMemberTaskReq struct {
	Current int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}
```


3. response definition



```golang
type ListMemberTaskResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Current int64 `json:"current,default=1"`
	Data []*ListtMemberTaskData `json:"data"`
	PageSize int64 `json:"pageSize,default=20"`
	Success bool `json:"success"`
	Total int64 `json:"total"`
}
```

### 178. N/A

1. route definition

- Url: /api/member/task/update
- Method: POST
- Request: `UpdateMemberTaskReq`
- Response: `UpdateMemberTaskResp`

2. request definition



```golang
type UpdateMemberTaskReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Growth int64 `json:"growth"` // 赠送成长值
	Intergration int64 `json:"intergration"` // 赠送积分
	Type int64 `json:"type"` // 任务类型：0-&gt;新手任务；1-&gt;日常任务
}
```


3. response definition



```golang
type UpdateMemberTaskResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

### 179. N/A

1. route definition

- Url: /api/member/task/delete
- Method: POST
- Request: `DeleteMemberTaskReq`
- Response: `DeleteMemberTaskResp`

2. request definition



```golang
type DeleteMemberTaskReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type DeleteMemberTaskResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
```

