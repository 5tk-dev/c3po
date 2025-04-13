package test

import (
	"testing"

	"5tk.dev/c3po"
)

func TestSliceInt(t *testing.T) {
	var s []int
	sch := c3po.Validate(s, []any{1, "2", 3, 4, "5", "6.5"})
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
	if len(sch.Value().([]int)) != 6 {
		t.Errorf("TestMapString got %v, want %v", sch.Value(), []int{1, 2, 3, 4, 5, 6})
	}
}

func TestSliceStruct(t *testing.T) {
	var s []*UserTest
	var data = []map[string]any{
		{"name": "etho", "age": 18},
		{"name": "etho", "age": 18},
		{"name": "etho", "age": 18},
		{"name": "etho", "age": 18},
	}
	sch := c3po.Validate(s, data)
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
	if len(sch.Value().([]*UserTest)) != 4 {
		t.Errorf("TestMapString got %v, want %v", sch.Value(), []*UserTest{})
	}
}
