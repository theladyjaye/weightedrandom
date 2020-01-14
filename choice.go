package weightedrandom

import (
	"math/rand"
	"time"
)

type Choice struct {
	numChoices float64
	aliases    []alias
	rand       *rand.Rand
}

type alias struct {
	odds  float64
	alias int
}

type record struct {
	index int
	value float64
}

// NewChoiceFromWeights initializes a new Choice from provided weights.
// Note: this will return the index of the weight
// Ported from Python
// https://blog.bruce-hill.com/a-faster-weighted-random-choice
//
// Ported directly from here:
// https://blog.bruce-hill.com/code?f=weighted-random/alias_with_index.py
// https://www.keithschwarz.com/darts-dice-coins/
//
// Additional Sources:
// http://blog.gainlo.co/index.php/2016/11/11/uber-interview-question-weighted-random-numbers/
func NewChoiceFromWeights(weights ...float64) Choice {
	numChoices := len(weights)
	sumChoices := sumFloat(weights)
	avgChoices := sumChoices / float64(numChoices)

	aliases := make([]alias, numChoices)
	for i := range aliases {
		aliases[i] = alias{1, -1}
	}

	smallIndex := 0
	bigIndex := 0

	for smallIndex < numChoices && weights[smallIndex] >= avgChoices {
		smallIndex += 1
	}

	// If all weights are the same, nothing to do
	if smallIndex < numChoices {
		small := &record{smallIndex, weights[smallIndex] / avgChoices}
		bigIndex = 0

		for bigIndex < numChoices && weights[bigIndex] < avgChoices {
			bigIndex += 1
		}

		big := &record{bigIndex, weights[bigIndex] / avgChoices}

		for big != nil && small != nil {
			aliases[small.index] = alias{small.value, big.index}
			big = &record{big.index, big.value - (1 - small.value)}

			if big.value < 1 {
				small = big
				bigIndex += 1

				for bigIndex < numChoices && weights[bigIndex] < avgChoices {
					bigIndex += 1
				}

				if bigIndex >= numChoices {
					break
				}
				big = &record{bigIndex, weights[bigIndex] / avgChoices}
			} else {
				smallIndex += 1
				for smallIndex < numChoices && weights[smallIndex] >= avgChoices {
					smallIndex += 1
				}

				if smallIndex >= numChoices {
					break
				}

				small = &record{smallIndex, weights[smallIndex] / avgChoices}
			}
		}
	}

	source := rand.NewSource(time.Now().UTC().UnixNano())
	randGenerator := rand.New(source)

	return Choice{
		numChoices: float64(numChoices),
		aliases:    aliases,
		rand:       randGenerator,
	}
}

func (this Choice) Next() int {
	floatIndex := this.rand.Float64() * this.numChoices
	intIndex := int(floatIndex)
	candidate := this.aliases[intIndex]
	odds := candidate.odds
	alias := candidate.alias

	if (floatIndex - float64(intIndex)) > odds {
		return alias
	} else {
		return intIndex
	}
}
