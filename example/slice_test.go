package example

import (
	"github.com/golyu/valid"
	"testing"
)

type A struct {
	B string `json:"b" valid:"Must;Min(2);ErrorCode(1001)"`
	C int    `json:"c" valid:"Must;Min(2);ErrorCode(1002)"`
}

func TestSlice(t *testing.T) {
	as := []A{
		{B: "ttt", C: 3},
		{B: "z", C: 2},
	}
	v := valid.NewValidation()
	t.Log(v.Valid(as))
}
