package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_ValidateEmail(t *testing.T) {
	user := User{Email: "test@example.com", Password: "password123"}

	valid := user.ValidateEmail()
	assert.True(t, valid, "Valid email should return true")
}

func TestUser_ValidateEmail_Invalid(t *testing.T) {
	user := User{Email: "invalid-email", Password: "password123"}

	valid := user.ValidateEmail()
	assert.False(t, valid, "Invalid email should return false")
}
