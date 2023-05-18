package main

import (
  "e-commerce-with-chatgpt/usermanagement"
  "fmt"
)

func main() {
	um := usermanagement.NewUserManager()

	// Test case 1: Creating a new user with valid data
	userData := usermanagement.UserData{
		Username: "john_doe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}

	createdUser, err := um.CreateUser(userData)
	if err != nil {
	  fmt.Println("create user fail")
	}
	fmt.Println(createdUser.Username)
	fmt.Println(createdUser.Email)
}
