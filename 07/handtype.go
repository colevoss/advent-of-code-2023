package main

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

func (ht HandType) Upgrade(jokerCount int) HandType {
	if jokerCount == 0 {
		return ht
	}

	switch ht {
	case HighCard:
		// cannot have more than one joker because each card is unique
		return OnePair

	case OnePair:
		// can either be one joker or two
		// pair can be non jokers so then there can be one joker
		// either way its three of a kind
		return ThreeKind

	case TwoPair:
		// Can have 1 joker (the card that isn't in either pair)
		// or two jokers, being one of the pair.

		// if 1 joker, then full house
		if jokerCount == 1 {
			return FullHouse
		}

		// if 2 jokers, then 4 of a kind
		if jokerCount == 2 {
			return FourKind
		}

	case ThreeKind:
		if jokerCount == 1 {
			return FourKind
		}

		if jokerCount == 2 {
			return FiveKind
		}

	case FullHouse:
		// can have 2 or three jokers
		// either way should be five kind
		return FiveKind

	case FourKind:
		return FiveKind
	}

	return ht
}
