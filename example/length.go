package example

import (
	"github.com/golyu/valid"
)

type TestLength struct {
	LengthString    string        `valid:"Length(4)"`
	LengthInterface []interface{} `valid:"Length(5)"`
}

func Test_Length(v valid.IValidation) {
	success := TestLength{
		LengthString:    "测试成功",
		LengthInterface: []interface{}{1, 2, 3, 4, 5},
	}
	Println(v.Valid(&success))
	Println(v.Valid(&TestLength{LengthString: "测试失败情况"}))
	Println(v.Valid(&TestLength{LengthInterface: []interface{}{1, 2}}))
}
