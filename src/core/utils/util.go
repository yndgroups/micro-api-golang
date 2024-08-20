package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// HmacSha512 返回一个加密的字符串
// @param str 要加密的字符串
// @param secret加密的密钥
// @author yangDaqiong
// @date  2021-10-24
func HmacSha512(str string, secret string) string {
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// SHa1 sha1加密
// @param str 加密的字符串
// @author yangDaqiong
// @date  2021-01-03
func SHa1(str string) string {
	o := sha1.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}

// GetTrueIp 获取真实内网IP地址或者错误信息
// @author yangDaqiong
// @date  2022-01-31
func GetTrueIp() (net.IP, error) {
	faces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, face := range faces {
		if face.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if face.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		adders, err := face.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range adders {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

// getIpFromAddr 获取id地址
// @author yangDaqiong
// @date  2022-01-31
func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		// 不是一个ipv4地址
		return nil
	}
	return ip
}

// GetIpV4 获取外网ipv4
// @author yangDaqiong
// @date  2022-01-31
func GetIpV4() string {
	resp, err := http.Get("https://ipv4.netarm.com")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}

// GetIpV6 获取外网ipv6
// @author yangDaqiong
// @date  2022-01-31
func GetIpV6() string {
	resp, err := http.Get("https://ipv6.netarm.com")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}

// GetIpV6Lan GO语言可以通过直接读取网卡的方式获取到外网ipv6。
// @author yangDaqiong
// @date  2022-01-31
func GetIpV6Lan() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		ipv6 := regexp.MustCompile(`(\w+:){7}\w+`).FindString(addr.String())
		if strings.Count(ipv6, ":") == 7 {
			return ipv6
		}
	}
	return ""
}

// ValidateErrorHandle ValidatErrorHandle 验证参数
// @author yangDaqiong
// @date  2022-01-31
func ValidateErrorHandle(err error) (map[string]interface{}, error) {
	info := make(map[string]interface{})
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(" ------------------ > ValidateParams")
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
			switch e.Tag() {
			case "required":
				info[e.Field()] = fmt.Sprintf("必传参数值%v校验失败%v", e.Value(), e.Param())
			case "min":
				info[e.Field()] = fmt.Sprintf("参数值%v小于设定的值%v", e.Value(), e.Param())
			case "max":
				info[e.Field()] = fmt.Sprintf("参数值%v大于设定的值%v", e.Value(), e.Param())
			}
		}
		return info, errors.New("verification failed")
	} else {
		return info, nil
	}
}

// FirstCharacterUppercase 首字符大写
func FirstCharacterUppercase(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]

}

// FirstCharacterLowercase 首字符小写
func FirstCharacterLowercase(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToLower(str[:1]) + str[1:]
}

// GetAvailablePort 获取可用端口
func GetAvailablePort() (uint64, error) {
	address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:0", "0.0.0.0"))
	if err != nil {
		return 0, err
	}
	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return uint64(listener.Addr().(*net.TCPAddr).Port), nil
}

// IsPortAvailable 判断端口是否可用
func IsPortAvailable(port int) bool {
	address := fmt.Sprintf("%s:%d", "0.0.0.0", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return false
	}
	defer listener.Close()
	return true
}

// GetRandom 随机获取一个服务的获取某个区间的随机数
func GetRandom(min, max int) int {
	num := rand.Intn(max-min) + min
	return num
}

// Base64EncodeToString base64加密
func Base64EncodeToString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64DecodeString base64解密
func Base64DecodeString(str string) string {
	s, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(s)
}

// ConfoundBase64EncodeToString 混淆base64加密
func ConfoundBase64EncodeToString(str string) string {
	l := len(str)
	if str == "" && l < 2 {
		return ""
	}
	bs := Base64EncodeToString(str)
	strArr := strings.Split(bs, "")
	ts := ""
	for i, v := range strArr {
		if i == 1 || i == 10 || i == 17 || i == 23 {
			ts += (v + randomString(1))
		} else {
			ts += v
		}
	}
	// println(fmt.Sprintf("加密前: %v, 加密前长度: %d, \n加密混淆后: %v, 加密混淆后长度为:%d", bs, len(bs), ts, len(ts)))
	return ts
}

// ConfoundBase64DecodeString 混淆base64解密
func ConfoundBase64DecodeString(str string) string {
	l := len(str)
	if str == "" || l < 3 {
		return ""
	}
	strArr := strings.Split(str, "")
	ts := ""
	for i, v := range strArr {
		if i != 2 && i != 12 && i != 20 && i != 27 {
			ts += v
		}
		// else {
		// 	println(v)
		// }
	}
	// println(fmt.Sprintf("解密前长度:%d, 混淆解密字符串:%s,\n解密字符串:%s, 解密字符串长度:%d", l, str, ts, len(ts)))
	return Base64DecodeString(ts)
}

// randomString 随机字符串
// length 长度
func randomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789"
	sb := strings.Builder{}
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
