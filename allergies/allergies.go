// Package allergies decodes an allergy score into concrete values
package allergies

const testVersion = 1

// A listing of allergen code values
const (
	eggs = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

// allergenToScore maps allergen names to score values
var allergenToScore = map[string]uint{
	"eggs":         eggs,
	"peanuts":      peanuts,
	"shellfish":    shellfish,
	"strawberries": strawberries,
	"tomatoes":     tomatoes,
	"chocolate":    chocolate,
	"pollen":       pollen,
	"cats":         cats,
}

// Allergies gives back a decoded list of allergens from a score
func Allergies(score uint) []string {
	var allergens []string
	for allergen, code := range allergenToScore {
		if code&score > 0 {
			allergens = append(allergens, allergen)
		}
	}
	return allergens
}

// AllergicTo tests if a specific allergen is reported by the test
func AllergicTo(score uint, allergen string) bool {
	return allergenToScore[allergen]&score > 0
}
