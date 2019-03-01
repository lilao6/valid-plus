package example

import (
	"github.com/golyu/valid/rule"
	"log"
)

func Println(v ...interface{}) {
	if rule.Model {
		log.Println(v...)
	}
}
