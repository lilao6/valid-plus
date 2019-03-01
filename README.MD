### valid
一个可扩展的golang请求数据校验框架,通过在struct的field上按照一定的规则编写tag,来校验请求的数据是否合法

### 安装或更新
```go
go get -u github.com/golyu/valid
```
### 注意

框架使用了反射,甚至递归反射,对性能苛刻的项目,请慎重考虑,使用了beego,echo等框架请求数据绑定的框架,对请求数据需要递归校验的,需要修改框架源码实现递归绑定,如果框架不支持递归绑定,递归校验可能会无法通过
尽量不要使用未命名结构体,否则规则不会做缓存,影响性能

如果定义Must必须放在最前,ErrorCode必须放在最后
### 示例
单个规则校验
```go
    type A struct{
        Value  string   `valid:"Must;Max(5);ErrorCode(1111)"`// 最大长度为5
    }
    v := new(valid.Validation)
    b, code, err:= v.Valid(&A{Value:"aaaaa"})
    // b 是否校验通过
    // code 错误码,如果校验通过,code为0,校验不通过,则为ErrorCode中标注的错误码,没有标注则为默认错误码1000
    // err 开发人员的错误,规则编写错误或者校验的tag不支持该数据类型
```
多个规则校验
```go
    type A struct {
        Value  string   `valid:"Must;Max(5);Min(1);ErrorCode(1111)"`// 最大长度为5,最小长度为1
    }
    v := new(valid.Validation)
    b, code, err:= v.Valid(&A{Value:"aaaaa"})
```
非必传数据校验
```go
    type A struct {
        Value  string   `valid:"Max(5);Min(1);ErrorCode(1111)"`// 最大长度为5,最小长度为1,可以为零值
    }
    v := new(valid.Validation)
    b, code, err:= v.Valid(&A{Value:"aaaaa"})
```
递归类型校验
```go
    type A struct {
    		Sex int `valid:"Must;Max(2);Min(1);ErrorCode(1111)"` // 最大数值为2,最小数值为1
    	}
    	type B struct {
    		Age int `valid:"Must;Max(23);Min(18);ErrorCode(1111)"` // 最大长度为23,最小长度为18
    		A
    	}
    	v := new(valid.Validation)
    	b, code, err := v.Valid(
    		&B{
    			Age: 15,
    			A:   A{Sex: 1},
    		})
```
组合校验Or
```go
    type TestOr struct {
    	Or string `valid:"Or<Phone Tel>"`//只要是手机号码或座机号码其中的一个
    }
```

#### Must
- 说明

    是否必传,如果校验中写了Must,则必须通过后面的所有校验规则,如果没有写Must,filed为零值时,不需要校验其它规则,field不为零值才校验其它规则
- 支持的数据类型

    所有其它校验的超集
#### Max(?)
- 说明

    最大不能大过?,?是数值
- 支持的数据类型
    int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64,string,slice
- 备注
    
    数值类型校验的是大小,字符串和切片类型,校验的是长度,float类型是近似值,精度是0.00001,可以根据具体情况调整

#### Min(?)
- 说明
    
    最小不能小过?,?是数值
- 支持的数据类型
    
    int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64,string,slice
- 备注
    
    数值类型校验的是大小,字符串和切片类型,校验的是长度,float类型是近似值,精度是0.00001,可以根据具体情况调整

#### Length(?)
- 说明

    指定大小为?,?是正整数
- 支持的数据类型

    string,slice
- 备注
    
    校验长度
    
#### In(?,?,...)
- 说明

    验证的数据必须是?...中的一个,用逗号隔开
- 支持的数据类型

    int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64
- 备注
    
    float类型是近似值,精度是0.00001,可以根据具体情况调整

#### Alpha
- 说明
    
    alpha字符
- 支持的数据类型

    string

#### Base64
- 说明
    
    Base64字符
- 支持的数据类型

    string
#### Email
- 说明
    
    邮箱字符串
- 支持的数据类型

    string    
#### IP
- 说明
    
    IP字符串
- 支持的数据类型

    string           
#### Number
- 说明
    
    字符串必须是可以转换成数值的
- 支持的数据类型

    string 
    
#### Phone
- 说明
    
    字符串必须是有效的手机号码,可以加86或者+86前缀
- 支持的数据类型

    string     
#### Tel
- 说明
    
    字符串必须是有效的固定座机电话,可以加区号前缀
- 支持的数据类型

    string 

#### ZipCode
- 说明
    
    字符串必须是有效的邮政编码
- 支持的数据类型

    string     
#### Or<? ? ? ...>
- 说明
    
    只要满足?中的任意一个规则,就可以通过,用空格隔开
- 支持的数据类型
    根据所有?支持的数据类型决定

### 拓展

valid还支持灵活的拓展,如果你有新的校验规则需要加入,不需要改动valid源代码,只需要在程序初始化的时候注册到valid就可以在后续使用该规则

例:增加一个qq号码校验规则
```go
    
    func customRule() {
    	qqRule := new(QQRule)
    	rule.RegisterRule(qqRule)
    
    	type Xx struct {
    		MyQQ string `valid:"Must;QQ;ErrorCode(1111)"`
    	}
    	v := new(valid.Validation)
    	fmt.Println(v.Valid(&Xx{MyQQ: "2131121212"}))
    }
    
    // qq号码校验规则
    type QQRule struct {
    	value string
    	rule.FullTag
    }
    
    var qqRegexp = regexp.MustCompile("[1-9][0-9]{4,14}")
    
    func (*QQRule) Tag() string {
    	return "QQ"
    }
    func (r *QQRule) Clone() rule.Rule {
    	clone := *r
    	return &clone
    }

    
    func (r *QQRule) Generate(value interface{}, tagValue string) error {
    	if value == nil {
    		return errors.New("Generate QQ:value is nil")
    	}
    	var ok bool
    	r.value, ok = value.(string)
    	if !ok {
    		return errors.New("Generate QQ:the value generate failed")
    	}
    	return nil
    }
    
    func (r *QQRule) Valid() error {
    	if r.value == "" {
    		return errors.New("Validation QQ:value is nil")
    	}
    	if !qqRegexp.MatchString(r.value) {
    		return errors.New("Validation QQ:the string value verification failed")
    	}
    	return nil
    }
```       

### Benchmark

- v0.8

>  BenchmarkMax-12                	   10000	    127881 ns/op	   38512 B/op	    1608 allocs/op
> BenchmarkMin-12                	   10000	    131214 ns/op	   43672 B/op	    1858 allocs/op
> BenchmarkLength-12             	  500000	      2546 ns/op	     920 B/op	      43 allocs/op
> BenchmarkAlpha-12              	 1000000	      1341 ns/op	     592 B/op	      22 allocs/op
> BenchmarkEmail-12              	 1000000	      1462 ns/op	     264 B/op	      18 allocs/op
> BenchmarkIp-12                 	 1000000	      1632 ns/op	     268 B/op	      18 allocs/op
> BenchmarkBase64-12             	 1000000	      2048 ns/op	     268 B/op	      18 allocs/op
> BenchmarkTel-12                	 1000000	      1316 ns/op	     268 B/op	      18 allocs/op
> BenchmarkZipCode-12            	 1000000	      1169 ns/op	     264 B/op	      18 allocs/op
> BenchmarkPhone-12              	  200000	      8472 ns/op	    1635 B/op	     104 allocs/op
> BenchmarkNumber-12             	 1000000	      1031 ns/op	     312 B/op	      19 allocs/op
> BenchmarkIn-12                 	  200000	      7962 ns/op	    3592 B/op	     117 allocs/op
> BenchmarkOr-12                 	  500000	      3377 ns/op	    1296 B/op	      58 allocs/op
> BenchmarkCharacter-12          	  200000	      7796 ns/op	    2520 B/op	     131 allocs/op
> BenchmarkCharacterNumber-12    	  200000	      6838 ns/op	    2016 B/op	     108 allocs/op

- v0.9

> BenchmarkMax-12                	   20000	     63141 ns/op	   27688 B/op	    1115 allocs/op
> BenchmarkMin-12                	   20000	     70907 ns/op	   31152 B/op	    1285 allocs/op
> BenchmarkLength-12             	 1000000	      1968 ns/op	     968 B/op	      34 allocs/op
> BenchmarkAlpha-12              	 1000000	      1742 ns/op	     816 B/op	      26 allocs/op
> BenchmarkEmail-12              	 1000000	      1311 ns/op	     320 B/op	      15 allocs/op
> BenchmarkIp-12                 	 1000000	      1498 ns/op	     325 B/op	      15 allocs/op
> BenchmarkBase64-12             	 1000000	      1925 ns/op	     325 B/op	      15 allocs/op
> BenchmarkTel-12                	 1000000	      1182 ns/op	     326 B/op	      15 allocs/op
> BenchmarkZipCode-12            	 1000000	      1007 ns/op	     352 B/op	      15 allocs/op
> BenchmarkPhone-12              	  200000	      6589 ns/op	    1735 B/op	      77 allocs/op
> BenchmarkNumber-12             	 2000000	       941 ns/op	     368 B/op	      16 allocs/op
> BenchmarkIn-12                 	  200000	      5840 ns/op	    2184 B/op	      92 allocs/op
> BenchmarkOr-12                 	  500000	      2885 ns/op	    1192 B/op	      48 allocs/op
> BenchmarkCharacter-12          	  200000	      5610 ns/op	    2344 B/op	      98 allocs/op
> BenchmarkCharacterNumber-12    	  300000	      5832 ns/op	    2136 B/op	      81 allocs/op