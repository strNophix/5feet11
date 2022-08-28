package logic

import (
	"context"
	"fmt"

	"5feet11/internal/db"
	"5feet11/internal/svc"
	"5feet11/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLinkLogic {
	return &GetLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLinkLogic) GetLink(req *types.ExpandReq) (resp *types.GetLinkResp, err error) {
	queryUrl := db.UrlTable.SelectQuery(l.svcCtx.DB)
	queryUrl.BindStruct(db.UrlModel{ID: req.ID})

	var urls []db.UrlModel
	if err := queryUrl.Select(&urls); err != nil {
		return nil, err
	}

	if len(urls) != 1 {
		return nil, fmt.Errorf("no URL found")
	}

	resp = &types.GetLinkResp{
		ID:        urls[0].ID,
		LongUrl:   urls[0].LongUrl,
		CreatedAt: urls[0].CreatedAt.String(),
		Lifespan:  urls[0].Lifespan,
	}

	return resp, err
}
