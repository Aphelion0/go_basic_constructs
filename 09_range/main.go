package main

import "fmt"

func main(){
	ids := []int{1,2,3,4,5,6,2,4,62,1,435632,342}
	for _,id := range ids{
		fmt.Printf("ID:%d\n",id)
	}
}