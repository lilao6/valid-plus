package rule

import (
	"errors"
	"regexp"
)

type AlphaDashRule struct {
	value string
	FullTag
}

var alphaDashRegexp = regexp.MustCompile(`[^\d\w-_]`)

func (r *AlphaDashRule) Tag() string {
	return "AlphaDash"
}
func (r *AlphaDashRule) Generate(value interface{}, tagValue string) error {
	if nil == value {
		return errors.New("Generate alphaDash:value is nil")
	}
	var ok bool
	r.value, ok = value.(string)
	if !ok {
		return errors.New("Generate alphaDash:the value generate failed")
	}
	return nil
}

func (r *AlphaDashRule) Valid() error {
	if "" == r.value {
		return errors.New("Generate alphaDash:value is nil")
	}
	if !alphaDashRegexp.MatchString(r.value) {
		return errors.New("Validation alphaDash:the string value verification failed")
	}
	return nil
}
