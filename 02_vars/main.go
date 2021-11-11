package main

import "fmt"

func main(){
	name, size := "Alfhiem",2.33
	var age int32 = 37
	const isCool = true
	fmt.Println(name,age,isCool)
	fmt.Printf("%T\n",size)
}