package c3po

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	`"`, "&#34;",
	"'", "&#39;",
)

func HtmlEscape(s string) string { return htmlReplacer.Replace(s) }

// return true if err == nil, else false
func try() bool {
	return recover() == nil
}

func makeFielderRules(ftags map[string]string) map[string]*Rule {
	rs := map[string]*Rule{}
	for ruleName, rule := range rules {
		if tagValue, ok := ftags[ruleName]; ok {
			rs[ruleName] = &Rule{
				Name:     ruleName,
				Value:    tagValue,
				Message:  rule.Message,
				Validate: rule.Validate,
			}
		}
	}
	return rs
}

func parseTags(tag string) map[string]string {
	kvTags := map[string]string{}
	pairs := strings.Split(tag, ",")
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		kv := strings.Split(pair, "=")
		key := strings.ToLower(kv[0])
		if len(kv) == 1 {
			kvTags[key] = "true"
		} else {
			kvTags[key] = kv[1]
		}
	}
	return kvTags
}

func RetMissing(f *Fielder) error {
	s := fmt.Sprintf(`{"field":"%s", "type": "%s","message": "missing"}`, f.Name, f.Type.String())
	return errors.New(s)
}

func RetInvalidType(f *Fielder) error {
	s := fmt.Sprintf(`{"field":"%s", "type": "%s","message": "invalid type"}`, f.Name, f.Type.String())
	return errors.New(s)
}

func RetInvalidValue(f *Fielder) error {
	s := fmt.Sprintf(`{"field":"%s", "type": "%s","message": "invalid value"}`, f.Name, f.Type.String())
	return errors.New(s)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(
		reflect.ValueOf(i).Pointer(),
	).Name()
}
