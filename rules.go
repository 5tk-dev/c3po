package c3po

import (
	"fmt"
	"reflect"
	"strings"
)

type typer int

const (
	Setter typer = iota // before
	Formatter
)

type Rule struct {
	Name     string //
	Typer    typer
	Value    string //
	Message  string // ex.: {field} require a value >= 18
	Validate func(rv reflect.Value, rule string) bool
}

func (r *Rule) ToMap() map[string]any {
	return map[string]any{
		"name":         r.Name,
		"value":        r.Value,
		"errorMessage": r.Message,
	}
}

var defaultRules = map[string]bool{
	"":         true,
	"required": true,
}

var (
	rules = map[string]*Rule{
		"required": {
			Name:     "required",
			Message:  "{field} is required",
			Validate: req,
		},
		"escape": {
			Name:     "escape",
			Message:  "{field} do not be replaced",
			Validate: escape,
		},
		"min": {
			Name:     "min",
			Message:  "{field} requires a value >= {value}",
			Validate: min,
		},
		"max": {
			Name:     "max",
			Message:  "{field} requires a value <= {value}",
			Validate: max,
		},
		"minlen": {
			Name:     "minlen",
			Message:  "{field} requires a length value >= {value}",
			Validate: minLen,
		},
		"maxlen": {
			Name:     "maxlen",
			Message:  "{field} requires a length value <= {value}",
			Validate: maxLen,
		},
	}
)

// ex:
//
//		c3po.SetRules("min",&c3po.Rule{Message:"min value: {value}"}) // '{value}' be replace by "Rule.Value"
//		c3po.SetRules("max",&c3po.Rule{Message:"max value: {value}"}) // '{value}' be replace by "Rule.Value"
//		c3po.SetRules("format",&c3po.Rule{Validate:func(value any) bool {...}})
//
//		type User struct {
//			Age int `c3po:"min=18"`
//	 }
func SetRule(field string, rule *Rule) {
	rule.Name = strings.ToLower(field)
	if _, ok := defaultRules[rule.Name]; ok {
		panic(fmt.Errorf("%q ia a invalid tag rule", rule.Name))
	}
	rules[rule.Name] = rule
}

func GetRule(rule string) *Rule { return rules[rule] }
