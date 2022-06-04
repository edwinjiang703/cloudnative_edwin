package main

import "fmt"

func main(){
	a1 := [...]int{1,2,5,3,78,9,4,9,23,32}
	s1 := a1[:] //得到切片
	fmt.Println(s1)

	//删掉索引为2和3的5，3
	s1 = append(s1[:2],s1[4:]...)
	fmt.Println(s1)
	fmt.Println(a1)
}

