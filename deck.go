package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a tyoe od deck
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"spades", "Diamonds", "hearts", "clubs"}
	cardValues := []string{"Ace", "two", "three", "Four"}
	for _, suite := range cardSuits {
		for _, vales := range cardValues {
			cards = append(cards, vales+" of "+suite)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFIle(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 06666)
}

func newDeckFromFile(filename string) deck {
	bs, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) //time.Now().UnixNano() is used to generate int64 every new time
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
