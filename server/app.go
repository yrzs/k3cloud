package server

import (
	"context"
	"errors"
	"github.com/yrzs/k3cloud/kernel"
	"github.com/yrzs/k3cloud/object"
	"github.com/yrzs/k3cloud/response"
	"time"
)

// K3Cloud. application
type K3Cloud struct {
	Config *K3Config
}

// K3Config. 金蝶云星空账号地址
type K3Config struct {
	Host     string
	AccID    string
	Username string
	Password string
	LcID     int
}

// initLogin.
func initLogin(b *kernel.Browser, c *K3Config) error {
	var parameters = make([]interface{}, 0, 4)
	parameters = append(parameters, c.AccID)
	parameters = append(parameters, c.Username)
	parameters = append(parameters, c.Password)
	parameters = append(parameters, c.LcID)
	var data = &object.HashMap{
		"format":     1,
		"useragent":  "ApiClient",
		"rid":        &object.HashMap{},
		"parameters": parameters,
		"timestamp":  time.Now().Format("2006-01-02"),
		"v":          "1.0",
	}
	ctx := context.Background()
	res, _ := b.PostJson(ctx, c.Host+kernel.LoginApi, data)
	var k3Response = &response.LoginResponse{}
	e := object.HashMapToStructure(res, k3Response)
	if e != nil {
		return errors.New("k3 cloud login fail")
	}
	if k3Response.LoginResultType == 0 {
		return errors.New(k3Response.Message)
	}
	return nil
}

// NewK3Cloud. new application
func NewK3Cloud(c *K3Config) (*K3Cloud, error) {
	browser := kernel.NewBrowser()
	err := initLogin(browser, c)
	if err != nil {
		return nil, err
	}
	app := &K3Cloud{
		Config: c,
	}
	return app, nil
}

func (k *K3Cloud) newBrowser() *kernel.Browser {
	browser := kernel.NewBrowser()
	err := initLogin(browser, k.Config)
	if err != nil {
		return nil
	}
	return browser
}

//Submit.  提交
//formId   表单ID
//data     数据
func (k *K3Cloud) Submit(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.SubmitApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//Save.  保存
//formId 表单ID
//data   数据
func (k *K3Cloud) Save(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.SaveApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//BatchSave.  批量保存
//formId      表单ID
//data        数据
func (k *K3Cloud) BatchSave(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.BatchSaveApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//Audit.  审核
//formId  表单ID
//data    数据
func (k *K3Cloud) Audit(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.AuditApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//UnAudit.  反审核
//formId    表单ID
//data      数据
func (k *K3Cloud) UnAudit(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.UnAuditApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//View.  详情
//formId 查询表单ID
//data   查询数据
func (k *K3Cloud) View(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.ViewApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

// ExecuteBillQuery. 单据查询
func (k *K3Cloud) ExecuteBillQuery(ctx context.Context, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.ExecuteBillQueryApi
	var postData = &object.HashMap{
		"data": data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//Operation. 操作
//formId     查询表单ID
//opNumber   操作标识
//data       查询数据
func (k *K3Cloud) Operation(ctx context.Context, formId, opNumber string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.ExecuteOperationApi
	var postData = &object.HashMap{
		"formid":   formId,
		"opNumber": opNumber,
		"data":     data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//Push.  下推
//formId 表单ID
//data   数据
func (k *K3Cloud) Push(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.PushApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//Draft. 暂存
//formId 表单ID
//data   数据
func (k *K3Cloud) Draft(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.DraftApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//Delete. 删除
//formId  表单ID
//data    数据
func (k *K3Cloud) Delete(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.DeleteApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//Allocate. 分配
//formId    表单ID
//data      数据
func (k *K3Cloud) Allocate(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.AllocateApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//FlexSave. 弹性域保存
//formId    表单ID
//data      数据
func (k *K3Cloud) FlexSave(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.DeleteApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//SendMsg. 发送消息
//data     数据
func (k *K3Cloud) SendMsg(ctx context.Context, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.SendMsgApi
	var postData = &object.HashMap{
		"data": data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//Disassembly. 拆单
//formId       表单ID
//data         数据
func (k *K3Cloud) Disassembly(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.DisassemblyApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//WorkflowAudit. 工作流审批
//data           数据
func (k *K3Cloud) WorkflowAudit(ctx context.Context, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.WorkflowAuditApi
	var postData = &object.HashMap{
		"data": data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//QueryGroupInfo. 查询分组信息
//data            数据
func (k *K3Cloud) QueryGroupInfo(ctx context.Context, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.QueryGroupInfoApi
	var postData = &object.HashMap{
		"data": data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}

//GroupDelete. 分组删除
//data         数据
func (k *K3Cloud) GroupDelete(ctx context.Context, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + kernel.GroupDeleteApi
	var postData = &object.HashMap{
		"data": data,
	}
	return k.newBrowser().PostJson(ctx, url, postData)
}
