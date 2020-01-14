package weightedrandom

func sumFloat(values []float64) float64 {
	result := 0.0
	for _, each := range values {
		result = result + each
	}

	return result
}
