package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str1 := "Hello, World!"
	str2 := "Hello, 世界!"
	str3 := `Have\n you thought about learning Mandarin Chinese?
When learning a new language, greetings are often the first thing to learn. Here’s my hello to the world in Mandarin Chinese: 你好，世界 (Nǐ Hǎo, Shì Jiè)
My name is Anna, I’m a Business and Technical Translator between English and Mandarin Chinese. The world’s 2 most spoken languages.
In my blog, I’d like to share a few tips with you on learning the language. First of all: Saying Hello! 你好 (Nǐ Hǎo) . Listen here.
你好 is the Mandarin Chinese writing, let’s call it The Characters. Nǐ Hǎo is the pronunciation using the modern Chinese called Pin Yin.
Chinese Characters are a visual and graphical writing language, the current square form has evolved from original drawing like form. In the image below it shows you how to write the word Hello 你好.`

	fmt.Println(str1, "len:", len(str1), "Size:", unsafe.Sizeof(str1))
	fmt.Println(str2, "len:", len(str2), "Size:", unsafe.Sizeof(str2))
	fmt.Println(str3, "len:", len(str3), "Size:", unsafe.Sizeof(str3))

	str1 = "Hello, wrold!"
	// deallocate old str1 memory
	fmt.Println(str1, "len:", len(str1))
	c := 0 // creates a new variable
	{      // a new stack frame is creaged
		a, b := 123, 23 // creates two new integers
		{               // creats a new stackframe
			d := a + b // creats a new variable
			c = d
		} // deallocates d
	} // dealloacates a,b
	println(c)

	var str4 string // dummy addresses are created and assigned to Ptr and assign Len=0

} // deallocates str1,str2,c

// utf8 --> What is the size of utf8? 1-4
// string
// Is String mutable?

/* String Header
Ptr *int // 8 bytes
Len int // 8 bytes
*/
