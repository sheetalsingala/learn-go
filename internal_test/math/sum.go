package main 

import(
	"fmt"
)

func main(){
	res:=sum(4, 5)
	fmt.Println(res)
}

func sum(a, b int) int {
  return a + b
}
