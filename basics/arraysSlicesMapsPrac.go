package main
import "fmt"

func main(){
	// var numbers= []int{1,2,3,4,5,6,7,8,9,0}
	// fmt.Println(numbers)

	// fruits := []string{"apple", "banana"}
	// fruits = append(fruits, "cherry", "khaman")
	// fmt.Println(fruits)

	student := map[string] int {
		"Kevin" : 105,
		"Mayank": 77,
		"Dhruv": 68,
		"Tirth": 103,
	}

	fmt.Println(student["Kevin"])

}