package c3po

import (
	"errors"
	"fmt"
	"reflect"
)

// use 'validate' in tags
//
//	type User struct{
//		Name string `validate:"minlen=4"`
//		Age int `validate:"min=18"`
//	}
//	sch := c3po.Validate(&User{},map[string]any{})
//	if sch.HasErrors(){
//		err := sch.Errors()
//		....
//	} else {
//		user := sch.Value().(*User)
//	}
func Validate(sch, data any) Schema {
	return ParseSchemaWithTag("validate", sch).Decode(data)
}

func UnmarshalValidate(dst, src any) error {
	sch := Validate(dst, src)
	if sch.HasErrors() {
		return errors.New(fmt.Sprint(sch.Errors()))
	}
	copyValue(dst, sch.Value())
	return nil
}

func copyValue(dst, src any) {
	rd := reflect.ValueOf(dst)
	rs := reflect.ValueOf(src)

	if rd.Kind() == reflect.Pointer {
		rd = rd.Elem()
	}

	if rs.Kind() == reflect.Pointer {
		rs = rs.Elem()
	}

	rd.Set(rs)
}
