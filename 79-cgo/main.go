package main

/*
#cgo LDFLAGS: -Lbuild/ -lmylib
#include <stdlib.h>

// Declare the C functions
int add(int a, int b);
char* concat(const char* str1, const char* str2);
typedef struct {
    int x;
    int y;
} Point;
void movePoint(Point* p, int dx, int dy);
void printArray(int* arr, int length);
*/
import "C" // psudo package
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
