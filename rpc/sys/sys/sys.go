// Code generated by goctl. DO NOT EDIT!
// Source: sys.proto

package sys

import (
	"context"

	"zero-admin-learn/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ConfigAddReq          = sysclient.ConfigAddReq
	ConfigAddResp         = sysclient.ConfigAddResp
	ConfigDeleteReq       = sysclient.ConfigDeleteReq
	ConfigDeleteResp      = sysclient.ConfigDeleteResp
	ConfigListData        = sysclient.ConfigListData
	ConfigListReq         = sysclient.ConfigListReq
	ConfigListResp        = sysclient.ConfigListResp
	ConfigUpdateReq       = sysclient.ConfigUpdateReq
	ConfigUpdateResp      = sysclient.ConfigUpdateResp
	DeptAddReq            = sysclient.DeptAddReq
	DeptAddResp           = sysclient.DeptAddResp
	DeptDeleteReq         = sysclient.DeptDeleteReq
	DeptDeleteResp        = sysclient.DeptDeleteResp
	DeptListData          = sysclient.DeptListData
	DeptListReq           = sysclient.DeptListReq
	DeptListResp          = sysclient.DeptListResp
	DeptUpdateReq         = sysclient.DeptUpdateReq
	DeptUpdateResp        = sysclient.DeptUpdateResp
	DictAddReq            = sysclient.DictAddReq
	DictAddResp           = sysclient.DictAddResp
	DictDeleteReq         = sysclient.DictDeleteReq
	DictDeleteResp        = sysclient.DictDeleteResp
	DictListData          = sysclient.DictListData
	DictListReq           = sysclient.DictListReq
	DictListResp          = sysclient.DictListResp
	DictUpdateReq         = sysclient.DictUpdateReq
	DictUpdateResp        = sysclient.DictUpdateResp
	InfoReq               = sysclient.InfoReq
	InfoResp              = sysclient.InfoResp
	JobAddReq             = sysclient.JobAddReq
	JobAddResp            = sysclient.JobAddResp
	JobDeleteReq          = sysclient.JobDeleteReq
	JobDeleteResp         = sysclient.JobDeleteResp
	JobListData           = sysclient.JobListData
	JobListReq            = sysclient.JobListReq
	JobListResp           = sysclient.JobListResp
	JobUpdateReq          = sysclient.JobUpdateReq
	JobUpdateResp         = sysclient.JobUpdateResp
	LoginLogAddReq        = sysclient.LoginLogAddReq
	LoginLogAddResp       = sysclient.LoginLogAddResp
	LoginLogDeleteReq     = sysclient.LoginLogDeleteReq
	LoginLogDeleteResp    = sysclient.LoginLogDeleteResp
	LoginLogListData      = sysclient.LoginLogListData
	LoginLogListReq       = sysclient.LoginLogListReq
	LoginLogListResp      = sysclient.LoginLogListResp
	LoginReq              = sysclient.LoginReq
	LoginResp             = sysclient.LoginResp
	MenuAddReq            = sysclient.MenuAddReq
	MenuAddResp           = sysclient.MenuAddResp
	MenuDeleteReq         = sysclient.MenuDeleteReq
	MenuDeleteResp        = sysclient.MenuDeleteResp
	MenuListData          = sysclient.MenuListData
	MenuListReq           = sysclient.MenuListReq
	MenuListResp          = sysclient.MenuListResp
	MenuListTree          = sysclient.MenuListTree
	MenuUpdateReq         = sysclient.MenuUpdateReq
	MenuUpdateResp        = sysclient.MenuUpdateResp
	QueryMenuByRoleIdReq  = sysclient.QueryMenuByRoleIdReq
	QueryMenuByRoleIdResp = sysclient.QueryMenuByRoleIdResp
	ReSetPasswordReq      = sysclient.ReSetPasswordReq
	ReSetPasswordResp     = sysclient.ReSetPasswordResp
	RoleAddReq            = sysclient.RoleAddReq
	RoleAddResp           = sysclient.RoleAddResp
	RoleDeleteReq         = sysclient.RoleDeleteReq
	RoleDeleteResp        = sysclient.RoleDeleteResp
	RoleListData          = sysclient.RoleListData
	RoleListReq           = sysclient.RoleListReq
	RoleListResp          = sysclient.RoleListResp
	RoleUpdateReq         = sysclient.RoleUpdateReq
	RoleUpdateResp        = sysclient.RoleUpdateResp
	SysLogAddReq          = sysclient.SysLogAddReq
	SysLogAddResp         = sysclient.SysLogAddResp
	SysLogDeleteReq       = sysclient.SysLogDeleteReq
	SysLogDeleteResp      = sysclient.SysLogDeleteResp
	SysLogListData        = sysclient.SysLogListData
	SysLogListReq         = sysclient.SysLogListReq
	SysLogListResp        = sysclient.SysLogListResp
	UpdateConfigRoleReq   = sysclient.UpdateConfigRoleReq
	UpdateConfigRoleResp  = sysclient.UpdateConfigRoleResp
	UpdateMenuRoleReq     = sysclient.UpdateMenuRoleReq
	UpdateMenuRoleResp    = sysclient.UpdateMenuRoleResp
	UpdateRoleRoleReq     = sysclient.UpdateRoleRoleReq
	UpdateRoleRoleResp    = sysclient.UpdateRoleRoleResp
	UserAddReq            = sysclient.UserAddReq
	UserAddResp           = sysclient.UserAddResp
	UserDeleteReq         = sysclient.UserDeleteReq
	UserDeleteResp        = sysclient.UserDeleteResp
	UserListData          = sysclient.UserListData
	UserListReq           = sysclient.UserListReq
	UserListResp          = sysclient.UserListResp
	UserStatusReq         = sysclient.UserStatusReq
	UserStatusResp        = sysclient.UserStatusResp
	UserUpdateReq         = sysclient.UserUpdateReq
	UserUpdateResp        = sysclient.UserUpdateResp

	Sys interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
		UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error)
		UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error)
		UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error)
		ReSetPassword(ctx context.Context, in *ReSetPasswordReq, opts ...grpc.CallOption) (*ReSetPasswordResp, error)
		UpdateUserStatus(ctx context.Context, in *UserStatusReq, opts ...grpc.CallOption) (*UserStatusResp, error)
		RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error)
		RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
		RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error)
		RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error)
		UpdateRoleRole(ctx context.Context, in *UpdateRoleRoleReq, opts ...grpc.CallOption) (*UpdateRoleRoleResp, error)
		QueryMenuByRoleId(ctx context.Context, in *QueryMenuByRoleIdReq, opts ...grpc.CallOption) (*QueryMenuByRoleIdResp, error)
		UpdateMenuRole(ctx context.Context, in *UpdateMenuRoleReq, opts ...grpc.CallOption) (*UpdateMenuRoleResp, error)
		MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error)
		MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
		MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error)
		MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error)
		DictAdd(ctx context.Context, in *DictAddReq, opts ...grpc.CallOption) (*DictAddResp, error)
		DictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error)
		DictUpdate(ctx context.Context, in *DictUpdateReq, opts ...grpc.CallOption) (*DictUpdateResp, error)
		DictDelete(ctx context.Context, in *DictDeleteReq, opts ...grpc.CallOption) (*DictDeleteResp, error)
		DeptAdd(ctx context.Context, in *DeptAddReq, opts ...grpc.CallOption) (*DeptAddResp, error)
		DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error)
		DeptUpdate(ctx context.Context, in *DeptUpdateReq, opts ...grpc.CallOption) (*DeptUpdateResp, error)
		DeptDelete(ctx context.Context, in *DeptDeleteReq, opts ...grpc.CallOption) (*DeptDeleteResp, error)
		LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error)
		LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error)
		LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error)
		SysLogAdd(ctx context.Context, in *SysLogAddReq, opts ...grpc.CallOption) (*SysLogAddResp, error)
		SysLogList(ctx context.Context, in *SysLogListReq, opts ...grpc.CallOption) (*SysLogListResp, error)
		SysLogDelete(ctx context.Context, in *SysLogDeleteReq, opts ...grpc.CallOption) (*SysLogDeleteResp, error)
		ConfigAdd(ctx context.Context, in *ConfigAddReq, opts ...grpc.CallOption) (*ConfigAddResp, error)
		ConfigList(ctx context.Context, in *ConfigListReq, opts ...grpc.CallOption) (*ConfigListResp, error)
		ConfigUpdate(ctx context.Context, in *ConfigUpdateReq, opts ...grpc.CallOption) (*ConfigUpdateResp, error)
		ConfigDelete(ctx context.Context, in *ConfigDeleteReq, opts ...grpc.CallOption) (*ConfigDeleteResp, error)
		UpdateConfigRole(ctx context.Context, in *UpdateConfigRoleReq, opts ...grpc.CallOption) (*UpdateConfigRoleResp, error)
		JobAdd(ctx context.Context, in *JobAddReq, opts ...grpc.CallOption) (*JobAddResp, error)
		JobList(ctx context.Context, in *JobListReq, opts ...grpc.CallOption) (*JobListResp, error)
		JobUpdate(ctx context.Context, in *JobUpdateReq, opts ...grpc.CallOption) (*JobUpdateResp, error)
		JobDelete(ctx context.Context, in *JobDeleteReq, opts ...grpc.CallOption) (*JobDeleteResp, error)
	}

	defaultSys struct {
		cli zrpc.Client
	}
)

func NewSys(cli zrpc.Client) Sys {
	return &defaultSys{
		cli: cli,
	}
}

func (m *defaultSys) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultSys) UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultSys) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UserAdd(ctx, in, opts...)
}

func (m *defaultSys) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}

func (m *defaultSys) UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UserUpdate(ctx, in, opts...)
}

func (m *defaultSys) UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UserDelete(ctx, in, opts...)
}

func (m *defaultSys) ReSetPassword(ctx context.Context, in *ReSetPasswordReq, opts ...grpc.CallOption) (*ReSetPasswordResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.ReSetPassword(ctx, in, opts...)
}

func (m *defaultSys) UpdateUserStatus(ctx context.Context, in *UserStatusReq, opts ...grpc.CallOption) (*UserStatusResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UpdateUserStatus(ctx, in, opts...)
}

func (m *defaultSys) RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.RoleAdd(ctx, in, opts...)
}

func (m *defaultSys) RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.RoleList(ctx, in, opts...)
}

func (m *defaultSys) RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in, opts...)
}

func (m *defaultSys) RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.RoleDelete(ctx, in, opts...)
}

func (m *defaultSys) UpdateRoleRole(ctx context.Context, in *UpdateRoleRoleReq, opts ...grpc.CallOption) (*UpdateRoleRoleResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UpdateRoleRole(ctx, in, opts...)
}

func (m *defaultSys) QueryMenuByRoleId(ctx context.Context, in *QueryMenuByRoleIdReq, opts ...grpc.CallOption) (*QueryMenuByRoleIdResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.QueryMenuByRoleId(ctx, in, opts...)
}

func (m *defaultSys) UpdateMenuRole(ctx context.Context, in *UpdateMenuRoleReq, opts ...grpc.CallOption) (*UpdateMenuRoleResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UpdateMenuRole(ctx, in, opts...)
}

func (m *defaultSys) MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.MenuAdd(ctx, in, opts...)
}

func (m *defaultSys) MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.MenuList(ctx, in, opts...)
}

func (m *defaultSys) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.MenuUpdate(ctx, in, opts...)
}

func (m *defaultSys) MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.MenuDelete(ctx, in, opts...)
}

func (m *defaultSys) DictAdd(ctx context.Context, in *DictAddReq, opts ...grpc.CallOption) (*DictAddResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.DictAdd(ctx, in, opts...)
}

func (m *defaultSys) DictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.DictList(ctx, in, opts...)
}

func (m *defaultSys) DictUpdate(ctx context.Context, in *DictUpdateReq, opts ...grpc.CallOption) (*DictUpdateResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.DictUpdate(ctx, in, opts...)
}

func (m *defaultSys) DictDelete(ctx context.Context, in *DictDeleteReq, opts ...grpc.CallOption) (*DictDeleteResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.DictDelete(ctx, in, opts...)
}

func (m *defaultSys) DeptAdd(ctx context.Context, in *DeptAddReq, opts ...grpc.CallOption) (*DeptAddResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.DeptAdd(ctx, in, opts...)
}

func (m *defaultSys) DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.DeptList(ctx, in, opts...)
}

func (m *defaultSys) DeptUpdate(ctx context.Context, in *DeptUpdateReq, opts ...grpc.CallOption) (*DeptUpdateResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.DeptUpdate(ctx, in, opts...)
}

func (m *defaultSys) DeptDelete(ctx context.Context, in *DeptDeleteReq, opts ...grpc.CallOption) (*DeptDeleteResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.DeptDelete(ctx, in, opts...)
}

func (m *defaultSys) LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.LoginLogAdd(ctx, in, opts...)
}

func (m *defaultSys) LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.LoginLogList(ctx, in, opts...)
}

func (m *defaultSys) LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.LoginLogDelete(ctx, in, opts...)
}

func (m *defaultSys) SysLogAdd(ctx context.Context, in *SysLogAddReq, opts ...grpc.CallOption) (*SysLogAddResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.SysLogAdd(ctx, in, opts...)
}

func (m *defaultSys) SysLogList(ctx context.Context, in *SysLogListReq, opts ...grpc.CallOption) (*SysLogListResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.SysLogList(ctx, in, opts...)
}

func (m *defaultSys) SysLogDelete(ctx context.Context, in *SysLogDeleteReq, opts ...grpc.CallOption) (*SysLogDeleteResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.SysLogDelete(ctx, in, opts...)
}

func (m *defaultSys) ConfigAdd(ctx context.Context, in *ConfigAddReq, opts ...grpc.CallOption) (*ConfigAddResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.ConfigAdd(ctx, in, opts...)
}

func (m *defaultSys) ConfigList(ctx context.Context, in *ConfigListReq, opts ...grpc.CallOption) (*ConfigListResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.ConfigList(ctx, in, opts...)
}

func (m *defaultSys) ConfigUpdate(ctx context.Context, in *ConfigUpdateReq, opts ...grpc.CallOption) (*ConfigUpdateResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.ConfigUpdate(ctx, in, opts...)
}

func (m *defaultSys) ConfigDelete(ctx context.Context, in *ConfigDeleteReq, opts ...grpc.CallOption) (*ConfigDeleteResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.ConfigDelete(ctx, in, opts...)
}

func (m *defaultSys) UpdateConfigRole(ctx context.Context, in *UpdateConfigRoleReq, opts ...grpc.CallOption) (*UpdateConfigRoleResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UpdateConfigRole(ctx, in, opts...)
}

func (m *defaultSys) JobAdd(ctx context.Context, in *JobAddReq, opts ...grpc.CallOption) (*JobAddResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.JobAdd(ctx, in, opts...)
}

func (m *defaultSys) JobList(ctx context.Context, in *JobListReq, opts ...grpc.CallOption) (*JobListResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.JobList(ctx, in, opts...)
}

func (m *defaultSys) JobUpdate(ctx context.Context, in *JobUpdateReq, opts ...grpc.CallOption) (*JobUpdateResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.JobUpdate(ctx, in, opts...)
}

func (m *defaultSys) JobDelete(ctx context.Context, in *JobDeleteReq, opts ...grpc.CallOption) (*JobDeleteResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.JobDelete(ctx, in, opts...)
}
