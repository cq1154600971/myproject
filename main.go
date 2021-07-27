package main

import (
	"fmt"
	"github.com/myproject/leetcode"
)

func main() {
	s := "hello world ok"
	fmt.Println(leetcode.ReplaceSpace(s))
	testmap := map[string]int64{}
	arr := []int{1,2,3,4}
	arr1 := [...]int{1,2,3}
	test(testmap, arr, arr1)
	fmt.Println(".......")
	fmt.Println(testmap)
	fmt.Println(&arr[0])
	fmt.Println(&arr1[0])
}

func test(testmap map[string]int64, arr []int, arr1 [3]int) {
	testmap["test1"] = 1
	testmap["test2"] = 2
	testmap["test3"] = 3
	testmap["test4"] = 4
	arr[1] = 5
	arr[2] = 3
	fmt.Println(&arr[0])
	arr1[0] = 5
	fmt.Println(&arr1[0])
}