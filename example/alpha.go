package example

import (
	"fmt"
	"github.com/golyu/valid"
)

type TestAlpha struct {
	Alpha string `valid:"Or<Tel Phone>;er,Alphat;Character(U);"`
}

func Test_Alpha(){
	v := new(valid.Validation)

	fmt.Println(v.Valid(&TestAlpha{Alpha:"success"}))
	fmt.Println(v.Valid(&TestAlpha{Alpha:"fail."}))
}