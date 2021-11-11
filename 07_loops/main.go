package main

import "fmt"

func main(){
	i := 1
	for i<=10 {
		fmt.Println(i);
		i++
	}

	for i := 1; i<=100 ; i++ {
		if(i%15==0){
			fmt.Printf("FIZZBUZZ\n")
		} else if(i%5==0){
			fmt.Printf("BUZZ\n")
		}else if(i%3==0){
			fmt.Printf("FIZZ\n")
		} else{
			fmt.Printf("Number%d\n",i)
		}
	}
}