package rule

import (
	"errors"
)

type CharacterNumberRule struct {
	value     string
	valueType int
	FullTag
}

func (r *CharacterNumberRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *CharacterNumberRule) Tag() string {
	return "CharacterNumber"
}

func (r *CharacterNumberRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate Character:value is nil")
	}
	var ok bool
	r.value, ok = value.(string)
	if !ok {
		return errors.New("Generate Character:the value generate failed")
	}
	tagLen := len(r.Tag())
	if "" == tagValue || len(tagValue) == tagLen {
		r.valueType = arbitrarily
	}
	//大小写标准获取
	tagValueLen := len(tagValue)
	if tagValueLen <= tagLen+2 { // 2的长度是()两个括号
		r.valueType = arbitrarily
	} else {
		// CharacterNumber(U)  -->   "U"
		tagValue = tagValue[tagLen+1 : tagValueLen-1]
	}

	switch tagValue {
	case "U":
		r.valueType = uppercase
	case "L":
		r.valueType = lowercase
	default:
		r.valueType = arbitrarily
	}
	return nil
}

func (r *CharacterNumberRule) Valid() error {
	if r.value == "" {
		return errors.New("Validation CharacterNumber:value is nil")
	}
	switch r.valueType {
	case arbitrarily:
		for _, v := range r.value {
			if !('A' <= v && v <= 'Z') && !('a' <= v && v <= 'z') && !('0' <= v && v <= '9') {
				return errors.New("Validation CharacterNumber:Data does not meet the goal")
			}
		}
	case uppercase:
		for _, v := range r.value {
			if !('A' <= v && v <= 'Z') && !('0' <= v && v <= '9') {
				return errors.New("Validation CharacterNumber:Data does not meet the goal")
			}
		}
	case lowercase:
		for _, v := range r.value {
			if !('a' <= v && v <= 'z') && !('0' <= v && v <= '9') {
				return errors.New("Validation CharacterNumber:Data does not meet the goal")
			}
		}
	default:
		return errors.New("Validation CharacterNumber:the Character type unkonwn")
	}
	return nil
}
