package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestNumeric struct {
	Numeric string `valid:"Numeric"`
}

func Test_Numeric(){
	v := new(valid.Validation)
	fmt.Println(v.Valid(&TestNumeric{Numeric:"100"}))
	fmt.Println(v.Valid(&TestNumeric{Numeric:"100."}))
}
