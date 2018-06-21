package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestAlphaDash struct {
	AlphaDash string `valid:"AlphaDash"`
}

func Test_AlphaDash(){
	v := new(valid.Validation)
	fmt.Println(v.Valid(&TestAlphaDash{AlphaDash:"1234aB-_"}))
	fmt.Println(v.Valid(&TestAlphaDash{AlphaDash:"a,1-@ $"}))
}
