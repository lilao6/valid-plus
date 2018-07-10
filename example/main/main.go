package main

import (
	"github.com/golyu/valid/rule"
	"github.com/golyu/valid/example"
)

func main() {
	rule.Model = rule.RULE_DEBUG
	example.Test_Max()
	example.Test_Min()
	example.Test_Length()
	example.Test_Alpha()
	example.Test_Email()
	example.Test_Ip()
	example.Test_Base64()
	example.Test_Tel()
	example.Test_ZipCode()
	example.Test_Phone()
	example.Test_Number()
	example.Test_In()
	example.Test_Or()
	example.Test_Character()
	example.Test_CharacterNumber()
}
