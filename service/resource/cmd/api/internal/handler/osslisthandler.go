package handler

import (
	"fmt"
	"go-zero-resource/common/api"
	"go-zero-resource/common/errorx"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-resource/service/resource/cmd/api/internal/logic"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"
)

func ossListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OssListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprint(err)))
			return
		}

		l := logic.NewOssListLogic(r.Context(), ctx)
		resp, err := l.OssList(req)
		if err != nil {
			httpx.Error(w, errorx.ParseError(err))
		} else {
			httpx.OkJson(w, api.OkWithData(resp))
		}
	}
}
