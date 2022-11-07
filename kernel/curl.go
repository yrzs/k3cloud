package kernel

/**
1.可设置代理
2.可设置 cookie
3.自动保存并应用响应的 cookie
4.自动为重新向的请求添加 cookie
*/

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/yrzs/k3cloud/object"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Browser struct {
	cookies []*http.Cookie
	client  *http.Client
}

//初始化
func NewBrowser() *Browser {
	hc := &Browser{}
	hc.client = &http.Client{}
	//为所有重定向的请求增加cookie
	hc.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > 0 {
			for _, v := range hc.GetCookie() {
				req.AddCookie(v)
			}
		}
		return nil
	}
	return hc
}

//设置代理地址
func (b *Browser) SetProxyUrl(proxyUrl string) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}
	transport := &http.Transport{Proxy: proxy}
	b.client.Transport = transport
}

//设置请求cookie
func (b *Browser) AddCookie(cookies []*http.Cookie) {
	b.cookies = append(b.cookies, cookies...)
}

//获取当前所有的cookie
func (b *Browser) GetCookie() []*http.Cookie {
	return b.cookies
}

//发送PostJson请求
func (b *Browser) PostJson(requestUrl string, params *object.HashMap) (*object.HashMap, error) {
	postData, _ := object.JsonEncode(params)
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(postData))
	request.Header.Set("Content-Type", "application/json")
	b.setRequestCookie(request)
	response, err := b.client.Do(request)
	if err != nil {
		return nil, errors.New("http post json fail")
	}
	defer response.Body.Close()
	//保存响应的 cookie
	respCks := response.Cookies()
	b.cookies = append(b.cookies, respCks...)
	data, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return nil, errors.New("http read io result fail")
	}
	m, ok := gjson.Parse(string(data)).Value().(map[string]interface{})
	if ok {
		o := object.HashMap(m)
		return &o, nil
	} else {
		mm, okk := gjson.Parse(string(data)).Value().([]interface{})
		if okk {
			var oo = object.HashMap{
				"data": mm,
			}
			return &oo, nil
		}
		return nil, nil
	}
}

//为请求设置 cookie
func (b *Browser) setRequestCookie(request *http.Request) {
	for _, v := range b.cookies {
		request.AddCookie(v)
	}
}

//参数 encode
func (b *Browser) encodeParams(params map[string]string) string {
	paramsData := url.Values{}
	for k, v := range params {
		paramsData.Set(k, v)
	}
	return paramsData.Encode()
}

//参数 json
func (b *Browser) jsonParams(params map[string]interface{}) string {
	j, _ := json.Marshal(params)
	return string(j)
}
