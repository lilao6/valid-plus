package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestBase64 struct {
	Base64 string `valid:"Base64"`
}

func Test_Base64(){
	v := new(valid.Validation)
	fmt.Println(v.Valid(&TestBase64{Base64:"c3VjaHVhbmdqaUBnbWFpbC5jb20="}))
	fmt.Println(v.Valid(&TestBase64{Base64:"suchuangji@gmail.com"}))
}
