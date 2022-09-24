### 1. N/A

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

### 2. N/A

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

### 3. N/A

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

### 4. N/A

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

