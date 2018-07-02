package rule

import (
	"errors"
	"regexp"
	"strconv"
)

type TelRule struct {
	value string
	FullTag
}

func (r *TelRule) Clone() Rule {
	clone := *r
	return &clone
}

/*
* <p>区号+座机号码</p>
*/
var telRegexp = regexp.MustCompile("^(0\\d{2,3}(\\-)?)?\\d{7,8}$")

func (r *TelRule) Tag() string {
	return "Tel"
}

func (r *TelRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate Tel:value is nil")
	}
	switch value.(type) {
	case string:
		r.value = value.(string)
	case int, int8, int16, int32, int64:
		r.value = strconv.FormatInt(intAll2int(value), 10)
	case uint, uint8, uint16, uint32, uint64:
		r.value = strconv.FormatUint(uintAll2uint(value), 10)
	default:
		return errors.New("Generate Tel:Only number and string data types are supported")
	}
	return nil
}

func (r *TelRule) Valid() error {
	if r.value == "" {
		return errors.New("Validation tel:value is nil")
	}
	if !telRegexp.MatchString(r.value) {
		return errors.New("Validation tel:the string value verification failed")
	}
	return nil
}
