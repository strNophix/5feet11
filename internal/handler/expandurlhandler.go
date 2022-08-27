package handler

import (
	"net/http"

	"5feet11/internal/logic"
	"5feet11/internal/svc"
	"5feet11/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ExpandUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewExpandUrlLogic(r.Context(), svcCtx)
		resp, err := l.ExpandUrl(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			http.Redirect(w, r, resp.LongUrl, http.StatusTemporaryRedirect)
		}
	}
}
