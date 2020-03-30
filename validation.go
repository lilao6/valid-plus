package valid

import (
	"fmt"
	"github.com/golyu/valid/rule"
	"log"
	"reflect"
	"strings"
	"sync"
)

type IValidation interface {
	Valid(obj interface{}) (b bool, code int64, err error) // -是否校验成功 -错误码 -错误信息
}

// 校验器对象
type validation struct {
	fieldTagRuleCache   map[string]RuleCache // 缓存了fields上的valid校验规则
	fieldTagRuleCacheRW sync.RWMutex         // 读写锁
}

func NewValidation() IValidation {
	//rule.Model = rule.RULE_DEBUG
	return &validation{
		fieldTagRuleCache: make(map[string]RuleCache),
	}
}

func (v *validation) Valid(obj interface{}) (b bool, code int64, err error) {
	objV := reflect.ValueOf(obj)
	objT := objV.Type()
	if isPtr(objT) {
		objT = objT.Elem()
		objV = objV.Elem()
	}
	switch {
	case isStruct(objT):
	case isSlice(objT):
		for i := 0; i < objV.Len(); i++ {
			indexObj := objV.Index(i)
			b, code, err = v.Valid(indexObj.Interface())
			if err != nil || !b {
				return
			}
		}
		return
	default:
		err = fmt.Errorf("%v must be a struct or a struct pointer", obj)
		return
	}
	//keyPrefix := objT.PkgPath() + "-" + objT.Name() + "-" // 给field缓存rule来使用
	keyPrefix := objV.Type().PkgPath() + "-" + objV.Type().Name() + "-" // 给field缓存rule来使用
	for i := 0; i < objT.NumField(); i++ {
		if !objV.Field(i).CanInterface() || strings.HasPrefix(objT.Field(i).Name, "XXX_") {
			continue
		}
		// 如果该field的类型是个struct或者struct的指针,则递归struct下的struct
		if isStruct(objT.Field(i).Type) || isStructPtr(objT.Field(i).Type) {
			if objV.Field(i).Interface() == nil { // 如果这个结果的指针为空,就跳过
				continue
			}
			b, code, err = v.Valid(objV.Field(i).Interface())
			if !b {
				return b, code, err
			}
		}
		ruleCache, err := v.genRule(keyPrefix, objT.Field(i))
		if err != nil {
			return false, 0, err
		}
		// 如果数据上没有加Must的tag,同时数据为零值,后面的校验都可以省略
		if IsReflectZeroValue(objV.Field(i)) {
			if ruleCache.isMust {
				return b, ruleCache.errCode, nil
			} else {
				continue
			}
		}
		//如果字段是私有 后面的省略
		if !objV.Field(i).CanInterface() {
			continue
		}
		for _, r := range ruleCache.rules {
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
				return false, ruleCache.errCode, nil
			}
		}
	}
	return true, 0, nil
}

func (v *validation) genRule(keyPrefix string, f reflect.StructField) (RuleCache, error) {
	if keyPrefix == "--" {
		// 使用了匿名结构体,反射不出pkg和struct name,对性能有影响,无法做缓存
		return genRule(f)
	}
	key := keyPrefix + f.Name
	// 缓存查找
	v.fieldTagRuleCacheRW.RLock()
	ruleCache, ok := v.fieldTagRuleCache[key]
	v.fieldTagRuleCacheRW.RUnlock()
	if ok {
		return ruleCache, nil
	}
	// 再次判断缓存,避免前一次等锁的时候,其它线程做了缓存功能,这次重复做工作
	v.fieldTagRuleCacheRW.RLock()
	ruleCache, ok = v.fieldTagRuleCache[key]
	v.fieldTagRuleCacheRW.RUnlock()
	if ok {
		return ruleCache, nil
	}
	ruleCache, err := genRule(f)
	if err != nil {
		return RuleCache{}, err
	}
	v.fieldTagRuleCacheRW.Lock()
	v.fieldTagRuleCache[key] = ruleCache
	v.fieldTagRuleCacheRW.Unlock()
	return ruleCache, nil
}
