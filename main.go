package main

import (
	"fmt"
	"health-assistant/bodyfat"
)

func main() {
	var (
		age    int     = 30
		height float64 = 1.7
		weight float64 = 70
		gender string  = "男"
	)
	fmt.Println("Hello health-assistant")
	fmt.Println("bmi:", bodyfat.CalBMI(height, weight))
	fmt.Println("bodyfat:", bodyfat.CalBodyFat(age, gender, bodyfat.CalBMI(height, weight)))
}
