package test

import (
	"testing"

	"5tk.dev/c3po"
)

func TestDecodeStruct(t *testing.T) {
	u := &UserTest{Name: "5tk", Age: 18}

	r, err := c3po.Encode(u)
	if err != nil {
		t.Error(err)
	}

	m := r.(map[string]any)

	if m["name"] != "5tk" {
		t.Errorf("map['name'] got %v, want: '5tk'", m["name"])
	}
	if m["age"] != 18 {
		t.Errorf("map['age'] got %v, want: 18", m["age"])
	}
}

func TestDecodeSlice(t *testing.T) {
	u := []*UserTest{
		{Name: "5", Age: 18},
		{Name: "5t", Age: 18},
		{Name: "5tk", Age: 18},
		{Name: "5tkG", Age: 18},
		{Name: "5tkGa", Age: 18},
		{Name: "5tkGar", Age: 18},
		{Name: "5tkGara", Age: 18},
		{Name: "5tkGarag", Age: 18},
		{Name: "5tkGarage", Age: 18},
	}

	r, err := c3po.Encode(u)
	if err != nil {
		t.Error(err)
	}

	m := r.([]any)

	if len(m) != 9 {
		t.Errorf("m got %v, want: []map[string]any{}9x", m)
	}
}
