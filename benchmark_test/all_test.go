package benchmark_test

import (
	"github.com/golyu/valid"
	"github.com/golyu/valid/example"
	"testing"
)

var v = valid.NewValidation()

func BenchmarkMax(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i != b.N; i++ {
		example.Test_Max(v)
	}
}

func BenchmarkMin(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_Min(v)
	}
}

func BenchmarkLength(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_Length(v)
	}
	// 未缓存校验规则前 100000	     15371 ns/op	     920 B/op	      43 allocs/op
}

func BenchmarkAlpha(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_Alpha(v)
	}
	// 100000	     10416 ns/op	     592 B/op	      22 allocs/op
}

func BenchmarkEmail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_Email(v)
	}
	// 100000	     10069 ns/op	     264 B/op	      18 allocs/op
}

func BenchmarkIp(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_Ip(v)
	}
}
func BenchmarkBase64(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_Base64(v)
	}
}
func BenchmarkTel(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_Tel(v)
	}
}

func BenchmarkZipCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_ZipCode(v)
	}
}
func BenchmarkPhone(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_Phone(v)
	}
}

func BenchmarkNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_Number(v)
	}
}

func BenchmarkIn(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_In(v)
	}
}

func BenchmarkOr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_Or(v)
	}
}

func BenchmarkCharacter(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_Character(v)
	}
}
func BenchmarkCharacterNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		example.Test_CharacterNumber(v)
	}
}

// -12表示12个CPU线程执行；5000表示总共执行了5000次；264963 ns/op，表示每次执行耗时264963纳秒；38533 B/op表示每次执行分配了38533字节内存；1608 allocs/op表示每次执行分配了1608次对象。
// 0.8版本
//BenchmarkMax-12                	   10000	    127881 ns/op	   38512 B/op	    1608 allocs/op
//BenchmarkMin-12                	   10000	    131214 ns/op	   43672 B/op	    1858 allocs/op
//BenchmarkLength-12             	  500000	      2546 ns/op	     920 B/op	      43 allocs/op
//BenchmarkAlpha-12              	 1000000	      1341 ns/op	     592 B/op	      22 allocs/op
//BenchmarkEmail-12              	 1000000	      1462 ns/op	     264 B/op	      18 allocs/op
//BenchmarkIp-12                 	 1000000	      1632 ns/op	     268 B/op	      18 allocs/op
//BenchmarkBase64-12             	 1000000	      2048 ns/op	     268 B/op	      18 allocs/op
//BenchmarkTel-12                	 1000000	      1316 ns/op	     268 B/op	      18 allocs/op
//BenchmarkZipCode-12            	 1000000	      1169 ns/op	     264 B/op	      18 allocs/op
//BenchmarkPhone-12              	  200000	      8472 ns/op	    1635 B/op	     104 allocs/op
//BenchmarkNumber-12             	 1000000	      1031 ns/op	     312 B/op	      19 allocs/op
//BenchmarkIn-12                 	  200000	      7962 ns/op	    3592 B/op	     117 allocs/op
//BenchmarkOr-12                 	  500000	      3377 ns/op	    1296 B/op	      58 allocs/op
//BenchmarkCharacter-12          	  200000	      7796 ns/op	    2520 B/op	     131 allocs/op
//BenchmarkCharacterNumber-12    	  200000	      6838 ns/op	    2016 B/op	     108 allocs/op

//BenchmarkMax-12                	   20000	     63141 ns/op	   27688 B/op	    1115 allocs/op
//BenchmarkMin-12                	   20000	     70907 ns/op	   31152 B/op	    1285 allocs/op
//BenchmarkLength-12             	 1000000	      1968 ns/op	     968 B/op	      34 allocs/op
//BenchmarkAlpha-12              	 1000000	      1742 ns/op	     816 B/op	      26 allocs/op
//BenchmarkEmail-12              	 1000000	      1311 ns/op	     320 B/op	      15 allocs/op
//BenchmarkIp-12                 	 1000000	      1498 ns/op	     325 B/op	      15 allocs/op
//BenchmarkBase64-12             	 1000000	      1925 ns/op	     325 B/op	      15 allocs/op
//BenchmarkTel-12                	 1000000	      1182 ns/op	     326 B/op	      15 allocs/op
//BenchmarkZipCode-12            	 1000000	      1007 ns/op	     352 B/op	      15 allocs/op
//BenchmarkPhone-12              	  200000	      6589 ns/op	    1735 B/op	      77 allocs/op
//BenchmarkNumber-12             	 2000000	       941 ns/op	     368 B/op	      16 allocs/op
//BenchmarkIn-12                 	  200000	      5840 ns/op	    2184 B/op	      92 allocs/op
//BenchmarkOr-12                 	  500000	      2885 ns/op	    1192 B/op	      48 allocs/op
//BenchmarkCharacter-12          	  200000	      5610 ns/op	    2344 B/op	      98 allocs/op
//BenchmarkCharacterNumber-12    	  300000	      5832 ns/op	    2136 B/op	      81 allocs/op
