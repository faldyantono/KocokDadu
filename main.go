package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollAllOnes() (numberOfRolls int) {
	roundCounter := 0
	for {
		roundCounter++
		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1
		die3 := rand.Intn(6) + 1
		die4 := rand.Intn(6) + 1
		if die1 == 1 && die2 == 1 && die3 == 1 && die4 == 1 {
			numberOfRolls = roundCounter
			return numberOfRolls
		}
	}
}

func main() {
	fmt.Println("Masukan jumlah pemain :")
	var pemain int
	fmt.Scanln(&pemain)
	if pemain < 2 {
		fmt.Println("Jumlah pemain harus lebih banyak dari 1!")
	} else {
		fmt.Println("Masukan jumlah giliran :")
		var rounds int
		fmt.Scanln(&rounds)
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Printf("%s", "===============")
		for round := 1; round <= rounds; round++ {
			fmt.Printf("\n%s %d %s", "Giliran", round, "lempar dadu:")
			for i := 1; i <= 1; i++ {
				for j, poin := 1, 0; j <= pemain; j++ {
					numberOfRolls := rollAllOnes()
					fmt.Printf("\n%s%d (%d): %d,", "	Pemain #", j, poin, numberOfRolls)
					for l := 1; l < 5; l++ {
					}
				}
			}
			fmt.Printf("\n%s", "Setelah evaluasi:")
			fmt.Printf("\n%s", "==============")
		}
	}
}
