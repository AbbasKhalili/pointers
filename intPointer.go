package main

import "fmt"

func Int_Pointer() {
	var age uint8 = 28
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
