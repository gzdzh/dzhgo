package config

import "github.com/gogf/gf/v2/frame/g"

func GetWxCon() (data g.Map) {

	type Config struct {
		RequestUrl string `json:"requestUrl"`
		Appid      string `json:"appid"`
		Secret     string `json:"secret"`
		JsCode     string `json:"js_code"`
		GrantType  string `json:"grant_type"`
	}

	data = g.Map{
		"requestUrl": "https://api.weixin.qq.com/sns/jscode2session",
		"appid":      "wx61d904e89976921e",
		"secret":     "adc9faa4a1abc916556b107dba43eaac",
		"jsCode":     "",
		"grantType":  "authorization_code",
	}

	return
}
