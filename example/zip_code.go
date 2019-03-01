package example

import (
	"github.com/golyu/valid"
)

type TestZipCode struct {
	ZipCode string `valid:"ZipCode"`
}

func Test_ZipCode(v valid.IValidation) {
	Println(v.Valid(&TestZipCode{ZipCode: "536000"}))
	Println(v.Valid(&TestZipCode{ZipCode: "00008888"}))
}
