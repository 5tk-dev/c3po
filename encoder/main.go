package encoder

import (
	"encoding/json"
	"reflect"
	"strings"
)

func encode(v any) any {
	if v == nil {
		return nil
	}

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
		return v
	case reflect.Invalid, reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return nil
	case reflect.Pointer:
		if rv.Elem().IsValid() {
			return Encode(rv.Elem())
		}
		return nil
	case reflect.Struct:
		d := map[string]any{}
		for i := 0; i < rv.NumField(); i++ {
			f := rv.Field(i)
			if !f.IsValid() || !f.CanInterface() {
				continue
			}
			ft := rt.Field(i)
			fv := Encode(f.Interface())
			d[strings.ToLower(ft.Name)] = fv
		}
		return d
	case reflect.Slice, reflect.Array:
		return encodeSlice(v)
	case reflect.Map:
		return encodeMap(v)
	}
}

func encodeSlice(v any) []any {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer && rv.IsValid() {
		if rv.Elem().IsValid() {
			rv = rv.Elem()
		}
	}
	d := reflect.MakeSlice(reflect.TypeOf([]any{}), rv.Len(), rv.Cap())
	for i := range rv.Len() {
		f := rv.Index(i)
		if !f.IsValid() || !f.CanInterface() {
			continue
		}
		fv := Encode(f.Interface())
		d.Index(i).Set(reflect.ValueOf(fv))
	}
	return d.Interface().([]any)
}

func encodeMap(v any) map[any]any {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer && rv.IsValid() {
		if rv.Elem().IsValid() {
			rv = rv.Elem()
		}
	}

	m := reflect.MakeMap(reflect.TypeOf(map[any]any{}))

	for _, key := range rv.MapKeys() {
		f := rv.MapIndex(key)
		if !f.IsValid() || !f.CanInterface() {
			continue
		}
		fdata := encode(f.Interface())
		elem := reflect.ValueOf(fdata)
		m.SetMapIndex(key, elem)
	}
	data := m.Interface().(map[any]any)
	return data
}

func Encode(v ...any) any {
	if v == nil {
		return nil
	}
	values := []any{}
	for _, val := range v {
		v2 := encode(val)
		values = append(values, v2)
	}

	if len(values) == 1 {
		return values[0]
	}
	return values
}

/*
EncodeToJSON is similar to Encode, but return a []byte{}

	c3po.EncodeToJSON(struct{Name:"Jão", Age:99}) => []byte{"{'Name':'jão', 'Age':99}"}
*/
func EncodeToBytes(v ...any) ([]byte, error) {
	d := Encode(v...)
	b, err := json.Marshal(d)
	if err == nil {
		return b, nil
	}
	return b, err
}

func EncodeToBytesWithIndent(indent string, v ...any) ([]byte, error) {
	d := Encode(v...)
	b, err := json.MarshalIndent(d, "", indent)
	if err == nil {
		return b, nil
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
