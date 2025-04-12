package test

import (
	"fmt"
	"testing"

	"github.com/5tkgarage/c3po"
)

type BStructValue struct {
	Age   int
	Name  string
	Float float32
}

func BenchmarkStruct(b *testing.B) {

	bs := &BStructValue{}
	sch := c3po.Validate(bs, map[string]any{
		"age":   21,
		"name":  "foo",
		"float": "12.1",
	})
	if sch.HasErrors() {
		b.Error(sch.Errors())
	}
	s, ok := sch.Value().(*BStructValue)
	fmt.Println(s, ok)
}
