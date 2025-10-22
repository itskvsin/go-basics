package main

import "fmt"

type Person struct{
	name string
	age int
}

func main(){
	p := Person{name: "kevin", age: 19}
	fmt.Println("Welcome",p.name," \nAt the age of:",p.age)
}