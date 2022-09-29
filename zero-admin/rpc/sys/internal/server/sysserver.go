// Code generated by goctl. DO NOT EDIT!
// Source: sys.proto

package server

import (
	"context"

	"zero-admin-learn/rpc/sys/internal/logic"
	"zero-admin-learn/rpc/sys/internal/svc"
	"zero-admin-learn/rpc/sys/sysclient"
)

type SysServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedSysServer
}

func NewSysServer(svcCtx *svc.ServiceContext) *SysServer {
	return &SysServer{
		svcCtx: svcCtx,
	}
}

func (s *SysServer) Login(ctx context.Context, in *sysclient.LoginReq) (*sysclient.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *SysServer) UserInfo(ctx context.Context, in *sysclient.InfoReq) (*sysclient.InfoResp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *SysServer) UserAdd(ctx context.Context, in *sysclient.UserAddReq) (*sysclient.UserAddResp, error) {
	l := logic.NewUserAddLogic(ctx, s.svcCtx)
	return l.UserAdd(in)
}

func (s *SysServer) UserList(ctx context.Context, in *sysclient.UserListReq) (*sysclient.UserListResp, error) {
	l := logic.NewUserListLogic(ctx, s.svcCtx)
	return l.UserList(in)
}

func (s *SysServer) UserUpdate(ctx context.Context, in *sysclient.UserUpdateReq) (*sysclient.UserUpdateResp, error) {
	l := logic.NewUserUpdateLogic(ctx, s.svcCtx)
	return l.UserUpdate(in)
}

func (s *SysServer) UserDelete(ctx context.Context, in *sysclient.UserDeleteReq) (*sysclient.UserDeleteResp, error) {
	l := logic.NewUserDeleteLogic(ctx, s.svcCtx)
	return l.UserDelete(in)
}

func (s *SysServer) ReSetPassword(ctx context.Context, in *sysclient.ReSetPasswordReq) (*sysclient.ReSetPasswordResp, error) {
	l := logic.NewReSetPasswordLogic(ctx, s.svcCtx)
	return l.ReSetPassword(in)
}

func (s *SysServer) UpdateUserStatus(ctx context.Context, in *sysclient.UserStatusReq) (*sysclient.UserStatusResp, error) {
	l := logic.NewUpdateUserStatusLogic(ctx, s.svcCtx)
	return l.UpdateUserStatus(in)
}

func (s *SysServer) RoleAdd(ctx context.Context, in *sysclient.RoleAddReq) (*sysclient.RoleAddResp, error) {
	l := logic.NewRoleAddLogic(ctx, s.svcCtx)
	return l.RoleAdd(in)
}

func (s *SysServer) RoleList(ctx context.Context, in *sysclient.RoleListReq) (*sysclient.RoleListResp, error) {
	l := logic.NewRoleListLogic(ctx, s.svcCtx)
	return l.RoleList(in)
}

func (s *SysServer) RoleUpdate(ctx context.Context, in *sysclient.RoleUpdateReq) (*sysclient.RoleUpdateResp, error) {
	l := logic.NewRoleUpdateLogic(ctx, s.svcCtx)
	return l.RoleUpdate(in)
}

func (s *SysServer) RoleDelete(ctx context.Context, in *sysclient.RoleDeleteReq) (*sysclient.RoleDeleteResp, error) {
	l := logic.NewRoleDeleteLogic(ctx, s.svcCtx)
	return l.RoleDelete(in)
}

func (s *SysServer) UpdateRoleRole(ctx context.Context, in *sysclient.UpdateRoleRoleReq) (*sysclient.UpdateRoleRoleResp, error) {
	l := logic.NewUpdateRoleRoleLogic(ctx, s.svcCtx)
	return l.UpdateRoleRole(in)
}

func (s *SysServer) QueryMenuByRoleId(ctx context.Context, in *sysclient.QueryMenuByRoleIdReq) (*sysclient.QueryMenuByRoleIdResp, error) {
	l := logic.NewQueryMenuByRoleIdLogic(ctx, s.svcCtx)
	return l.QueryMenuByRoleId(in)
}

func (s *SysServer) UpdateMenuRole(ctx context.Context, in *sysclient.UpdateMenuRoleReq) (*sysclient.UpdateMenuRoleResp, error) {
	l := logic.NewUpdateMenuRoleLogic(ctx, s.svcCtx)
	return l.UpdateMenuRole(in)
}

func (s *SysServer) MenuAdd(ctx context.Context, in *sysclient.MenuAddReq) (*sysclient.MenuAddResp, error) {
	l := logic.NewMenuAddLogic(ctx, s.svcCtx)
	return l.MenuAdd(in)
}

func (s *SysServer) MenuList(ctx context.Context, in *sysclient.MenuListReq) (*sysclient.MenuListResp, error) {
	l := logic.NewMenuListLogic(ctx, s.svcCtx)
	return l.MenuList(in)
}

func (s *SysServer) MenuUpdate(ctx context.Context, in *sysclient.MenuUpdateReq) (*sysclient.MenuUpdateResp, error) {
	l := logic.NewMenuUpdateLogic(ctx, s.svcCtx)
	return l.MenuUpdate(in)
}

func (s *SysServer) MenuDelete(ctx context.Context, in *sysclient.MenuDeleteReq) (*sysclient.MenuDeleteResp, error) {
	l := logic.NewMenuDeleteLogic(ctx, s.svcCtx)
	return l.MenuDelete(in)
}

func (s *SysServer) DictAdd(ctx context.Context, in *sysclient.DictAddReq) (*sysclient.DictAddResp, error) {
	l := logic.NewDictAddLogic(ctx, s.svcCtx)
	return l.DictAdd(in)
}

func (s *SysServer) DictList(ctx context.Context, in *sysclient.DictListReq) (*sysclient.DictListResp, error) {
	l := logic.NewDictListLogic(ctx, s.svcCtx)
	return l.DictList(in)
}

func (s *SysServer) DictUpdate(ctx context.Context, in *sysclient.DictUpdateReq) (*sysclient.DictUpdateResp, error) {
	l := logic.NewDictUpdateLogic(ctx, s.svcCtx)
	return l.DictUpdate(in)
}

func (s *SysServer) DictDelete(ctx context.Context, in *sysclient.DictDeleteReq) (*sysclient.DictDeleteResp, error) {
	l := logic.NewDictDeleteLogic(ctx, s.svcCtx)
	return l.DictDelete(in)
}

func (s *SysServer) DeptAdd(ctx context.Context, in *sysclient.DeptAddReq) (*sysclient.DeptAddResp, error) {
	l := logic.NewDeptAddLogic(ctx, s.svcCtx)
	return l.DeptAdd(in)
}

func (s *SysServer) DeptList(ctx context.Context, in *sysclient.DeptListReq) (*sysclient.DeptListResp, error) {
	l := logic.NewDeptListLogic(ctx, s.svcCtx)
	return l.DeptList(in)
}

func (s *SysServer) DeptUpdate(ctx context.Context, in *sysclient.DeptUpdateReq) (*sysclient.DeptUpdateResp, error) {
	l := logic.NewDeptUpdateLogic(ctx, s.svcCtx)
	return l.DeptUpdate(in)
}

func (s *SysServer) DeptDelete(ctx context.Context, in *sysclient.DeptDeleteReq) (*sysclient.DeptDeleteResp, error) {
	l := logic.NewDeptDeleteLogic(ctx, s.svcCtx)
	return l.DeptDelete(in)
}

func (s *SysServer) LoginLogAdd(ctx context.Context, in *sysclient.LoginLogAddReq) (*sysclient.LoginLogAddResp, error) {
	l := logic.NewLoginLogAddLogic(ctx, s.svcCtx)
	return l.LoginLogAdd(in)
}

func (s *SysServer) LoginLogList(ctx context.Context, in *sysclient.LoginLogListReq) (*sysclient.LoginLogListResp, error) {
	l := logic.NewLoginLogListLogic(ctx, s.svcCtx)
	return l.LoginLogList(in)
}

func (s *SysServer) LoginLogDelete(ctx context.Context, in *sysclient.LoginLogDeleteReq) (*sysclient.LoginLogDeleteResp, error) {
	l := logic.NewLoginLogDeleteLogic(ctx, s.svcCtx)
	return l.LoginLogDelete(in)
}

func (s *SysServer) SysLogAdd(ctx context.Context, in *sysclient.SysLogAddReq) (*sysclient.SysLogAddResp, error) {
	l := logic.NewSysLogAddLogic(ctx, s.svcCtx)
	return l.SysLogAdd(in)
}

func (s *SysServer) SysLogList(ctx context.Context, in *sysclient.SysLogListReq) (*sysclient.SysLogListResp, error) {
	l := logic.NewSysLogListLogic(ctx, s.svcCtx)
	return l.SysLogList(in)
}

func (s *SysServer) SysLogDelete(ctx context.Context, in *sysclient.SysLogDeleteReq) (*sysclient.SysLogDeleteResp, error) {
	l := logic.NewSysLogDeleteLogic(ctx, s.svcCtx)
	return l.SysLogDelete(in)
}

func (s *SysServer) ConfigAdd(ctx context.Context, in *sysclient.ConfigAddReq) (*sysclient.ConfigAddResp, error) {
	l := logic.NewConfigAddLogic(ctx, s.svcCtx)
	return l.ConfigAdd(in)
}

func (s *SysServer) ConfigList(ctx context.Context, in *sysclient.ConfigListReq) (*sysclient.ConfigListResp, error) {
	l := logic.NewConfigListLogic(ctx, s.svcCtx)
	return l.ConfigList(in)
}

func (s *SysServer) ConfigUpdate(ctx context.Context, in *sysclient.ConfigUpdateReq) (*sysclient.ConfigUpdateResp, error) {
	l := logic.NewConfigUpdateLogic(ctx, s.svcCtx)
	return l.ConfigUpdate(in)
}

func (s *SysServer) ConfigDelete(ctx context.Context, in *sysclient.ConfigDeleteReq) (*sysclient.ConfigDeleteResp, error) {
	l := logic.NewConfigDeleteLogic(ctx, s.svcCtx)
	return l.ConfigDelete(in)
}

func (s *SysServer) UpdateConfigRole(ctx context.Context, in *sysclient.UpdateConfigRoleReq) (*sysclient.UpdateConfigRoleResp, error) {
	l := logic.NewUpdateConfigRoleLogic(ctx, s.svcCtx)
	return l.UpdateConfigRole(in)
}

func (s *SysServer) JobAdd(ctx context.Context, in *sysclient.JobAddReq) (*sysclient.JobAddResp, error) {
	l := logic.NewJobAddLogic(ctx, s.svcCtx)
	return l.JobAdd(in)
}

func (s *SysServer) JobList(ctx context.Context, in *sysclient.JobListReq) (*sysclient.JobListResp, error) {
	l := logic.NewJobListLogic(ctx, s.svcCtx)
	return l.JobList(in)
}

func (s *SysServer) JobUpdate(ctx context.Context, in *sysclient.JobUpdateReq) (*sysclient.JobUpdateResp, error) {
	l := logic.NewJobUpdateLogic(ctx, s.svcCtx)
	return l.JobUpdate(in)
}

func (s *SysServer) JobDelete(ctx context.Context, in *sysclient.JobDeleteReq) (*sysclient.JobDeleteResp, error) {
	l := logic.NewJobDeleteLogic(ctx, s.svcCtx)
	return l.JobDelete(in)
}