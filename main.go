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
	// noch mehr Champions einfügen
	if x < 9 {
		return "Groot 🪐 ✯ ✯ ✯"
	}
	if x < 10 {
		return "Groot 🪐 ✯ ✯ ✯ ✯"
	}
	if x < 19 {
		return "King Groot 🪐 ✯ ✯ ✯"
	}
	if x < 20 {
		return "King Groot 🪐 ✯ ✯ ✯ ✯"
	}
	if x < 29 {
		return "Sabretooth 🧬 ✯ ✯ ✯"
	}
	if x < 30 {
		return "Sabretooth 🧬 ✯ ✯ ✯ ✯"
	}
	if x < 39 {
		return "Domino 🧬 ✯ ✯ ✯"
	}
	if x < 40 {
		return " 🍀 Domino 🧬 ✯ ✯ ✯ ✯"
	}
	if x < 49 {
		return "Sunspot 🧬 ✯ ✯ ✯"
	}
	if x < 50 {
		return " 🔥 Sunspot 🧬 ✯ ✯ ✯ ✯"
	}
	if x < 59 {
		return "Guillotine 🔮 ✯ ✯ ✯"
	}
	if x < 60 {
		return "Guillotine 🔮 ✯ ✯ ✯ ✯"
	}
	if x < 69 {
		return "Howard the Duck ⚙️  ✯ ✯ ✯"
	}
	if x < 70 {
		return "Howard the Duck ⚙️  ✯ ✯ ✯ ✯"
	}
	if x < 79 {
		return "Ronan 🪐 ✯ ✯ ✯"
	}
	if x < 80 {
		return "Ronan 🪐 ✯ ✯ ✯ ✯"
	}
	if x < 89 {
		return "Iron Patriot ⚙️  ✯ ✯ ✯"
	}
	if x < 90 {
		return "Iron Patriot ⚙️  ✯ ✯ ✯ ✯"
	}
	if x < 99 {
		return "Omega Red 🧬 ✯ ✯ ✯"
	}
	if x < 100 {
		return "Omega Red 🧬 ✯ ✯ ✯ ✯"
	}
	if x < 109 {
		return "Black Panther ✊ ✯ ✯ ✯"
	}
	if x < 110 {
		return "Black Panther ✊ ✯ ✯ ✯ ✯"
	}
	if x < 119 {
		return "Deadpool 🧬 ✯ ✯ ✯"
	}
	if x < 120 {
		return "Deadpool 🧬 ✯ ✯ ✯ ✯"
	}
	if x < 129 {
		return "Colossus 🧬 ✯ ✯ ✯"
	}
	if x < 130 {
		return "Colossus 🧬 ✯ ✯ ✯ ✯"
	}
	if x < 139 {
		return "Thor 🪐 ✯ ✯ ✯"
	}
	if x < 140 {
		return " 🪧  Thor 🪐 ✯ ✯ ✯ ✯"
	}
	if x < 149 {
		return "Scarlet Witch 🔮 ✯ ✯ ✯"
	}
	if x < 150 {
		return "Scarlet Witch 🔮 ✯ ✯ ✯ ✯"
	}
	if x < 159 {
		return "Wolverine 🧬 ✯ ✯ ✯"
	}
	if x < 160 {
		return "Wolverine 🧬 ✯ ✯ ✯ ✯"
	}
	if x < 169 {
		return "Spider-man 🔬 ✯ ✯ ✯"
	}
	if x < 170 {
		return " 🕷 Spider-man 🔬 ✯ ✯ ✯ ✯"
	}
	if x < 179 {
		return "Hulk 🔬 ✯ ✯ ✯"
	}
	if x < 180 {
		return "Hulk 🔬 ✯ ✯ ✯ ✯"
	}
	if x < 189 {
		return "Captain America 🔬 ✯ ✯ ✯"
	}
	if x < 190 {
		return "Captain America 🔬 ✯ ✯ ✯ ✯"
	}
	if x < 199 {
		return "Iron Man ⚙️  ✯ ✯ ✯"
	}
	return "Iron ⚙️  ✯ ✯ ✯ ✯"

}

func main() {
	var noc int

	fmt.Print("\n")
	fmt.Print("💎 Hello and welcome to the cristal opening party\n")

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
			fmt.Println(" 👋 Bye bye 👋 ")
			os.Exit(0)
		}

	}

}
