package main

import "fmt"

func main(){
	hasDuplicates([]int{1,2,4,4})
}

func hasDuplicates(nums []int){
	var i int;
	for i=0; i<len(nums);i++{
		fmt.Println(nums[i])
	}
}
