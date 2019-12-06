package main

import (
	"fmt"
)

func main() {
	/*var initialEnergy int
	var noOfEnergies int
	var noOfCoins int

	fmt.Scanln(&initialEnergy)

	fmt.Scanln(&noOfEnergies)
	energyValues := make([]int, noOfEnergies)
	var temp int
	for a := 0; a < noOfEnergies; a++ {
		fmt.Scanln(&temp)
		energyValues[a] = temp
	}

	fmt.Scanln(&noOfCoins)
	if noOfCoins != noOfEnergies {
		panic("noOfCoins and noOfEnergies must be same")
	}

	coinsValues := make([]int, noOfCoins)
	for a := 0; a < noOfEnergies; a++ {
		fmt.Scanln(&temp)
		coinsValues[a] = temp
	}*/

	initialEnergy := 1
	//noOfEnergies := 5
	noOfCoins := 5

	coinsValues := []int{1, 5, 3, 3, 1}
	energyValues := []int{3, 23, 9, 2, 2}

	bestProfit := 0
	if noOfCoins <= initialEnergy {
		bestProfit = getEnoughEnergy(initialEnergy, coinsValues, energyValues, noOfCoins, 0)
		fmt.Println(bestProfit)
		return
	}

	bestProfit = getBestProfit(initialEnergy, coinsValues, energyValues, noOfCoins)
	fmt.Println(bestProfit)
}

func getBestProfit(initialEnergy int, coinsValues []int, energyValues []int, noOfHouses int) int {
	bestProfit := 0
	i := 0
	for i < noOfHouses {
		if initialEnergy <= 0 {
			initialEnergy = energyValues[i]
			i++
		}

		for initialEnergy > 0 && initialEnergy < noOfHouses-i {

		}
	}
	return bestProfit
}

func getEnoughEnergy(initialEnergy int, coinsValues []int, energyValues []int, noOfHouses int, readIndex int) int {
	bestProfit := 0
	for a := readIndex; a < noOfHouses; a++ {
		bestProfit = bestProfit + coinsValues[a]
	}
	return bestProfit
}
