package logic


type BaseHead struct {
	Ver    string `json:"h_av"`
	Dt     int32  `json:"h_dt"`
	Did    string `json:"h_did"`
	Ch     string `json:"h_ch"`
	Net    int32  `json:"h_nt"`
	Source int32  `json:"h_src"`
	Ip     string `json:"-"`
	Region string `json:"-"`
}

type UidBaseHead struct {
	BaseHead
	Uid int64 `json:"h_m"`
}

type ApiReturn struct {
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type ApiExt struct {
	Users     interface{} `json:"users,omitempty"`
	Userviews interface{} `json:"userviews,omitempty"`
}

type ApiData struct {
	Ent interface{} `json:"ent,omitempty"`
	Ext *ApiExt     `json:"ext,omitempty"`
}

type ApiReturnPlus struct {
	Ret  int      `json:"ret"`
	Msg  string   `json:"msg,omitempty"`
	Data *ApiData `json:"data,omitempty"`
}

type ApiData2 struct {
	Ent interface{} `json:"ent,omitempty"`
	Ext interface{} `json:"ext,omitempty"`
}

type ApiReturnPlus2 struct {
	Ret  int       `json:"ret"`
	Msg  string    `json:"msg,omitempty"`
	Data *ApiData2 `json:"data,omitempty"`
}
