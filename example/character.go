package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestCharacter struct {
	Up           string `valid:"Character(U)"` //大写
	Low          string `valid:"Character(L)"` //小写
	Arbitrarily  string `valid:"Character()"`  //任意大小写
	Arbitrarily1 string `valid:"Character"`    //任意大小写
}

func Test_Character() {
	success := TestCharacter{
		Up:           "SUCCESS",
		Low:          "success",
		Arbitrarily:  "Success",
		Arbitrarily1: "Success",
	}
	v := new(valid.Validation)
	fmt.Println(v.Valid(&success))
	fmt.Println(v.Valid(&TestCharacter{Up: "fail"}))
	fmt.Println(v.Valid(&TestCharacter{Low: "FAIL"}))
	fmt.Println(v.Valid(&TestCharacter{Up: "Fail"}))
	fmt.Println(v.Valid(&TestCharacter{Low: "Fail"}))
	fmt.Println(v.Valid(&TestCharacter{Arbitrarily: "Fail1"}))
	fmt.Println(v.Valid(&TestCharacter{Arbitrarily1: "Fail1"}))
}
