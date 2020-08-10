package main

import (
	"fmt"
)

func main() {
	//使用示例
	arr := []int{1, 2, 3, 4}
	in, err := InSlice(1, arr)
	if err == nil {
		fmt.Println(in)
	} else {
		fmt.Println(err)
	}

	s := []int{1, 1, 2, 2, 3}
	err = RemoveDuplicate(&s)
	if err == nil {
		fmt.Println(s)
	} else {
		fmt.Println(err)
	}
}
