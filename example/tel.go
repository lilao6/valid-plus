package example

import (
	"github.com/golyu/valid"
)

type TestTel struct {
	Tel string `valid:"Tel"`
}

func Test_Tel(v valid.IValidation) {
	Println(v.Valid(&TestTel{Tel: "022-00008888"}))
	Println(v.Valid(&TestTel{Tel: "222-70008888"}))
}
