package main

import "fmt"

func main(){
	x:=5
	y:=10
	if x< y {
		fmt.Printf("%d is less than %d\n",x,y)
	} else {
		fmt.Printf("%d is less than %d\n",y,x)
	}

	color:="green"
	switch color {
	case "red":
		fmt.Println("RED\n")
	case "green":
		fmt.Println("GREEN\n")
	case "blue":
		fmt.Println("BLUE\n")
	default:
		fmt.Println("NOT BLUE OR RED OR GREEN\n")
	}
}