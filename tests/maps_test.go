package test

import (
	"testing"

	"5tk.dev/c3po"
)

func TestMapsStringString(t *testing.T) {
	var m1 = map[string]string{}
	var m2 = c3po.Validate(m1, map[any]any{"hello": "world"})
	if m2.HasErrors() {
		t.Error(m2.Errors())
	}
	h := m2.Value().(map[string]string)["hello"]
	if h != "world" {
		t.Errorf("TestMapString got %q, want %q", h, "world")
	}
}

func TestMapsStringInt(t *testing.T) {
	var m1 = map[string]int{}
	var m2 = c3po.Validate(m1, map[any]any{"int": "1"})
	if m2.HasErrors() {
		t.Error(m2.Errors())
	}

	h := m2.Value().(map[string]int)["int"]
	if h != 1 {
		t.Errorf("TestMapString got %q, want %q", h, 1)
	}
}

func TestMapsStringFloat32(t *testing.T) {
	var m1 = map[string]float32{}
	var m2 = c3po.Validate(m1, map[any]any{"float": "1.123"})
	if m2.HasErrors() {
		t.Error(m2.Errors())
	}

	h := m2.Value().(map[string]float32)["float"]
	if h != 1.123 {
		t.Errorf("TestMapString got %v, want %v", h, 1.123)
	}
}

func TestMapsStringBool(t *testing.T) {
	var m1 = map[string]bool{}
	var m2 = c3po.Validate(m1, map[any]any{"bool": true})
	if m2.HasErrors() {
		t.Error(m2.Errors())
	}

	h, ok := m2.Value().(map[string]bool)["bool"]
	if !ok {
		t.Errorf("TestMapString got %v, want %v", h, true)
	}
}

func TestMapsStringStruct(t *testing.T) {
	var m1 = map[string]*UserTest{}
	var data = map[any]any{
		"1": &UserTest{Name: "etho"},
		"2": &UserTest{Name: "etho2"},
		"3": &UserTest{Name: "etho3"},
		"4": &UserTest{Name: "etho4"},
		"5": &UserTest{Name: "etho5"},
		"6": &UserTest{Name: "etho6"},
	}
	var m2 = c3po.Validate(m1, data)

	if m2.HasErrors() {
		t.Error(m2.Errors())
	}

	m3 := m2.Value().(map[string]*UserTest)
	if len(m3) != 6 {
		t.Errorf("TestMapString got %v, want %v", m3, data)
	}
}
