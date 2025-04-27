package test

import (
	"testing"

	"5tk.dev/c3po"
)

func TestUnmarshalSliceOfStrings(t *testing.T) {
	s := []string{}
	c3po.UnmarshalValidate(&s, []any{"batata", 1, "3456"})
	if len(s) != 3 {
		t.Errorf("got %d, want len==3", len(s))
		return
	}
	if s[0] != "batata" {
		t.Errorf("got %q, want %s", s[0], "batata")
	}
	if s[1] != "1" {
		t.Errorf("got %q, want %s", s[1], "1")
	}
	if s[2] != "3456" {
		t.Errorf("got %q, want %s", s[2], "3456")
	}
}
func TestUnmarshalSliceOfInts(t *testing.T) {
	s := []int{}
	c3po.UnmarshalValidate(&s, []any{1, "2", "4849849"})
	if len(s) != 3 {
		t.Errorf("got %d, want len==3", len(s))
		return
	}
	if s[0] != 1 {
		t.Errorf("got %d, want %s", s[0], "1")
	}
	if s[1] != 2 {
		t.Errorf("got %d, want %s", s[1], "2")
	}
	if s[2] != 4849849 {
		t.Errorf("got %d, want %s", s[2], "4849849")
	}
}

func TestUnmarshalSliceOfFloats(t *testing.T) {
	s := []float32{}
	c3po.UnmarshalValidate(&s, []any{1.2, "2.3", "56"})
	if len(s) != 3 {
		t.Errorf("got %d, want len==3", len(s))
		return
	}
	if s[0] != 1.2 {
		t.Errorf("got %f, want %s", s[0], "1")
	}
	if s[1] != 2.3 {
		t.Errorf("got %f, want %s", s[1], "2")
	}
	if s[2] != 56 {
		t.Errorf("got %f, want %s", s[2], "4849849")
	}
}

func TestUnmarshalSliceOfMaps(t *testing.T) {
	s := []map[string]any{}

	c3po.UnmarshalValidate(&s, []any{
		map[string]any{"hello": "world"},
		map[int]any{1: "world"},
		map[any]any{true: "false"},
	})

	if len(s) != 3 {
		t.Errorf("got %d, want len==3", len(s))
		return
	}
}

func TestUnmarshalSliceOfStructs(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	s := []User{}
	c3po.UnmarshalValidate(&s, []map[string]any{
		{"name": "etho", "age": "1234"},
		{"name": "joao", "age": "21"},
		{"name": "frederico", "age": "56"},
	})
	if len(s) != 3 {
		t.Errorf("got %d, want len==3", len(s))
		return
	}
}

func TestUnmarshalSliceOfStructsPtr(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	s := []*User{}
	c3po.UnmarshalValidate(&s, []map[string]any{
		{"name": "etho", "age": "1234"},
		{"name": "joao", "age": "21"},
		{"name": "frederico", "age": "56"},
	})
	if len(s) != 3 {
		t.Errorf("got %d, want len==3", len(s))
		return
	}
}

func TestUnmarshalSliceOfComplexData(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	type Employer struct {
		User
		Wage float32
	}
	type Firm struct {
		Owner     User
		Cashier   float32
		Employees []*Employer
	}

	groupCompany := []*Firm{}

	firm := map[string]any{
		"owner": map[string]any{
			"name": "etho",
			"age":  "99",
		},
		"cashier": "-1234.56",
		"employees": []map[string]any{
			{
				"wage": 1510.1,
				"user": map[string]any{
					"name": "etho1",
					"age":  "18",
				},
			},
			{
				"wage": 1510.1,
				"user": map[string]any{
					"name": "etho2",
					"age":  "22",
				},
			},
			{
				"wage": 1510.1,
				"user": map[string]any{
					"name": "etho3",
					"age":  "33",
				},
			},
		},
	}

	c3po.UnmarshalValidate(&groupCompany, []map[string]any{firm, firm, firm})
	if len(groupCompany) != 3 {
		t.Errorf("got %d, want len==3", len(groupCompany))
		return
	}
	if len(groupCompany[0].Employees) != 3 {
		t.Errorf("got %d, want len==3", len(groupCompany))
	}
	if groupCompany[1].Owner.Name != "etho" {
		t.Errorf("got %s, want etho", groupCompany[1].Owner.Name)
	}
	if groupCompany[2].Cashier != -1234.56 {
		t.Errorf("got %f, want -1234.56", groupCompany[2].Cashier)
	}
}
