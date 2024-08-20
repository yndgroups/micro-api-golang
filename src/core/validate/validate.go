package validate

import (
	"core/utils"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

// Validated 校验请求参数模型
func Validated[T any](modelParams T) (map[string]interface{}, error) {
	// 校验分页参数
	err := validator.New().Struct(modelParams)
	info := make(map[string]interface{})
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			// 将首字符变为小写保持与前端一致
			field := utils.FirstCharacterLowercase(e.Field())
			switch e.Tag() {
			case "required":
				fmt.Println("e.Value() == 0", e.Value() == 0)
				switch e.Value() {
				case 0:
					if e.Field() == "PageIndex" || e.Field() == "PageSize" {
						info[field] = "参数校验失败,请传入必须大于0且必须为数字类型的值"
					}
					// else {
					//	info[field] = fmt.Sprintf("参数值%v校验失败, %v", e.Value(), "请传入大于0的数值")
					//}
					break
				case "", "null", "undefined", "0":
					info[field] = fmt.Sprintf("参数值%v校验失败, %v", e.Value(), "必传参数不能为空")
					break
				}
				break
			case "min":
				fmt.Println(fmt.Printf("%v min => ", e))
				info[field] = fmt.Sprintf("参数值%v小于设定的值%v", e.Value(), e.Param())
				break
			case "max":
				fmt.Println(fmt.Printf("%v max => ", e))
				info[field] = fmt.Sprintf("参数值%v大于设定的值%v", e.Value(), e.Param())
				break
			}
			fmt.Println("Namespace:", e.Namespace())
			fmt.Println("Field:", e.Field())
			fmt.Println("StructNamespace:", e.StructNamespace())
			fmt.Println("StructField:", e.StructField())
			fmt.Println("Tag:", e.Tag())
			fmt.Println("ActualTag:", e.ActualTag())
			fmt.Println("Kind:", e.Kind())
			fmt.Println("Type:", e.Type())
			fmt.Println("Value:", e.Value())
			fmt.Println("Param:", e.Param())
		}
		if len(info) > 0 {
			return info, errors.New("请求参数效验失败")
		} else {
			return info, nil
		}
	} else {
		return info, nil
	}
}

// ValidatedAuthorized 校验借口权限合法性
func ValidatedAuthorized(authorizedValue string, apiAuthorizedValue string) bool {
	return strings.Contains(authorizedValue, apiAuthorizedValue)
}
