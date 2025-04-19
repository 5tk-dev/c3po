package c3po

import (
	"errors"
	"fmt"
	"reflect"
)

func UnmarshalValidate(dst, src any) error {
	sch := Validate(dst, src)
	if sch.HasErrors() {
		return errors.New(fmt.Sprint(sch.Errors()))
	}
	copyValue(dst, sch.Value())
	return nil
}

func copyValue(dst, src any) {
	rd := reflect.ValueOf(dst).Elem()
	rs := reflect.ValueOf(src).Elem()
	rd.Set(rs)
}
