package c3po

import (
	"errors"
	"fmt"
	"reflect"
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

func (f *Fielder) checkSchPtr(r reflect.Value) any {

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
	for rn := range f.Tags {
		if r := GetRule(rn); r != nil {
			f.Rules[rn] = r
		}
	}
}

func (f *Fielder) execRules(sch reflect.Value) (reflect.Value, any) {
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

func (f *Fielder) decodePrimitive(rv reflect.Value) (sch reflect.Value, err any) {
	if f.Type == reflect.Interface {
		sch = rv
	} else {
		sch = GetReflectElem(f.New())
		if !SetReflectValue(sch, rv) {
			if !f.SkipError {
				return reflect.Value{}, RetInvalidType(f)
			}
		}
	}
	return sch, nil
}

func (f *Fielder) decodeSlice(rv reflect.Value) (sch reflect.Value, err any) {
	sliceOf := reflect.TypeOf(f.Schema)
	lenSlice := rv.Len()
	capSlice := rv.Cap()

	sch = reflect.MakeSlice(sliceOf, lenSlice, capSlice)

	errs := []any{}
	for i := 0; i < lenSlice; i++ {
		var (
			s       = GetReflectElem(rv.Index(i))
			sf      = f.SliceType
			err     any
			slicSch reflect.Value
		)

		if f.Walk {
			slicSch, err = sf.decodeSchema(s.Interface())
		} else {
			if sliceOf == s.Type() {
				slicSch = s
			} else {
				err = RetInvalidType(f.SliceType)
			}
		}

		if err != nil {
			errs = append(errs, err)
			continue
		}
		sIndex := sch.Index(i)
		if f.SliceType.IsPointer {
			if slicSch.Kind() != reflect.Ptr && slicSch.CanAddr() {
				slicSch = slicSch.Addr()
			}
		} else {
			if slicSch.Kind() == reflect.Ptr {
				slicSch = slicSch.Elem()
			}
		}
		sIndex.Set(slicSch)
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
		m.SetMapIndex(mkey, mval)
	}
	sch = m
	return
}

func (f *Fielder) decodeStruct(v any) (sch reflect.Value, err any) {
	errs := []any{}
	data, ok := v.(map[string]any)
	if !ok {
		_data, _err := Encode(v)
		if _err != nil {
			err = _err
			return
		}
		data, ok = _data.(map[string]any)
		if !ok {
			if f.SkipError {
				return f.New(), nil
			}
			err = RetInvalidType(f)
			return
		}
	}
	sch = f.New().Elem()

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
			if value, ok = data[fielder.Name]; !ok {
				if fielder.Default != nil {
					value = fielder.Default
				}
			}
			if value == nil && !fielder.Nullable {
				if fielder.Required {
					errs = append(errs, ValidationError{
						Rule: &Rule{
							Name:    "required",
							Message: "{field} is required",
						},
						Field: fielder.Name,
					})
				}
				fielder.execRules(schF)
				continue
			}
		}

		var rvF reflect.Value

		if fielder.SkipValidate {
			SetReflectValue(schF,
				reflectOf(value))
			continue
		} else if fielder.Walk {
			_rvF, e := fielder.decodeSchema(value)
			if e != nil {
				errs = append(errs, e)
				continue
			}
			rvF = _rvF
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
				errs = append(errs, map[string]any{fielder.Name: RetInvalidType(fielder)})
			}
			continue
		}
	}

	if len(errs) == 1 {
		err = errs[0]
	} else if len(errs) > 0 {
		err = errs
	}
	return
}

func (f *Fielder) decodeSchema(v any) (reflect.Value, any) {
	if v == "" && f.Type != reflect.String { // if v == a string (nil or null), v = nil
		v = nil
	}
	if v == nil {
		if f.Default != nil {
			return f.decodeSchema(f.Default)
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
		} else {
			return f.execRules(reflectOf(nil))
		}
	}

	var rfVal = reflectOf(v)

	switch f.Type {
	default:
		r, err := f.decodePrimitive(rfVal)
		if err != nil {
			return r, err
		}
		return f.execRules(r)
	case reflect.Map:
		r, err := f.decodeMap(rfVal)
		if err != nil {
			return r, err
		}
		return f.execRules(r)
	case reflect.Array, reflect.Slice:
		r, err := f.decodeSlice(rfVal)
		if err != nil {
			return r, err
		}
		return f.execRules(r)
	case reflect.Struct:
		r, err := f.decodeStruct(v)
		if err != nil {
			return r, err
		}
		return f.execRules(r)
	}
}

func (f *Fielder) Decode(data any) Schema {
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

	s.val = f.checkSchPtr(sch)
	return s
}

func (f *Fielder) New() reflect.Value {
	rs := reflect.ValueOf(f.Schema)

	if f.IsSlice {
		t := reflect.TypeOf(f.SliceType.Schema)
		t = reflect.SliceOf(t)
		rs = reflect.MakeSlice(t, 0, 0)
	}

	t := GetReflectTypeElem(rs.Type())
	v := reflect.New(t)
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

	return fieldMap
}

func (f *Fielder) String() string {
	return EncodeToStringIndent("  ", f.ToMap())
}
