package main

import (
	"errors"
	"fmt"
	"reflect"
)

func convertToInterfaceSlice(slice interface{}) ([]interface{}, error) {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return nil, errors.New("input is not a slice")
	}

	length := rv.Len()
	result := make([]interface{}, length)

	for i := 0; i < length; i++ {
		result[i] = rv.Index(i).Interface()
	}

	return result, nil
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	interfaceSlice, err := convertToInterfaceSlice(intSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(interfaceSlice)
}
