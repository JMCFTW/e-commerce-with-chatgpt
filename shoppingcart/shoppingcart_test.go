package shoppingcart

import (
	"errors"
	"testing"
)

type MockCartRepository struct {
	carts map[string]Cart
}

func (r *MockCartRepository) Create(cart Cart) error {
	// Check if the cart with the same ID already exists
	_, ok := r.carts[cart.ID]
	if ok {
		return errors.New("cart with the same ID already exists")
	}

	r.carts[cart.ID] = cart
	return nil
}

func (r *MockCartRepository) FindByID(id string) (Cart, error) {
	cart, ok := r.carts[id]
	if !ok {
		return Cart{}, errors.New("cart not found")
	}

	return cart, nil
}

func TestCreateCart(t *testing.T) {
	repo := &MockCartRepository{
		carts: make(map[string]Cart),
	}
	uc := NewCartUseCase(repo)

	// Test case 1: Creating a new cart with valid data
	cartData := CartData{
		ID:     "c1",
		UserID: "user123",
	}

	createdCart, err := uc.CreateCart(cartData)
	if err != nil {
		t.Errorf("Failed to create cart: %v", err)
	}

	if createdCart.ID != cartData.ID {
		t.Errorf("Expected cart ID: %s, got: %s", cartData.ID, createdCart.ID)
	}

	// Test case 2: Creating a cart with an existing ID
	duplicateCartData := CartData{
		ID:     "c1",
		UserID: "user456",
	}

	_, err = uc.CreateCart(duplicateCartData)
	if err == nil {
		t.Error("Expected an error when creating a cart with an existing ID")
	}
}

func TestGetCart(t *testing.T) {
	repo := &MockCartRepository{
		carts: make(map[string]Cart),
	}
	uc := NewCartUseCase(repo)

	// Create a test cart
	cartData := CartData{
		ID:     "c1",
		UserID: "user123",
	}

	uc.CreateCart(cartData)

	// Test case 1: Getting an existing cart
	retrievedCart, err := uc.GetCart(cartData.ID)
	if err != nil {
		t.Errorf("Failed to get cart: %v", err)
	}

	if retrievedCart.ID != cartData.ID {
		t.Errorf("Expected cart ID: %s, got: %s", cartData.ID, retrievedCart.ID)
	}

	// Test case 2: Getting a non-existing cart
	_, err = uc.GetCart("nonexistentcart")
	if err == nil {
		t.Error("Expected an error when getting a non-existing cart")
	}
}
