package c3po

import (
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
	v := &ValidationError{
		Field: f.Name,
		Rule:  rules["required"],
	}
	return v
}

func RetInvalidType(f *Fielder) error {
	v := &ValidationError{
		Field: f.Name,
		Rule: &Rule{
			Message: fmt.Sprintf("{field} require a type: %s", f.Type),
		},
	}
	return v
}

func RetInvalidValue(f *Fielder) error {
	v := &ValidationError{
		Field: f.Name,
		Rule: &Rule{
			Message: "{field} require a value",
		},
	}
	return v
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(
		reflect.ValueOf(i).Pointer(),
	).Name()
}
