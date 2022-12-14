package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func getnoc() int {
	var noc int
	fmt.Print("How many cristals do you want to open?\n")
	fmt.Print("Type a number: ")
	fmt.Scan(&noc)
	fmt.Println("Your number is:", noc)
	return noc
}

func openCristals(x int) {
	fmt.Printf("I will open your %d cristal/s in:\n", x)

	countdown()

	time.Sleep(1 * time.Second)

	var champ string
	for i := 1; i <= x; i++ {
		time.Sleep(1 * time.Second)
		// Champions
		champ = champion(rand.Intn(200))

		fmt.Printf("Your %d. champion is %s\n", i, champ)

	}
	fmt.Print("This are your champion/s\n")
	fmt.Print("Do you want to continue?\n")
}

func countdown() {
	time.Sleep(1 * time.Second)
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
func champion(x int) string {
	// noch mehr Champions einfÃ¼gen
	if x < 9 {
		return "Groot ðª â¯ â¯ â¯"
	}
	if x < 10 {
		return "Groot ðª â¯ â¯ â¯ â¯"
	}
	if x < 19 {
		return "King Groot ðª â¯ â¯ â¯"
	}
	if x < 20 {
		return "King Groot ðª â¯ â¯ â¯ â¯"
	}
	if x < 29 {
		return "Sabretooth ð§¬ â¯ â¯ â¯"
	}
	if x < 30 {
		return "Sabretooth ð§¬ â¯ â¯ â¯ â¯"
	}
	if x < 39 {
		return "Domino ð§¬ â¯ â¯ â¯"
	}
	if x < 40 {
		return " ð Domino ð§¬ â¯ â¯ â¯ â¯"
	}
	if x < 49 {
		return "Sunspot ð§¬ â¯ â¯ â¯"
	}
	if x < 50 {
		return " ð¥ Sunspot ð§¬ â¯ â¯ â¯ â¯"
	}
	if x < 59 {
		return "Guillotine ð® â¯ â¯ â¯"
	}
	if x < 60 {
		return "Guillotine ð® â¯ â¯ â¯ â¯"
	}
	if x < 69 {
		return "Howard the Duck âï¸  â¯ â¯ â¯"
	}
	if x < 70 {
		return "Howard the Duck âï¸  â¯ â¯ â¯ â¯"
	}
	if x < 79 {
		return "Ronan ðª â¯ â¯ â¯"
	}
	if x < 80 {
		return "Ronan ðª â¯ â¯ â¯ â¯"
	}
	if x < 89 {
		return "Iron Patriot âï¸  â¯ â¯ â¯"
	}
	if x < 90 {
		return "Iron Patriot âï¸  â¯ â¯ â¯ â¯"
	}
	if x < 99 {
		return "Omega Red ð§¬ â¯ â¯ â¯"
	}
	if x < 100 {
		return "Omega Red ð§¬ â¯ â¯ â¯ â¯"
	}
	if x < 109 {
		return "Black Panther â â¯ â¯ â¯"
	}
	if x < 110 {
		return "Black Panther â â¯ â¯ â¯ â¯"
	}
	if x < 119 {
		return "Deadpool ð§¬ â¯ â¯ â¯"
	}
	if x < 120 {
		return "Deadpool ð§¬ â¯ â¯ â¯ â¯"
	}
	if x < 129 {
		return "Colossus ð§¬ â¯ â¯ â¯"
	}
	if x < 130 {
		return "Colossus ð§¬ â¯ â¯ â¯ â¯"
	}
	if x < 139 {
		return "Thor ðª â¯ â¯ â¯"
	}
	if x < 140 {
		return " ðª§  Thor ðª â¯ â¯ â¯ â¯"
	}
	if x < 149 {
		return "Scarlet Witch ð® â¯ â¯ â¯"
	}
	if x < 150 {
		return "Scarlet Witch ð® â¯ â¯ â¯ â¯"
	}
	if x < 159 {
		return "Wolverine ð§¬ â¯ â¯ â¯"
	}
	if x < 160 {
		return "Wolverine ð§¬ â¯ â¯ â¯ â¯"
	}
	if x < 169 {
		return "Spider-man ð¬ â¯ â¯ â¯"
	}
	if x < 170 {
		return " ð· Spider-man ð¬ â¯ â¯ â¯ â¯"
	}
	if x < 179 {
		return "Hulk ð¬ â¯ â¯ â¯"
	}
	if x < 180 {
		return "Hulk ð¬ â¯ â¯ â¯ â¯"
	}
	if x < 189 {
		return "Captain America ð¬ â¯ â¯ â¯"
	}
	if x < 190 {
		return "Captain America ð¬ â¯ â¯ â¯ â¯"
	}
	if x < 199 {
		return "Iron Man âï¸  â¯ â¯ â¯"
	}
	return "Iron âï¸  â¯ â¯ â¯ â¯"

}

func main() {
	var noc int

	fmt.Print("\n")
	fmt.Print("ð Hello and welcome to the cristal opening party\n")

	time.Sleep(3 * time.Second)

	rand.Seed(time.Now().Unix())

	var i string
	for {
		noc = getnoc()
		openCristals(noc)

		fmt.Print("Type Yes or No: ")
		fmt.Scan(&i)
		fmt.Println("Your anser is:", i)

		if i != "Yes" {
			fmt.Println(" ð Bye bye ð ")
			os.Exit(0)
		}

	}

}
