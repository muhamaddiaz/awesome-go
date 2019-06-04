package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println()
}

func newDeck() deck {
	cards := deck{}
	greetCard := []string{"Hi, ", "Hola, ", "Wow, "}
	nameCard := []string{"Diaz", "Stephan", "Abigail"}

	for _, greet := range greetCard {
		for _, name := range nameCard {
			cards = append(cards, greet + name)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ". ")
}

func (d deck) writeDeckToFile(filename string) {
	ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	stringBs := string(bs)
	return strings.Split(stringBs, ".")
}

func (d deck) shuffle() {
	for i := range d {
		nextState := rand.Intn(len(d) - 1)
		d[nextState], d[i] = d[i], d[nextState]
	}
}