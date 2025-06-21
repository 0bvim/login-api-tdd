package main

import "regexp"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) ValidateEmail() bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(u.Email)
}
