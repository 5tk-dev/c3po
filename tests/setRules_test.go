package test

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"5tk.dev/c3po"
)

type SetRuleNowStruct struct {
	Now time.Time `validate:"now"`
}

func TestRuleNow(t *testing.T) {
	c3po.SetRule("now", &c3po.Rule{
		Validate: func(rv reflect.Value, rule string) bool {
			v := reflect.ValueOf(time.Now())
			rv.Set(v)
			return true
		},
	})
	s := &SetRuleNowStruct{}
	sch := c3po.Validate(s, map[string]any{})
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
}

type SetRuleCpfStruct struct {
	Cpf string `validate:"cpf,minlength=11,maxlength=14"`
}

func TestRuleCpf(t *testing.T) {
	c3po.SetRule("cpf", &c3po.Rule{
		Validate: func(rv reflect.Value, rule string) bool {
			v := rv.Interface().(string)
			if len(v) != 9 {
				v = strings.ReplaceAll(v, ".", "")
				v = strings.ReplaceAll(v, "-", "")
			}
			rv.Set(reflect.ValueOf(v))
			return true
		},
	})
	sr := &SetRuleCpfStruct{}

	cpf0 := map[string]string{"cpf": "000.000.000-00"}
	cpf1 := map[string]string{"cpf": "11111111111"}

	sch0 := c3po.Validate(sr, cpf0)
	sch1 := c3po.Validate(sr, cpf1)

	if sch0.HasErrors() {
		t.Error(sch0.Errors())
	}
	if sch1.HasErrors() {
		t.Error(sch1.Errors())
	}
	if v, ok := sch0.Value().(*SetRuleCpfStruct); ok {
		if v.Cpf != "00000000000" {
			t.Errorf("SetRuleCpfStruct got %v, want %v", v.Cpf, "00000000000")

		}
	}
	if v, ok := sch1.Value().(*SetRuleCpfStruct); ok {
		if v.Cpf != "11111111111" {
			t.Errorf("SetRuleCpfStruct got %v, want %v", v.Cpf, "11111111111")

		}
	}
}

type SetRuleEmailStruct struct {
	Email string `validate:"email,minlength=10,maxlength=64"`
}

func TestRuleEmail(t *testing.T) {
	c3po.SetRule("email", &c3po.Rule{
		Validate: func(rv reflect.Value, rule string) bool {
			v := rv.Interface().(string)
			strs := strings.Split(v, "@")
			if len(strs) != 2 {
				return false
			}
			if len(strs[0]) < 4 {
				return false
			}
			if len(strs[1]) < 4 {
				return false
			}
			return true
		},
	})

	sr := &SetRuleEmailStruct{}
	sch0 := c3po.Validate(sr, map[string]string{"email": "foobar@example.com"})
	if sch0.HasErrors() {
		t.Error(sch0.Errors())
	}
	if v := sch0.Value().(*SetRuleEmailStruct); v.Email != "foobar@example.com" {
		t.Errorf("SetRuleEmailStruct got %v, want %v", v.Email, "foobar@example.com")
	}

	sch1 := c3po.Validate(sr, map[string]string{"email": "foo.com"})
	if !sch1.HasErrors() {
		v := sch1.Value().(*SetRuleEmailStruct)
		t.Errorf("SetRuleEmailStruct got %v, want %v", v.Email, "nil")
	}

}
