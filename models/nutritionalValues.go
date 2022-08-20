// besin değerlerinin tutulduğu tablo
package models

type NutritionalValueType int8

const (
	Fish                 NutritionalValueType = iota + 1
	Offal                                     // sakatat
	MeatProducts                              // et ürünleri
	VegetablesAndLegumes                      // sebze ve baklagiller
	Fruit                                     // meyve
	DairyProducts                             // süt ürünleri
	Spices                                    // baharatlar
	FlourProducts                             // un ürünleri
	Desserts                                  // tatlılar
	Nuts                                      // kuru yemişler
	Others                                    // diğerleri
)

type NutritionalValue struct {
	ID           uint   `gorm:"primarykey"`
	Name         string `json:"name"`
	Calorie      int    `json:"calorie"`
	Protein      int    `json:"protein"`
	Oil          int    `json:"oil"`
	Carbohydrate int    `json:"carbohydrate"`
}
