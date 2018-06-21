package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestNumber struct {
	Number string `valid:"Number"`
}

func Test_Number(){
	v := new(valid.Validation)
	fmt.Println(v.Valid(&TestNumber{Number:"1234"}))
	fmt.Println(v.Valid(&TestNumber{Number:"1234a"}))
}