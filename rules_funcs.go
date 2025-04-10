package c3po

import (
	"reflect"
	"strconv"
)

/*
	MIN
*/

var minKindFunc = map[reflect.Kind]func(reflect.Value, string) bool{
	reflect.Int:     minInt,
	reflect.Int8:    minInt8,
	reflect.Int16:   minInt16,
	reflect.Int32:   minInt32,
	reflect.Int64:   minInt64,
	reflect.Float32: minFloat32,
	reflect.Float64: minFloat64,
}

func minInt(sch reflect.Value, m string) bool {
	min, _ := strconv.Atoi(m)
	return sch.Interface().(int) < min
}

func minInt8(sch reflect.Value, m string) bool {
	min, _ := strconv.Atoi(m)
	return sch.Interface().(int8) < int8(min)
}

func minInt16(sch reflect.Value, m string) bool {
	min, _ := strconv.Atoi(m)
	return sch.Interface().(int16) < int16(min)
}
func minInt32(sch reflect.Value, m string) bool {
	min, _ := strconv.Atoi(m)
	return sch.Interface().(int32) < int32(min)
}
func minInt64(sch reflect.Value, m string) bool {
	min, _ := strconv.Atoi(m)
	return sch.Interface().(int64) < int64(min)
}

func minFloat32(sch reflect.Value, m string) bool {
	min, _ := strconv.ParseFloat(m, 32)
	return sch.Interface().(float32) < float32(min)
}
func minFloat64(sch reflect.Value, m string) bool {
	min, _ := strconv.ParseFloat(m, 64)
	return sch.Interface().(float64) < float64(min)
}

func min(rv reflect.Value, ruleValue string) bool {
	return !minKindFunc[rv.Kind()](rv, ruleValue)
}

/*
	MAX
*/

var maxKindFunc = map[reflect.Kind]func(reflect.Value, string) bool{
	reflect.Int:     maxInt,
	reflect.Int8:    maxInt8,
	reflect.Int16:   maxInt16,
	reflect.Int32:   maxInt32,
	reflect.Int64:   maxInt64,
	reflect.Float32: maxFloat32,
	reflect.Float64: maxFloat64,
}

func maxInt(sch reflect.Value, m string) bool {
	min, _ := strconv.Atoi(m)
	return sch.Interface().(int) > min
}

func maxInt8(sch reflect.Value, m string) bool {
	min, _ := strconv.Atoi(m)
	return sch.Interface().(int8) > int8(min)
}

func maxInt16(sch reflect.Value, m string) bool {
	min, _ := strconv.Atoi(m)
	return sch.Interface().(int16) > int16(min)
}
func maxInt32(sch reflect.Value, m string) bool {
	min, _ := strconv.Atoi(m)
	return sch.Interface().(int32) > int32(min)
}
func maxInt64(sch reflect.Value, m string) bool {
	min, _ := strconv.Atoi(m)
	return sch.Interface().(int64) > int64(min)
}

func maxFloat32(sch reflect.Value, m string) bool {
	min, _ := strconv.ParseFloat(m, 32)
	return sch.Interface().(float32) > float32(min)
}
func maxFloat64(sch reflect.Value, m string) bool {
	min, _ := strconv.ParseFloat(m, 64)
	return sch.Interface().(float64) > float64(min)
}

func max(rv reflect.Value, ruleValue string) bool {
	return !maxKindFunc[rv.Kind()](rv, ruleValue)
}

/*
	REQUIRED
*/

func req(rv reflect.Value, ruleValue string) bool {
	return !rv.IsZero()
}
