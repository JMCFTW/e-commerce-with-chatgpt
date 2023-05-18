package main

import (
	"e-commerce-with-chatgpt/usermanagement"
	"fmt"
)

func main() {
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
		fmt.Println("create user fail")
	}
	fmt.Println(createdUser.Username)
	fmt.Println(createdUser.Email)
}
