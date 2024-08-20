package ksii

import (
	"errors"
	"reflect"

	"github.com/shopspring/decimal"
)

// Add 加法
func Add(d1 interface{}, d2 interface{}) (decimal.Decimal, error) {
	m, me := number(d1)
	if me != nil {
		return m, me
	}
	n, ne := number(d2)
	if ne != nil {
		return m, ne
	}
	return m.Add(n), nil
}

// Sub 减法
func Sub(d1 interface{}, d2 interface{}) (decimal.Decimal, error) {
	m, me := number(d1)
	if me != nil {
		return m, me
	}
	n, ne := number(d2)
	if ne != nil {
		return m, ne
	}
	return m.Sub(n), nil
}

// Mul 乘法
func Mul(d1 interface{}, d2 interface{}) (decimal.Decimal, error) {
	m, me := number(d1)
	if me != nil {
		return m, me
	}
	n, ne := number(d2)
	if ne != nil {
		return m, ne
	}
	return m.Mul(n), nil
}

// Div 除法
func Div(d1 interface{}, d2 interface{}) (decimal.Decimal, error) {
	m, me := number(d1)
	if me != nil {
		return m, me
	}
	n, ne := number(d2)
	if ne != nil {
		return m, ne
	}
	return m.Div(n), nil
}

// number 将 interface 进行类型转换成 decimal
func number(num interface{}) (decimal.Decimal, error) {
	d := decimal.Decimal{}
	switch num.(type) {
	case int:
		d = decimal.NewFromInt(int64(num.(int)))
		break
	case int64:
		d = decimal.NewFromInt(num.(int64))
		break
	case string:
		if sd, e := decimal.NewFromString(num.(string)); e != nil {
			return sd, errors.New("参数转换decimal失败，请开发人员进行排查！")
		} else {
			d = sd
		}
		break
	case float32:
		d = decimal.NewFromFloat32(num.(float32))
		break
	case float64:
		d = decimal.NewFromFloat(num.(float64))
		break
	default:
		t := reflect.TypeOf(num)
		if t.String() != "decimal.Decimal" {
			return d, errors.New("未知的数据类型")
		}
	}
	return d, nil
}

// ChangeDecimal 改变decimal的类型
func ChangeDecimal(num interface{}) (decimal.Decimal, error) {
	d := decimal.Decimal{}
	switch num.(type) {
	case int:
		d = decimal.NewFromInt(int64(num.(int)))
		break
	case int8:
		d = decimal.NewFromInt(int64(num.(int8)))
		break
	case int16:
		d = decimal.NewFromInt(int64(num.(int16)))
		break
	case int32:
		d = decimal.NewFromInt(int64(num.(int32)))
		break
	case int64:
		d = decimal.NewFromInt(num.(int64))
		break
	case string:
		if sd, e := decimal.NewFromString(num.(string)); e != nil {
			return sd, errors.New("参数转换decimal失败，请开发人员进行排查！")
		} else {
			d = sd
		}
		break
	case float32:
		d = decimal.NewFromFloat32(num.(float32))
		break
	case float64:
		d = decimal.NewFromFloat(num.(float64))
		break
	default:
		t := reflect.TypeOf(num)
		if t.String() != "decimal.Decimal" {
			return d, errors.New("未知的数据类型")
		}
	}
	return d, nil
}
