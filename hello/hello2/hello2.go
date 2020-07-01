package main

import "fmt"

func Hello2(name string) string{
	if name == ""{
		name="World"
	}
	return "Hello, " + name
}
func main(){
	fmt.Println(Hello2("Hello"))
}
