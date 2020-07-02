package main

func Sum(numbers []int) int{
	sum:=0
	for _, item:= range numbers {
		sum += item
	}
	return sum
}	
