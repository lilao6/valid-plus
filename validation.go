package valid

import (
	"reflect"
	"github.com/golyu/valid/rule"
	"strings"
	"strconv"
	"errors"
	"log"
	"fmt"
)

//
type Validation struct {
}

func (v *Validation) Valid(obj interface{}) (b bool, code int64, err error) {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	switch {
	case isStruct(objT):
	case isStructPtr(objT):
		objT = objT.Elem()
		objV = objV.Elem()
	default:
		err = fmt.Errorf("%v must be a struct or a struct pointer", obj)
		return
	}

	for i := 0; i < objT.NumField(); i++ {
		// 递归struct下的struct
		if isStruct(objT.Field(i).Type) || isStructPtr(objT.Field(i).Type) {
			b, code, err = v.Valid(objV.Field(i).Interface())
			if !b {
				return b, code, err
			}
		}
		isMust, errCode, rules, err := genRule(objT.Field(i))
		if err != nil {
			return false, 0, err
		}
		// 如果数据上没有加Must的tag,同时数据为零值,后面的校验都可以省略
		if IsReflectZeroValue(objV.Field(i)) {
			if isMust {
				return b, errCode, nil
			} else {
				continue
			}
		}
		//如果字段是私有 后面的省略
		if !objV.Field(i).CanInterface() {
			continue
		}
		for _, r := range rules {
			// 生成规则
			err := r.Generate(objV.Field(i).Interface(), r.GetFullTag())
			if err != nil {
				return false, 0, err
			}
			// 校验规则
			err = r.Valid()
			if err != nil {
				if rule.Model {
					log.Println("rule valid err:", err.Error())
				}
				return false, errCode, nil
			}
		}
	}
	return true, 0, nil
}

// genRule 切割并生成各个规则,并获取是否存在必传和错误码
// bool 是否必传
// int64 错误码
// []rule.Rule 规则
// error 解析tag出错
func genRule(f reflect.StructField) (bool, int64, []rule.Rule, error) {
	var defaultCode int64 = 1000
	var isMust bool // 是否必传
	var rules = make([]rule.Rule, 0)
	tag := f.Tag.Get(VALID_TAG)
	if len(tag) > 0 {
		fs := strings.Split(tag, SPLIT_SEP)
		if len(fs) > 0 {
			for i, vTag := range fs {
				switch i {
				case 0: //首个,是否必传
					if fs[0] == MUST {
						isMust = true
						continue
					}
				case len(fs) - 1: // 最后一个,校验码
					if strings.HasPrefix(vTag, "ErrorCode(") && strings.HasSuffix(vTag, ")") {
						codeStr := vTag[10 : len(vTag)-1]
						var err error
						defaultCode, err = strconv.ParseInt(codeStr, 10, 64)
						if err != nil {
							return false, 0, nil, errors.New("Parse ErrorCode tag failure:" + vTag)
						}
						continue
					}
				}
				// 解析单个
				r, err := rule.GetRule(vTag)
				if err != nil {
					return false, 0, nil, errors.New("Parse rule tag failure:" + vTag)
				}
				rules = append(rules, r)
			}

		}
	}
	return isMust, defaultCode, rules, nil
}
