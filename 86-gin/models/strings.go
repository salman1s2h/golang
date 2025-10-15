package models

import (
	"fmt"
	"strings"
)

func ReverseStr(str string) string {
	rstr := ""
	for _, v := range str {
		rstr = string(v) + rstr
	}
	return rstr
}

func ConcatStr() {
	str := ""
	for i := 1; i <= 100000; i++ {
		str += fmt.Sprint(i)
	}
}

func ConcatStrBuilder() {
	str := strings.Builder{}
	for i := 1; i <= 100000; i++ {
		str.Write([]byte(fmt.Sprint(i)))
	}
}
