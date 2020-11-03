package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ballAndStrike struct {
	strike int
	ball   int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// create three random number
	numbers := MakeNumbers()
	cnt := 0

	for {
		cnt++
		// input numbers
		inputNumbers := InputNumbers()

		//결과를 비교
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(CompareNumbers(numbers, inputNumbers))
		//3s?
		if IsGameEnd(result) {
			//게임끝
			break

		}
	}

	//게임끝. 몇번만에 맞췄는지 출력
	fmt.Printf("%d번만에 맞추었습니다", cnt)
}

func MakeNumbers() [3]int {
	// 0~9 사이의 겹치지 않는 무작위 숫자 3개를
	var rst [3]int

	rst[0] = rand.Intn(9) + 1
	rst[1] = rand.Intn(10)
	rst[2] = rand.Intn(10)

	for {

		if rst[1] == rst[0] || rst[1] == rst[2] {
			rst[1] = rand.Intn(10)

		} else {
			break
		}

	}

	for {

		if rst[2] == rst[1] || rst[2] == rst[0] {
			rst[2] = rand.Intn(10)

		} else {
			break
		}

	}
	fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int {
	// 키보드로부터 0~9 사이의 겹치지 않는 숫자 3개를 입력후 반환
	var rst [3]int
	for i := 0; i < 3; i++ {
		fmt.Println("Enter a numbe
		
		
		
		r")
		fmt.Scanf("%d", &rst[i])

	}

	return rst
}

func CompareNumbers(numbers, InputNumbers [3]int) ballAndStrike {

	strike := 0
	ball := 0

	for i := 0; i < 3; i++ {
		if numbers[i] == InputNumbers[i] {
			strike++
		}

		for j := 0; j < 3; j++ {
			if i != j {
				if numbers[i] == numbers[j] {
					ball++
				}

			}

		}
	}

	return ballAndStrike{strike, ball}
}

func PrintResult(result ballAndStrike) {
	fmt.Printf("ball is %d\n", result.ball)
	fmt.Printf("strike is %d\n", result.strike)

}

func IsGameEnd(result ballAndStrike) bool {

	if result.strike == 3 {
		return true
	}

	return false
}
