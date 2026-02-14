package deck

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of CLubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestLoadDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck, err := loadDeckFromFile("_decktesting")
	if err != nil {
		t.Errorf("Expected no error loading deck, but got %v", err)
	}

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in loaded deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDealFileInvalidHand(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	_, _, err := DealFile("_decktesting", 0)
	if err == nil {
		t.Errorf("Expected error for hand size 0, got nil")
	}

	_, _, err = DealFile("_decktesting", 17)
	if err == nil {
		t.Errorf("Expected error for hand size larger than deck, got nil")
	}

	os.Remove("_decktesting")
}
