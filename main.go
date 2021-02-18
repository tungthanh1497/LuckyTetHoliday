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
			// randomize 3 cards
			card := getRandomCard(&existedCards)
			existedCards[card] = true
			(*listPlayers)[position].listCard[i] = card
		}
		// check sum and type won for each person
		calSum(&((*listPlayers)[position]))
	}
	sort.SliceStable(*listPlayers, func(x, y int) bool {
		xPlayer := (*listPlayers)[x]
		yPlayer := (*listPlayers)[y]
		if xPlayer.typeWon != yPlayer.typeWon {
			// sort type won: triple > sum 10 > normal
			return xPlayer.typeWon > yPlayer.typeWon
		} else if 3 == xPlayer.typeWon {
			// case triple: sort by value
			xCheckValue := getCheckValue(xPlayer.listCard[0])
			yCheckValue := getCheckValue(yPlayer.listCard[0])
			return xCheckValue > yCheckValue
		} else if xPlayer.sum != yPlayer.sum {
			// case sum 10 or normal: sort by sum
			return xPlayer.sum > yPlayer.sum
		} else {
			// case sum equal
			xMaxSuit := getMaxSuit(xPlayer.listCard)
			yMaxSuit := getMaxSuit(yPlayer.listCard)
			if xMaxSuit != yMaxSuit {
				// sort by largest suit
				return xMaxSuit > yMaxSuit
			} else {
				// sort by largest value of largest suit
				xMaxValue := getMaxValue(xPlayer.listCard, xMaxSuit)
				yMaxValue := getMaxValue(yPlayer.listCard, yMaxSuit)
				return xMaxValue > yMaxValue
			}
		}
	})
}

/**
get max value in list card (by max suit)
*/
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

/**
get max suit in list card
*/
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

/**
case triple, get check value (any value in list cards)
*/
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
			// first card
			checkingValue = value
		} else if checkingValue != value {
			// if 2 cards have different value -> can not be triple
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

/**
randomize a card for user contain
* value: from 1-9
* suit: from 0 to 3
   + 0: spade
   + 1: club
   + 2: heart
   + 3: diamond
*/
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

/**
init players dex
*/
func initPlayers() []Player {
	numPlayers := inputNumPlayers()
	var listPlayers []Player
	for i := 0; i < numPlayers; i++ {
		name := inputName(i)
		listPlayers = append(listPlayers, Player{name: name})
	}
	return listPlayers
}

/**
get name of user by position
name can not be empty
*/
func inputName(position int) string {
	var name string
	fmt.Printf("Enter name of player %d: ", position+1)
	for {
		_, err := fmt.Scanf("%s", &name)
		if err != nil || len(name) == 0 {
			fmt.Printf("%s. Try again: ", err)
		} else {
			return name
		}
	}
}

/**
get number of players from input
must be a number larger than 0 and smaller than 12
*/
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
