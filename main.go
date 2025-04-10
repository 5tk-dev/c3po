package c3po

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

func init() {
	once.Do(initRules)
	once.Do(initRegistry)
}

// return a []map, map or a simple value. Depende doq vc passou como argumento
func encode(v any) (any, error) {
	if v == nil {
		return nil, errors.New("'v' is nil")
	}
	if r, ok := v.(reflect.Value); ok {
		return encode(r.Interface())
	}

	errs := bytes.NewBufferString("")
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		rv2 := rv.Elem()
		if rv2.IsValid() {
			rv = rv2
		}
	}
	rt := rv.Type()
	switch rv.Kind() {
	default:
		return v, nil
	case reflect.Func:
		return "what the fuck, it's a func?????", errors.New("unsupported type" + rv.Kind().String())
	case reflect.Pointer:
		if rv.Elem().IsValid() {
			return Encode(rv.Elem())
		} else if rv.CanInterface() {
			return rv.Interface(), nil
		} else {
			return nil, errors.New("sei la, só não consegui pegar o valor de 'v'")
		}
	case reflect.Struct:
		d := map[string]any{}
		for i := 0; i < rv.NumField(); i++ {
			f := rv.Field(i)
			if !f.IsValid() || !f.CanInterface() {
				continue
			}
			ft := rt.Field(i)
			if fv, err := Encode(f.Interface()); err == nil {
				d[strings.ToLower(ft.Name)] = fv
			} else {
				errs.WriteString(err.Error())
			}
		}
		return d, nil
	case reflect.Slice, reflect.Array:
		return encodeSlice(v)
	case reflect.Map:
		return encodeMap(v)
	}
}

func encodeSlice(v any) (any, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer && rv.IsValid() {
		if rv.Elem().IsValid() {
			rv = rv.Elem()
		}
	}
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, errors.New("'v' not is a Slice or Array ")
	}

	errs := bytes.NewBufferString("")
	d := reflect.MakeSlice(reflect.TypeOf([]any{}), rv.Len(), rv.Cap())
	for i := 0; i < rv.Len(); i++ {
		f := rv.Index(i)
		if !f.IsValid() || !f.CanInterface() {
			continue
		}
		if fv, err := Encode(f.Interface()); err == nil {
			d.Index(i).Set(reflect.ValueOf(fv))
		} else {
			errs.WriteString(err.Error())
		}
	}
	data := d.Interface()
	if errs.Len() > 0 {
		return data, errors.New(errs.String())
	}
	return data, nil
}

func encodeMap(v any) (any, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer && rv.IsValid() {
		if rv.Elem().IsValid() {
			rv = rv.Elem()
		}
	}
	if rv.Kind() != reflect.Map {
		return nil, errors.New("'v' not is a Map")
	}

	errs := bytes.NewBufferString("")
	m := reflect.MakeMap(reflect.TypeOf(map[string]any{}))

	for _, key := range rv.MapKeys() {
		f := rv.MapIndex(key)
		if !f.IsValid() || !f.CanInterface() {
			continue
		}
		fdata, err := encode(f.Interface())
		if err != nil {
			panic(err)
		}
		elem := reflect.ValueOf(fdata)
		m.SetMapIndex(key, elem)
	}
	data := m.Interface()
	if errs.Len() > 0 {
		return data, errors.New(errs.String())
	}
	return data, nil
}

/*
Transforms a complex struct into a map or a []map

	c3po.Encode(struct{Name:"Jão"}) => map[string]any{"Name":"jão"}
	c3po.Encode(struct{Name:"Jão", Age:99}) => map[string]any{"Name":"jão", "Age":99}
	c3po.Encode([]struct{Name:"Jão", Age:99}) => []map[string]any{"Name":"jão", "Age":99}
*/
func Encode(v ...any) (any, error) {
	if v == nil {
		panic("'v' is nil")
	}

	vals := []any{}
	errs := bytes.NewBufferString("")
	for _, val := range v {

		if _v, err := encode(val); err == nil {
			vals = append(vals, _v)
		} else {
			errs.WriteString(err.Error())
		}
	}

	var _e error
	if errs.Len() > 0 {
		_e = errors.New(errs.String())
	}
	if len(vals) == 1 {
		return vals[0], _e
	}
	return vals, _e
}

/*
EncodeToJSON is similar to Encode, but return a []byte{}

	c3po.EncodeToJSON(struct{Name:"Jão", Age:99}) => []byte{"{'Name':'jão', 'Age':99}"}
*/
func EncodeToBytes(v ...any) ([]byte, error) {
	d, err := Encode(v...)
	if err == nil {
		var b []byte
		b, err = json.Marshal(d)
		if err == nil {
			return b, nil
		}
	}
	return []byte{}, err
}

func EncodeToBytesWithIndent(indent string, v ...any) ([]byte, error) {
	d, err := Encode(v...)
	if err == nil {
		var b []byte
		b, err = json.MarshalIndent(d, "", indent)
		if err == nil {
			return b, nil
		}
	}
	return []byte{}, err
}

/*
EncodeToString return a string representation if ok, else empty string

	c3po.EncodeToString(struct{Name:"Jão", Age:99}) => "{'Name':'jão', 'Age':99}"
*/
func EncodeToString(v ...any) string {
	d, err := EncodeToBytes(v...)
	if err == nil {
		return string(d)
	}
	return ""
}

func EncodeToStringIndent(indent string, v ...any) string {
	d, err := EncodeToBytesWithIndent(indent, v...)
	if err == nil {
		return string(d)
	}
	return ""
}
