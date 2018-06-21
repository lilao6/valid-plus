package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestLength struct {
	LengthString    string        `valid:"Length(4)"`
	LengthInterface []interface{} `valid:"Length(5)"`
}

func Test_Length() {
	v := new(valid.Validation)
	success := TestLength{
		LengthString: "测试成功",
		LengthInterface:[]interface{}{1,2,3,4,5},
	}
	fmt.Println(v.Valid(&success))
	fmt.Println(v.Valid(&TestLength{LengthString:"测试失败情况"}))
	fmt.Println(v.Valid(&TestLength{LengthInterface:[]interface{}{1,2}}))
}
