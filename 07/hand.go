package main

import (
	"strconv"
	"strings"
)

type HandType int

const (
	HighCard  HandType = iota // 0
	OnePair                   // 1
	TwoPair                   // 2
	ThreeKind                 // 3
	FullHouse                 // 4
	FourKind                  // 5
	FiveKind                  // 6
)

type Hand struct {
	Cards []Card
	Type  HandType
	Bid   int
}

type SortByHand []*Hand

func (h SortByHand) Len() int      { return len(h) }
func (h SortByHand) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h SortByHand) Less(i, j int) bool {
	if h[i].Type > h[j].Type {
		return true
	}

	if h[i].Type < h[j].Type {
		return false
	}

	for cardIdx, iCard := range h[i].Cards {
		jCard := h[j].Cards[cardIdx]

		if iCard == jCard {
			continue
		}

		if iCard > jCard {
			return true
		}

		if iCard < jCard {
			return false
		}
	}

	return true
}

func NewHand(cardStr string, bidStr string) *Hand {
	cards := parseCards(cardStr)
	bid, err := parseBid(bidStr)

	if err != nil {
		panic(err)
	}

	hand := &Hand{
		Cards: cards,
		Bid:   bid,
	}

	handType := hand.FigureHand()
	hand.Type = handType

	return hand
}

func (h *Hand) FigureHand() HandType {
	cardCounts := make([]int, 13)

	for _, c := range h.Cards {
		cardCounts[c] += 1
	}

	hasPair := false
	hasThreeKind := false

	for _, count := range cardCounts {
		switch count {
		case 5:
			return FiveKind
		case 4:
			return FourKind
		case 3:
			hasThreeKind = true
		case 2:
			if hasPair {
				return TwoPair
			}

			hasPair = true
		}
	}

	if hasThreeKind && !hasPair {
		return ThreeKind
	}

	if hasThreeKind && hasPair {
		return FullHouse
	}

	if hasPair {
		return OnePair
	}

	return HighCard
}

func parseCards(cardsStr string) HandCards {
	split := strings.Split(strings.TrimSpace(cardsStr), "")

	cards := make([]Card, len(split))

	for i, c := range split {
		cards[i] = parseCard(c)
	}

	return cards
}

func parseBid(bidStr string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(bidStr))
}
