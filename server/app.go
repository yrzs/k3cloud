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
	res, _ := browser.PostJson(c.Host+kernel.LOGIN_API, formParams)
	var k3Response = &response.K3LoginResponseStruct{}
	e := object.HashMapToStructure(res, k3Response)
	if e != nil {
		fmt.Println("eeee", e)
		return nil, errors.New("k3 cloud login fail")
	}
	if k3Response.LoginResultType == 0 {
		return nil, errors.New(k3Response.Message)
	}
	app := &K3Cloud{
		Config:  c,
		Browser: browser,
	}
	return app, nil
}

//Submit.  提交接口
//formId   表单ID
//data     数据
func (k *K3Cloud) Submit(formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.SUBMIT_API
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Browser.PostJson(url, postData)
}

//Save.  保存接口
//formId 表单ID
//data   数据
func (k *K3Cloud) Save(formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.SAVE_API
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Browser.PostJson(url, postData)
}

//BatchSave.  批量接口
//formId      表单ID
//data        数据
func (k *K3Cloud) BatchSave(formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.BATCHSAVE_API
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Browser.PostJson(url, postData)
}

//Audit.  审核接口
//formId  表单ID
//data    数据
func (k *K3Cloud) Audit(formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.AUDIT_API
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Browser.PostJson(url, postData)
}

//UnAudit.  反审核
//formId    表单ID
//data      数据
func (k *K3Cloud) UnAudit(formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.UNAUDIT_API
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Browser.PostJson(url, postData)
}

//View.  详情接口
//formId 查询表单ID
//data   查询数据
func (k *K3Cloud) View(formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.VIEW_API
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Browser.PostJson(url, postData)
}

// ExecuteBillQuery. 单据查询
func (k *K3Cloud) ExecuteBillQuery(data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.EXECUTEBILLQUERY_API
	var postData = &object.HashMap{
		"data": data,
	}
	return k.Browser.PostJson(url, postData)
}

//Operation. 操作接口
//formId     查询表单ID
//opNumber   操作标识
//data       查询数据
func (k *K3Cloud) Operation(formId, opNumber string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.EXCUTEOPERATION_API
	var postData = &object.HashMap{
		"formid":   formId,
		"opNumber": opNumber,
		"data":     data,
	}
	return k.Browser.PostJson(url, postData)
}

//Push.   下推接口
//formId  表单ID
//data    数据
func (k *K3Cloud) Push(formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.PUSH_API
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Browser.PostJson(url, postData)
}
