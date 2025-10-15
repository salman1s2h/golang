package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	str1 := "Hello"
	str2 := " World"

	str3 := str1 + str2 // just for understanding that + operator can work on strings
	// str3 = str1 - str2  // - operator cannot work on string
	fmt.Println(str3)

	// This is a monomorphization way of doing things but Go does not do that
	// AddG(10, (20.34))                  // AddG1(int,int)int
	// AddG(float32(10.4), float32(12.4)) // AddG2(float32,float32)float32
	// AddG(10.4, 12.4)                   // AddG3(float64,float64)float64
	// AddG("Hello ", "World")            // AddG4(string,string)string
	// type eraser way.. Go's way of doing things
	// AddG(10, 20)                       // AddG(any,any)any
	// AddG(float32(10.4), float32(12.4)) // AddG(any,any)any
	if r, err := AddTS(10, 20); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("result of AddTS:", r)
	}
}

func Add(i any, j any) any {
	v1, ok1 := i.(int)
	v2, ok2 := j.(int)

	if ok1 && ok2 {
		return v1 + v2 // Compiler undderstands it
	}
	v3, ok1 := i.(float32)
	v4, ok2 := j.(float32)
	if ok1 && ok2 {
		return v3 + v4 // Compiler undderstands it
	}
	v5, ok1 := i.(float64)
	v6, ok2 := j.(float64)
	if ok1 && ok2 {
		return v5 + v6 // Compiler undderstands it
	}
	return nil
}

func AddTS(i any, j any) (any, error) {
	if IsNumber(i) && IsNumber(j) {
		if reflect.TypeOf(i) != reflect.TypeOf(j) {
			return nil, errors.New("i and j must be the same type")
		}

		switch v := i.(type) {
		case int:
			return i.(int) + j.(int), nil
		case uint:
			return i.(uint) + j.(uint), nil
		case uint8:
			return uint16(i.(uint8) + j.(uint8)), nil
		case uint16:
			return uint32(i.(uint16) + j.(uint16)), nil
		case uint32:
			return uint64(i.(uint32) + j.(uint32)), nil
		case uint64:
			return i.(uint64) + j.(uint64), nil
		case int8:
			return int16(i.(int8) + j.(int8)), nil
		case int16:
			return int32(i.(int16) + j.(int16)), nil
		case int32:
			return int64(i.(int32) + j.(int32)), nil
		case int64:
			return i.(int64) + j.(int64), nil
		case float32:
			return float64(i.(float32) + j.(float32)), nil
		case float64:
			return v + j.(float64), nil
		}
	} else {
		return nil, fmt.Errorf("i or j is not a number.Type of i:%T and type of j: %T", i, j)
		//return nil, errors.New("i or j is not a number.So cannot perform arithmetic operation")
	}

	return nil, nil
}

func AddG[T int | float32 | float64 | string](i, j T) T {
	fmt.Printf("type of i %T and type of j %T\n", i, j)
	return i + j
}

func IsNumber(i any) bool {
	switch i.(type) {
	case int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}

// Type Eraser
