package main

func main() {
	funcslice := make([]func(), 0)
	j := 0
loop1:
	if j < 10 {
		funcslice = append(funcslice, func() {
			println("calling a funcSlice func", j)
		})
		j++
		goto loop1
	}
	for _, v := range funcslice {
		v()
	}

	// solve the above problem

	funcslice1 := make([]func(int), 0)
	k := 0
loop2:
	if k < 10 {
		fn := func(i int) {
			println("funcslice1:", i)
		}
		funcslice1 = append(funcslice1, fn)

		k++
		goto loop2
	}
	for i, v := range funcslice1 {
		v(i)
	}

	// using loop

	funcslice2 := make([]func(), 0)
	for j := 0; j < 10; j++ {
		funcslice2 = append(funcslice2, func() {
			println("calling a funcSlice2 func", j)
		})
	}

	for _, v := range funcslice2 {
		v()
	}
}
