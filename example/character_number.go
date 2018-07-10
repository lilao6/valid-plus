package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestCharacterNumber struct {
	Value string `json:"value" valid:"CharacterNumber"`
	Up    string `json:"value" valid:"CharacterNumber(U)"`
	Low   string `json:"value" valid:"CharacterNumber(L)"`
}

func Test_CharacterNumber() {
	v := new(valid.Validation)
	fmt.Println(v.Valid(&TestCharacterNumber{Value: "asZ09"}))
	fmt.Println(v.Valid(&TestCharacterNumber{Up: "AZS09"}))
	fmt.Println(v.Valid(&TestCharacterNumber{Low: "asz09"}))
	fmt.Println(v.Valid(&TestCharacterNumber{Value: "asz,09"}))
	fmt.Println(v.Valid(&TestCharacterNumber{Low: "Asz09"}))
	fmt.Println(v.Valid(&TestCharacterNumber{Up: "AZSs09"}))
}
