package main

import "fmt"

func greet(name string){
	fmt.Println("Hi", name)
}

func add(a ,b int) int {
	return a+b
}

func dm(a , b int) (int , int){
	return a*b, a%b
}

func main(){

	greet("Kevin")

	result:= add(1,1)
	fmt.Println(result)

	mul, div:= dm(12,24)
	fmt.Println(mul,div)

}