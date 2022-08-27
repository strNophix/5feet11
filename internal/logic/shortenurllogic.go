package logic

import (
	"context"
	"time"

	"5feet11/internal/db"
	"5feet11/internal/svc"
	"5feet11/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenUrlLogic {
	return &ShortenUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenUrlLogic) ShortenUrl(req *types.ShortenReq) (resp *types.ShortenResp, err error) {
	id := l.svcCtx.Snowflake.Generate().Base58()

	insertUrl := db.UrlTable.InsertBuilder().TTL(30 * time.Second).Query(l.svcCtx.DB)
	insertUrl.BindStruct(db.UrlModel{
		ID:          id,
		RedirectUrl: req.RedirectUrl,
	})

	if err := insertUrl.ExecRelease(); err != nil {
		return nil, err
	}

	resp = &types.ShortenResp{
		ID: id,
	}
	return resp, nil
}
