package processor

import "github.com/julienschmidt/httprouter"

func InitRouter() (string,interface{}){

	router := httprouter.New()

	//router.POST("/corpus/user/login",logic)



	return ":9101",router
}
