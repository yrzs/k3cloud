package kernel

/**
1.可设置代理
2.可设置 cookie
3.自动保存并应用响应的 cookie
4.自动为重新向的请求添加 cookie
*/

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/yrzs/k3cloud/object"
	resp "github.com/yrzs/k3cloud/response"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Browser.
type Browser struct {
	cookies []*http.Cookie
	client  *http.Client
}

type httpTransport struct {
	*http.Transport
}

func (t *httpTransport) SetProxyUrl(proxyUrl string) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}
	t.Transport.Proxy = proxy
}

// NewBrowserWithTransport.
func NewBrowserWithTransport() *Browser {
	hc := &Browser{}
	hc.client = &http.Client{
		Transport: &httpTransport{
			&http.Transport{
				MaxIdleConns:        100,              // 最大空闲连接数
				MaxIdleConnsPerHost: 10,               // 每个目标主机最大空闲连接数
				IdleConnTimeout:     90 * time.Second, // 空闲连接超时时间
				DisableKeepAlives:   true,             // 关闭 keep-alive
			},
		},
	}
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

// SetProxyUrl. 设置代理地址
func (b *Browser) SetProxyUrl(proxyUrl string) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}
	transport := &http.Transport{Proxy: proxy}
	b.client.Transport = transport
}

// AddCookie. 设置请求cookie
func (b *Browser) AddCookie(cookies []*http.Cookie) {
	b.cookies = append(b.cookies, cookies...)
}

// appendCookie. 追加请求cookie
func (b *Browser) appendCookie(cookies []*http.Cookie) {
	if len(b.cookies) > 2 {
		b.cookies = b.cookies[:len(b.cookies)-1]
	}
	b.cookies = append(b.cookies, cookies...)
}

// GetCookie. 获取当前所有的cookie
func (b *Browser) GetCookie() []*http.Cookie {
	return b.cookies
}

// setRequestCookie. 为请求设置cookie
func (b *Browser) setRequestCookie(request *http.Request) {
	for _, v := range b.cookies {
		request.AddCookie(v)
	}
}

// initLogin.
func (b *Browser) InitLogin(c *K3Config) error {
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
	res, _ := b.PostJson(ctx, c, c.Host+LoginApi, data)
	var k3Response = &resp.LoginResponse{}
	e := object.HashMapToStructure(res, k3Response)
	if e != nil {
		return errors.New("k3 cloud login fail")
	}
	if k3Response.LoginResultType == 0 {
		return errors.New(k3Response.Message)
	}
	return nil
}

// PostJson. 发送Post请求Json格式数据
func (b *Browser) PostJson(ctx context.Context, c *K3Config, requestUrl string, params *object.HashMap) (*object.HashMap, error) {
	// 设置日志前缀和格式
	log.SetPrefix("[Info]")
	log.SetFlags(log.Ldate | log.Ltime)
	postData, _ := object.JsonEncode(params)
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(postData))
	// 携带ctx
	request = request.WithContext(ctx)
	request.Header.Set("Content-Type", "application/json")
	b.setRequestCookie(request)
	response, err := b.client.Do(request)
	if err != nil {
		return nil, errors.New("http post json fail")
	}
	defer response.Body.Close()
	//保存响应的 cookie
	b.cookies = response.Cookies()
	data, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return nil, errors.New("http read io result fail")
	}
	//fmt.Println(string(data))
	var res object.HashMap
	res, ok := gjson.Parse(string(data)).Value().(map[string]interface{})
	if !ok {
		var k3Response [][]*resp.K3Response
		if e = json.Unmarshal(data, &k3Response); e == nil {
			responseStatus := k3Response[0][0].Result.Status
			if responseStatus.ErrorCode == http.StatusInternalServerError && responseStatus.Errors[0].Message == "会话信息已丢失，请重新登录" {
				log.Println("没登录!")
				if e = b.InitLogin(c); e == nil {
					log.Println("登录成功!")
					log.Println("重放请求!")
					return b.PostJson(ctx, c, requestUrl, params)
				}
			}
		}
		mm, okk := gjson.Parse(string(data)).Value().([]interface{})
		if okk {
			res = object.HashMap{
				"data": mm,
			}
		}
	}
	return &res, nil
}
