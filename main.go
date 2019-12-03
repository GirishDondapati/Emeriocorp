package main

import (
	"fmt"
)

func main() {
	var initialEnergy int
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
	}

	bestProfit := 0
	if noOfCoins <= initialEnergy {
		for a := 0; a < noOfEnergies; a++ {
			bestProfit = bestProfit + coinsValues[a]
		}
		fmt.Println(bestProfit)
		return
	}

	bestProfit = getBestProfit(initialEnergy, coinsValues, energyValues, noOfCoins)
	fmt.Println(bestProfit)
}

func getBestProfit(initialEnergy int, coinsValues []int, energyValues []int, noOfHouses int) int {
	bestProfit := 0
	for i := 0; i < noOfHouses; i++ {
		curHouseE := energyValues[i]
		curHouseC := coinsValues[i]

		if initialEnergy <= 0 
			initialEnergy = curHouseE

	}
	return 1
}
