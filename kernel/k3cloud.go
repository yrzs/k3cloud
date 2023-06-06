package kernel

import (
	"context"
	"github.com/yrzs/k3cloud/object"
)

// K3Config. 金蝶云星空账号地址
type K3Config struct {
	Host     string
	AccID    string
	Username string
	Password string
	LcID     int
}

// K3Cloud. application
type K3Cloud struct {
	Config *K3Config
	Client *Browser
}

//Submit.  提交
//formId   表单ID
//data     数据
func (k *K3Cloud) Submit(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + SubmitApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//Save.  保存
//formId 表单ID
//data   数据
func (k *K3Cloud) Save(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + SaveApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//BatchSave.  批量保存
//formId      表单ID
//data        数据
func (k *K3Cloud) BatchSave(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + BatchSaveApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//Audit.  审核
//formId  表单ID
//data    数据
func (k *K3Cloud) Audit(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + AuditApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//UnAudit.  反审核
//formId    表单ID
//data      数据
func (k *K3Cloud) UnAudit(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + UnAuditApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//View.  详情
//formId 查询表单ID
//data   查询数据
func (k *K3Cloud) View(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + ViewApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

// ExecuteBillQuery. 单据查询
func (k *K3Cloud) ExecuteBillQuery(ctx context.Context, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + ExecuteBillQueryApi
	var postData = &object.HashMap{
		"data": data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//Operation. 操作
//formId     查询表单ID
//opNumber   操作标识
//data       查询数据
func (k *K3Cloud) Operation(ctx context.Context, formId, opNumber string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + ExecuteOperationApi
	var postData = &object.HashMap{
		"formid":   formId,
		"opNumber": opNumber,
		"data":     data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//Push.  下推
//formId 表单ID
//data   数据
func (k *K3Cloud) Push(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + PushApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//Draft. 暂存
//formId 表单ID
//data   数据
func (k *K3Cloud) Draft(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + DraftApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//Delete. 删除
//formId  表单ID
//data    数据
func (k *K3Cloud) Delete(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + DeleteApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//Allocate. 分配
//formId    表单ID
//data      数据
func (k *K3Cloud) Allocate(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + AllocateApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//FlexSave. 弹性域保存
//formId    表单ID
//data      数据
func (k *K3Cloud) FlexSave(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + DeleteApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//SendMsg. 发送消息
//data     数据
func (k *K3Cloud) SendMsg(ctx context.Context, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + SendMsgApi
	var postData = &object.HashMap{
		"data": data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//Disassembly. 拆单
//formId       表单ID
//data         数据
func (k *K3Cloud) Disassembly(ctx context.Context, formId string, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + DisassemblyApi
	var postData = &object.HashMap{
		"formid": formId,
		"data":   data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//WorkflowAudit. 工作流审批
//data           数据
func (k *K3Cloud) WorkflowAudit(ctx context.Context, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + WorkflowAuditApi
	var postData = &object.HashMap{
		"data": data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//QueryGroupInfo. 查询分组信息
//data            数据
func (k *K3Cloud) QueryGroupInfo(ctx context.Context, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + QueryGroupInfoApi
	var postData = &object.HashMap{
		"data": data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}

//GroupDelete. 分组删除
//data         数据
func (k *K3Cloud) GroupDelete(ctx context.Context, data *object.HashMap) (*object.HashMap, error) {
	url := k.Config.Host + GroupDeleteApi
	var postData = &object.HashMap{
		"data": data,
	}
	return k.Client.PostJson(ctx, k.Config, url, postData)
}
