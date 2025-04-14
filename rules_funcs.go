package c3po

import (
	"reflect"
	"strconv"
)

/*
	MIN
*/

var minKindFunc = map[reflect.Kind]func(reflect.Value, string) bool{
	reflect.String:  minInt,
	reflect.Int:     minInt,
	reflect.Int8:    minInt8,
	reflect.Int16:   minInt16,
	reflect.Int32:   minInt32,
	reflect.Int64:   minInt64,
	reflect.Float32: minFloat32,
	reflect.Float64: minFloat64,
}

func minInt(sch reflect.Value, m string) bool {
	var v int

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(int)
	}

	min, _ := strconv.Atoi(m)

	return v >= min
}

func minInt8(sch reflect.Value, m string) bool {
	var v int8

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(int8)
	}

	min, _ := strconv.Atoi(m)
	return v >= int8(min)
}

func minInt16(sch reflect.Value, m string) bool {
	var v int16

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(int16)
	}
	min, _ := strconv.Atoi(m)
	return v >= int16(min)
}

func minInt32(sch reflect.Value, m string) bool {
	var v int32

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(int32)
	}

	min, _ := strconv.Atoi(m)
	return v >= int32(min)
}

func minInt64(sch reflect.Value, m string) bool {
	var v int64

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(int64)
	}

	min, _ := strconv.Atoi(m)
	return v >= int64(min)
}

func minFloat32(sch reflect.Value, m string) bool {
	var v float32

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(float32)
	}

	min, _ := strconv.ParseFloat(m, 32)
	return v >= float32(min)
}

func minFloat64(sch reflect.Value, m string) bool {
	var v float64

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(float64)
	}

	min, _ := strconv.ParseFloat(m, 64)
	return v >= float64(min)
}

func min(rv reflect.Value, ruleValue string) bool {
	return minKindFunc[rv.Kind()](rv, ruleValue)
}

/*
	MAX
*/

var maxKindFunc = map[reflect.Kind]func(reflect.Value, string) bool{
	reflect.String:  maxInt,
	reflect.Int:     maxInt,
	reflect.Int8:    maxInt8,
	reflect.Int16:   maxInt16,
	reflect.Int32:   maxInt32,
	reflect.Int64:   maxInt64,
	reflect.Float32: maxFloat32,
	reflect.Float64: maxFloat64,
}

func maxInt(sch reflect.Value, m string) bool {
	var v int

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(int)
	}

	min, _ := strconv.Atoi(m)
	return v <= int(min)
}

func maxInt8(sch reflect.Value, m string) bool {
	var v int8

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(int8)
	}

	min, _ := strconv.Atoi(m)
	return v <= int8(min)
}

func maxInt16(sch reflect.Value, m string) bool {
	var v int16

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(int16)
	}

	min, _ := strconv.Atoi(m)
	return v <= int16(min)
}

func maxInt32(sch reflect.Value, m string) bool {
	var v int32

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(int32)
	}

	min, _ := strconv.Atoi(m)
	return v <= int32(min)
}

func maxInt64(sch reflect.Value, m string) bool {
	var v int64

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(int64)
	}

	min, _ := strconv.Atoi(m)
	return v <= int64(min)
}

func maxFloat32(sch reflect.Value, m string) bool {
	var v float32

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(float32)
	}

	min, _ := strconv.ParseFloat(m, 32)
	return v <= float32(min)
}

func maxFloat64(sch reflect.Value, m string) bool {
	var v float64

	to := reflect.TypeOf(v)
	if sch.CanConvert(to) {
		sch2 := sch.Convert(to)
		v = sch2.Interface().(float64)
	}

	min, _ := strconv.ParseFloat(m, 64)
	return v <= float64(min)
}

func max(rv reflect.Value, ruleValue string) bool {
	return maxKindFunc[rv.Kind()](rv, ruleValue)
}

/*
	REQUIRED
*/

func req(rv reflect.Value, ruleValue string) bool { return !rv.IsZero() }

/*
	MINLEN
*/

func canLen(rv reflect.Value) bool {
	switch rv.Kind() {
	default:
		return false
	case reflect.Array, reflect.Slice, reflect.String, reflect.Map, reflect.Chan:
		return true
	}
}

func minLen(rv reflect.Value, ruleValue string) bool {
	if canLen(rv) {
		l, _ := strconv.Atoi(ruleValue)
		if rv.Len() >= l {
			return true
		}
	}
	return false
}

func maxLen(rv reflect.Value, ruleValue string) bool {
	if canLen(rv) {
		l, _ := strconv.Atoi(ruleValue)
		if rv.Len() <= l {
			return true
		}
	}
	return false
}

/*
	ESCAPE
*/

func escape(rv reflect.Value, _ string) bool {
	if str, ok := rv.Interface().(string); ok {
		rv.Set(reflect.ValueOf(HtmlEscape(str)))
	}
	return true
}
