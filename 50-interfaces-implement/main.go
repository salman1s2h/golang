package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Fprintln(os.Stdout, "Hello World!")
	fo := NewFileOps("data.txt")

	fmt.Fprintln(fo, "Hello World!")

	var fo1 *FileOps
	_, err := fmt.Fprintln(fo1, "Hello World")
	if err != nil {
		fmt.Println(err.Error())
		ferr, ok := err.(FileError)
		if ok {
			fmt.Println(ferr.Code)
		}
	}

}

type FileOps struct {
	FileName string
}

func NewFileOps(fn string) *FileOps {
	return &FileOps{FileName: fn}
}

func (fo *FileOps) Write(p []byte) (n int, err error) {
	if fo == nil {
		return -1, FileError{Code: 101, Msg: "FileOps object is nil"}
	}
	if fo.FileName == "" {
		return -1, FileError{Code: 102, Msg: "FileOps object FileName is empty"}
	}

	f, err := os.OpenFile(fo.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return -1, err
	}

	defer f.Close() // will discuss later
	return f.Write(p)
}

type FileError struct {
	Code int
	Msg  string
}

func (f FileError) Error() string {
	return fmt.Sprintf("Error Code:%d Message:%s", f.Code, f.Msg)
}

// var Status uint8

// 0 0 0 0 0 0 0 0
// 1 0 0 0 0 0 0 0  --> active
// 0 1 0 0 0 0 0 0  --> inactive
// 0 0 1 0 0 0 0 0  --> running
// 0 0 0 1 0 0 0 0  --> paused
// 0 0 0 0 1 0 0 0  -->resumed

// Write 0
// Read 1
// Execute 2
// View 3
// Delete 4
// _ 5
// _ 6
// _ 7
// 1 1 0 0 0 0 0 0
// 1 1 1 1 1 1 0 0
//

// 4
// 2
// 1
// 0 6 4 4
