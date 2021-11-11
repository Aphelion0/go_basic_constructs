package main

import "fmt"

func main(){
	var fruitArr[2] string
	fruitArr[0] = "Apple"
	fruitArr[0] = "JackFruit"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])

	fruitArr2 := [2]string{"Apple","JackFruit"}
	fmt.Println(fruitArr2)
	fmt.Println(fruitArr2[0])

	fruitSlice := []string{"Apple","JackFruit","DragonFruit"}
	fmt.Println(fruitSlice[1:3])
	fmt.Println(fruitSlice[0:2])
}