package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		num0   uint    = 12312
		num1   uint8   = 12
		num2   uint16  = 1233
		num3   uint32  = 12312
		num4   uint64  = 12321312
		num5   int     = 12312131
		num6   int8    = 10
		num7   int16   = 12312
		num8   int32   = 32432
		num9   int64   = 9394433
		float1 float32 = 43434.34
		float2 float64 = 34324324.23432
		any1   any     = 234234234234
		any2   any     = 234323.234234
		//any3   any     = "12312.123"
		str1 string = "1231334"
		str2 string = "123123.1231"
		ok1  bool   = true
		sum  float64
	)

	sum = float64(num0) + float64(num1) + float64(num2) + float64(num3) + float64(num4) + float64(num5) + float64(num6) + float64(num7) + float64(num8) + float64(num9) + float64(float1) + float2 + float64(any1.(int)) + any2.(float64)
	n1, err := strconv.ParseInt(str1, 10, 64)
	if err == nil {
		sum += float64(n1)
	}
	n2, _ := strconv.ParseFloat(str2, 64)
	sum += n2
	if ok1 {
		sum += 1
	}
	fmt.Printf("Sum of all types:%.4f\n", sum)

}
