package lasagna

// PreparationTime
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}

	return len(layers) * time
}

// Quantities
func Quantities(quantity []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, quant := range quantity {
		if quant == "noodles" {
			noodles += 50
		} else if quant == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// AddSecretIngredient
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// ScaleRecipe
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaledAmounts := make([]float64, len(amounts))
	for index := range amounts {
		scaledAmounts[index] = float64(amounts[index] * float64(portions) / 2)
	}
	return scaledAmounts
}
