package productcatalog

import (
	"errors"
	"testing"
)

type MockProductRepository struct {
	products map[string]Product
}

func (r *MockProductRepository) Create(product Product) error {
	// Check if the product with the same ID already exists
	_, ok := r.products[product.ID]
	if ok {
		return errors.New("product with the same ID already exists")
	}

	r.products[product.ID] = product
	return nil
}

func (r *MockProductRepository) FindByID(id string) (Product, error) {
	product, ok := r.products[id]
	if !ok {
		return Product{}, errors.New("product not found")
	}

	return product, nil
}

func TestCreateProduct(t *testing.T) {
	repo := &MockProductRepository{
		products: make(map[string]Product),
	}
	uc := NewProductUseCase(repo)

	// Test case 1: Creating a new product with valid data
	productData := ProductData{
		ID:       "p1",
		Name:     "Product 1",
		Price:    9.99,
		Quantity: 100,
	}

	createdProduct, err := uc.CreateProduct(productData)
	if err != nil {
		t.Errorf("Failed to create product: %v", err)
	}

	if createdProduct.ID != productData.ID {
		t.Errorf("Expected product ID: %s, got: %s", productData.ID, createdProduct.ID)
	}

	// Test case 2: Creating a product with an existing ID
	duplicateProductData := ProductData{
		ID:       "p1",
		Name:     "Duplicate Product",
		Price:    19.99,
		Quantity: 50,
	}

	_, err = uc.CreateProduct(duplicateProductData)
	if err == nil {
		t.Error("Expected an error when creating a product with an existing ID")
	}

	// Test case 3: Creating a product with an invalid price
	invalidPriceData := ProductData{
		ID:       "p2",
		Name:     "Product 2",
		Price:    -9.99,
		Quantity: 100,
	}

	_, err = uc.CreateProduct(invalidPriceData)
	if err == nil {
		t.Error("Expected an error when creating a product with an invalid price")
	}
}

func TestGetProduct(t *testing.T) {
	repo := &MockProductRepository{
		products: make(map[string]Product),
	}
	uc := NewProductUseCase(repo)

	// Create a test product
	productData := ProductData{
		ID:       "p1",
		Name:     "Product 1",
		Price:    9.99,
		Quantity: 100,
	}

	uc.CreateProduct(productData)

	// Test case 1: Getting an existing product
	retrievedProduct, err := uc.GetProduct(productData.ID)
	if err != nil {
		t.Errorf("Failed to get product: %v", err)
	}

	if retrievedProduct.ID != productData.ID {
		t.Errorf("Expected product ID: %s, got: %s", productData.ID, retrievedProduct.ID)
	}

	// Test case 2: Getting a non-existing product
	_, err = uc.GetProduct("nonexistentproduct")
	if err == nil {
		t.Error("Expected an error when getting a non-existing product")
	}
}
