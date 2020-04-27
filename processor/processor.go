package processor

import (
	"github.com/Snowlights/router/factory"
	"github.com/Snowlights/router/logic"
	"github.com/julienschmidt/httprouter"
)

func InitRouter() (string,interface{}){

	router := httprouter.New()

	router.POST("/corpus/user/login",logic.InitStandardReq(factory.FactoryLoginUser))
	router.POST("/corpus/sendmessage",logic.InitStandardReq(factory.FactorySendMessage))


	return ":9101",router
}

