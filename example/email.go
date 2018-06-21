package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestEmail struct {
	Email string `valid:"Email"`
}

func Test_Email(){
	v := new(valid.Validation)
	fmt.Println(v.Valid(&TestEmail{Email:"11@qq.com"}))
	fmt.Println(v.Valid(&TestEmail{Email:"11@qqcom"}))
}