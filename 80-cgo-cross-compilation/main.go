package main

/*
#cgo darwin LDFLAGS: -L${SRCDIR}/lib/darwin -lmylib
#cgo linux LDFLAGS: -L${SRCDIR}/lib/linux -lmylib
#cgo windows LDFLAGS: -L${SRCDIR}/lib/windows -lmylib
#cgo CFLAGS: -I${SRCDIR}/include
#include <stdlib.h>
#include "mylib.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// Call the add function
	sum := C.add(10, 20)
	fmt.Println("Sum:", sum)

	// Call the concat function
	str1 := C.CString("Hello, ")
	str2 := C.CString("CGo!")
	defer C.free(unsafe.Pointer(str1)) // Free allocated memory
	defer C.free(unsafe.Pointer(str2)) // Free allocated memory
	result := C.concat(str1, str2)
	defer C.free(unsafe.Pointer(result)) // Free allocated memory
	fmt.Println("Concatenated string:", C.GoString(result))

	// Call the movePoint function
	point := C.Point{x: 10, y: 20}
	C.movePoint(&point, 5, 5)
	fmt.Printf("Moved point: (%d, %d)\n", point.x, point.y)

	// Call the printArray function
	goArray := []int32{1, 2, 3, 4, 5}
	cArray := (*C.int)(unsafe.Pointer(&goArray[0]))
	length := C.int(len(goArray))
	C.printArray(cArray, length)
}
