package test

import (
	"reflect"
	"testing"

	"github.com/5tkgarage/c3po"
)

var (
	i   int
	i8  int8
	i16 int16
	i32 int32
	i64 int64
)

// int.Decode(int) -> !Schema.hasError()
func TestIntegersWithInts(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"i":   c3po.ParseSchema(i),
		"i8":  c3po.ParseSchema(i8),
		"i16": c3po.ParseSchema(i16),
		"i32": c3po.ParseSchema(i32),
		"i64": c3po.ParseSchema(i64),
	}
	for k, fielder := range fielders {
		switch k {
		case "i":
			sch := fielder.Decode(9999999)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int {
				t.Errorf("got %q, want %q", schT, reflect.Int)
			}
		case "i8":
			sch := fielder.Decode(127)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int8 {
				t.Errorf("got %q, want %q", schT, reflect.Int8)
			}
		case "i16":
			sch := fielder.Decode(32767)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int16 {
				t.Errorf("got %q, want %q", schT, reflect.Int16)
			}
		case "i32":
			sch := fielder.Decode(2147483647)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int32 {
				t.Errorf("got %q, want %q", schT, reflect.Int32)
			}
		case "i64":
			sch := fielder.Decode(99999999)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int64 {
				t.Errorf("got %q, want %q", schT, reflect.Int64)
			}
		}
	}
}

// int.Decode("int") -> !Schema.hasError()
func TestIntegersWithStringInts(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"i":   c3po.ParseSchema(i),
		"i8":  c3po.ParseSchema(i8),
		"i16": c3po.ParseSchema(i16),
		"i32": c3po.ParseSchema(i32),
		"i64": c3po.ParseSchema(i64),
	}
	for k, fielder := range fielders {
		switch k {
		case "i":
			sch := fielder.Decode("9999999")
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int {
				t.Errorf("got %q, want %q", schT, reflect.Int)
			}
		case "i8":
			sch := fielder.Decode("127")
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int8 {
				t.Errorf("got %q, want %q", schT, reflect.Int8)
			}
		case "i16":
			sch := fielder.Decode("32767")
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int16 {
				t.Errorf("got %q, want %q", schT, reflect.Int16)
			}
		case "i32":
			sch := fielder.Decode("2147483647")
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int32 {
				t.Errorf("got %q, want %q", schT, reflect.Int32)
			}
		case "i64":
			sch := fielder.Decode("99999999")
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int64 {
				t.Errorf("got %q, want %q", schT, reflect.Int64)
			}
		}
	}
}

// int.Decode(float) -> !Schema.hasError()
func TestIntegersWithFloats(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"i":   c3po.ParseSchema(i),
		"i8":  c3po.ParseSchema(i8),
		"i16": c3po.ParseSchema(i16),
		"i32": c3po.ParseSchema(i32),
		"i64": c3po.ParseSchema(i64),
	}
	for k, fielder := range fielders {
		switch k {
		case "i":
			sch := fielder.Decode(123.123)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int {
				t.Errorf("got %q, want %q", schT, reflect.Int)
			}
		case "i8":
			sch := fielder.Decode(123.123)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int8 {
				t.Errorf("got %q, want %q", schT, reflect.Int8)
			}
		case "i16":
			sch := fielder.Decode(123.123)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int16 {
				t.Errorf("got %q, want %q", schT, reflect.Int16)
			}
		case "i32":
			sch := fielder.Decode(123.123)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int32 {
				t.Errorf("got %q, want %q", schT, reflect.Int32)
			}
		case "i64":
			sch := fielder.Decode(123.123)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Int64 {
				t.Errorf("got %q, want %q", schT, reflect.Int64)
			}
		}
	}
}

// int.Decode("string") -> Schema.hasError()
func TestIntegersWithStringValues(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"i":   c3po.ParseSchema(i),
		"i8":  c3po.ParseSchema(i8),
		"i16": c3po.ParseSchema(i16),
		"i32": c3po.ParseSchema(i32),
		"i64": c3po.ParseSchema(i64),
	}
	for k, fielder := range fielders {
		switch k {
		case "i":
			sch := fielder.Decode("fdagaj156171")
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i8":
			sch := fielder.Decode("a")
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i16":
			sch := fielder.Decode("cdfc")
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i32":
			sch := fielder.Decode("\t\t")
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i64":
			sch := fielder.Decode("\r\n")
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}

// int.Decode(false) -> Schema.hasError()
func TestIntegersWithBooleans0(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"i":   c3po.ParseSchema(i),
		"i8":  c3po.ParseSchema(i8),
		"i16": c3po.ParseSchema(i16),
		"i32": c3po.ParseSchema(i32),
		"i64": c3po.ParseSchema(i64),
	}
	for k, fielder := range fielders {
		switch k {
		case "i":
			sch := fielder.Decode(false)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i8":
			sch := fielder.Decode(false)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i16":
			sch := fielder.Decode(false)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i32":
			sch := fielder.Decode(false)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i64":
			sch := fielder.Decode(false)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}

// int.Decode(true) -> Schema.hasError()
func TestIntegersWithBooleans1(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"i":   c3po.ParseSchema(i),
		"i8":  c3po.ParseSchema(i8),
		"i16": c3po.ParseSchema(i16),
		"i32": c3po.ParseSchema(i32),
		"i64": c3po.ParseSchema(i64),
	}
	for k, fielder := range fielders {
		switch k {
		case "i":
			sch := fielder.Decode(true)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i8":
			sch := fielder.Decode(true)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i16":
			sch := fielder.Decode(true)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i32":
			sch := fielder.Decode(true)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i64":
			sch := fielder.Decode(true)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}

// int.Decode(struct) -> Schema.hasError()
func TestIntegersWithStructs(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"i":   c3po.ParseSchema(i),
		"i8":  c3po.ParseSchema(i8),
		"i16": c3po.ParseSchema(i16),
		"i32": c3po.ParseSchema(i32),
		"i64": c3po.ParseSchema(i64),
	}
	s := struct{}{}
	for k, fielder := range fielders {
		switch k {
		case "i":
			sch := fielder.Decode(s)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i8":
			sch := fielder.Decode(s)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i16":
			sch := fielder.Decode(s)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i32":
			sch := fielder.Decode(s)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i64":
			sch := fielder.Decode(s)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}

// int.Decode(map) -> Schema.hasError()
func TestIntegersWithMaps(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"i":   c3po.ParseSchema(i),
		"i8":  c3po.ParseSchema(i8),
		"i16": c3po.ParseSchema(i16),
		"i32": c3po.ParseSchema(i32),
		"i64": c3po.ParseSchema(i64),
	}
	m := map[string]any{}
	for k, fielder := range fielders {
		switch k {
		case "i":
			sch := fielder.Decode(m)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i8":
			sch := fielder.Decode(m)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i16":
			sch := fielder.Decode(m)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i32":
			sch := fielder.Decode(m)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i64":
			sch := fielder.Decode(m)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}

// int.Decode(func) -> Schema.hasError()
func TestIntegersWithFuncs(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"i":   c3po.ParseSchema(i),
		"i8":  c3po.ParseSchema(i8),
		"i16": c3po.ParseSchema(i16),
		"i32": c3po.ParseSchema(i32),
		"i64": c3po.ParseSchema(i64),
	}
	f := func() {}
	for k, fielder := range fielders {
		switch k {
		case "i":
			sch := fielder.Decode(f)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i8":
			sch := fielder.Decode(f)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i16":
			sch := fielder.Decode(f)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i32":
			sch := fielder.Decode(f)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "i64":
			sch := fielder.Decode(f)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}
