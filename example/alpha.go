package example

import (
	"github.com/golyu/valid"
)

type TestAlpha struct {
	Alpha string `valid:"Or<Tel Phone>;er,Alphat;Character(U);"`
}

func Test_Alpha(v valid.IValidation) {
	Println(v.Valid(&TestAlpha{Alpha: "success"}))
	Println(v.Valid(&TestAlpha{Alpha: "fail."}))
}
