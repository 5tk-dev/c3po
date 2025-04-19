package encoder

import (
	"testing"
)

func TestPrimitives(t *testing.T) {
	if v := Encode("batata"); v != "batata" {
		t.Errorf("got %q, want %q", v, "batata")
	}
	if v := Encode(1); v != 1 {
		t.Errorf("got %q, want %v", v, 1)
	}
	if v := Encode(1.2); v != 1.2 {
		t.Errorf("got %q, want %v", v, 1.2)
	}
	if v := Encode(true); v != true {
		t.Errorf("got %q, want %v", v, true)
	}
	if v := Encode(false); v != false {
		t.Errorf("got %q, want %v", v, false)
	}
	if v := Encode(nil); v != nil {
		t.Errorf("got %q, want %v", v, nil)
	}
}

func TestSlicePrimitives(t *testing.T) {

	if v, ok := Encode([]string{"batata", "frita"}).([]any); !ok || len(v) != 2 {
		t.Errorf("got %v, want %v", v, []any{"batata", "frita"})
	}

	if v, ok := Encode([]int{1, 2, 3, 4, 5, 6}).([]any); !ok || len(v) != 6 {
		t.Errorf("got %v, want %v", v, []any{1, 2, 3, 4, 5, 6})
	}
	if v, ok := Encode([]float64{1.2, 3.4}).([]any); !ok || len(v) != 2 {
		t.Errorf("got %f, want %v", v, []any{1.2, 3.4})
	}
	if v, ok := Encode([]bool{true, false}).([]any); !ok || len(v) != 2 {
		t.Errorf("got %v, want %v", v, []any{true, false})
	}
	if v, ok := Encode([]any{"batata", 1, 1.2, true}).([]any); !ok || len(v) != 4 {
		t.Errorf("got %v, want %v", v, []any{"batata", 1, 1.2, true})
	}
}

func TestMapStruct(t *testing.T) {
	a := struct {
		Name string
		Age  int
	}{"etho", 40}
	v := Encode(a)
	if v2, ok := v.(map[string]any); !ok {
		t.Errorf("got %v, want map[string]any", v)
	} else if v2["name"] != "etho" && v2["age"] != 40 {
		t.Errorf("got %v, want map[string]any{name: etho, age:40}", v2)
	}
}

type User struct {
	Age  int
	Name string
}

type Employe struct {
	*User
	Wage float64
}

type Firm struct {
	Employees []*Employe
}

func TestMapStructComplex(t *testing.T) {
	firm := &Firm{
		Employees: []*Employe{
			{
				User: &User{
					Name: "etho",
					Age:  52,
				},
				Wage: 100000.01,
			},
			{
				User: &User{
					Name: "etho2",
					Age:  52,
				},
				Wage: 100000.03,
			},
			{
				User: &User{
					Name: "etho",
					Age:  22,
				},
				Wage: 100.0,
			},
		},
	}

	m := Encode(firm)
	v, ok := m.(map[string]any)
	if !ok {
		t.Errorf("got %v, want map[string]any", v)
		return
	}
	es, ok := v["employees"]
	if !ok {
		t.Errorf("got %v, want map[string]any", es)
	}

	e := es.([]any)[0].(map[string]any)
	w := e["wage"]
	if w2, ok := w.(float64); !ok || w2 != 100000.01 {
		t.Errorf("got %v, want 100000.01", w2)
	}

	u := e["user"]
	if u == nil {
		t.Errorf("got %v, want map[string]any", u)
	} else if u2, ok := u.(map[string]any); !ok {
		t.Errorf("got %v, want map[string]any", u2)
	} else if n := u2["name"]; n != "etho" {
		t.Errorf("got %v, want etho", n)
	}
}
