package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestTel struct {
	Tel string `valid:"Tel"`
}

func Test_Tel(){
	v := new(valid.Validation)
	fmt.Println(v.Valid(&TestTel{Tel:"022-00008888"}))
	fmt.Println(v.Valid(&TestTel{Tel:"222-70008888"}))
}