package main

import "fmt"

func Type_Pointer() {
	var car Car = Car{
		brandId:   125,
		brandName: "BMW",
		modelId:   1450,
		modelName: "720i",
	}
	fmt.Println("Car Info: ", car)

	var carPtr *Car = &car
	fmt.Println("Car address: ", carPtr)

	carPtr.brandId = 141
	fmt.Println("Car Info: ", car)

	(*carPtr).brandId = 100
	fmt.Println("Car Info: ", car)

	SetModelInfo(&car)

	fmt.Println("Car Info After SetModelInfo: ", car)

	car.SetBrandInfo(909, "Fiat")

	fmt.Println("Car Info After SetBrandInfo: ", car)
}
func SetModelInfo(car *Car) {
	car.modelId = 707
	car.modelName = "i8"
	fmt.Println("Car Info in SetModelInfo: ", car)
}
func (car *Car) SetBrandInfo(brandId uint16, brandName string) {
	car.brandId = brandId
	car.brandName = brandName
	fmt.Println("Car Info in SetBrandInfo: ", *car)
}

type Car struct {
	brandId   uint16
	brandName string
	modelId   uint32
	modelName string
}
