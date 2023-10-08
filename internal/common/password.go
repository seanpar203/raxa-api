package common

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a hashed password.
//
// Takes a password string as input and returns the hashed password string and an error.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

// CompareHashAndPassword compares a bcrypt hashed password with its possible
// plaintext equivalent. Returns nil on success, or an error on failure.
//
// hash: The bcrypt hashed password.
// password: The possible plaintext equivalent of the hashed password.
// error: A possible error returned if the comparison fails.
func PasswordsMatch(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
