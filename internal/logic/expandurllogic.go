package logic

import (
	"context"
	"fmt"

	"5feet11/internal/db"
	"5feet11/internal/svc"
	"5feet11/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandUrlLogic {
	return &ExpandUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandUrlLogic) ExpandUrl(req *types.ExpandReq) (resp *types.ExpandResp, err error) {
	queryUrl := db.UrlTable.SelectBuilder("redirect_url").Query(l.svcCtx.DB)
	queryUrl.BindStruct(db.UrlModel{ID: req.ID})

	var urls []db.UrlModel
	if err := queryUrl.Select(&urls); err != nil {
		return nil, err
	}

	if len(urls) != 1 {
		return nil, fmt.Errorf("no URL found")
	}

	resp = &types.ExpandResp{
		LongUrl: urls[0].LongUrl,
	}

	return resp, err
}
