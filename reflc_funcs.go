package c3po

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func reflectOf(v any) reflect.Value {
	var rv = reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		if rv2 := rv.Elem(); rv2.Kind() == reflect.Pointer {
			rv = rv2
		}
	}
	return rv
}

func GetReflectTypeElem(t reflect.Type) reflect.Type {
	for c := 0; c < 10; c++ {
		if t.Kind() != reflect.Ptr {
			break
		}
		t = t.Elem()
	}
	return t
}

func GetReflectElem(r reflect.Value) reflect.Value {
	for c := 0; c < 10; c++ {
		if r.Kind() != reflect.Ptr {
			break
		}
		r = r.Elem()
	}
	return r
}

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

func SetReflectValue(r reflect.Value, v reflect.Value) bool {
	defer try()
	if v.IsValid() {
		c := convert(&v, r.Type())
		if c {
			if r.Kind() == reflect.Pointer && v.Kind() != reflect.Pointer {
				v = v.Addr()
			} else if r.Kind() != reflect.Pointer && v.Kind() == reflect.Pointer {
				v = v.Elem()
			}
			r.Set(v)
			return true
		}
	}
	return false
}
