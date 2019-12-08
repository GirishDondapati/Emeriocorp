package main

import "fmt"

type Node struct {
	left          *Node
	right         *Node
	nodeStatus    bool
	totalCoins    int
	initialEnergy int
	indexLevel    int
}

var noOfHouses int
var coinsValues []int

var result int = 0

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

	coinsList := make([]int, noOfCoins)
	for a := 0; a < noOfEnergies; a++ {
		fmt.Scanln(&temp)
		coinsList[a] = temp
	}
	coinsValues = coinsList[0:]

	noOfHouses = len(energyValues)

	if noOfHouses <= initialEnergy {
		result = getEnoughEnergy(coinsValues, noOfHouses, 0)
		fmt.Println(result)
		return
	}

	if initialEnergy == 0 && energyValues[0] == 0 {
		result = coinsValues[0]
		fmt.Println(result)
		return
	}

	rootNode := Node{left: nil, right: nil, nodeStatus: false, totalCoins: 0, initialEnergy: initialEnergy}
	firstHouseCoins := coinsValues[0]
	firstHouseEnergy := energyValues[0]
	rootL := Node{left: nil, right: nil, nodeStatus: true, totalCoins: firstHouseCoins, initialEnergy: initialEnergy, indexLevel: 0}
	rootR := Node{left: nil, right: nil, nodeStatus: true, totalCoins: 0, initialEnergy: initialEnergy + firstHouseEnergy, indexLevel: 0}
	rootNode = Node{left: &rootL, right: &rootR, nodeStatus: false, totalCoins: 0, initialEnergy: initialEnergy, indexLevel: 0}

	var finalNodes Node
	for i := 1; i < noOfHouses; i++ {
		currHouseCoins := coinsValues[i]
		currHouseEnergy := energyValues[i]

		finalNodes = *makeTree(i, currHouseCoins, currHouseEnergy, &rootNode)
	}

	finalNodes.PrintPre()
	fmt.Println(result)
}

func makeTree(indexVal int, currHouseCoins int, currHouseEnergy int, rootNode *Node) (root *Node) {
	if rootNode == nil {
		//	fmt.Println("nil return%%%%%%%%%%")
		return rootNode
	} else if rootNode.left != nil || rootNode.right != nil {
		//	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
		makeTree(indexVal, currHouseCoins, currHouseEnergy, rootNode.left)
		makeTree(indexVal, currHouseCoins, currHouseEnergy, rootNode.right)
	} else if rootNode.left == nil && rootNode.right == nil && rootNode.nodeStatus == true {
		//fmt.Println("--------------------------------------------------")
		//fmt.Println("--", rootNode.totalCoins, rootNode.initialEnergy, "--")
		energy := rootNode.initialEnergy - 1
		coins := rootNode.totalCoins
		rootNode.nodeStatus = false

		rootL := Node{left: nil, right: nil, nodeStatus: true, totalCoins: (coins + currHouseCoins), initialEnergy: energy, indexLevel: indexVal}
		if energy == 0 {
			rootL = Node{left: nil, right: nil, nodeStatus: false, totalCoins: (coins + currHouseCoins), initialEnergy: energy, indexLevel: indexVal}
		}
		//fmt.Println("--rootL", rootL.totalCoins, rootL.initialEnergy, "--")

		rootR := Node{left: nil, right: nil, nodeStatus: true, totalCoins: coins, initialEnergy: energy + currHouseEnergy, indexLevel: indexVal}
		if energy+currHouseEnergy >= noOfHouses-indexVal {
			totalCoins := getEnoughEnergy(coinsValues, noOfHouses, indexVal+1)
			rootR = Node{left: nil, right: nil, nodeStatus: false, totalCoins: (coins + totalCoins), initialEnergy: 0, indexLevel: indexVal}
		}
		//fmt.Println("--rootR", rootR.totalCoins, rootR.initialEnergy, "--")

		rootNode.left = &rootL
		rootNode.right = &rootR
		rootNode.nodeStatus = false
		rootNode.totalCoins = coins
	}
	return rootNode
}

func getEnoughEnergy(coinsValues []int, noOfHouses int, readIndex int) int {
	bestProfit := 0
	for a := readIndex; a < noOfHouses; a++ {
		bestProfit = bestProfit + coinsValues[a]
	}
	return bestProfit
}

func (root *Node) PrintPre() {
	//fmt.Println("[", root.totalCoins, root.initialEnergy, "]")
	if root.totalCoins > result {
		result = root.totalCoins
	}
	if root.left != nil {
		root.left.PrintPre()
	}
	if root.right != nil {
		root.right.PrintPre()
	}
}
