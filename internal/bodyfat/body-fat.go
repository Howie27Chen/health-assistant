package bodyfat

func CalBMI(height, weight float64) float64 {
	return weight / (height * height)
}

func CalBodyFat(age int, gender string, bmi float64) float64 {
	genderWeight := 0.0
	if gender == "男" {
		genderWeight = 1
	} else {
		genderWeight = 0
	}
	return 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*genderWeight
}
