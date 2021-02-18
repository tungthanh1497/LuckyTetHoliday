package main

import "fmt"

func main() {
	var numPlayers int
	var listPlayers []Player

	numPlayers, listPlayers = initPlayers()

	fmt.Println(numPlayers)
	fmt.Println(listPlayers)
}

func initPlayers() (int, []Player) {
	numPlayers := inputNumPlayers()
	var listPlayers []Player
	for i := 0; i < numPlayers; i++ {
		name := inputName(i)
		listPlayers = append(listPlayers, Player{name: name})
	}
	return numPlayers, listPlayers
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
		if err != nil {
			fmt.Printf("%s. Try again: ", err)
		} else {
			return numPlayer
		}
	}

}

type Player struct {
	name     string
	listCard []string
	sum      int
}
