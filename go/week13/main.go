package main

import (
	"fmt"
)

func main() {
	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.23}
	gpa_slice := gpa[1:4] // 4.1, 4.5, 3.9
	//gpa_slice[1] = 2.71
	gpa[1] = 2.71
	fmt.Println(gpa_slice, gpa)
}
