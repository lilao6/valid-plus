package rule

import "errors"

type AlphaRule struct {
	value interface{}
	FullTag
}

func (r *AlphaRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *AlphaRule) Tag() string {
	return "Alpha"
}

func (r *AlphaRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate alphe:value is nil")
	}
	var ok bool
	r.value, ok = value.(string)
	if !ok {
		return errors.New("Generate alphe:the value generate failed")
	}
	return nil
}

func (r *AlphaRule) Valid() error {
	if str, ok := r.value.(string); ok {
		for _, v := range str {
			if ('Z' < v || v < 'A') && ('z' < v || v < 'a') {
				return errors.New("Validation alpha:Data does not meet the goal")
			}
		}
		return nil
	}
	return errors.New("Validation alpha:Only string data types are supported")
}
