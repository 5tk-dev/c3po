package c3po

import (
	"fmt"
	"sync"
)

var (
	once             sync.Once
	muFielder        sync.RWMutex
	fielderResgistry map[string]*Fielder
)

func initRegistry() {
	once.Do(func() {
		fielderResgistry = map[string]*Fielder{}
	})
}

func RegisterFielder(key string, schema *Fielder) error {
	initRegistry()

	muFielder.Lock()
	defer muFielder.Unlock()
	_, ok := fielderResgistry[key]
	if ok {
		return fmt.Errorf("fielder %q already registered", key)
	}
	fielderResgistry[key] = schema
	return nil
}

func GetFielder(key string) *Fielder {
	muFielder.RLock()
	defer muFielder.RUnlock()
	return fielderResgistry[key]
}

func ValidateSchema(key string, data any) Schema {
	f := GetFielder(key)
	if f == nil {
		return &schema{
			errors: []error{
				fmt.Errorf("fielder %q is undefined", key),
			},
		}
	}
	return f.Decode(data)
}
