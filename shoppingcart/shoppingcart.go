package shoppingcart

import "errors"

// CartData represents the data required to create a new shopping cart
type CartData struct {
	ID     string
	UserID string
}

// Cart represents a shopping cart
type Cart struct {
	ID     string
	UserID string
}

// CartRepository defines the interface for interacting with the shopping cart repository
type CartRepository interface {
	Create(cart Cart) error
	FindByID(id string) (Cart, error)
}

// CartUseCase represents the shopping cart use cases
type CartUseCase struct {
	cartRepository CartRepository
}

// NewCartUseCase creates a new instance of CartUseCase
func NewCartUseCase(cartRepository CartRepository) *CartUseCase {
	return &CartUseCase{
		cartRepository: cartRepository,
	}
}

// CreateCart creates a new shopping cart with the provided data
func (uc *CartUseCase) CreateCart(cartData CartData) (Cart, error) {
	// Check if the cart ID already exists
	_, err := uc.cartRepository.FindByID(cartData.ID)
	if err == nil {
		return Cart{}, errors.New("cart with the same ID already exists")
	}

	// Create the shopping cart
	cart := Cart{
		ID:     cartData.ID,
		UserID: cartData.UserID,
	}

	// Save the shopping cart
	err = uc.cartRepository.Create(cart)
	if err != nil {
		return Cart{}, errors.New("failed to create cart")
	}

	return cart, nil
}

// GetCart retrieves a shopping cart by its ID
func (uc *CartUseCase) GetCart(id string) (Cart, error) {
	cart, err := uc.cartRepository.FindByID(id)
	if err != nil {
		return Cart{}, errors.New("cart not found")
	}

	return cart, nil
}
