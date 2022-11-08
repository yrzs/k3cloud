package response

type K3ResponseStruct struct {
	Result Result `json:"Result"`
}

type Result struct {
	ResponseStatus K3ResponseStatus `json:"ResponseStatus"`
	ID             int              `json:"Id"`
	Number         string           `json:"Number"`
	NeedReturnData []NeedReturnData `json:"NeedReturnData"`
}
type NeedReturnData struct {
	FVOUCHERGROUPNO string `json:"FVOUCHERGROUPNO"`
	FVOUCHERID      string `json:"FVOUCHERID"`
}

type K3ResponseStatus struct {
	IsSuccess       bool              `json:"IsSuccess"`
	Errors          []interface{}     `json:"Errors"`
	SuccessEntitys  []SuccessEntities `json:"SuccessEntitys"`
	SuccessMessages []interface{}     `json:"SuccessMessages"`
	MsgCode         int               `json:"MsgCode"`
}

type SuccessEntities struct {
	ID     int    `json:"Id"`
	Number string `json:"Number"`
	DIndex int    `json:"DIndex"`
}
