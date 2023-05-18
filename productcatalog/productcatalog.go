package productcatalog

import "errors"

// ProductData represents the data required to create a new product
type ProductData struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

// Product represents a product in the catalog
type Product struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

// ProductRepository defines the interface for interacting with the product repository
type ProductRepository interface {
	Create(product Product) error
	FindByID(id string) (Product, error)
}

// ProductUseCase represents the product catalog use cases
type ProductUseCase struct {
	productRepository ProductRepository
}

// NewProductUseCase creates a new instance of ProductUseCase
func NewProductUseCase(productRepository ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepository: productRepository,
	}
}

// CreateProduct creates a new product with the provided data
func (uc *ProductUseCase) CreateProduct(productData ProductData) (Product, error) {
	// Check if the product ID already exists
	_, err := uc.productRepository.FindByID(productData.ID)
	if err == nil {
		return Product{}, errors.New("product with the same ID already exists")
	}

	// Validate the price
	if productData.Price <= 0 {
		return Product{}, errors.New("invalid product price")
	}

	// Create the product
	product := Product{
		ID:       productData.ID,
		Name:     productData.Name,
		Price:    productData.Price,
		Quantity: productData.Quantity,
	}

	// Save the product
	err = uc.productRepository.Create(product)
	if err != nil {
		return Product{}, errors.New("failed to create product")
	}

	return product, nil
}

// GetProduct retrieves a product by its ID
func (uc *ProductUseCase) GetProduct(id string) (Product, error) {
	product, err := uc.productRepository.FindByID(id)
	if err != nil {
		return Product{}, errors.New("product not found")
	}

	return product, nil
}
