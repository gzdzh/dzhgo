package util

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var ProxyOpen bool
var ProxyURL string

func init() {

	ctx := gctx.New()

	proxy_open, err := g.Cfg().Get(ctx, "http.proxy_open")
	if err != nil {
		g.Log().Error(ctx, err)
	}
	ProxyOpen = proxy_open.Bool()

	proxy_url, err := g.Cfg().Get(ctx, "http.proxy_url")
	if err != nil {
		g.Log().Error(ctx, err)
	}

	ProxyURL = proxy_url.String()
}

func HttpGet(ctx context.Context, url string, header map[string]string, data g.Map, result interface{}, cookies ...map[string]string) error {

	g.Log().Debugf(ctx, "HttpGet url: %s, header: %+v, data: %+v", url, header, data)

	client := g.Client().Timeout(60 * time.Second)
	if header != nil {
		client.SetHeaderMap(header)
	}

	if len(cookies) > 0 {
		client.Cookie(cookies[0])
	}

	response, err := client.Get(ctx, url, data)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	defer func() {
		err = response.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}()

	bytes := response.ReadAll()
	g.Log().Debugf(ctx, "HttpGet url: %s, header: %+v, data: %+v, response: %s", url, header, data, string(bytes))

	if bytes != nil && len(bytes) > 0 {
		err = gjson.Unmarshal(bytes, result)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	}

	return nil
}

func HttpPost(ctx context.Context, url string, header map[string]string, data, result interface{}) error {

	// g.Log().Debugf(ctx, "HttpPost url: %s, header: %+v, data: %+v", url, header, data)

	client := g.Client().Timeout(60 * time.Second)
	if header != nil {
		client.SetHeaderMap(header)
	}

	if ProxyOpen && len(ProxyURL) > 0 {
		client.SetProxy(ProxyURL)
	}

	response, err := client.ContentJson().Post(ctx, url, data)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	defer func() {
		err = response.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}()

	bytes := response.ReadAll()
	g.Log().Debugf(ctx, "HttpPost url: %s, header: %+v, data: %+v, response: %s", url, header, data, string(bytes))

	if bytes != nil && len(bytes) > 0 {
		err = gjson.Unmarshal(bytes, result)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	}

	return nil
}

func HttpDownloadFile(ctx context.Context, fileURL string, useProxy ...bool) []byte {

	g.Log().Debugf(ctx, "HttpDownloadFile fileURL: %s", fileURL)

	client := g.Client().Timeout(600 * time.Second)

	transport := &http.Transport{}

	if ProxyOpen && len(ProxyURL) > 0 && (len(useProxy) == 0 || useProxy[0]) {

		g.Log().Debugf(ctx, "HttpDownloadFile ProxyURL: %s", ProxyURL)

		proxyUrl, err := url.Parse(ProxyURL)
		if err != nil {
			g.Log().Error(ctx, err)
		}

		transport.Proxy = http.ProxyURL(proxyUrl)
		client.Transport = transport
	}

	return client.GetBytes(ctx, fileURL)
}

func GetProxy(ctx context.Context) func(*http.Request) (*url.URL, error) {

	var proxy func(*http.Request) (*url.URL, error)

	if ProxyOpen && len(ProxyURL) > 0 {

		g.Log().Debugf(ctx, "ProxyURL: %s", ProxyURL)

		proxyURL, err := url.Parse(ProxyURL)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil
		}

		return http.ProxyURL(proxyURL)
	}

	return proxy
}

func GetProxyTransport(ctx context.Context) *http.Transport {

	transport := &http.Transport{}

	transport.Proxy = GetProxy(ctx)

	return transport
}
