package main

import "fmt"

type User struct{
	Name string
	Age int
}

func greet(user User){
	if user.Age >= 18 {
		fmt.Println("Welcome",user.Name)
	} else {
		fmt.Println("sorry",user.Name,"You should be 18+")
	}
}

func main(){
	users := []User{
		{"Kevin", 10},
		{"Tirth", 19},
	}

	for _,user := range users{
		greet(user)
	}
}