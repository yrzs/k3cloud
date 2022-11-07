package response

import (
	"encoding/json"
)

type K3LoginResponseStruct struct {
	Message           string      `json:"Message"`
	MessageCode       string      `json:"MessageCode"`
	LoginResultType   int         `json:"LoginResultType"`
	Context           interface{} `json:"Context"`
	KDSVCSessionID    string      `json:"KDSVCSessionId"`
	FormID            interface{} `json:"FormId"`
	RedirectFormParam interface{} `json:"RedirectFormParam"`
	FormInputObject   interface{} `json:"FormInputObject"`
	ErrorStackTrace   interface{} `json:"ErrorStackTrace"`
	Lcid              int         `json:"Lcid"`
	AccessToken       interface{} `json:"AccessToken"`
	KdAccessResult    interface{} `json:"KdAccessResult"`
	IsSuccessByAPI    bool        `json:"IsSuccessByAPI"`
}

type K3ResponseStruct struct {
	Result Result `json:"Result"`
}

type Result struct {
	ResponseStatus ResponseStatus   `json:"ResponseStatus"`
	ID             int              `json:"Id"`
	Number         string           `json:"Number"`
	NeedReturnData []NeedReturnData `json:"NeedReturnData"`
}
type NeedReturnData struct {
	FVOUCHERGROUPNO string `json:"FVOUCHERGROUPNO"`
	FVOUCHERID      string `json:"FVOUCHERID"`
}

type ResponseStatus struct {
	IsSuccess       bool             `json:"IsSuccess"`
	Errors          []interface{}    `json:"Errors"`
	SuccessEntitys  []SuccessEntitys `json:"SuccessEntitys"`
	SuccessMessages []interface{}    `json:"SuccessMessages"`
	MsgCode         int              `json:"MsgCode"`
}

type SuccessEntitys struct {
	ID     int    `json:"Id"`
	Number string `json:"Number"`
	DIndex int    `json:"DIndex"`
}

func K3LoginResponseToStruct(data []byte) K3LoginResponseStruct {
	var response K3LoginResponseStruct
	_ = json.Unmarshal(data, &response)
	return response
}

func K3ResponseToStruct(data []byte) K3ResponseStruct {
	var response K3ResponseStruct
	_ = json.Unmarshal(data, &response)
	return response
}
