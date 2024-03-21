package main

import "fmt"

type person struct {
	name   string
	weight float64
	height float64
}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func main() {
	mark := person{"Mark", 78, 1.69}
	john := person{"John", 92, 1.95}

	markBMI := calculateBMI(mark.weight, mark.height)
	johnBMI := calculateBMI(john.weight, john.height)

	fmt.Printf("BMI Mark : %.2f\n", markBMI)
	fmt.Printf("BMI John : %.2f\n", johnBMI)

	markHigherBMI := markBMI > johnBMI
	fmt.Println("Is Mark's BMI higher than John's BMI?", markHigherBMI)

	fmt.Print("---------------------------------------------------------\n")

	mark2 := person{"Mark", 95, 1.88}
	john2 := person{"John", 85, 1.76}

	mark2BMI := calculateBMI(mark2.weight, mark2.height)
	john2BMI := calculateBMI(john2.weight, john2.height)

	fmt.Printf("BMI Mark : %.2f\n", mark2BMI)
	fmt.Printf("BMI John : %.2f\n", john2BMI)

	mark2HigherBMI := mark2BMI > john2BMI
	fmt.Println("Is Mark's BMI higher than John's BMI?", mark2HigherBMI)
}
