package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qq-bot-307/service/http/internal/logic"
	"qq-bot-307/service/http/internal/svc"
	"qq-bot-307/service/http/internal/types"
)

func MessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMessageLogic(r.Context(), svcCtx)
		resp, err := l.Message(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
