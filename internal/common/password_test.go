package common

import (
	"testing"
)

func TestPasswordsMatch(t *testing.T) {
	t.Parallel()

	// Test case 1: Matching passwords
	hash := "$2a$10$e0RqJtK1KQqVc5PnYQBM/OlL9Q20o/6sVWqJt1eNfMlE5n5Y3uxaG"
	password := "password123"
	expected := true

	result := PasswordsMatch(hash, password)

	if result != expected {
		t.Errorf("PasswordsMatch() = %v, expected %v", result, expected)
	}

	// Test case 2: Non-matching passwords
	hash = "$2a$10$e0RqJtK1KQqVc5PnYQBM/OlL9Q20o/6sVWqJt1eNfMlE5n5Y3uxaG"
	password = "wrongpassword"
	expected = false

	result = PasswordsMatch(hash, password)

	if result != expected {
		t.Errorf("PasswordsMatch() = %v, expected %v", result, expected)
	}

	// Test case 3: Empty password
	hash = "$2a$10$e0RqJtK1KQqVc5PnYQBM/OlL9Q20o/6sVWqJt1eNfMlE5n5Y3uxaG"
	password = ""
	expected = false

	result = PasswordsMatch(hash, password)

	if result != expected {
		t.Errorf("PasswordsMatch() = %v, expected %v", result, expected)
	}
}
