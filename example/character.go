package example

import (
	"github.com/golyu/valid"
)

type TestCharacter struct {
	Up           string `valid:"Character(U)"` //大写
	Low          string `valid:"Character(L)"` //小写
	Arbitrarily  string `valid:"Character()"`  //任意大小写
	Arbitrarily1 string `valid:"Character"`    //任意大小写
}

func Test_Character(v valid.IValidation) {
	success := TestCharacter{
		Up:           "SUCCESS",
		Low:          "success",
		Arbitrarily:  "Success",
		Arbitrarily1: "Success",
	}
	Println(v.Valid(&success))
	Println(v.Valid(&TestCharacter{Up: "fail"}))
	Println(v.Valid(&TestCharacter{Low: "FAIL"}))
	Println(v.Valid(&TestCharacter{Up: "Fail"}))
	Println(v.Valid(&TestCharacter{Low: "Fail"}))
	Println(v.Valid(&TestCharacter{Arbitrarily: "Fail1"}))
	Println(v.Valid(&TestCharacter{Arbitrarily1: "Fail1"}))
}
