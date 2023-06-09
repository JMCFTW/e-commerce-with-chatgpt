package usermanagement

import (
	"errors"
	"testing"
)

type MockUserRepository struct {
	users map[string]User
}

func (r *MockUserRepository) Create(user User) error {
	// Check if the username already exists
	_, ok := r.users[user.Username]
	if ok {
		return errors.New("username already exists")
	}

	r.users[user.Username] = user
	return nil
}

func (r *MockUserRepository) FindByUsername(username string) (User, error) {
	user, ok := r.users[username]
	if !ok {
		return User{}, errors.New("user not found")
	}

	return user, nil
}

func TestCreateUser(t *testing.T) {
	repo := &MockUserRepository{
		users: make(map[string]User),
	}
	uc := NewUserUseCase(repo)

	// Test case 1: Creating a new user with valid data
	userData := UserData{
		Username: "john_doe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}

	createdUser, err := uc.CreateUser(userData)
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}

	if createdUser.Username != userData.Username {
		t.Errorf("Expected username: %s, got: %s", userData.Username, createdUser.Username)
	}

	// Test case 2: Creating a user with an existing username
	duplicateUserData := UserData{
		Username: "john_doe",
		Email:    "johndoe@example.com",
		Password: "newpassword",
	}

	_, err = uc.CreateUser(duplicateUserData)
	if err == nil {
		t.Error("Expected an error when creating a user with an existing username")
	}

	// Test case 3: Creating a user with an invalid email
	invalidEmailData := UserData{
		Username: "jane_smith",
		Email:    "invalidemail",
		Password: "password123",
	}

	_, err = uc.CreateUser(invalidEmailData)
	if err != nil {
		t.Error("Expected an error when creating a user with an invalid email")
	}
}

func TestAuthenticateUser(t *testing.T) {
	repo := &MockUserRepository{
		users: make(map[string]User),
	}
	uc := NewUserUseCase(repo)

	// Create a test user
	userData := UserData{
		Username: "johndoe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}

	uc.CreateUser(userData)

	// Test case 1: Authenticating a user with valid credentials
	authenticatedUser, err := uc.AuthenticateUser(userData.Username, userData.Password)
	if err != nil {
		t.Errorf("Failed to authenticate user: %v", err)
	}

	if authenticatedUser.Username != userData.Username {
		t.Errorf("Expected username: %s, got: %s", userData.Username, authenticatedUser.Username)
	}

	// Test case 2: Authenticating a user with an incorrect password
	_, err = uc.AuthenticateUser(userData.Username, "incorrectpassword")
	if err == nil {
		t.Error("Expected an error when authenticating a user with an incorrect password")
	}

	// Test case 3: Authenticating a non-existing user
	_, err = uc.AuthenticateUser("nonexistentuser", "password")
	if err == nil {
		t.Error("Expected an error when authenticating a non-existing user")
	}
}
