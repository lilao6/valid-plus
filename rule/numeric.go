package rule

import "errors"

type NumericRule struct {
	value interface{}
	FullTag
}

func (r *NumericRule) Tag() string {
	return "Numeric"
}
func (r *NumericRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate numeric:value is nil")
	}
	var ok bool
	r.value, ok = value.(string)
	if !ok {
		return errors.New("Generate numeric:the value generate failed")
	}
	return nil
}

func (r *NumericRule) Valid() error {
	if str, ok := r.value.(string); ok {
		for _, v := range str {
			if '9' < v || v < '0' {
				return errors.New("Validation numeric:Data does not meet the goal")
			}
		}
		return nil
	}
	return errors.New("Validation numeric:Only string data types are supported")
}
