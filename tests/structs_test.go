package test

import (
	"testing"

	"5tk.dev/c3po"
)

type TagOmit struct {
	Field string `validate:"-"`
}

// tag "-"
func TestStructTag_omit(t *testing.T) {
	var s = &TagOmit{}
	sch := c3po.Validate(s, map[string]string{"field": "batata"})
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
	if v := sch.Value().(*TagOmit); v.Field != "" {
		t.Errorf("TagOmit.Field: got %q, want %q", v.Field, "")
	}
}

type TagName struct {
	Field string `validate:"name=uuid"`
}

// tag "name"
func TestStructTag_name(t *testing.T) {
	var s = &TagName{}
	uid := "000-000"
	sch := c3po.Validate(s, map[string]any{"uuid": uid})
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
	if v := sch.Value().(*TagName); v.Field != uid {
		t.Errorf("TagName.Field: got %q, want %q", v.Field, uid)
	}
}

type TagEscape struct {
	Field string `validate:"escape"`
}

// tag "escape"
func TestStructTag_escape(t *testing.T) {
	var s = &TagEscape{}
	html := "<h1>Hello</h1>"
	sch := c3po.Validate(s, map[string]string{"field": html})
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
	if v := sch.Value().(*TagEscape); v.Field == html {
		t.Errorf("TagEscape.Field: got %q, want %q", v.Field, html)
	}
}

type TagRequired struct {
	Field string `validate:"required"`
}

// tag "required"
func TestStructTag_required(t *testing.T) {
	var s = &TagRequired{}
	sch := c3po.Validate(s, map[string]string{"field": ""}) // same result map[string]string{}
	if !sch.HasErrors() {
		v := sch.Value().(*TagRequired)
		t.Errorf("TagRequired.Field: got %q, want %q", v.Field, "error")
	}
}

type TagNullable struct {
	Field *TagEscape `validate:"nullable"`
}

// tag "nullable"
func TestStructTag_nullable(t *testing.T) {
	var s = &TagNullable{}
	sch := c3po.Validate(s, map[string]any{})
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
	v := sch.Value().(*TagNullable)
	if v.Field != nil {
		t.Errorf("TagNullable.Field: got %v, want \"nil\"", v.Field)
	}
}

type TagSkiperr struct {
	Field int `validate:"skiperr"`
}

// tag "skiperr"
func TestStructTag_skiperr(t *testing.T) {
	var s = &TagSkiperr{}
	sch := c3po.Validate(s, map[string]any{"field": nil})
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
}

type UserTest struct {
	Name string
	Age  int
}

type TagRecursive struct {
	User UserTest `validate:"recursive"`
	Role int
}

// tag "recursive"
func TestStructTag_recurcive(t *testing.T) {
	var s = &TagRecursive{}
	sch := c3po.Validate(s, map[string]any{"name": "r2d2", "age": 18, "role": 1})
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
	v := sch.Value().(*TagRecursive)
	if v.Role != 1 || v.User.Age != 18 || v.User.Name != "r2d2" {
		t.Errorf("TagRecursive.Field: got %v, want %v", v, TagRecursive{
			User: UserTest{
				Name: "r2d2",
				Age:  18,
			},
			Role: 1,
		})
	}
}

type TagMinMax struct {
	Num int `validate:"min=1,max=2"`
}

func TestTagStructTag_minMax(t *testing.T) {
	s := &TagMinMax{}
	sch0 := c3po.Validate(s, map[string]string{})
	sch1 := c3po.Validate(s, map[string]string{"num": "1"})
	sch2 := c3po.Validate(s, map[string]string{"num": "2"})
	sch3 := c3po.Validate(s, map[string]string{"num": "3"})

	if !sch0.HasErrors() {
		t.Errorf("TagMinMax.Field: got %v, want min error", sch0.Value().(*TagMinMax).Num)
	}
	if sch1.HasErrors() {
		t.Error(sch1.Errors())
	}
	if sch2.HasErrors() {
		t.Error(sch2.Errors())
	}
	if !sch3.HasErrors() {
		t.Errorf("TagMinMax.Field: got %v, want min error", sch1.Value().(*TagMinMax).Num)
	}
}
