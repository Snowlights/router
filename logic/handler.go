package logic

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


type ReqHeader struct {
	Token string `json:"token"`
	Uid   int64  `json:"h_m"`

	Did     string `json:"h_did"`
	Ver     string `json:"h_av"`
	Dt      int32  `json:"h_dt"`
	Dtsub   int32  `json:"h_dt_sub"`
	Ch      string `json:"h_ch"`
	Net     int32  `json:"h_nt"`
	Unionid string `json:"h_unionid"`

	HLc  string `json:"h_lc"`
	Cate int32  `json:"cate"`

	Source    int32 `json:"h_src"`
	Zone      int32 `json:"zone"`
	Subsource int32 `json:"h_sub_src"`

	Ip     string `json:"-"`
	Region string `json:"-"`
	Host   string `json:"-"`
	Level  int32  `json:"-"`

}

type LogicHandle interface {
	Handle(h *ReqHeader, r *HttpRequest) HttpResponse
}

type factoryHandle func() LogicHandle

type StandardReq struct {
	h        ReqHeader
	fac      factoryHandle
	authType int32
}

func InitStandardReq(fac factoryHandle) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return HttpRequestWrapper(func() HandleRequest {
		return &StandardReq{
			fac:      fac,
			authType: 0,
		}
	})
}


func (m *StandardReq) getByteResp(resp HttpResponse) (HttpResponse, int, int) {
	//status, rbody := resp.Marshal()
	//return snetutil.NewHttpRespBytes(status, rbody), status, len(rbody)
	return resp, 0, 0

}

func (m *StandardReq) Handle(r *HttpRequest) HttpResponse {
	fun := "StandardReq.Handle -->"

	var resp HttpResponse
	var status, lenrbody int

	host := r.Headers().String("X-Host")
	m.h.Host = host
	defer func() {
		// 时延统计
		_ = status
		_ = lenrbody
		deviceid := r.Cookies().String("ipalfish_device_id")
		deviceidheader := ""
		if len(deviceid) == 0 {
			deviceidheader = r.Headers().String("ipalfish-device-id")
			deviceid = deviceidheader
		}
	}()

	gzip := r.Query().Int("gzip") == 1
	var err error
	if gzip {
		err = r.Body().JsonUnGzip(&m.h)

	} else {
		err = r.Body().Json(&m.h)
	}

	if err != nil {
		fmt.Printf("%s jsonunmarhead url:%s body:%s err:%s", fun, r.URL(), r.Body().Binary(), err)
		resp, status, lenrbody = m.getByteResp(NewHttpRespString(400, "json unmarshal err:"+err.Error()))
		return resp
	}


	bh := m.fac()
	if gzip {
		err = r.Body().JsonUnGzip(bh)
	} else {
		err = r.Body().Json(bh)
	}
	if err != nil {
		fmt.Printf("%s jsonunmarhand url:%s body:%s err:%s", fun, r.URL(), r.Body().Binary(), err)
		resp, status, lenrbody = m.getByteResp(NewHttpRespString(400, "json unmarshal err:"+err.Error()))
		return resp
	}


	resp, status, lenrbody = m.getByteResp(bh.Handle(&m.h, r))

	return resp
}
