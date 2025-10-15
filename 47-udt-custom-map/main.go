package main

import (
	"errors"
	"fmt"
)

func main() {

	//m1 := MyMap(make(map[string]any))
	m1 := make(MyMap)
	m1["560081"] = "Blr-1"
	m1["560083"] = "Blr-2"
	m1["560084"] = "Blr-3"
	m1["done"] = true

	if keys, err := m1.GetKeys(); err != nil {
		println(err.Error())
	} else {
		fmt.Println(keys)
	}

	if values, err := m1.GetValues(); err != nil {
		println(err.Error())
	} else {
		fmt.Println(values)
	}

	if err := m1.Delete("done"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("key successfully deleted from the map")
	}

	if keys, values, err := m1.GetKeysAndValues(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(keys)
		fmt.Println(values)
	}

	m2 := make(map[string]any)
	m2["560081"] = "Blr-1"
	m2["560083"] = "Blr-2"
	m2["560084"] = "Blr-3"
	m2["done"] = "true"

	if keys, values, err := MyMap(m2).GetKeysAndValues(); err != nil {
		println(err.Error())
	} else {
		fmt.Println(keys)
		fmt.Println(values)
	}

}

type MyMap map[string]any

func (m MyMap) GetKeys() ([]string, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys, nil
}

func (m MyMap) GetValues() ([]any, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	values := make([]any, len(m)) // 5 [nil nil nil nil nil]
	c := 0
	for _, v := range m {
		values[c] = v
		c++
	}
	return values, nil
}

func (m MyMap) Delete(k string) error {
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

func (m MyMap) GetKeysAndValues() (keys []string, values []any, err error) {
	if m == nil {
		return nil, nil, errors.New("nil map")
	}
	values = make([]any, len(m))  // 5 [nil nil nil nil nil]
	keys = make([]string, len(m)) // 5 [nil nil nil nil nil]
	c := 0
	for k, v := range m {
		keys[c] = k
		values[c] = v
		c++
	}
	return keys, values, nil
}
