package models

import (
	"encoding/json"
	"testing"
)

func TestAdminUnmarshalAndValidation(t *testing.T) {
	t.Run("Valid User Token", func(t *testing.T) {
		input := `{
			"email": "user@example.com",
			"first_name": "John",
			"last_name": "Doe",
            "tags": [
                "mist-customer"
            ],
            "privileges": [
                {
                    "scope": "org",
                    "org_id": "00000000-0000-0000-0000-000000000000",
                    "role": "admin",
                    "name": "Test Org"
                }
            ]
		}`
		var admin Admin
		err := json.Unmarshal([]byte(input), &admin)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

	})

	t.Run("Valid Org Token", func(t *testing.T) {
		input := `{
			"name": "test-org",
            "privileges": [
                {
                    "scope": "org",
                    "org_id": "00000000-0000-0000-0000-000000000000",
                    "role": "admin",
                    "name": "Test Org"
                }
            ]
		}`
		var admin Admin
		err := json.Unmarshal([]byte(input), &admin)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

	})

	t.Run("Invalid Token - Missing Email", func(t *testing.T) {
		input := `{
			"first_name": "John",
			"last_name": "Doe"
		}`
		var admin Admin
		err := json.Unmarshal([]byte(input), &admin)
		if err == nil {
			t.Fatal("expected error, got none")
		}
		expectedError := "required field `email` is missing for type `Admin`"
		if err.Error() != expectedError {
			t.Fatalf("expected error '%s', got '%v'", expectedError, err)
		}
	})

	t.Run("Invalid Token - Missing First Name", func(t *testing.T) {
		input := `{
			"email": "user@example.com",
			"last_name": "Doe"
		}`
		var admin Admin
		err := json.Unmarshal([]byte(input), &admin)
		if err == nil {
			t.Fatal("expected error, got none")
		}
		expectedError := "required field `first_name` is missing for type `Admin`"
		if err.Error() != expectedError {
			t.Fatalf("expected error '%s', got '%v'", expectedError, err)
		}
	})

	t.Run("Invalid Token - Missing Last Name", func(t *testing.T) {
		input := `{
			"email": "user@example.com",
			"first_name": "Doe"
		}`
		var admin Admin
		err := json.Unmarshal([]byte(input), &admin)
		if err == nil {
			t.Fatal("expected error, got none")
		}
		expectedError := "required field `last_name` is missing for type `Admin`"
		if err.Error() != expectedError {
			t.Fatalf("expected error '%s', got '%v'", expectedError, err)
		}
	})

	t.Run("Invalid Token - Missing Name and User Info", func(t *testing.T) {
		input := `{
			"tags": ["tag1", "tag2"]
		}`
		var admin Admin
		err := json.Unmarshal([]byte(input), &admin)
		if err == nil {
			t.Fatal("expected error, got none")
		}
		expectedError := "required field `email` is missing for type `Admin`\nrequired field `first_name` is missing for type `Admin`\nrequired field `last_name` is missing for type `Admin`"
		if err.Error() != expectedError {
			t.Fatalf("expected error '%s', got '%v'", expectedError, err)
		}
	})
}