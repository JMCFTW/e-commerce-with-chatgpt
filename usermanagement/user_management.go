package usermanagement

import "errors"

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

// UserManager is responsible for user management operations
type UserManager struct {
	users map[string]User // Assuming username is unique
}

// NewUserManager creates a new instance of UserManager
func NewUserManager() *UserManager {
	return &UserManager{
		users: make(map[string]User),
	}
}

// CreateUser creates a new user with the provided data
func (um *UserManager) CreateUser(userData UserData) (User, error) {
	// Check if the username already exists
	if _, ok := um.users[userData.Username]; ok {
		return User{}, errors.New("username already exists")
	}

	// Validate the email format
	// Note: This is a simple validation, you can use a more comprehensive email validation library
	// or implement your own validation logic
	if !isValidEmail(userData.Email) {
		return User{}, errors.New("invalid email format")
	}

	// Create the user
	user := User{
		Username: userData.Username,
		Email:    userData.Email,
	}

	// Save the user
	um.users[userData.Username] = user

	return user, nil
}

// AuthenticateUser authenticates a user with the provided credentials
func (um *UserManager) AuthenticateUser(username, password string) (User, error) {
	user, ok := um.users[username]
	if !ok {
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
