package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestZipCode struct {
	ZipCode string `valid:"ZipCode"`
}

func Test_ZipCode(){
	v := new(valid.Validation)
	fmt.Println(v.Valid(&TestZipCode{ZipCode:"536000"}))
	fmt.Println(v.Valid(&TestZipCode{ZipCode:"00008888"}))
}