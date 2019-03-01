package example

import (
	"github.com/golyu/valid"
)

type TestOr struct {
	Or string `valid:"Or<Number Alpha>"`
}

func Test_Or(v valid.IValidation) {
	Println(v.Valid(&TestOr{Or: "123"}))
	Println(v.Valid(&TestOr{Or: "aaaa"}))
	Println(v.Valid(&TestOr{Or: "123a"}))
}
