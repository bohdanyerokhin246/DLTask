package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Laptop struct {
	Gain   int
	Price  int
	Volume int
}

func multiDimensionalKnapsack(amountOfCapital, amountOfLaptops int, laptops []Laptop) int {
	dp := make([][][]int, len(laptops)+1)
	for i := range dp {
		dp[i] = make([][]int, amountOfCapital+1)
		for w := range dp[i] {
			dp[i][w] = make([]int, amountOfLaptops+1)
		}
	}

	for i := 1; i <= len(laptops); i++ {
		for w := 0; w <= amountOfCapital; w++ {
			for v := 0; v <= amountOfLaptops; v++ {
				if laptops[i-1].Price <= w && laptops[i-1].Volume <= v {
					dp[i][w][v] = Max(
						dp[i-1][w][v],
						dp[i-1][w-laptops[i-1].Price][v-laptops[i-1].Volume]+laptops[i-1].Gain,
					)
				} else {
					dp[i][w][v] = dp[i-1][w][v]
				}
			}
		}
	}

	return amountOfCapital + dp[len(laptops)][amountOfCapital][amountOfLaptops]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func inputData() (int, int, []Laptop) {
	var amountOfLaptops, amountOfCapital int
	laptop := Laptop{}
	laptops := make([]Laptop, 0)

	fmt.Println("Input N")
	_, err := fmt.Scanln(&amountOfLaptops)
	if err != nil {
		fmt.Printf("Error with scaning N value. Error: %v", err)
	}

	fmt.Println("Input C")
	_, err = fmt.Scanln(&amountOfCapital)
	if err != nil {
		fmt.Printf("Error with scaning N value. Error: %v", err)
	}

	fmt.Println("Input price and gain for", len(laptops), "laptop. Example:150 350")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println("If you want to stop input press Ctrl+Z")
		fmt.Println("Input price and gain for", len(laptops)+1, "laptop. Example:150 350")

		str := strings.Split(scanner.Text(), " ")
		if len(str) == 2 {
			laptop.Price, _ = strconv.Atoi(str[0])
			laptop.Gain, _ = strconv.Atoi(str[1])
			laptop.Volume = 1

			laptops = append(laptops, laptop)
		} else {
			fmt.Println("\nInput price and gain correctly. Example:150 350")
			continue
		}
	}
	return amountOfCapital, amountOfLaptops, laptops
}

func main() {
	amountOfCapital, amountOfLaptops, laptops := inputData()

	maxValue := multiDimensionalKnapsack(amountOfCapital, amountOfLaptops, laptops)
	fmt.Println("\nCapital at the end of the summer:", maxValue)
}
