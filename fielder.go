package c3po

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"5tk.dev/c3po/encoder"
)

var (
	_skipTag = map[string]any{
		"realName": nil,
		"strType":  nil,
	}
)

type Fielder struct {
	Name     string
	Type     reflect.Kind
	Tags     map[string]string
	Default  any
	Schema   any
	RealName string
	Rules    map[string]*Rule

	IsMAP,
	IsSlice,
	IsStruct,
	IsPointer bool

	SliceType,
	MapKeyType,
	MapValueType *Fielder

	SkipValidate bool // default: false
	Walk         bool // default: true -> deep validation
	Recurcive    bool // default: false -> for embed struct

	Children      map[string]*Fielder //
	FieldsByIndex map[int]string      //
	SuperIndex    *int                // if a field to a struct

	Required bool // default: false

	Nullable bool // default: true

	SkipError bool // default: false
	OmitEmpty bool // default: false
}

func (f *Fielder) CheckSchPtr(r reflect.Value) any {
	if f.IsPointer && (r.CanAddr() && r.Kind() != reflect.Pointer) {
		return r.Addr().Interface()
	} else if !f.IsPointer && r.Kind() == reflect.Pointer {
		return r.Elem().Interface()
	}
	return r.Interface()
}

func (f *Fielder) parseRules() {
	if f.Rules == nil {
		f.Rules = map[string]*Rule{}
	}
	for rn, v := range f.Tags {
		if r := GetRule(rn); r != nil {
			nr := &Rule{
				Name:     r.Name,
				Message:  r.Message,
				Validate: r.Validate,
			}
			if r.Value == "" {
				nr.Value = v
			}
			f.Rules[rn] = nr
		}
	}
}

func (f *Fielder) ExecRules(sch reflect.Value) (reflect.Value, any) {
	for _, r := range f.Rules {
		if !r.Validate(sch, r.Value) {
			err := ValidationError{
				Rule:  r,
				Field: f.Name,
			}
			return reflect.Value{}, err
		}
	}
	return sch, nil
}

func (f *Fielder) decodePrimitive(rv reflect.Value) (reflect.Value, any) {
	if f.Type == reflect.Interface {
		return rv, nil
	}
	sch := f.New()
	if !SetReflectValue(sch, rv) {
		if !f.SkipError {
			return reflect.Value{}, RetInvalidType(f)
		}
	}
	if f.IsPointer && sch.CanAddr() {
		return sch.Addr(), nil
	}
	return sch, nil
}

func (f *Fielder) decodeSlice(rv reflect.Value) (sch reflect.Value, err any) {

	errs := []any{}
	sliceOf := reflect.TypeOf(f.Schema)
	lenSlice := rv.Len()

	if f.IsPointer {
		sch = reflect.MakeSlice(sliceOf.Elem(), lenSlice, rv.Cap())
	} else {
		sch = reflect.MakeSlice(sliceOf, lenSlice, rv.Cap())
	}

	for i := range lenSlice {
		var (
			is       = rv.Index(i)
			sf       = f.SliceType
			err      any
			sliceSch reflect.Value
		)

		if f.Walk {
			sliceSch, err = sf.decodeSchema(is.Interface())
		} else {
			sliceSch = is
		}

		if err != nil {
			errs = append(errs, err)
			continue
		}

		schIndex := sch.Index(i)
		if f.SliceType.IsPointer {
			if sliceSch.Kind() != reflect.Ptr && sliceSch.CanAddr() {
				sliceSch = sliceSch.Addr()
			}
		} else {
			if sliceSch.Kind() == reflect.Ptr {
				sliceSch = sliceSch.Elem()
			}
		}
		schIndex.Set(sliceSch)
	}

	if sch.Len() == 0 {
		if f.Required {
			errs = append(errs, RetMissing(f))
		}
	}

	if len(errs) == 1 {
		err = errs[0]
	} else if len(errs) > 0 {
		err = errs
	}

	return
}

func (f *Fielder) decodeMap(rv reflect.Value) (sch reflect.Value, err any) {
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		err = map[string]any{f.Name: RetInvalidType(f)}
		return
	}
	mt := reflect.TypeOf(f.Schema)
	m := reflect.MakeMap(mt)
	for _, key := range rv.MapKeys() {
		mindex := rv.MapIndex(key)

		mkey, _err := f.MapKeyType.decodeSchema(key.Interface())
		if _err != nil {
			err = _err
			return
		}
		mval, _err := f.MapValueType.decodeSchema(mindex.Interface())
		if _err != nil {
			err = _err
			return
		}
		if f.MapValueType.IsPointer && mval.CanAddr() {
			mval = mval.Addr()
		}
		if f.MapKeyType.IsPointer && mkey.CanAddr() {
			mkey = mkey.Addr()
		}
		m.SetMapIndex(mkey, mval)
	}
	sch = m
	return
}

func (f *Fielder) decodeStruct(v any) (reflect.Value, any) {
	errs := []any{}
	data, ok := v.(map[string]any)
	if !ok {
		_data := encoder.Encode(v)

		data, ok = _data.(map[string]any)
		if !ok {
			if d, ok := _data.(map[any]any); ok {
				data = map[string]any{}
				for k, v := range d {
					data[fmt.Sprint(k)] = v
				}
			} else if f.SkipError {
				return f.New(), nil
			} else {
				return reflect.Value{}, RetInvalidType(f)
			}
		}
	}

	sch := f.New()

	for i := range sch.NumField() {
		schF := sch.Field(i)
		if !schF.CanInterface() {
			continue
		}

		var value any
		fName := f.FieldsByIndex[i]
		fielder, ok := f.Children[fName]
		if !ok {
			continue
		}

		if fielder.Recurcive {
			value = data
		} else {
			fn := fName
			if value, ok = data[fn]; !ok {
				if fielder.Default != nil {
					value = fielder.Default
				}
			}

			if value == nil && !fielder.Nullable {
				if fielder.Required {
					errs = append(errs, RetMissing(fielder))
				}
				fielder.ExecRules(schF)
				continue
			}
		}

		var rvF reflect.Value
		var err any

		if fielder.SkipValidate {
			SetReflectValue(schF,
				reflect.ValueOf(value))
			continue
		}

		if fielder.Walk {
			if rvF, err = fielder.decodeSchema(value); err != nil {
				errs = append(errs, err)
				continue
			}
		} else {
			rvF = reflect.ValueOf(value)
		}

		if rvF.Kind() == reflect.Invalid {
			if fielder.Nullable {
				continue
			} else if !fielder.SkipError {
				errs = append(errs, map[string]any{fielder.Name: RetInvalidType(fielder)})
			}
		} else if !SetReflectValue(schF, rvF) {
			if !fielder.SkipError {
				errs = append(errs, RetInvalidType(fielder))
			}
			continue
		}
	}
	var err any
	if len(errs) == 1 {
		err = errs[0]
	} else if len(errs) > 0 {
		err = errs
	}

	return sch, err
}

func (f *Fielder) decodeSchema(v any) (reflect.Value, any) {

	if v == "" && f.Type != reflect.String { // if v == a string (nil or null), v = nil
		v = nil
	}
	var (
		rfVal = reflect.ValueOf(v)
		sch   any
		err   any
	)

	if v == nil {
		if f.Default != nil {
			sch, err = f.decodeSchema(f.Default)
		} else if f.Required {
			errs := map[string]any{}
			if len(f.Children) > 0 {
				for _, c := range f.Children {
					if c.Required {
						errs[c.Name] = RetMissing(c)
					}
				}
				return reflect.Value{}, errs
			} else {
				return reflect.Value{}, map[string]any{
					f.Name: RetMissing(f),
				}
			}
		} else if f.Nullable {
			return f.ExecRules(reflect.ValueOf(nil))
		}
	}
	if err == nil {
		switch f.Type {
		default:
			sch, err = f.decodePrimitive(rfVal)
		case reflect.Map:
			sch, err = f.decodeMap(rfVal)
		case reflect.Array, reflect.Slice:
			sch, err = f.decodeSlice(rfVal)
		case reflect.Struct:
			sch, err = f.decodeStruct(v)
		}
	}
	if err != nil {
		return reflect.Value{}, err
	}
	return f.ExecRules(sch.(reflect.Value))
}

func (f *Fielder) Decode(data any) Schema {
	if d, ok := data.(string); ok {
		if (f.Type == reflect.Map) || (f.Type == reflect.Struct) || (f.Type == reflect.Slice) {
			var m any
			err := json.Unmarshal([]byte(d), &m)
			if err == nil {
				data = m
			}
		}
	}
	sch, err := f.decodeSchema(data)
	s := &schema{}
	if err != nil {
		if e, ok := err.(error); ok {
			s.errors = append(s.errors, e)
			return s
		}
		if str, ok := err.(string); ok {
			s.errors = append(s.errors, errors.New(str))
			return s
		}
		s.errors = append(s.errors, errors.New(fmt.Sprint(err)))
		return s
	}
	s.val = f.CheckSchPtr(sch)
	return s
}

func (f *Fielder) New() reflect.Value {
	rs := reflect.ValueOf(f.Schema)

	if f.IsSlice {
		t := reflect.TypeOf(f.SliceType.Schema)
		t = reflect.SliceOf(t)
		rs = reflect.MakeSlice(t, 0, 0)
	}

	t := rs.Type()
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	v := reflect.New(t).Elem()
	return v
}

func (f *Fielder) ToMap() map[string]any {
	fieldMap := map[string]any{}
	for t, v := range f.Tags {
		if _, ok := _skipTag[t]; ok {
			continue
		}
		fieldMap[t] = v
	}

	st := f.Tags["strType"]
	if st == "" {
		st = f.Type.String()
	}

	if st != "struct" {
		fieldMap["type"] = st
	}

	if len(f.Children) > 0 {
		for cn, cv := range f.Children {
			fieldMap[cn] = cv.ToMap()
		}
	} else if f.IsSlice {
		fieldMap["fields"] = f.SliceType.ToMap()
	}

	mapRules := []map[string]any{}
	for _, r := range f.Rules {
		mapRules = append(mapRules, r.ToMap())
	}
	fieldMap["rules"] = mapRules

	return fieldMap
}

func (f *Fielder) String() string {
	return encoder.EncodeToStringIndent("  ", f.ToMap())
}
