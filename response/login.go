package response

type LoginResponse struct {
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
