package usermanagement

import (
	"errors"
)

// UserData represents the data required to create a new user
type UserData struct {
	Username string
	Email    string
	Password string
}

// User represents a user in the system
type User struct {
	Username string
	Email    string
}

// UserRepository defines the interface for interacting with the user repository
type UserRepository interface {
	Create(user User) error
	FindByUsername(username string) (User, error)
}

// UserUseCase represents the user management use cases
type UserUseCase struct {
	userRepository UserRepository
}

// NewUserUseCase creates a new instance of UserUseCase
func NewUserUseCase(userRepository UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

// CreateUser creates a new user with the provided data
func (uc *UserUseCase) CreateUser(userData UserData) (User, error) {
	// Check if the username already exists
	_, err := uc.userRepository.FindByUsername(userData.Username)
	if err == nil {
		return User{}, errors.New("username already exists")
	}

	// Validate the email format
	if !isValidEmail(userData.Email) {
		return User{}, errors.New("invalid email format")
	}

	// Create the user
	user := User{
		Username: userData.Username,
		Email:    userData.Email,
	}

	// Save the user
	err = uc.userRepository.Create(user)
	if err != nil {
		return User{}, errors.New("failed to create user")
	}

	return user, nil
}

// AuthenticateUser authenticates a user with the provided credentials
func (uc *UserUseCase) AuthenticateUser(username, password string) (User, error) {
	user, err := uc.userRepository.FindByUsername(username)
	if err != nil {
		return User{}, errors.New("user not found")
	}

	// Check the password
	if password != "password123" { // Replace with your password validation logic
		return User{}, errors.New("invalid password")
	}

	return user, nil
}

// isValidEmail checks if the provided email address has a valid format
func isValidEmail(email string) bool {
	// Add your email validation logic here
	// This is a simplified example, you can use a regular expression or a library for better validation
	// e.g., you can use a library like github.com/badoux/checkmail
	return true
}
