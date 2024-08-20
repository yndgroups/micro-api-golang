package utils

// AccessToken 微信AccessToken信息
// "{\"access_token\":\"53_qjyVwyO-aAbNo0ZjT1PUynSnWdGHbaKK8kjJVGg17aiCUMMC5X7XvRT5vRuJUuT6asgMfgQwaBWbOz_zj6gs5vavWvzY14ZICqV-YhAr2oF6EGiC9mwsybGBLfj4ssfqrMbFwl99A-n7unykNZZjAJAQCC\",\"expires_in\":7200}"
// {"errcode":40164,"Msg":"invalid ip 125.33.201.21 ipv6 ::ffff:125.33.201.21, not in whitelist rid: 6207df1a-08ef8516-2a6c18c6"}
type AccessToken struct {
	// AccessToken值
	AccessToken string `json:"access_token"`
	// 过期时间
	ExpiresIn int `json:"expires_in"`
	// 错误编码
	ErrCode int `json:"errcode"`
	// 错误信息
	Msg string `json:"Msg"`
}
