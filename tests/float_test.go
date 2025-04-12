package test

import (
	"reflect"
	"testing"

	"github.com/5tkgarage/c3po"
)

var (
	f32 float32
	f64 float64
)

// floats.Decode(float) -> !Schema.hasError()
func TestFloatsWithFloats(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"f32": c3po.ParseSchema(f32),
		"f64": c3po.ParseSchema(f64),
	}
	for k, fielder := range fielders {
		switch k {
		case "f32":
			sch := fielder.Decode(123.123)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Float32 {
				t.Errorf("got %q, want %q", schT, reflect.Float32)
			}
		case "f64":
			sch := fielder.Decode(123.123)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Float64 {
				t.Errorf("got %q, want %q", schT, reflect.Float64)
			}
		}
	}
}

// floats.Decode(int) -> !Schema.hasError()
func TestFloatsWithInts(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"f32": c3po.ParseSchema(f32),
		"f64": c3po.ParseSchema(f64),
	}
	for k, fielder := range fielders {
		switch k {
		case "f32":
			sch := fielder.Decode(2147483647)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Float32 {
				t.Errorf("got %q, want %q", schT, reflect.Float32)
			}
		case "f64":
			sch := fielder.Decode(99999999)
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Float64 {
				t.Errorf("got %q, want %q", schT, reflect.Float64)
			}
		}
	}
}

// floats.Decode("int") -> !Schema.hasError()
func TestFloatsWithStringInts(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"f32": c3po.ParseSchema(f32),
		"f64": c3po.ParseSchema(f64),
	}
	for k, fielder := range fielders {
		switch k {
		case "f32":
			sch := fielder.Decode("214748364")
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Float32 {
				t.Errorf("got %q, want %q", schT, reflect.Float32)
			}
		case "f64":
			sch := fielder.Decode("99999999")
			if sch.HasErrors() {
				t.Error(sch.Errors())
			}
			if schT := reflect.ValueOf(sch.Value()).Kind(); schT != reflect.Float64 {
				t.Errorf("got %q, want %q", schT, reflect.Float64)
			}
		}
	}
}

// floats.Decode("string") -> Schema.hasError()
func TestFloatsWithStringValues(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"f32": c3po.ParseSchema(f32),
		"f64": c3po.ParseSchema(f64),
	}
	for k, fielder := range fielders {
		switch k {
		case "f32":
			sch := fielder.Decode("\t\t")
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "f64":
			sch := fielder.Decode("\r\n")
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}

// floats.Decode(false) -> Schema.hasError()
func TestFloatsWithBooleans0(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"f32": c3po.ParseSchema(f32),
		"f64": c3po.ParseSchema(f64),
	}
	for k, fielder := range fielders {
		switch k {
		case "f32":
			sch := fielder.Decode(false)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "f64":
			sch := fielder.Decode(false)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}

// floats.Decode(true) -> Schema.hasError()
func TestFloatsWithBooleans1(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"f32": c3po.ParseSchema(f32),
		"f64": c3po.ParseSchema(f64),
	}
	for k, fielder := range fielders {
		switch k {
		case "f32":
			sch := fielder.Decode(true)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "f64":
			sch := fielder.Decode(true)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}

// floats.Decode(struct) -> Schema.hasError()
func TestFloatsWithStructs(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"f32": c3po.ParseSchema(f32),
		"f64": c3po.ParseSchema(f64),
	}
	s := struct{}{}
	for k, fielder := range fielders {
		switch k {
		case "f32":
			sch := fielder.Decode(s)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "f64":
			sch := fielder.Decode(s)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}

// floats.Decode(map) -> Schema.hasError()
func TestFloatsWithMaps(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"f32": c3po.ParseSchema(f32),
		"f64": c3po.ParseSchema(f64),
	}
	m := map[string]any{}
	for k, fielder := range fielders {
		switch k {
		case "f32":
			sch := fielder.Decode(m)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "f64":
			sch := fielder.Decode(m)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}

// floats.Decode(func) -> Schema.hasError()
func TestFloatsWithFuncs(t *testing.T) {
	fielders := map[string]*c3po.Fielder{
		"f32": c3po.ParseSchema(f32),
		"f64": c3po.ParseSchema(f64),
	}
	f := func() {}
	for k, fielder := range fielders {
		switch k {
		case "f32":
			sch := fielder.Decode(f)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		case "f64":
			sch := fielder.Decode(f)
			if !sch.HasErrors() {
				t.Errorf("got %q, want %q", sch.Value(), "error")
			}
		}
	}
}
