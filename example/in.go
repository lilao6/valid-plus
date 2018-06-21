package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestIn struct {
	InInt      int      `valid:"In(1,2,3)"`
	InString   string   `valid:"In(1,2,3)"`
	InInt8     int8     `valid:"In(1,2,3)"`
	InInt16    int16    `valid:"In(1,2,3)"`
	InInt32    int32    `valid:"In(1,2,3)"`
	InInt64    int64    `valid:"In(1,2,3)"`
	InUint     uint     `valid:"In(1,2,3)"`
	InUint8    uint8    `valid:"In(1,2,3)"`
	InUint16   uint16   `valid:"In(1,2,3)"`
	InUint32   uint32   `valid:"In(1,2,3)"`
	InUint64   uint64   `valid:"In(1,2,3)"`
	InFloat32  float32  `valid:"In(1,2.0,3)"`
	InFloat64  float64  `valid:"In(1,2.0,3)"`
}

func Test_In() {
	v := new(valid.Validation)
	success := TestIn{
		InInt: 3,
		InString:"3",
		InInt8:3,
		InInt16:3,
		InInt32:3,
		InInt64:3,
		InUint:3,
		InUint8:3,
		InUint16:3,
		InUint32:3,
		InUint64:3,
		InFloat32:2.0,
		InFloat64:2.0,
	}
	fmt.Println(v.Valid(&success))
}
