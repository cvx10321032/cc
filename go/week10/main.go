package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%f\n", math.Sqrt(19.0))
	fmt.Print("Input number : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}

	//counts := 0
	var isPrime bool = true // 가독성 up, 메모리 down
	if n <= 1 {
		isPrime = false
	} else if n == 2 {
		isPrime = true
	} else if n%2 == 0 { //2를 제외한 짝수는 모두 소수가 아님
		isPrime = false
	}

	j := 3
	for j <= int(math.Sqrt(float64(n))) {
		if n%j == 0 {
			//counts = counts + 1
			isPrime = false // 더하기 연산 제거
			break
		}
		fmt.Printf("%d ", j)
		j += 2
	}

	//if counts == 0 {
	if isPrime { // == 비교 연산 제거
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is NOT prime number.", n)
	}
}
