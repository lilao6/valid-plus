package example

import (
	"github.com/golyu/valid"
	"fmt"
)

type TestMin struct {
	MinString  string   `valid:"Min(5)"`
	MinInt     int      `valid:"Min(5)"`
	MinInt8    int8     `valid:"Min(5)"`
	MinInt16   int16    `valid:"Min(5)"`
	MinInt32   int32    `valid:"Min(5)"`
	MinInt64   int64    `valid:"Min(5)"`
	MinUint    uint     `valid:"Min(5)"`
	MinUint8   uint8    `valid:"Min(5)"`
	MinUint16  uint16   `valid:"Min(5)"`
	MinUint32  uint32   `valid:"Min(5)"`
	MinUint64  uint64   `valid:"Min(5)"`
	MinInts    []int    `valid:"Min(5)"`
	MinInts8   []int8   `valid:"Min(5)"`
	MinInts16  []int16  `valid:"Min(5)"`
	MinInts32  []int32  `valid:"Min(5)"`
	MinInts64  []int64  `valid:"Min(5)"`
	MinUints   []uint   `valid:"Min(5)"`
	MinUints8  []uint8  `valid:"Min(5)"`
	MinUints16 []uint16 `valid:"Min(5)"`
	MinUints32 []uint32 `valid:"Min(5)"`
	MinUints64 []uint64 `valid:"Min(5)"`
}

func Test_Min() {
	v := new(valid.Validation)
	success := TestMin{
		MinString:  "测试成功情况",
		MinInt:     6,
		MinInt8:    6,
		MinInt16:   6,
		MinInt32:   6,
		MinInt64:   6,
		MinUint:    6,
		MinUint8:   6,
		MinUint16:  6,
		MinUint32:  6,
		MinUint64:  6,
		MinInts:    []int{1, 2, 3, 4, 5, 6},
		MinInts8:   []int8{1, 2, 3, 4, 5, 6},
		MinInts16:  []int16{1, 2, 3, 4, 5, 6},
		MinInts32:  []int32{1, 2, 3, 4, 5, 6},
		MinInts64:  []int64{1, 2, 3, 4, 5, 6},
		MinUints:   []uint{1, 2, 3, 4, 5, 6},
		MinUints8:  []uint8{1, 2, 3, 4, 5, 6},
		MinUints16: []uint16{1, 2, 3, 4, 5, 6},
		MinUints32: []uint32{1, 2, 3, 4, 5, 6},
		MinUints64: []uint64{1, 2, 3, 4, 5, 6},
	}
	fmt.Println(v.Valid(&success))
	fmt.Println(v.Valid(&TestMin{MinString: "测试失败"}))
	fmt.Println(v.Valid(&TestMin{MinInt: 4}))
	fmt.Println(v.Valid(&TestMin{MinInt8: 4}))
	fmt.Println(v.Valid(&TestMin{MinInt16: 4}))
	fmt.Println(v.Valid(&TestMin{MinInt32: 4}))
	fmt.Println(v.Valid(&TestMin{MinInt64: 4}))

	fmt.Println(v.Valid(&TestMin{MinUint: 4}))
	fmt.Println(v.Valid(&TestMin{MinInt8: 4}))
	fmt.Println(v.Valid(&TestMin{MinUint16: 4}))
	fmt.Println(v.Valid(&TestMin{MinUint32: 4}))
	fmt.Println(v.Valid(&TestMin{MinUint64: 4}))

	fmt.Println(v.Valid(&TestMin{MinInts: []int{1, 2, 3, 4}}))
	fmt.Println(v.Valid(&TestMin{MinInts8: []int8{1, 2, 3, 4}}))
	fmt.Println(v.Valid(&TestMin{MinInts16: []int16{1, 2, 3, 4}}))
	fmt.Println(v.Valid(&TestMin{MinInts32: []int32{1, 2, 3, 4}}))
	fmt.Println(v.Valid(&TestMin{MinInts64: []int64{1, 2, 3, 4}}))

	fmt.Println(v.Valid(&TestMin{MinUints: []uint{1, 2, 3, 4}}))
	fmt.Println(v.Valid(&TestMin{MinUints8: []uint8{1, 2, 3, 4}}))
	fmt.Println(v.Valid(&TestMin{MinUints16: []uint16{1, 2, 3, 4}}))
	fmt.Println(v.Valid(&TestMin{MinUints32: []uint32{1, 2, 3, 4}}))
	fmt.Println(v.Valid(&TestMin{MinUints64: []uint64{1, 2, 3, 4}}))
}
