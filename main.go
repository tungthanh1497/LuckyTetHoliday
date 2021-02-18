package main

import (
	"fmt"
	"math/rand"
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
	listCard [3]string
	sum      int
	typeWon  int // 1: normal, 2: sum equal 10, 3: triple
}
