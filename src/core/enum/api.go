package enum

type ApiModel struct {
	ApiName    string `json:"apiName"`
	Permission string `json:"permission"`
}

// ApiEnum 构造函数
func ApiEnum(apiName string, permission string) ApiModel {
	return ApiModel{
		ApiName:    apiName,
		Permission: permission,
	}
}
