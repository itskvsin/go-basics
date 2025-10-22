package main

import "fmt"

func main(){

	//for array
	fruits:= []string{"apple","banana","cherry"}

	for i,fruit := range fruits{
		fmt.Println(i,fruit)
	}

	//for MAP
	scores := map[string] int{"alice": 90, "john": 100}

	for name,score := range scores{
		fmt.Println(name,": ",score)
	}
}