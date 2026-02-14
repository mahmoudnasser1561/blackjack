package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
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

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func loadDeckFromFile(filename string) (deck, error) {
	bs, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	s := strings.Split(string(bs), ",")
	return deck(s), nil
}

func newDeckFromFile(filename string) deck {
	d, err := loadDeckFromFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return d
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func NewToFile(filename string) error {
	cards := newDeck()
	return cards.saveToFile(filename)
}

func ShuffleFile(inFilename, outFilename string) error {
	cards, err := loadDeckFromFile(inFilename)
	if err != nil {
		return err
	}

	cards.shuffle()
	return cards.saveToFile(outFilename)
}

func DealFile(filename string, handSize int) ([]string, []string, error) {
	cards, err := loadDeckFromFile(filename)
	if err != nil {
		return nil, nil, err
	}

	if handSize <= 0 {
		return nil, nil, fmt.Errorf("hand size must be greater than 0")
	}

	if handSize > len(cards) {
		return nil, nil, fmt.Errorf("hand size %d exceeds deck size %d", handSize, len(cards))
	}

	hand, remaining := deal(cards, handSize)
	return []string(hand), []string(remaining), nil
}

func Run() {
	cards := newDeck()

	cards.saveToFile("my_cards.txt")
	cards.shuffle()
	cards.print()
}
