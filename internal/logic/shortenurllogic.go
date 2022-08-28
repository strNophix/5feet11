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
	insertBuilder := db.UrlTable.InsertBuilder()

	logx.Info(req.ExpiresAfter)
	if req.ExpiresAfter != 0 {
		insertBuilder.TTL(time.Second * time.Duration(req.ExpiresAfter))
	}

	insertUrl := insertBuilder.Query(l.svcCtx.DB)
	insertUrl.BindStruct(db.UrlModel{
		ID:        id,
		LongUrl:   req.LongUrl,
		Lifespan:  req.ExpiresAfter,
		CreatedAt: time.Now(),
	})

	if err := insertUrl.ExecRelease(); err != nil {
		return nil, err
	}

	resp = &types.ShortenResp{
		ID: id,
	}

	return resp, nil
}
