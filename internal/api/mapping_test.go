package api

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"

	"github.com/seanpar203/go-api/internal/api/oas"
	"github.com/seanpar203/go-api/internal/models"
)

func TestV1User(t *testing.T) {
	t.Parallel()

	// Test case 1: Mapping user model with valid data
	user := &models.User{
		ID:    "123456",
		Name:  null.String{String: "John Doe", Valid: true},
		Email: "john.doe@example.com",
	}
	expected := &oas.V1User{
		ID:    oas.UUID(uuid.MustParse("123456")),
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	result := V1User(user)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Expected: %+v, but got: %+v", expected, result)
	}

	// Test case 2: Mapping user model with empty name
	user = &models.User{
		ID:    "789012",
		Name:  null.String{String: "", Valid: false},
		Email: "jane.doe@example.com",
	}
	expected = &oas.V1User{
		ID:    oas.UUID(uuid.MustParse("789012")),
		Name:  "",
		Email: "jane.doe@example.com",
	}
	result = V1User(user)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 2 failed. Expected: %+v, but got: %+v", expected, result)
	}
}

func TestV1LoginUserResponse(t *testing.T) {
	t.Parallel()

	// Test case 1: Testing with a valid user, access token, and refresh token
	user := &models.User{ /* create a valid user object */ }
	at := &models.AccessToken{ /* create a valid access token object */ }
	rt := &models.RefreshToken{ /* create a valid refresh token object */ }
	expectedUser := V1User(user)
	expectedAccessToken := oas.UUID(uuid.MustParse(at.Token))
	expectedRefreshToken := oas.UUID(uuid.MustParse(rt.Token))

	result := V1LoginUserResponse(user, at, rt)

	if !reflect.DeepEqual(result.User, expectedUser) {
		t.Errorf("Expected User to be %v, but got %v", expectedUser, result.User)
	}

	if result.AccessToken != expectedAccessToken {
		t.Errorf("Expected AccessToken to be %v, but got %v", expectedAccessToken, result.AccessToken)
	}

	if result.RefreshToken != expectedRefreshToken {
		t.Errorf("Expected RefreshToken to be %v, but got %v", expectedRefreshToken, result.RefreshToken)
	}
}

func TestV1AuthRefreshResponse(t *testing.T) {
	t.Parallel()

	at := &models.AccessToken{Token: "access_token"}

	expectedAccessToken := oas.UUID(uuid.MustParse(at.Token))

	resp := V1AuthRefreshResponse(at)

	if resp.AccessToken != expectedAccessToken {
		t.Errorf("Expected access token: %v, but got: %v", expectedAccessToken, resp.AccessToken)
	}

}
