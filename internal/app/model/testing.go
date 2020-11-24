package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "testuser@test.com",
		Password: "111111",
	}
}
