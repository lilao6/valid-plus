package example

import (
	"github.com/golyu/valid"
)

type TestNumber struct {
	Number string `valid:"Number"`
}

func Test_Number(v valid.IValidation) {
	Println(v.Valid(&TestNumber{Number: "1234"}))
	Println(v.Valid(&TestNumber{Number: "1234a"}))
}
