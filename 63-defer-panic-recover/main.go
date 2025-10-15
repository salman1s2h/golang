package main

import (
	"fmt"
	"log"
	"os"
)

// panic panics both fo1 and fo2 objects due to panic panics the whole call stack
func main() {
	defer println("end of main")

	var fo0 *FileOps // It is nil
	if _, err := fo0.Write([]byte("Hello World!")); err != nil {
		println(err.Error())
	} else {
		println("data successfully written to the file")
	}

	fo1 := NewFileOps("")
	if _, err := fo1.Write([]byte("Hello World!")); err != nil {
		println(err.Error())
	} else {
		println("data successfully written to the file")
	}

	fo2 := NewFileOps("data.text")
	if _, err := fo2.Write([]byte("Hello World!")); err != nil {
		println(err.Error())
	} else {
		println("data successfully written to the file")
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()

		slice := []int{1, 2, 3, 4, 5}

		for i := 0; i <= len(slice); i++ {
			println(slice[i])
		}
	}()

}

type FileOps struct {
	FileName string
}

func NewFileOps(fn string) *FileOps {
	return &FileOps{FileName: fn}
}

func (fo *FileOps) Write(p []byte) (n int, err error) {
	defer RecoverMe()
	defer println("trying to write to the file")
	if fo == nil {
		//panic("FileOps object is nil")
		panic(PanicInfo{Obj: fo, Msg: "FileOps object is nil"})
		//return -1, errors.New("FileOps object is nil")
	}
	if fo.FileName == "" {
		panic(PanicInfo{Obj: fo, Msg: "FileOps object FileName is empty"})
		//return -1, errors.New("FileOps object FileName is empty")
	}

	f, err := os.OpenFile(fo.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return -1, err
	}

	defer f.Close() // will discuss later
	return f.Write(p)
}

func RecoverMe() {
	if r := recover(); r != nil {
		panicinfo := r.(PanicInfo) // assert it back
		log.Println(panicinfo.Msg)
		log.Println(panicinfo.Obj)
	}
}

type PanicInfo struct {
	Obj any
	Msg string
}
