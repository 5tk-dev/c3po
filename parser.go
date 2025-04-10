package c3po

import (
	"log"
	"reflect"
	"strings"
)

/*
usage:

	c3po.ParseSchema(struct{}) => struct{}
	c3po.ParseSchema(&struct{}) => *struct{}
	c3po.ParseSchema(&struct{Field:Value}) => *struct{Field: value} // with default value

	type Schema struct{
		Field `c3po:"-"`			// omit this field
		Field `c3po:"name"`			// string: name of validation		(default realName)
		Field `c3po:"escape"`		// bool: escape html value			(default false)
		Field `c3po:"required"`		// bool:		...			 		(default false)
		Field `c3po:"nullable"`		// bool: if true, allow nil value	(default false)
		Field `c3po:"deep"`			// bool: deep validation			(default true)
		Field `c3po:"skiponerr"`	// bool: omit on valid. error		(default false)
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

	var rt reflect.Type
	rv := reflect.ValueOf(schema)
	f := &Fielder{}
	f.parseHeaders(tags)

	if schema != nil {
		if rv.Kind() == reflect.Ptr {
			f.IsPointer = true
		}

		rt = rv.Type()
		rv = GetReflectElem(rv)
		if rt.Kind() == reflect.Ptr {
			rt = rt.Elem()
		}

		f.Type = rt.Kind()
		f.Children = map[string]*Fielder{}
	} else {
		f.Type = reflect.Interface
	}
	f.Schema = schema

	f.RealName = tags["realName"]
	if f.RealName == "" && f.Type != reflect.Interface {
		f.RealName = rt.Name()
	}

	if rv.Kind() == reflect.Pointer {
		f.IsPointer = true
		rTmp := rv.Elem()
		if rTmp.Kind() == reflect.Pointer {
			rv = rTmp
			rt = rt.Elem()
		}
	}
	if !rv.IsValid() && f.Type != reflect.Interface {
		rv = reflect.New(rt).Elem()
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
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
					childTags["name"] = v
				} else {
					cname = strings.ToLower(ft.Name)
					childTags["name"] = cname
				}
				var finter any
				if fv.Kind() != reflect.Interface {
					finter = fv.Interface()
				}
				child := parseSchema(finter, tagKey, childTags)
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
		f.Walk = true
		f.Type = rt.Kind()
		f.IsSlice = true
		sliceObjet := reflect.New(rv.Type().Elem()).Elem()
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
			def, err := encode(rv.Interface())
			if err != nil {
				n := f.Name
				if n == "" {
					n = f.RealName
				}
				log.Printf("warn: Default Value is invalid into fielder: '%s'\n", n)
				log.Println(err)
			} else {
				f.Default = def
			}
		}
	}
	return f
}

func (f *Fielder) parseHeaders(tags map[string]string) {
	f.Tags = tags
	f.parseRules()

	if v, ok := tags["name"]; ok && v != "" {
		f.Name = v
	} else {
		f.Name = f.RealName
	}

	if r, ok := tags["recursive"]; ok {
		f.Recurcive = r != "false"
	}

	v, ok := tags["escape"] // default false
	f.Escape = (ok && (strings.ToLower(v) == "true"))

	v, ok = tags["required"] // default false
	f.Required = ok && (strings.ToLower(v) == "true")

	v, ok = tags["walk"] // default true
	f.Walk = !ok || strings.ToLower(v) != "false"

	v, ok = tags["nullable"] // default true
	f.Nullable = ok && strings.ToLower(v) == "true" && !f.Required

	v, ok = tags["nonzero"] // default false
	f.NonZero = ok && (strings.ToLower(v) == "true")

	v, ok = tags["skiperror"] // skip field on err - default false
	f.SkipError = ok && (strings.ToLower(v) == "true") && !f.Required
}
