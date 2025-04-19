package c3po

import (
	"reflect"
	"strings"
)

/*
usage:

	c3po.ParseSchema(struct{}) => struct{}
	c3po.ParseSchema(&struct{}) => *struct{}
	c3po.ParseSchema(&struct{Field:Value}) => *struct{Field: value} // with default value

	type Schema struct{
		Field `c3po:"-"`				// omit this field
		Field `c3po:"name"`				// string: name of validation		(default realName)
		Field `c3po:"walk"`				// bool: deep validation			(default true)
		Field `c3po:"escape"`			// bool: escape html value			(default false)
		Field `c3po:"required"`			// bool:		...			 		(default false)
		Field `c3po:"nullable"`			// bool: if true, allow nil value	(default true)
		Field `c3po:"recursive"`		// bool: for embbed data 			(default false)
		Field `c3po:"skiperr"`			// bool: omit on error				(default false)
		Field `c3po:"skip"`				// bool: set value without validate	(default false)
		Field `c3po:"min=18"`			// numbers only (int8, 16..., float32, ...)
		Field `c3po:"max=65"`			// numbers only (int8, 16..., float32, ...)
		Field `c3po:"minlength=1"`		// if a value can len, is valid. else skip
		Field `c3po:"maxlength=100"`	// if a value can len, is valid. else skip
	}
*/
func ParseSchema(schema any) *Fielder {
	return ParseSchemaWithTag("c3po", schema)
}

func ParseSchemaWithTag(tagKey string, schema any) *Fielder {
	tags := map[string]string{}
	if rn := reflect.TypeOf(schema).Name(); rn != "" {
		tags["realName"] = rn
	}
	return parseSchema(schema, tagKey, tags)
}

func parseSchema(schema any, tagKey string, tags map[string]string) *Fielder {
	if _, ok := tags["-"]; ok {
		return nil
	}
	var (
		f  = &Fielder{Schema: schema}
		rv = reflect.ValueOf(schema)
	)
	f.RealName = tags["realName"]

	f.parseHeaders(tags)
	f.parseRules()

	rt := rv.Type()
	if rv.Kind() == reflect.Pointer {
		f.IsPointer = true
		rv = rv.Elem()
		rt = rt.Elem()
	}

	if f.RealName == "" && f.Type != reflect.Interface {
		f.RealName = rt.Name()
	}

	if schema == nil {
		f.Type = reflect.Interface
	} else {
		f.Type = rv.Kind()
		f.Children = map[string]*Fielder{}
	}

	switch f.Type {
	case reflect.Struct:
		f.IsStruct = true
		f.FieldsByIndex = map[int]string{}
		for i := 0; i < rt.NumField(); i++ {
			fv := rv.Field(i)
			if fv.CanInterface() {
				ft := rt.Field(i)
				childTags := parseTags(ft.Tag.Get(tagKey))
				if _, ok := childTags["-"]; ok {
					continue
				}

				cname := ""
				childTags["realName"] = ft.Name
				if v, ok := childTags["name"]; ok && v != "" {
					cname = v
				} else {
					cname = ft.Name
				}
				childTags["name"] = cname

				var fi any
				if fv.Kind() != reflect.Interface {
					fi = fv.Interface()
				}

				child := parseSchema(fi, tagKey, childTags)
				f.FieldsByIndex[i] = cname
				if child != nil {
					child.SuperIndex = &i
					f.Children[cname] = child
					if v, ok := childTags["heritage"]; ok && strings.ToLower(v) == "true" {
						child.Recurcive = true
					}
				}
			}
		}
	case reflect.Slice, reflect.Array:
		f.Type = rt.Kind()
		f.IsSlice = true
		rvt := rv.Type().Elem()
		sliceObjet := reflect.New(rvt).Elem()
		f.SliceType = parseSchema(sliceObjet.Interface(), tagKey, map[string]string{"realName": ""})
	case reflect.Map:
		f.IsMAP = true
		mapKey := reflect.New(rt.Key()).Elem()
		mapValue := reflect.New(rt.Elem()).Elem()
		f.MapKeyType = parseSchema(mapKey.Interface(), tagKey, map[string]string{"realName": ""})
		f.MapValueType = parseSchema(mapValue.Interface(), tagKey, map[string]string{"realName": ""})
	}
	if rv.IsValid() {
		if rv.CanInterface() && !rv.IsZero() {
			f.Default = schema
		}
	}
	return f
}

func (f *Fielder) parseHeaders(tags map[string]string) {
	f.Tags = tags

	if v, ok := tags["name"]; ok && v != "" {
		f.Name = v
	} else {
		f.Name = f.RealName
	}

	if r, ok := tags["recursive"]; ok {
		f.Recurcive = r != "false"
	}

	v, ok := tags["walk"] // default true
	f.Walk = !ok || strings.ToLower(v) != "false"

	v, ok = tags["skip"] // default false
	f.SkipValidate = ok && (strings.ToLower(v) != "false")

	v, ok = tags["required"] // default false
	f.Required = ok && (strings.ToLower(v) == "true")

	_, ok = tags["min"] // default false
	if ok {
		f.Required = ok
	}

	v, ok = tags["nullable"] // default true
	if ok {
		if strings.ToLower(v) == "false" {
			f.Required = true
		} else {
			f.Nullable = true
		}
	}

	if f.Nullable && f.Required {
		f.Nullable = false
	}

	v, ok = tags["skiperr"] // skip field on err - default false
	f.SkipError = ok && (strings.ToLower(v) == "true") && !f.Required
}
