// Code generated by goctl. DO NOT EDIT!
// Source: sms.proto

package server

import (
	"context"

	"zero-admin-learn/rpc/sms/internal/logic"
	"zero-admin-learn/rpc/sms/internal/svc"
	"zero-admin-learn/rpc/sms/smsclient"
)

type SmsServer struct {
	svcCtx *svc.ServiceContext
	smsclient.UnimplementedSmsServer
}

func NewSmsServer(svcCtx *svc.ServiceContext) *SmsServer {
	return &SmsServer{
		svcCtx: svcCtx,
	}
}

func (s *SmsServer) CouponAdd(ctx context.Context, in *smsclient.CouponAddReq) (*smsclient.CouponAddResp, error) {
	l := logic.NewCouponAddLogic(ctx, s.svcCtx)
	return l.CouponAdd(in)
}

func (s *SmsServer) CouponList(ctx context.Context, in *smsclient.CouponListReq) (*smsclient.CouponListResp, error) {
	l := logic.NewCouponListLogic(ctx, s.svcCtx)
	return l.CouponList(in)
}

func (s *SmsServer) CouponUpdate(ctx context.Context, in *smsclient.CouponUpdateReq) (*smsclient.CouponUpdateResp, error) {
	l := logic.NewCouponUpdateLogic(ctx, s.svcCtx)
	return l.CouponUpdate(in)
}

func (s *SmsServer) CouponDelete(ctx context.Context, in *smsclient.CouponDeleteReq) (*smsclient.CouponDeleteResp, error) {
	l := logic.NewCouponDeleteLogic(ctx, s.svcCtx)
	return l.CouponDelete(in)
}

func (s *SmsServer) CouponHistoryAdd(ctx context.Context, in *smsclient.CouponHistoryAddReq) (*smsclient.CouponHistoryAddResp, error) {
	l := logic.NewCouponHistoryAddLogic(ctx, s.svcCtx)
	return l.CouponHistoryAdd(in)
}

func (s *SmsServer) CouponHistoryList(ctx context.Context, in *smsclient.CouponHistoryListReq) (*smsclient.CouponHistoryListResp, error) {
	l := logic.NewCouponHistoryListLogic(ctx, s.svcCtx)
	return l.CouponHistoryList(in)
}

func (s *SmsServer) CouponHistoryUpdate(ctx context.Context, in *smsclient.CouponHistoryUpdateReq) (*smsclient.CouponHistoryUpdateResp, error) {
	l := logic.NewCouponHistoryUpdateLogic(ctx, s.svcCtx)
	return l.CouponHistoryUpdate(in)
}

func (s *SmsServer) CouponHistoryDelete(ctx context.Context, in *smsclient.CouponHistoryDeleteReq) (*smsclient.CouponHistoryDeleteResp, error) {
	l := logic.NewCouponHistoryDeleteLogic(ctx, s.svcCtx)
	return l.CouponHistoryDelete(in)
}

func (s *SmsServer) CouponProductCategoryRelationAdd(ctx context.Context, in *smsclient.CouponProductCategoryRelationAddReq) (*smsclient.CouponProductCategoryRelationAddResp, error) {
	l := logic.NewCouponProductCategoryRelationAddLogic(ctx, s.svcCtx)
	return l.CouponProductCategoryRelationAdd(in)
}

func (s *SmsServer) CouponProductCategoryRelationList(ctx context.Context, in *smsclient.CouponProductCategoryRelationListReq) (*smsclient.CouponProductCategoryRelationListResp, error) {
	l := logic.NewCouponProductCategoryRelationListLogic(ctx, s.svcCtx)
	return l.CouponProductCategoryRelationList(in)
}

func (s *SmsServer) CouponProductCategoryRelationUpdate(ctx context.Context, in *smsclient.CouponProductCategoryRelationUpdateReq) (*smsclient.CouponProductCategoryRelationUpdateResp, error) {
	l := logic.NewCouponProductCategoryRelationUpdateLogic(ctx, s.svcCtx)
	return l.CouponProductCategoryRelationUpdate(in)
}

func (s *SmsServer) CouponProductCategoryRelationDelete(ctx context.Context, in *smsclient.CouponProductCategoryRelationDeleteReq) (*smsclient.CouponProductCategoryRelationDeleteResp, error) {
	l := logic.NewCouponProductCategoryRelationDeleteLogic(ctx, s.svcCtx)
	return l.CouponProductCategoryRelationDelete(in)
}

func (s *SmsServer) CouponProductRelationAdd(ctx context.Context, in *smsclient.CouponProductRelationAddReq) (*smsclient.CouponProductRelationAddResp, error) {
	l := logic.NewCouponProductRelationAddLogic(ctx, s.svcCtx)
	return l.CouponProductRelationAdd(in)
}

func (s *SmsServer) CouponProductRelationList(ctx context.Context, in *smsclient.CouponProductRelationListReq) (*smsclient.CouponProductRelationListResp, error) {
	l := logic.NewCouponProductRelationListLogic(ctx, s.svcCtx)
	return l.CouponProductRelationList(in)
}

func (s *SmsServer) CouponProductRelationUpdate(ctx context.Context, in *smsclient.CouponProductRelationUpdateReq) (*smsclient.CouponProductRelationUpdateResp, error) {
	l := logic.NewCouponProductRelationUpdateLogic(ctx, s.svcCtx)
	return l.CouponProductRelationUpdate(in)
}

func (s *SmsServer) CouponProductRelationDelete(ctx context.Context, in *smsclient.CouponProductRelationDeleteReq) (*smsclient.CouponProductRelationDeleteResp, error) {
	l := logic.NewCouponProductRelationDeleteLogic(ctx, s.svcCtx)
	return l.CouponProductRelationDelete(in)
}

func (s *SmsServer) FlashPromotionLogAdd(ctx context.Context, in *smsclient.FlashPromotionLogAddReq) (*smsclient.FlashPromotionLogAddResp, error) {
	l := logic.NewFlashPromotionLogAddLogic(ctx, s.svcCtx)
	return l.FlashPromotionLogAdd(in)
}

func (s *SmsServer) FlashPromotionLogList(ctx context.Context, in *smsclient.FlashPromotionLogListReq) (*smsclient.FlashPromotionLogListResp, error) {
	l := logic.NewFlashPromotionLogListLogic(ctx, s.svcCtx)
	return l.FlashPromotionLogList(in)
}

func (s *SmsServer) FlashPromotionLogUpdate(ctx context.Context, in *smsclient.FlashPromotionLogUpdateReq) (*smsclient.FlashPromotionLogUpdateResp, error) {
	l := logic.NewFlashPromotionLogUpdateLogic(ctx, s.svcCtx)
	return l.FlashPromotionLogUpdate(in)
}

func (s *SmsServer) FlashPromotionLogDelete(ctx context.Context, in *smsclient.FlashPromotionLogDeleteReq) (*smsclient.FlashPromotionLogDeleteResp, error) {
	l := logic.NewFlashPromotionLogDeleteLogic(ctx, s.svcCtx)
	return l.FlashPromotionLogDelete(in)
}

func (s *SmsServer) FlashPromotionAdd(ctx context.Context, in *smsclient.FlashPromotionAddReq) (*smsclient.FlashPromotionAddResp, error) {
	l := logic.NewFlashPromotionAddLogic(ctx, s.svcCtx)
	return l.FlashPromotionAdd(in)
}

func (s *SmsServer) FlashPromotionList(ctx context.Context, in *smsclient.FlashPromotionListReq) (*smsclient.FlashPromotionListResp, error) {
	l := logic.NewFlashPromotionListLogic(ctx, s.svcCtx)
	return l.FlashPromotionList(in)
}

func (s *SmsServer) FlashPromotionUpdate(ctx context.Context, in *smsclient.FlashPromotionUpdateReq) (*smsclient.FlashPromotionUpdateResp, error) {
	l := logic.NewFlashPromotionUpdateLogic(ctx, s.svcCtx)
	return l.FlashPromotionUpdate(in)
}

func (s *SmsServer) FlashPromotionDelete(ctx context.Context, in *smsclient.FlashPromotionDeleteReq) (*smsclient.FlashPromotionDeleteResp, error) {
	l := logic.NewFlashPromotionDeleteLogic(ctx, s.svcCtx)
	return l.FlashPromotionDelete(in)
}

func (s *SmsServer) FlashPromotionProductRelationAdd(ctx context.Context, in *smsclient.FlashPromotionProductRelationAddReq) (*smsclient.FlashPromotionProductRelationAddResp, error) {
	l := logic.NewFlashPromotionProductRelationAddLogic(ctx, s.svcCtx)
	return l.FlashPromotionProductRelationAdd(in)
}

func (s *SmsServer) FlashPromotionProductRelationList(ctx context.Context, in *smsclient.FlashPromotionProductRelationListReq) (*smsclient.FlashPromotionProductRelationListResp, error) {
	l := logic.NewFlashPromotionProductRelationListLogic(ctx, s.svcCtx)
	return l.FlashPromotionProductRelationList(in)
}

func (s *SmsServer) FlashPromotionProductRelationUpdate(ctx context.Context, in *smsclient.FlashPromotionProductRelationUpdateReq) (*smsclient.FlashPromotionProductRelationUpdateResp, error) {
	l := logic.NewFlashPromotionProductRelationUpdateLogic(ctx, s.svcCtx)
	return l.FlashPromotionProductRelationUpdate(in)
}

func (s *SmsServer) FlashPromotionProductRelationDelete(ctx context.Context, in *smsclient.FlashPromotionProductRelationDeleteReq) (*smsclient.FlashPromotionProductRelationDeleteResp, error) {
	l := logic.NewFlashPromotionProductRelationDeleteLogic(ctx, s.svcCtx)
	return l.FlashPromotionProductRelationDelete(in)
}

func (s *SmsServer) FlashPromotionSessionAdd(ctx context.Context, in *smsclient.FlashPromotionSessionAddReq) (*smsclient.FlashPromotionSessionAddResp, error) {
	l := logic.NewFlashPromotionSessionAddLogic(ctx, s.svcCtx)
	return l.FlashPromotionSessionAdd(in)
}

func (s *SmsServer) FlashPromotionSessionList(ctx context.Context, in *smsclient.FlashPromotionSessionListReq) (*smsclient.FlashPromotionSessionListResp, error) {
	l := logic.NewFlashPromotionSessionListLogic(ctx, s.svcCtx)
	return l.FlashPromotionSessionList(in)
}

func (s *SmsServer) FlashPromotionSessionUpdate(ctx context.Context, in *smsclient.FlashPromotionSessionUpdateReq) (*smsclient.FlashPromotionSessionUpdateResp, error) {
	l := logic.NewFlashPromotionSessionUpdateLogic(ctx, s.svcCtx)
	return l.FlashPromotionSessionUpdate(in)
}

func (s *SmsServer) FlashPromotionSessionDelete(ctx context.Context, in *smsclient.FlashPromotionSessionDeleteReq) (*smsclient.FlashPromotionSessionDeleteResp, error) {
	l := logic.NewFlashPromotionSessionDeleteLogic(ctx, s.svcCtx)
	return l.FlashPromotionSessionDelete(in)
}

func (s *SmsServer) HomeAdvertiseAdd(ctx context.Context, in *smsclient.HomeAdvertiseAddReq) (*smsclient.HomeAdvertiseAddResp, error) {
	l := logic.NewHomeAdvertiseAddLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseAdd(in)
}

func (s *SmsServer) HomeAdvertiseList(ctx context.Context, in *smsclient.HomeAdvertiseListReq) (*smsclient.HomeAdvertiseListResp, error) {
	l := logic.NewHomeAdvertiseListLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseList(in)
}

func (s *SmsServer) HomeAdvertiseUpdate(ctx context.Context, in *smsclient.HomeAdvertiseUpdateReq) (*smsclient.HomeAdvertiseUpdateResp, error) {
	l := logic.NewHomeAdvertiseUpdateLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseUpdate(in)
}

func (s *SmsServer) HomeAdvertiseDelete(ctx context.Context, in *smsclient.HomeAdvertiseDeleteReq) (*smsclient.HomeAdvertiseDeleteResp, error) {
	l := logic.NewHomeAdvertiseDeleteLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseDelete(in)
}

func (s *SmsServer) HomeBrandAdd(ctx context.Context, in *smsclient.HomeBrandAddReq) (*smsclient.HomeBrandAddResp, error) {
	l := logic.NewHomeBrandAddLogic(ctx, s.svcCtx)
	return l.HomeBrandAdd(in)
}

func (s *SmsServer) HomeBrandList(ctx context.Context, in *smsclient.HomeBrandListReq) (*smsclient.HomeBrandListResp, error) {
	l := logic.NewHomeBrandListLogic(ctx, s.svcCtx)
	return l.HomeBrandList(in)
}

func (s *SmsServer) HomeBrandUpdate(ctx context.Context, in *smsclient.HomeBrandUpdateReq) (*smsclient.HomeBrandUpdateResp, error) {
	l := logic.NewHomeBrandUpdateLogic(ctx, s.svcCtx)
	return l.HomeBrandUpdate(in)
}

func (s *SmsServer) HomeBrandDelete(ctx context.Context, in *smsclient.HomeBrandDeleteReq) (*smsclient.HomeBrandDeleteResp, error) {
	l := logic.NewHomeBrandDeleteLogic(ctx, s.svcCtx)
	return l.HomeBrandDelete(in)
}

func (s *SmsServer) HomeNewProductAdd(ctx context.Context, in *smsclient.HomeNewProductAddReq) (*smsclient.HomeNewProductAddResp, error) {
	l := logic.NewHomeNewProductAddLogic(ctx, s.svcCtx)
	return l.HomeNewProductAdd(in)
}

func (s *SmsServer) HomeNewProductList(ctx context.Context, in *smsclient.HomeNewProductListReq) (*smsclient.HomeNewProductListResp, error) {
	l := logic.NewHomeNewProductListLogic(ctx, s.svcCtx)
	return l.HomeNewProductList(in)
}

func (s *SmsServer) HomeNewProductUpdate(ctx context.Context, in *smsclient.HomeNewProductUpdateReq) (*smsclient.HomeNewProductUpdateResp, error) {
	l := logic.NewHomeNewProductUpdateLogic(ctx, s.svcCtx)
	return l.HomeNewProductUpdate(in)
}

func (s *SmsServer) HomeNewProductDelete(ctx context.Context, in *smsclient.HomeNewProductDeleteReq) (*smsclient.HomeNewProductDeleteResp, error) {
	l := logic.NewHomeNewProductDeleteLogic(ctx, s.svcCtx)
	return l.HomeNewProductDelete(in)
}

func (s *SmsServer) HomeRecommendProductAdd(ctx context.Context, in *smsclient.HomeRecommendProductAddReq) (*smsclient.HomeRecommendProductAddResp, error) {
	l := logic.NewHomeRecommendProductAddLogic(ctx, s.svcCtx)
	return l.HomeRecommendProductAdd(in)
}

func (s *SmsServer) HomeRecommendProductList(ctx context.Context, in *smsclient.HomeRecommendProductListReq) (*smsclient.HomeRecommendProductListResp, error) {
	l := logic.NewHomeRecommendProductListLogic(ctx, s.svcCtx)
	return l.HomeRecommendProductList(in)
}

func (s *SmsServer) HomeRecommendProductUpdate(ctx context.Context, in *smsclient.HomeRecommendProductUpdateReq) (*smsclient.HomeRecommendProductUpdateResp, error) {
	l := logic.NewHomeRecommendProductUpdateLogic(ctx, s.svcCtx)
	return l.HomeRecommendProductUpdate(in)
}

func (s *SmsServer) HomeRecommendProductDelete(ctx context.Context, in *smsclient.HomeRecommendProductDeleteReq) (*smsclient.HomeRecommendProductDeleteResp, error) {
	l := logic.NewHomeRecommendProductDeleteLogic(ctx, s.svcCtx)
	return l.HomeRecommendProductDelete(in)
}

func (s *SmsServer) HomeRecommendSubjectAdd(ctx context.Context, in *smsclient.HomeRecommendSubjectAddReq) (*smsclient.HomeRecommendSubjectAddResp, error) {
	l := logic.NewHomeRecommendSubjectAddLogic(ctx, s.svcCtx)
	return l.HomeRecommendSubjectAdd(in)
}

func (s *SmsServer) HomeRecommendSubjectList(ctx context.Context, in *smsclient.HomeRecommendSubjectListReq) (*smsclient.HomeRecommendSubjectListResp, error) {
	l := logic.NewHomeRecommendSubjectListLogic(ctx, s.svcCtx)
	return l.HomeRecommendSubjectList(in)
}

func (s *SmsServer) HomeRecommendSubjectUpdate(ctx context.Context, in *smsclient.HomeRecommendSubjectUpdateReq) (*smsclient.HomeRecommendSubjectUpdateResp, error) {
	l := logic.NewHomeRecommendSubjectUpdateLogic(ctx, s.svcCtx)
	return l.HomeRecommendSubjectUpdate(in)
}

func (s *SmsServer) HomeRecommendSubjectDelete(ctx context.Context, in *smsclient.HomeRecommendSubjectDeleteReq) (*smsclient.HomeRecommendSubjectDeleteResp, error) {
	l := logic.NewHomeRecommendSubjectDeleteLogic(ctx, s.svcCtx)
	return l.HomeRecommendSubjectDelete(in)
}
