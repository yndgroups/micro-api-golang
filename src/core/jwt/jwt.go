package jwt

import (
	"core/config"
	"core/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	golangJwt "github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
)

type TokenInfo struct {
	UserId        string `json:"userId"`        // 用户ID
	AppId         string `json:"appId"`         // 应用ID
	LoginRedisKey string `json:"loginRedisKey"` //登陆的rediskey
	ConfId        string `json:"confId"`        // 配置id
	PowerKey      string `json:"powerKey"`      // 权限值
	RegType       int64  `json:"regType"`       // 注册类型
	SubId         string `json:"subId"`         // 用户openId
}

// userId 用户id
// appId 应用id
// regType 来源分类(1:后台端 2:前端 3: 微信公众号 4:微信小程序）
func CreateToken(tk TokenInfo) (string, error) {
	// userId string, appId string, regType int32
	rules := golangJwt.NewWithClaims(golangJwt.SigningMethodHS256, golangJwt.MapClaims{
		"userId":  utils.ConfoundBase64EncodeToString(tk.UserId),
		"subId":   utils.ConfoundBase64EncodeToString(tk.SubId),
		"appId":   utils.ConfoundBase64EncodeToString(tk.AppId),
		"regType": tk.RegType,
		"exp":     time.Now().Add(time.Minute * 60 * 24 * 7 * 365 * 10).Unix(),
	})
	token, err := rules.SignedString([]byte(config.Global.ServerConfig().Secret))
	if err != nil {
		return "", err
	}
	return token, err
}

// 解析token
func ParseToken(token string) (string, error) {
	if !strings.Contains(token, config.Global.ServerConfig().TokenPrefix) {
		return "", errors.New("token格式不正确")
	}
	claim, err := golangJwt.Parse(strings.Replace(token, config.Global.ServerConfig().TokenPrefix, "", 1), func(token *golangJwt.Token) (any, error) {
		return []byte(config.Global.ServerConfig().Secret), nil
	})
	if err != nil {
		return "", errors.New("token解析失败")
	}
	if claim.Claims.(golangJwt.MapClaims)["userId"] != nil {
		return utils.ConfoundBase64DecodeString(claim.Claims.(golangJwt.MapClaims)["userId"].(string)), nil
	}
	return "", errors.New("非法伪造token")
}

// 解析token中的UserId
func GetUserIdByToken(token string) (string, error) {
	if !strings.Contains(token, config.Global.ServerConfig().TokenPrefix) {
		return "", errors.New("token格式不正确")
	}
	claim, err := golangJwt.Parse(strings.Replace(token, config.Global.ServerConfig().TokenPrefix, "", 1), func(token *golangJwt.Token) (any, error) {
		return []byte(config.Global.ServerConfig().Secret), nil
	})
	if err != nil {
		if err.Error() == "Token is expired" {
			return "", errors.New("token已失效")
		}
		if err.Error() == "token contains an invalid number of segments" {
			return "", errors.New("token不正确")
		}
		return "", errors.New(err.Error())
	}
	if claim.Claims.(golangJwt.MapClaims)["userId"] != nil {
		return utils.ConfoundBase64DecodeString(claim.Claims.(golangJwt.MapClaims)["userId"].(string)), nil
	}
	return "非法伪造token", errors.New("非法伪造token")
}

func ParseTokenToLoginInfo(token string) (*TokenInfo, error) {
	if !strings.Contains(token, config.Global.ServerConfig().TokenPrefix) {
		return nil, errors.New("token格式不正确")
	}
	if claim, err := golangJwt.Parse(strings.Replace(token, config.Global.ServerConfig().TokenPrefix, "", 1), func(token *golangJwt.Token) (any, error) {
		return []byte(config.Global.ServerConfig().Secret), nil
	}); err != nil {
		if err.Error() == "Token is expired" {
			return nil, errors.New("token已失效")
		}
		if err.Error() == "token contains an invalid number of segments" {
			return nil, errors.New("token不正确")
		}
		return nil, errors.New("非法伪造token")
	} else {
		var tokenInfo TokenInfo
		// 解析用户id
		if claim.Claims.(golangJwt.MapClaims)["userId"] != nil {
			tokenInfo.UserId = utils.ConfoundBase64DecodeString(claim.Claims.(golangJwt.MapClaims)["userId"].(string))
		} else {
			return nil, errors.New("token解析userId失败")
		}
		// 解析token是否有appId
		if claim.Claims.(golangJwt.MapClaims)["appId"] != nil {
			tokenInfo.AppId = utils.ConfoundBase64DecodeString(claim.Claims.(golangJwt.MapClaims)["appId"].(string))
		} else {
			return nil, errors.New("token解析appId失败")
		}
		// 解析是否包含注册类型
		if claim.Claims.(golangJwt.MapClaims)["regType"] != nil {
			regType := fmt.Sprintf("%v", claim.Claims.(golangJwt.MapClaims)["regType"].(float64))
			if regType != "" {
				tokenInfo.RegType = cast.ToInt64(regType)
			}
		}
		if claim.Claims.(golangJwt.MapClaims)["subId"] != nil {
			tokenInfo.SubId = utils.ConfoundBase64DecodeString(claim.Claims.(golangJwt.MapClaims)["subId"].(string))
		}
		// 用户的redis登陆的key
		tokenInfo.LoginRedisKey = fmt.Sprintf("%s%s_%s", config.Global.ServerConfig().LoginPrefix, tokenInfo.AppId, tokenInfo.UserId)
		// 用户的redis权限值key
		tokenInfo.PowerKey = fmt.Sprintf("%s%s_%s", config.Global.ServerConfig().PowerPrefix, tokenInfo.AppId, tokenInfo.UserId)
		return &tokenInfo, nil
	}
}

// 解析微信返回数据
func Base64Encoding(str string) string { //Base64编码
	res, _ := base64.StdEncoding.DecodeString(str)
	return string(res)
}
