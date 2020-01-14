package weightedrandom

import (
	"fmt"
	"testing"
)

func TestRandomChoice(t *testing.T) {
	choice := NewChoiceFromWeights(0.01, 0.02, 0.1, 1.0)
	for range make([]struct{}, 20) {
		next := choice.Next()
		fmt.Println(next)
	}
}
