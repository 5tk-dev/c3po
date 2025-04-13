package test

import (
	"testing"

	"5tk.dev/c3po"
)

type TagMaxInt8 struct {
	Value int8 `validate:"max=0"`
}

func TestMaxInt8_ok(t *testing.T) {
	m := &TagMaxInt8{}
	sch := c3po.Validate(m, map[string]any{"value": "-1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMaxInt8); !ok {
		t.Errorf("got: %v, want: *TagMaxInt8{}", sch.Value())
	} else if v.Value != -1 {
		t.Errorf("got: %v, want: -1", sch.Value())
	}
}

func TestMaxInt8_max(t *testing.T) {
	m := &TagMaxInt8{}
	sch := c3po.Validate(m, map[string]any{"value": "1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}

type TagMaxInt16 struct {
	Value int16 `validate:"max=0"`
}

func TestMaxInt16_ok(t *testing.T) {
	m := &TagMaxInt16{}
	sch := c3po.Validate(m, map[string]any{"value": "-1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMaxInt16); !ok {
		t.Errorf("got: %v, want: *TagMaxInt16{}", sch.Value())
	} else if v.Value != -1 {
		t.Errorf("got: %v, want: -1", sch.Value())
	}
}

func TestMaxInt16_max(t *testing.T) {
	m := &TagMaxInt16{}
	sch := c3po.Validate(m, map[string]any{"value": "1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}

type TagMaxInt32 struct {
	Value int32 `validate:"max=0"`
}

func TestMaxInt32_ok(t *testing.T) {
	m := &TagMaxInt32{}
	sch := c3po.Validate(m, map[string]any{"value": "-1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMaxInt32); !ok {
		t.Errorf("got: %v, want: *TagMaxInt32{}", sch.Value())
	} else if v.Value != -1 {
		t.Errorf("got: %v, want: -1", sch.Value())
	}
}

func TestMaxInt32_max(t *testing.T) {
	m := &TagMaxInt32{}
	sch := c3po.Validate(m, map[string]any{"value": "1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}

type TagMaxInt64 struct {
	Value int64 `validate:"max=0"`
}

func TestMaxInt64_ok(t *testing.T) {
	m := &TagMaxInt64{}
	sch := c3po.Validate(m, map[string]any{"value": "-1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMaxInt64); !ok {
		t.Errorf("got: %v, want: *TagMaxInt64{}", sch.Value())
	} else if v.Value != -1 {
		t.Errorf("got: %v, want: -1", sch.Value())
	}
}

func TestMaxInt64_max(t *testing.T) {
	m := &TagMaxInt64{}
	sch := c3po.Validate(m, map[string]any{"value": "1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}

type TagMaxInt struct {
	Value int `validate:"max=0"`
}

func TestMaxInt_ok(t *testing.T) {
	m := &TagMaxInt{}
	sch := c3po.Validate(m, map[string]any{"value": "-1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMaxInt); !ok {
		t.Errorf("got: %v, want: *TagMaxInt{}", sch.Value())
	} else if v.Value != -1 {
		t.Errorf("got: %v, want: -1", sch.Value())
	}
}

func TestMaxInt_max(t *testing.T) {
	m := &TagMaxInt{}
	sch := c3po.Validate(m, map[string]any{"value": "1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}

/*
	FLOATS
*/

type TagMaxFloat32 struct {
	Value float32 `validate:"max=0"`
}

func TestMaxFloat32_ok(t *testing.T) {
	m := &TagMaxFloat32{}
	sch := c3po.Validate(m, map[string]any{"value": "-1.32"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMaxFloat32); !ok {
		t.Errorf("got: %v, want: *TagMaxFloat32{}", sch.Value())
	} else if v.Value != -1.32 {
		t.Errorf("got: %v, want: -1.32", sch.Value())
	}
}

func TestMaxFloat32_max(t *testing.T) {
	m := &TagMaxFloat32{}
	sch := c3po.Validate(m, map[string]any{"value": "1.32"})

	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}

type TagMaxFloat64 struct {
	Value float64 `validate:"max=0"`
}

func TestMaxFloat64_ok(t *testing.T) {
	m := &TagMaxFloat64{}
	sch := c3po.Validate(m, map[string]any{"value": "-1.64"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMaxFloat64); !ok {
		t.Errorf("got: %v, want: *TagMaxFloat64{}", sch.Value())
	} else if v.Value != -1.64 {
		t.Errorf("got: %v, want: -1.64", sch.Value())
	}
}

func TestMaxFloat64_max(t *testing.T) {
	m := &TagMaxFloat64{}
	sch := c3po.Validate(m, map[string]any{"value": "1.64"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}
