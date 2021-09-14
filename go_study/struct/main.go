package main

import (
	"fmt"
	"sort"
)

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) []student {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	ce = append(ce, student{3, "xiaowang", 56})
	return ce
}
//func main() {
//	var ce []student  //定义一个切片类型的结构体
//	ce = []student{
//		student{1, "xiaoming", 22},
//		student{2, "xiaozhang", 33},
//	}
//	fmt.Println(ce)
//	fmt.Printf("len:%v, cap:%v\n", len(ce), cap(ce))
//	ce = demo(ce)
//	fmt.Println(ce)
//	fmt.Printf("len:%v, cap:%v\n", len(ce), cap(ce))
//}
func main() {
	ce := make(map[int]student)
	ce[1] = student{1, "xiaolizi", 22}
	ce[2] = student{2, "wang", 23}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
	// 顺序打印map
	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"
	sli := []int{}
	for k, _ := range map1 {
		sli = append(sli, k)
	}
	sort.Ints(sli)
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[sli[i]])
	}
}