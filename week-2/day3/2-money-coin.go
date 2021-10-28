package main

import "fmt"

func moneyCoins(money int) []int {
	availableCoins := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	changes := []int{}

	for i := len(availableCoins) - 1; i >= 0; i-- {
		coin := availableCoins[i]
		if coin > money {
			continue
		}

		numberOfCoin := money / coin
		for i := 0; i < numberOfCoin; i++ {
			changes = append(changes, coin)
		}

		money -= numberOfCoin * coin
	}

	return changes
}

func main() {
	fmt.Println(moneyCoins(123))   // [100 20 1 1 1]
	fmt.Println(moneyCoins(432))   // [200 200 20 10 1 1]
	fmt.Println(moneyCoins(543))   // [500, 20, 20, 1, 1, 1]
	fmt.Println(moneyCoins(7752))  // [5000, 2000, 500, 200, 50, 1, 1]
	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]
}
