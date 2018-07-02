package rule

import (
	"errors"
	"regexp"
)

var base64Regexp = regexp.MustCompile(`^(?:[A-Za-z0-99+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`)

type Base64Rule struct {
	value string
	FullTag
}

func (r *Base64Rule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *Base64Rule) Tag() string {
	return "Base64"
}
func (r *Base64Rule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate base64:value is nil")
	}
	var ok bool
	r.value, ok = value.(string)
	if !ok {
		return errors.New("Generate base64:the value generate failed")
	}
	return nil
}
func (r *Base64Rule) Valid() error {
	if r.value == "" {
		return errors.New("Validation base64:value is nil")
	}
	if !base64Regexp.MatchString(r.value) {
		return errors.New("Validation base64:the string value verification failed")
	}
	return nil
}
