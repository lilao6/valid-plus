package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestMax struct {
	MaxString  string   `valid:"Max(5)"`
	MaxInt     int      `valid:"Max(5)"`
	MaxInt8    int8     `valid:"Max(5)"`
	MaxInt16   int16    `valid:"Max(5)"`
	MaxInt32   int32    `valid:"Max(5)"`
	MaxInt64   int64    `valid:"Max(5)"`
	MaxUint    uint     `valid:"Max(5)"`
	MaxUint8   uint8    `valid:"Max(5)"`
	MaxUint16  uint16   `valid:"Max(5)"`
	MaxUint32  uint32   `valid:"Max(5)"`
	MaxUint64  uint64   `valid:"Max(5)"`
	MaxInts    []int    `valid:"Max(5)"`
	MaxInts8   []int8   `valid:"Max(5)"`
	MaxInts16  []int16  `valid:"Max(5)"`
	MaxInts32  []int32  `valid:"Max(5)"`
	MaxInts64  []int64  `valid:"Max(5)"`
	MaxUints   []uint   `valid:"Max(5)"`
	MaxUints8  []uint8  `valid:"Max(5)"`
	MaxUints16 []uint16 `valid:"Max(5)"`
	MaxUints32 []uint32 `valid:"Max(5)"`
	MaxUints64 []uint64 `valid:"Max(5)"`
}

func Test_Max() {
	v := new(valid.Validation)
	success := TestMax{
		MaxString:  "测试",
		MaxInt:     1,
		MaxInt8:    2,
		MaxInt16:   3,
		MaxInt32:   4,
		MaxInt64:   5,
		MaxUint:    1,
		MaxUint8:   2,
		MaxUint16:  3,
		MaxUint32:  4,
		MaxUint64:  5,
		MaxInts:    []int{1, 2, 3, 4},
		MaxInts8:   []int8{1, 2, 3, 4},
		MaxInts16:  []int16{1, 2, 3, 4},
		MaxInts32:  []int32{1, 2, 3, 4},
		MaxInts64:  []int64{1, 2, 3, 4},
		MaxUints:   []uint{1, 2, 3, 4},
		MaxUints8:  []uint8{1, 2, 3, 4},
		MaxUints16: []uint16{1, 2, 3, 4},
		MaxUints32: []uint32{1, 2, 3, 4},
		MaxUints64: []uint64{1, 2, 3, 4},
	}
	fmt.Println(v.Valid(&success))
	fmt.Println(v.Valid(&TestMax{MaxString: "测试不通过情况"}))
	fmt.Println(v.Valid(&TestMax{MaxInt: 6}))
	fmt.Println(v.Valid(&TestMax{MaxInt8: 6}))
	fmt.Println(v.Valid(&TestMax{MaxInt16: 6}))
	fmt.Println(v.Valid(&TestMax{MaxInt32: 6}))
	fmt.Println(v.Valid(&TestMax{MaxInt64: 6}))

	fmt.Println(v.Valid(&TestMax{MaxUint: 6}))
	fmt.Println(v.Valid(&TestMax{MaxInt8: 6}))
	fmt.Println(v.Valid(&TestMax{MaxUint16: 6}))
	fmt.Println(v.Valid(&TestMax{MaxUint32: 6}))
	fmt.Println(v.Valid(&TestMax{MaxUint64: 6}))

	fmt.Println(v.Valid(&TestMax{MaxInts: []int{1, 2, 3, 4, 5, 6}}))
	fmt.Println(v.Valid(&TestMax{MaxInts8: []int8{1, 2, 3, 4, 5, 6}}))
	fmt.Println(v.Valid(&TestMax{MaxInts16: []int16{1, 2, 3, 4, 5, 6}}))
	fmt.Println(v.Valid(&TestMax{MaxInts32: []int32{1, 2, 3, 4, 5, 6}}))
	fmt.Println(v.Valid(&TestMax{MaxInts64: []int64{1, 2, 3, 4, 5, 6}}))

	fmt.Println(v.Valid(&TestMax{MaxUints: []uint{1, 2, 3, 4, 5, 6}}))
	fmt.Println(v.Valid(&TestMax{MaxUints8: []uint8{1, 2, 3, 4, 5, 6}}))
	fmt.Println(v.Valid(&TestMax{MaxUints16: []uint16{1, 2, 3, 4, 5, 6}}))
	fmt.Println(v.Valid(&TestMax{MaxUints32: []uint32{1, 2, 3, 4, 5, 6}}))
	fmt.Println(v.Valid(&TestMax{MaxUints64: []uint64{1, 2, 3, 4, 5, 6}}))
}
