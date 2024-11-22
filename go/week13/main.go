package main

import (
	"fmt"
)

func main() {
	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.23}
	gpa_slice := gpa[1:4] // 4.1, 4.5, 3.9
	//gpa_slice[1] = 2.71
	gpa[2] = 2.71
	//gpa_slice = append(gpa_slice, 4.3) gpa[]의 길이 만큼만 넣으면 4.23이 4.3으로 바뀜
	gpa_slice = append(gpa_slice, 4.3, 5.55) // 초과해서 넣으면 4.23을 바꾸지 않고 gpa_slice에만 추가
	fmt.Println(len(gpa_slice), gpa_slice, gpa)
}
