package main

import "fmt"

func main() {

	map1 := make(map[[3]string][3][3]int)

	map1[[3]string{"a", "b", "c"}] = [3][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(map1)
	map2 := make(map[string]map[string]any)

	map2["m1"] = map[string]any{"1": 1, "2": 2}
	map2["m2"] = map[string]any{"name": "Jiten", "Phone": "9618558500"}

	v, ok := map2["m2"]
	if ok {
		for k2, v2 := range v {
			fmt.Println("Key:", k2, "Value:", v2)
		}
	}

	for _, v := range map2 {
		for k2, v2 := range v {
			fmt.Println("Key:", k2, "Value:", v2)
		}
	}

}
