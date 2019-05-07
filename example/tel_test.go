package example

import (
	"github.com/golyu/valid"
	"testing"
)

type TestTel struct {
	Tel string `valid:"Tel"`
}

func Test_Tel(t *testing.T) {
	v := valid.NewValidation()
	t.Log(v.Valid(&TestTel{Tel: "022-00008888"}))
	t.Log(v.Valid(&TestTel{Tel: "222-70008888"}))
	t.Log(v.Valid(&TestTel{Tel: "777777777"}))
}
