package main

import "fmt"

func main() {

	students := [...][3]float64{
		{3, 4, 5},
		{6, 7, 8},
	}

	var sum float64
	for _, grades := range students {

		for _, grade := range grades {

			sum += grade
		}
	}

	const N = float64(len(students) * len(students[0]))

	fmt.Printf("Avg Grade : %g\n", sum/N)

}
