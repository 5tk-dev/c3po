package test

import (
	"testing"

	"5tk.dev/c3po"
)

type TagMinInt8 struct {
	Value int8 `validate:"min=0"`
}

func TestMinInt8_ok(t *testing.T) {
	m := &TagMinInt8{}
	sch := c3po.Validate(m, map[string]any{"value": "1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMinInt8); !ok {
		t.Errorf("got: %v, want: *TagMinInt8{}", sch.Value())
	} else if v.Value != 1 {
		t.Errorf("got: %v, want: 1", sch.Value())
	}
}

func TestMinInt8_minor(t *testing.T) {
	m := &TagMinInt8{}
	sch := c3po.Validate(m, map[string]any{"value": "-1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}

type TagMinInt16 struct {
	Value int16 `validate:"min=0"`
}

func TestMinInt16_ok(t *testing.T) {
	m := &TagMinInt16{}
	sch := c3po.Validate(m, map[string]any{"value": "1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMinInt16); !ok {
		t.Errorf("got: %v, want: *TagMinInt8{}", sch.Value())
	} else if v.Value != 1 {
		t.Errorf("got: %v, want: 1", sch.Value())
	}
}

func TestMinInt16_minor(t *testing.T) {
	m := &TagMinInt16{}
	sch := c3po.Validate(m, map[string]any{"value": "-1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}

type TagMinInt32 struct {
	Value int32 `validate:"min=0"`
}

func TestMinInt32_ok(t *testing.T) {
	m := &TagMinInt32{}
	sch := c3po.Validate(m, map[string]any{"value": "1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMinInt32); !ok {
		t.Errorf("got: %v, want: *TagMinInt8{}", sch.Value())
	} else if v.Value != 1 {
		t.Errorf("got: %v, want: 1", sch.Value())
	}
}

func TestMinInt32_minor(t *testing.T) {
	m := &TagMinInt32{}
	sch := c3po.Validate(m, map[string]any{"value": "-1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}

type TagMinInt64 struct {
	Value int64 `validate:"min=0"`
}

func TestMinInt64_ok(t *testing.T) {
	m := &TagMinInt64{}
	sch := c3po.Validate(m, map[string]any{"value": "1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMinInt64); !ok {
		t.Errorf("got: %v, want: *TagMinInt8{}", sch.Value())
	} else if v.Value != 1 {
		t.Errorf("got: %v, want: 1", sch.Value())
	}
}

func TestMinInt64_minor(t *testing.T) {
	m := &TagMinInt64{}
	sch := c3po.Validate(m, map[string]any{"value": "-1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}

type TagMinInt struct {
	Value int `validate:"min=0"`
}

func TestMinInt_ok(t *testing.T) {
	m := &TagMinInt{}
	sch := c3po.Validate(m, map[string]any{"value": "1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMinInt); !ok {
		t.Errorf("got: %v, want: *TagMinInt8{}", sch.Value())
	} else if v.Value != 1 {
		t.Errorf("got: %v, want: 1", sch.Value())
	}
}

func TestMinInt_minor(t *testing.T) {
	m := &TagMinInt{}
	sch := c3po.Validate(m, map[string]any{"value": "-1"})
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

type TagMinFloat32 struct {
	Value float32 `validate:"min=1.32"`
}

func TestMinFloat32_ok(t *testing.T) {
	m := &TagMinFloat32{}
	sch := c3po.Validate(m, map[string]any{"value": "1.32"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMinFloat32); !ok {
		t.Errorf("got: %v, want: *TagMinFloat32{}", sch.Value())
	} else if v.Value != 1.32 {
		t.Errorf("got: %v, want: 1.32", sch.Value())
	}
}

func TestMinFloat32_minor(t *testing.T) {
	m := &TagMinFloat32{}
	sch := c3po.Validate(m, map[string]any{"value": 1.31})

	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}

type TagMinFloat64 struct {
	Value float64 `validate:"min=0"`
}

func TestMinFloat64_ok(t *testing.T) {
	m := &TagMinFloat64{}
	sch := c3po.Validate(m, map[string]any{"value": "1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
	}

	if v, ok := sch.Value().(*TagMinFloat64); !ok {
		t.Errorf("got: %v, want: *TagMinFloat64{}", sch.Value())
	} else if v.Value != 1 {
		t.Errorf("got: %v, want: 1", sch.Value())
	}
}

func TestMinFloat64_minor(t *testing.T) {
	m := &TagMinFloat64{}
	sch := c3po.Validate(m, map[string]any{"value": "-1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
	}
}
