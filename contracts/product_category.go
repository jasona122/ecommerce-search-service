package contracts

type ProductCategory string

const (
	Electronics   ProductCategory = "Electronics"
	FoodAndDrinks ProductCategory = "Food & Drinks"
	Cosmetics     ProductCategory = "Cosmetics"
	Health        ProductCategory = "Health and Wellness"
	Furniture     ProductCategory = "Furniture and Decor"
	Media         ProductCategory = "Media"
	PetCare       ProductCategory = "Pet Care"
	Office        ProductCategory = "Office Equipment"
	Clothing      ProductCategory = "Clothing"
)

func IsValidProductCategory(category string) bool {
	switch ProductCategory(category) {
	case Electronics, FoodAndDrinks, Cosmetics, Health, Furniture, Media, PetCare, Office, Clothing:
		return true
	default:
		return false
	}
}
