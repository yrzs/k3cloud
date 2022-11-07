package server

import (
	"errors"
	"fmt"
	"github.com/yrzs/k3cloud/kernel"
	"github.com/yrzs/k3cloud/object"
	"github.com/yrzs/k3cloud/response"
)

// K3Cloud. application
type K3Cloud struct {
	Config  *K3Config
	Browser *kernel.Browser
}

// K3Config. 金蝶云星空账号地址
type K3Config struct {
	Host     string
	AccID    string
	Username string
	Password string
	LcID     int
}

// NewK3Cloud. new application
func NewK3Cloud(c *K3Config) (*K3Cloud, error) {
	browser := kernel.NewBrowser()
	formParams := kernel.CreateLoginPostData(c.AccID, c.Username, c.Password, c.LcID)
	res := browser.PostJson(c.Host+kernel.LOGIN_API, formParams)
	k3Response := response.K3LoginResponseToStruct(res)
	if k3Response.LoginResultType == 0 {
		return nil, errors.New(k3Response.Message)
	}
	app := &K3Cloud{
		Config:  c,
		Browser: browser,
	}
	return app, nil
}

//View.  详情接口
//formId 查询表单ID
//data   查询数据
//func (k *K3Cloud) View(formId string, data map[string]interface{}) {
//	url := k.Config.Host + kernel.VIEW_API
//	var d = make(map[string]interface{}, 0)
//	var postData = make(map[string]interface{}, 0)
//	d["FormId"] = formId
//	d["FieldKeys"] = "FNUMBER,Fname"
//	postData["data"] = d
//	res := k.Browser.PostJson(url, postData)
//	fmt.Println(string(res))
//	var v interface{}
//	_ = object.JsonDecode(res, v)
//	fmt.Println(v)
//}

// ExecuteBillQuery. 单据查询
func (k *K3Cloud) ExecuteBillQuery(data *object.HashMap) {
	url := k.Config.Host + kernel.EXECUTEBILLQUERY_API
	var postData = &object.HashMap{
		"data":data,
	}
	res := k.Browser.PostJson(url, postData)
	fmt.Println(string(res))
	var v interface{}
	_ = object.JsonDecode(res, v)
	fmt.Println(v)
}
