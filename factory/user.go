package factory

import (
	"github.com/Snowlights/pub/adapter"
	corpus "github.com/Snowlights/pub/grpc"
	"github.com/Snowlights/router/logic"
)

type LoginUser struct {
	corpus.LoginUserReq
}

func FactoryLoginUser() logic.LogicHandle {
	return new(LoginUser)
}

func (m *LoginUser) Handle(h *logic.ReqHeader, r *logic.HttpRequest ) logic.HttpResponse{
	ctx := r.Request().Context()
	req := &m.LoginUserReq
	res := adapter.LoginUser(ctx,req)
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