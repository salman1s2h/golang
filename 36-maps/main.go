package main

import (
	"errors"
	"fmt"
)

func main() {

	var map1 map[string]any // declaration of the map
	if map1 == nil {
		println("map is nil")
	}
	map1 = make(map[string]any)

	map1["560086"] = "Bangalore-1"
	map1["560096"] = "Bangalore-2"
	map1["560034"] = "Bangalore-3"
	map1["560006"] = "Bangalore-4"
	map1["560001"] = 1

	val, ok := map1["560086"]
	if ok {
		fmt.Println(val)
	} else {
		println("key does not exist")
	}

	val, ok = map1["560066"]
	if ok {
		fmt.Println(val)
	} else {
		println("key does not exist")
	}

	for key, value := range map1 {
		fmt.Println("Key:", key, "Value:", value)
	}

	delete(map1, "560001") // delete function does not return any thing.If the map is nil delete does not do anything
	// if the key does not exist delete does nothing
	delete(map1, "abcd")
	fmt.Println("After deleting a key")
	for key, value := range map1 {
		fmt.Println("Key:", key, "Value:", value)
	}

	if err := DeleteElem(map1, "asda"); err != nil {
		println(err.Error())
	} else {
		println("key successfully deleted")
	}

	keys := GetKeys(map1)
	values := GetValues(map1)
	fmt.Println(keys)
	fmt.Println(values)

	keys1, values1 := GetKeysAndValues(map1)

	fmt.Println(keys1)
	fmt.Println(values1)

}

// What type can be the key?
// any type that satisifies == operator can be the key
// a key cannot be duplicated, if to same key a value is assigneed multiple times.. the value gets updated on the key
// map is not ordered
// map is not thread safe

func DeleteElem(m map[string]any, k string) error {
	if m == nil {
		return errors.New("nil map")
	}
	_, ok := m[k]
	if !ok {
		return fmt.Errorf("The following key %v does not exist", k)
	}
	delete(m, k)
	return nil
}

func GetKeys(m map[string]any) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func GetValues(m map[string]any) []any {
	values := make([]any, len(m)) // 5 [nil nil nil nil nil]
	c := 0
	for _, v := range m {
		values[c] = v
		c++
	}
	return values
}

func GetKeysAndValues(m map[string]any) (keys []string, values []any) {
	values = make([]any, len(m))  // 5 [nil nil nil nil nil]
	keys = make([]string, len(m)) // 5 [nil nil nil nil nil]
	c := 0
	for k, v := range m {
		keys[c] = k
		values[c] = v
		c++
	}
	return keys, values
}
