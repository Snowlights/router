package factory

import (
	"github.com/Snowlights/pub/adapter"
	corpus "github.com/Snowlights/pub/grpc"
	"github.com/Snowlights/router/logic"
)

type SendMessage struct {
	corpus.SendMessageReq
}

func FactorySendMessage() logic.LogicHandle {
	return new(SendMessage)
}

func (m *SendMessage) Handle(h *logic.ReqHeader, r *logic.HttpRequest ) logic.HttpResponse{
	ctx := r.Request().Context()
	req := &m.SendMessageReq
	res := adapter.SendMessage(ctx,req)
	if res.Errinfo != nil{
		return logic.NewHttpRespJson200(&logic.ApiReturn{
			Ret:  int(res.Errinfo.Ret),
			Msg:  res.Errinfo.Msg,
		})
	}
	return logic.NewHttpRespJson200(&logic.ApiReturnPlus2{
		Ret:  1,
		Data: &logic.ApiData2{
			Ent: res.Data,
		},
	})
}