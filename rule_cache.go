package valid

import (
	"errors"
	"github.com/golyu/valid/rule"
	"reflect"
	"strconv"
	"strings"
)

// 缓存field上面的tag规则
type RuleCache struct {
	isMust  bool        // 是否必传
	errCode int64       // 错误码
	rules   []rule.Rule // 规则
}

// genRule 切割并生成各个规则,并获取是否存在必传和错误码
// error 解析tag出错
func genRule(f reflect.StructField) (RuleCache, error) {
	ruleCache := RuleCache{}
	ruleCache.rules = make([]rule.Rule, 0)

	tag := f.Tag.Get(VALID_TAG)
	if len(tag) > 0 {
		fs := strings.Split(tag, SPLIT_SEP)
		if len(fs) > 0 {
			for i, vTag := range fs {
				switch i {
				case 0: //首个,是否必传
					if fs[0] == MUST {
						ruleCache.isMust = true
						continue
					}
				case len(fs) - 1: // 最后一个,校验码
					if strings.HasPrefix(vTag, "ErrorCode(") && strings.HasSuffix(vTag, ")") {
						codeStr := vTag[10 : len(vTag)-1]
						var err error
						ruleCache.errCode, err = strconv.ParseInt(codeStr, 10, 64)
						if err != nil {
							return ruleCache, errors.New("Parse ErrorCode tag failure:" + vTag)
						}
						continue
					}
				}
				// 解析单个
				r, err := rule.GetRule(vTag)
				if err != nil {
					return ruleCache, errors.New("Parse rule tag failure:" + vTag)
				}
				ruleCache.rules = append(ruleCache.rules, r)
			}

		}
	}
	return ruleCache, nil
}
