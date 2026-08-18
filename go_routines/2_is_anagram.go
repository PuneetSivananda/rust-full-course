package main

import "fmt"

func main(){
	s:= "orange"
	dict:= make(map[string]int)
	for _,ch:= range s{
		dict[string(ch)]++
	}
	
	for ch, count:= range dict{
		fmt.Printf("%s: %d\n", ch, count)	
	}
}