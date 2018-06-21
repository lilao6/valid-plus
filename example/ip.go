package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestIp struct {
	Ip string `valid:"IP"`
}

func Test_Ip(){
	v := valid.Validation{}
	fmt.Println(v.Valid(&TestIp{Ip:"218.88.124.41"}))
	fmt.Println(v.Valid(&TestIp{Ip:"218.256.124.41"}))
}
