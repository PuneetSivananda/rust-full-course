package main

import "fmt"

func main(){
	fmt.Println(hasDuplicates([]int{1,2,3,4}))
}

func hasDuplicates(nums []int) bool {
	var i int;
	seen:= make(map[int]bool);
	for i=0; i<len(nums);i++{
		if seen[nums[i]]{
			return true
		}
		seen[nums[i]] = true
	}
	return false
}
