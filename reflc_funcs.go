package c3po

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func convert(v *reflect.Value, t reflect.Type) bool {
	defer try()
	if v.Kind() == t.Kind() {
		return true
	}
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		switch v.Kind() {
		case reflect.String:
			i, err := strconv.ParseFloat(v.Interface().(string), 64)
			if err != nil {
				return false
			}
			*v = reflect.ValueOf(i).Convert(t)
		case
			reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8,
			reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			*v = v.Convert(t)
		case reflect.Float32, reflect.Float64:
			*v = v.Convert(t)
		default:
			return false
		}
	case
		reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8,
		reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		switch v.Kind() {
		case t.Kind():
			return true
		case reflect.String:
			val, err := strconv.ParseFloat(v.Interface().(string), 64)
			if err != nil {
				return false
			}
			*v = reflect.ValueOf(val).Convert(t)
		case reflect.Float32, reflect.Float64:
			*v = v.Convert(t)
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8,
			reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			*v = v.Convert(t)
		default:
			return false
		}
	case reflect.Bool:
		if v.Kind() != reflect.String {
			return false
		}

		b := strings.ToLower(v.Interface().(string))
		if b == "true" {
			*v = reflect.ValueOf(true)
		} else if b == "false" {
			*v = reflect.ValueOf(false)
		} else {
			return false
		}
	case reflect.String:
		str := v.Interface()
		if str == nil {
			str = ""
		} else if s, ok := str.(fmt.Stringer); ok {
			str = s.String()
		}
		*v = reflect.ValueOf(fmt.Sprint(str))
	}
	return true
}

// try convert and set a value of v in rcv
func SetReflectValue(rcv reflect.Value, v reflect.Value) bool {
	defer try()
	if v.IsValid() {
		c := convert(&v, rcv.Type())
		if c {
			if rcv.Kind() == reflect.Pointer && v.Kind() != reflect.Pointer {
				v = v.Addr()
			} else if rcv.Kind() != reflect.Pointer && v.Kind() == reflect.Pointer {
				v = v.Elem()
			}
			rcv.Set(v)
			return true
		}
	}
	return false
}
