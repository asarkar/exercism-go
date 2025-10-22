package lasagna

import "slices"

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
	n := float64(numPortions)
	// Go >= 1.23.
	// https://stackoverflow.com/a/78185810/839733
	// https://bitfieldconsulting.com/posts/iterators
	return slices.Collect(
		func(yield func(float64) bool) {
			for i := range quantities {
				if !yield(quantities[i] / 2 * n) {
					return
				}
			}
		},
	)
}
