package handler

import (
	"github.com/SliverFlow/minidouyin/common/result"
	"net/http"

	"github.com/SliverFlow/minidouyin/app/user/cmd/api/internal/logic"
	"github.com/SliverFlow/minidouyin/app/user/cmd/api/internal/svc"
	"github.com/SliverFlow/minidouyin/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		result.HttpResult(r, w, resp, err)
	}
}
