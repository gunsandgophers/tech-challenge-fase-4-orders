package entities

type ProductCategory string

func (s ProductCategory) String() string {
	return string(s)
}

const (
	PRODUCT_CATEGORY_SANDWICH   ProductCategory = "SANDWICH"
	PRODUCT_CATEGORY_SIDEDISHES ProductCategory = "SIDEDISHES"
	PRODUCT_CATEGORY_DRINKS     ProductCategory = "DRINKS"
	PRODUCT_CATEGORY_DESSERTS   ProductCategory = "DESSETS"
)

type Product struct {
	id          string
	name        string
	category    ProductCategory
	price       float64
	description string
	image       string
}

func RestoreProduct(
	id string,
	name string,
	category ProductCategory,
	price float64,
	description string,
	image string,
) *Product {
	return &Product{
		id:          id,
		name:        name,
		category:    category,
		price:       price,
		description: description,
		image:       image,
	}
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetId() string {
	return p.id
}

func (p *Product) GetCategory() ProductCategory {
	return p.category
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) GetDescription() string {
	return p.description
}

func (p *Product) GetImage() string {
	return p.image
}
