package main

import "fmt"

func main(){
	emails := make(map[string]string)
	emails["Entropy"] = "asgard.com"
	emails["Kratos"] = "midgard.com"
	fmt.Println(emails)
	fmt.Println(len(emails))

	delete(emails,"Entropy")
	fmt.Println(emails)

}