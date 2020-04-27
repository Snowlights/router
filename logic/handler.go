package logic

import (
	"encoding/json"
	"fmt"
	corpus "github.com/Snowlights/pub/grpc"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := &corpus.LoginUserReq{}

	if r.Method != "POST"{
		http.Error(w,"method error",405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		er := fmt.Sprintf("body read err:%s", err)
		http.Error(w, er, 500)
		return
	}

	if len(body) == 0{
		http.Error(w,"data empty",400)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "body invalid", 400)
		return
	}


}
