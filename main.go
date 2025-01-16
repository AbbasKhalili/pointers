package main

import "fmt"

type Car struct {
	brandId uint16
	modelId uint32
	trimId  *uint32
}

func main() {
	var age uint8 = 40
	var agePtr *uint8 = &age

	fmt.Println("Age is : ", age)
	fmt.Println("Age ptr is : ", agePtr)

	PrintAge(&age)

	fmt.Println("Age is : ", age)
	fmt.Println("Age ptr is : ", agePtr)
}

func PrintAge(age *uint8) {
	fmt.Println("Age is : ", *age)
	fmt.Println("Age address is : ", age)
	fmt.Println("Age ptr is : ", &age)

	*age = 45

	fmt.Println("Age is : ", *age)
	fmt.Println("Age address is : ", age)
	fmt.Println("Age ptr is : ", &age)
}
