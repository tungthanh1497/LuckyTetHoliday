package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func main() {

	listPlayers := initPlayers()

	fmt.Println("Start game")
	startGame(&listPlayers)
	fmt.Println(listPlayers)

}

func startGame(listPlayers *[]Player) {
	rand.Seed(time.Now().UnixNano())
	existedCards := make(map[string]bool)
	for position := 0; position < len(*listPlayers); position++ {
		for i := 0; i < 3; i++ {
			card := getRandomCard(&existedCards)
			existedCards[card] = true
			(*listPlayers)[position].listCard[i] = card
		}
		calSum(&((*listPlayers)[position]))
	}
	sort.SliceStable(*listPlayers, func(x, y int) bool {
		xPlayer := (*listPlayers)[x]
		yPlayer := (*listPlayers)[y]
		if xPlayer.typeWon != yPlayer.typeWon {
			return xPlayer.typeWon > yPlayer.typeWon
		} else if 3 == xPlayer.typeWon {
			xCheckValue := getCheckValue(xPlayer.listCard[0])
			yCheckValue := getCheckValue(yPlayer.listCard[0])
			return xCheckValue > yCheckValue
		} else if xPlayer.sum != yPlayer.sum {
			return xPlayer.sum > yPlayer.sum
		} else {
			xMaxSuit := getMaxSuit(xPlayer.listCard)
			yMaxSuit := getMaxSuit(yPlayer.listCard)
			if xMaxSuit != yMaxSuit {
				return xMaxSuit > yMaxSuit
			} else {
				xMaxValue := getMaxValue(xPlayer.listCard, xMaxSuit)
				yMaxValue := getMaxValue(yPlayer.listCard, yMaxSuit)
				return xMaxValue > yMaxValue
			}
		}
	})
}

func getMaxValue(listCard [3]string, maxSuit int) int {
	maxValue := 0
	for i := 0; i < len(listCard); i++ {
		iCard, _ := strconv.Atoi(listCard[i])
		value := iCard / 10
		suit := iCard % 10
		if 1 == value {
			value = 10
		}
		if suit == maxSuit && value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func getMaxSuit(listCard [3]string) int {
	maxSuit := 0
	for i := 0; i < len(listCard); i++ {
		iCard, _ := strconv.Atoi(listCard[i])
		suit := iCard % 10
		if suit > maxSuit {
			maxSuit = suit
		}
	}
	return maxSuit
}

func getCheckValue(card string) int {
	if 1 == card[0] {
		return 10
	} else {
		iCard, _ := strconv.Atoi(card)
		return iCard / 10
	}
}

func calSum(player *Player) {
	sum := 0
	isTriple := true
	checkingValue := 0
	for i := 0; i < len(player.listCard); i++ {
		value, _ := strconv.Atoi((*player).listCard[i][0:1])
		sum += value
		if 0 == checkingValue {
			checkingValue = value
		} else if checkingValue != value {
			isTriple = false
		}
	}
	switch sum % 10 {
	case 0:
		(*player).sum = 10
		(*player).typeWon = 2
	default:
		(*player).sum = sum % 10
		(*player).typeWon = 1
	}
	if isTriple {
		(*player).typeWon = 3
	}
}

func getRandomCard(existedCards *map[string]bool) string {
	for {
		value := rand.Intn(9) + 1
		suit := rand.Intn(4)
		card := fmt.Sprintf("%d%d", value, suit)
		if _, found := (*existedCards)[card]; found {
			continue
		} else {
			return card
		}
	}
}

func initPlayers() []Player {
	numPlayers := inputNumPlayers()
	var listPlayers []Player
	for i := 0; i < numPlayers; i++ {
		name := inputName(i)
		listPlayers = append(listPlayers, Player{name: name})
	}
	return listPlayers
}

func inputName(position int) string {
	var name string
	fmt.Printf("Enter name of player %d: ", position+1)
	for {
		_, err := fmt.Scanf("%s", &name)
		if err != nil {
			fmt.Printf("%s. Try again: ", err)
		} else {
			return name
		}
	}
}

func inputNumPlayers() int {
	var numPlayer int
	fmt.Print("Enter number of players: ")
	for {
		_, err := fmt.Scanf("%d", &numPlayer)
		if err != nil || numPlayer < 0 || numPlayer > 12 {
			fmt.Printf("%s. Try again: ", err)
		} else {
			return numPlayer
		}
	}

}

type Player struct {
	name     string
	listCard [3]string // string <value-suit>: value 1-9 ; suit 0-3 (spade, club, heart, diamond)
	sum      int
	typeWon  int // 1: normal, 2: sum equal 10, 3: triple
}
