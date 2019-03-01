package example

import (
	"github.com/golyu/valid"
)

type TestCharacterNumber struct {
	Value string `json:"value" valid:"CharacterNumber"`
	Up    string `json:"value" valid:"CharacterNumber(U)"`
	Low   string `json:"value" valid:"CharacterNumber(L)"`
}

func Test_CharacterNumber(v valid.IValidation) {
	Println(v.Valid(&TestCharacterNumber{Value: "asZ09"}))
	Println(v.Valid(&TestCharacterNumber{Up: "AZS09"}))
	Println(v.Valid(&TestCharacterNumber{Low: "asz09"}))
	Println(v.Valid(&TestCharacterNumber{Value: "asz,09"}))
	Println(v.Valid(&TestCharacterNumber{Low: "Asz09"}))
	Println(v.Valid(&TestCharacterNumber{Up: "AZSs09"}))
}
