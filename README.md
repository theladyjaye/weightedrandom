# Weight Random Selection
Ported from Python
https://blog.bruce-hill.com/a-faster-weighted-random-choice

Ported directly from here:
https://blog.bruce-hill.com/code?f=weighted-random/alias_with_index.py
https://www.keithschwarz.com/darts-dice-coins/

Additional Sources:
http://blog.gainlo.co/index.php/2016/11/11/uber-interview-question-weighted-random-numbers/

## Usage Example

```go
gear := []WeightedGear{
    WeightedGear{Weight: 0.05, Item: theItem},
    WeightedGear{Weight: 0.25, Item: theNextItem},
    WeightedGear{Weight: 1.0, Item: theLastItem},
}

weights := make(float64[], len(gear))

for i, g := range gear{
    weights[i] = g.Weights,
}

choices := weightedrandom.NewChoiceFromWeights(weights...)
nextItem := gear[choices.Next()].Item
```