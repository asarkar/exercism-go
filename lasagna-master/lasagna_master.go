package lasagna

func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}
	return len(layers) * prepTimePerLayer
}

func Quantities(layers []string) (int, float64) {
	var noodlesAmountInGram int
	var sauceAmountInLitres float64
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodlesAmountInGram += 50
		case "sauce":
			sauceAmountInLitres += 0.2
		}
	}
	return noodlesAmountInGram, sauceAmountInLitres
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	n := float64(numPortions)
	for i := range quantities {
		scaledQuantities[i] = quantities[i] / 2 * n
	}
	return scaledQuantities
}
