package main

import "fmt"

type Card int

type HandCards []Card

func (h HandCards) Len() int           { return len(h) }
func (h HandCards) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h HandCards) Less(i, j int) bool { return h[i] > h[j] }

func (h HandCards) HighCard() Card {
	high := TwoCard

	for _, h := range h {
		if h > TwoCard {
			high = h
		}
	}

	return high
}

const (
	TwoCard Card = iota
	ThreeCard
	FourCard
	FiveCard
	SixCard
	SevenCard
	EightCard
	NineCard
	TenCard
	JackCard
	QueenCard
	KingCard
	AceCard
)

func parseCard(cardStr string) Card {
	switch cardStr {
	case "2":
		return TwoCard
	case "3":
		return ThreeCard
	case "4":
		return FourCard
	case "5":
		return FiveCard
	case "6":
		return SixCard
	case "7":
		return SevenCard
	case "8":
		return EightCard
	case "9":
		return NineCard
	case "T":
		return TenCard
	case "J":
		return JackCard
	case "Q":
		return QueenCard
	case "K":
		return KingCard
	case "A":
		return AceCard
	}

	panic(fmt.Sprintf("Bad card: %s", cardStr))
}
