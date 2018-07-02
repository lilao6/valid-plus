package rule

import (
	"regexp"
	"errors"
	"strconv"
)

var zipCodeRegexp = regexp.MustCompile(`^[1-9]\d{5}$`)

type ZipCodeRule struct {
	value string
	FullTag
}

func (r *ZipCodeRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *ZipCodeRule) Tag() string {
	return "ZipCode"
}

func (r *ZipCodeRule) Generate(value interface{}, tagValue string) error {
	if nil == value {
		return errors.New("Generate ZipCode:value is nil")
	}
	switch value.(type) {
	case string:
		r.value = value.(string)
	case int, int8, int16, int32, int64:
		r.value = strconv.FormatInt(intAll2int(value), 10)
	case uint, uint8, uint16, uint32, uint64:
		r.value = strconv.FormatUint(uintAll2uint(value), 10)
	default:
		return errors.New("Generate ZipCode:Only number and string data types are supported")
	}
	return nil
}

func (r *ZipCodeRule) Valid() error {
	if r.value == "" {
		return errors.New("Validation ZipCode:value is nil")
	}
	if !zipCodeRegexp.MatchString(r.value) {
		return errors.New("Validation ZipCode:the string value verification failed")
	}
	return nil
}
