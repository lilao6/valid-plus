package rule

import (
	"errors"
)

const (
	arbitrarily = 0 //任意大小写
	uppercase   = 1 //小写
	lowercase   = 2 //大写
)

type CharacterRule struct {
	character     string
	characterType int
	FullTag
}

func (r *CharacterRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *CharacterRule) Tag() string {
	return "Character"
}

// 	大写:Character(U) 小写:Character(L) //任意大小写:Character or Character()
func (r *CharacterRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate Character:value is nil")
	}
	var ok bool
	r.character, ok = value.(string)
	if !ok {
		return errors.New("Generate Character:the value generate failed")
	}
	tagLen := len(r.Tag())
	if "" == tagValue || len(tagValue) == tagLen {
		r.characterType = arbitrarily
	}
	//大小写标准获取
	tagValueLen := len(tagValue)
	if tagValueLen <= tagLen+2 { // 2的长度是()两个括号
		r.characterType = arbitrarily
	} else {
		// Character(U)  -->   "U"
		tagValue = tagValue[tagLen+1 : tagValueLen-1]
	}

	switch tagValue {
	case "U":
		r.characterType = uppercase
	case "L":
		r.characterType = lowercase
	default:
		r.characterType = arbitrarily
	}
	return nil
}

func (r *CharacterRule) Valid() error {
	if r.character == "" {
		return errors.New("Validation Character:value is nil")
	}
	switch r.characterType {
	case arbitrarily:
			for _, v := range r.character {
				if !('A' <= v && v <= 'Z') && !('a' <= v && v <= 'z') {
					return errors.New("Validation Character:Data does not meet the goal")
				}
			}
	case uppercase:
		for _, v := range r.character {
			if !('A' <= v && v <= 'Z') {
				return errors.New("Validation Character:Data does not meet the goal")
			}
		}
	case lowercase:
		for _, v := range r.character {
			if !('a' <= v && v <= 'z') {
				return errors.New("Validation Character:Data does not meet the goal")
			}
		}
	default:
		return errors.New("Validation Character:the Character type unkonwn")
	}
	return nil
}
