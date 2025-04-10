package c3po

import (
	"reflect"
	"strings"
	"sync"
)

type Rule struct {
	Name     string //
	Value    string //
	Message  string // ex.: {field} require a value >= 18
	Validate func(rv reflect.Value, rule string) bool
}

var (
	rules map[string]*Rule
	once  sync.Once
)

func initRules() {
	rules = map[string]*Rule{
		"required": {
			Name:     "required",
			Message:  "{field} is required",
			Validate: req,
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
	}
}

// ex:
//
//		c3po.SetRules("min",&c3po.Rule{Message:"min value: {value}"}) // '{value}' be replace by Rule.Value
//		c3po.SetRules("max",&c3po.Rule{Message:"max value: {value}"}) // '{value}' be replace by Rule.Value
//		c3po.SetRules("format",&c3po.Rule{Validate:func(value any) bool {...}})
//
//		type User struct {
//			Age int `c3po:"min=18"`
//	 }
func SetRule(field string, rule *Rule) {
	once.Do(initRules)
	rule.Name = strings.ToLower(field)
	rules[field] = rule
}

func GetRule(rule string) *Rule {
	once.Do(initRules)
	return rules[rule]
}
