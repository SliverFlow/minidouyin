package handler

import (
	"github.com/SliverFlow/minidouyin/common/result"
	"net/http"

	"github.com/SliverFlow/minidouyin/app/user/cmd/api/internal/logic"
	"github.com/SliverFlow/minidouyin/app/user/cmd/api/internal/svc"
	"github.com/SliverFlow/minidouyin/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPingReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping(&req)
		result.HttpResult(r, w, resp, err)
	}
}
