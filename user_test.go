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

func TestUser_HashPassword(t *testing.T) {
	user := User{Email: "test@example.com", Password: "plainpassword"}
	originalPassword := user.Password

	err := user.HashPassword()

	assert.NoError(t, err, "Hashing password should not return error")
	assert.NotEqual(t, originalPassword, user.Password, "Password should be hashed")
	assert.True(t, len(user.Password) > 0, "Hashed password should not be empty")
}

func TestUser_CheckPassword(t *testing.T) {
	user := User{Email: "test@example.com", Password: "plainpassword"}
	user.HashPassword()

	valid := user.checkPassword("plainpassword")
	assert.True(t, valid, "Correct password should return true")

	invalid := user.checkPassword("notplainpassword")
	assert.False(t, invalid, "Invalid password should return false")
}
