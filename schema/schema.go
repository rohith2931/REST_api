package schema

import "github.com/jinzhu/gorm"

func SetUpDB(db *gorm.DB) {
	db.DropTable(&Product{})
	db.CreateTable(&Product{})
	db.DropTable(&Rating{})
	db.CreateTable(&Rating{})
	db.DropTable(&Variant{})
	db.CreateTable(&Variant{})
	db.Model(&Rating{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	db.Model(&Variant{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	products_list := []Product{
		{Name: "Iphone14", Description: "latest iphone", Category: "Mobile", Quantity: 1, Price: 3999, Image: "iphone Image",
			Rating: []Rating{
				{Name: "ramu", Review: "Good", Rating: 4},
				{Name: "raj", Review: "Nice phone", Rating: 5},
				{Name: "jhon", Review: "excellent", Rating: 4},
			},
			Variants: []Variant{
				{Color: "blue", Image: "blue Image"},
				{Color: "black", Image: "black Image"},
			}},
		{Name: "One plus", Description: "latest oneplus modile", Category: "Mobile", Quantity: 2, Price: 2999, Image: "OnePlus Image",
			Rating: []Rating{
				{Name: "vicky", Review: "Nice camera quality", Rating: 4},
				{Name: "vikram", Review: "Good OS", Rating: 5},
			},
			Variants: []Variant{
				{Color: "blue", Image: "blue Image"},
				{Color: "white", Image: "white Image"},
			}},
	}
	for _, i := range products_list {
		db.Create(&i)
	}
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Category    string
	Quantity    int
	Price       int
	Image       string
	Rating      []Rating
	Variants    []Variant
}

type Rating struct {
	gorm.Model
	ProductID uint
	Name      string
	Review    string
	Rating    int
}

type Variant struct {
	gorm.Model
	ProductID uint
	Color     string
	Image     string
}
