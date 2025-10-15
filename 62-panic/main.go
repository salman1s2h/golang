package main

import (
	"os"
)

// panic panics both fo1 and fo2 objects due to panic panics the whole call stack
func main() {
	defer println("end of main")
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
}

type FileOps struct {
	FileName string
}

func NewFileOps(fn string) *FileOps {
	return &FileOps{FileName: fn}
}

func (fo *FileOps) Write(p []byte) (n int, err error) {
	defer println("trying to write to the file")
	if fo == nil {
		panic("FileOps object is nil")
		//return -1, errors.New("FileOps object is nil")
	}
	if fo.FileName == "" {
		panic("FileOps object FileName is empty")
		//return -1, errors.New("FileOps object FileName is empty")
	}

	f, err := os.OpenFile(fo.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return -1, err
	}

	defer f.Close() // will discuss later
	return f.Write(p)
}
