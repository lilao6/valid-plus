package example

import (
	"github.com/golyu/valid"
)

type TestPhone struct {
	Phone       string `valid:"Phone"`
	PhoneInt    int    `valid:"Phone"`
	PhoneInt64  int64  `valid:"Phone"`
	PhoneUint   uint   `valid:"Phone"`
	PhoneUint64 uint64 `valid:"Phone"`
}

func Test_Phone(v valid.IValidation) {
	success := TestPhone{
		Phone:       "+8615182552591",
		PhoneInt:    15182652591,
		PhoneInt64:  15182652591,
		PhoneUint:   15182652591,
		PhoneUint64: 15182652591,
	}
	Println(v.Valid(&success))
	Println(v.Valid(&TestPhone{Phone: "+8615182552591"}))
	Println(v.Valid(&TestPhone{Phone: "8615182552591"}))
}
