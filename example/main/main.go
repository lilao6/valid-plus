package main

import (
	"github.com/golyu/valid"
	"github.com/golyu/valid/example"
	"github.com/golyu/valid/rule"
)

func main() {
	v := valid.NewValidation()
	rule.Model = rule.RULE_DEBUG
	example.Test_Max(v)
	example.Test_Min(v)
	example.Test_Length(v)
	example.Test_Alpha(v)
	example.Test_Email(v)
	example.Test_Ip(v)
	example.Test_Base64(v)
	example.Test_Tel(v)
	example.Test_ZipCode(v)
	example.Test_Phone(v)
	example.Test_Number(v)
	example.Test_In(v)
	example.Test_Or(v)
	example.Test_Character(v)
	example.Test_CharacterNumber(v)
}
