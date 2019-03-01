package example

import (
	"github.com/golyu/valid"
)

type TestBase64 struct {
	Base64 string `valid:"Base64"`
}

func Test_Base64(v valid.IValidation) {
	Println(v.Valid(&TestBase64{Base64: "c3VjaHVhbmdqaUBnbWFpbC5jb20="}))
	Println(v.Valid(&TestBase64{Base64: "suchuangji@gmail.com"}))
}
