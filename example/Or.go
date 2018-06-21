package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestOr struct {
	Or string `valid:"Or<Numeric Alpha>"`
}

func Test_Or() {
	v := new(valid.Validation)
	fmt.Println(v.Valid(&TestOr{Or: "123"}))
	fmt.Println(v.Valid(&TestOr{Or: "aaaa"}))
	fmt.Println(v.Valid(&TestOr{Or: "123a"}))
}
