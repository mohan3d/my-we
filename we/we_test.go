package we

import (
	"os"
	"testing"
)

func unAuthenticatedClient() *Client {
	return New("testName", "testPassword")
}

func authenticatedClient() (*Client, error) {
	email := os.Getenv("WE_EMAIL")
	password := os.Getenv("WE_PASSWORD")

	client := New(email, password)
	if _, err := client.Login(); err != nil {
		return nil, err
	}
	return client, nil
}

func TestNewRequest(t *testing.T) {
	tests := []struct {
		key   string
		value string
	}{
		{"Content-Type", "application/json"},
		// base64("abc@xyz.com:abcdef12345") = "YWJjQHh5ei5jb206YWJjZGVmMTIzNDU="
		{"Authorization", "Basic YWJjQHh5ei5jb206YWJjZGVmMTIzNDU="},
	}
	client := New("abc@xyz.com", "abcdef12345")
	r, err := client.newRequest("", "", nil)
	if err != nil {
		t.Error(err)
	}
	headers := r.Header
	for _, test := range tests {
		if v := headers.Get(test.key); v != test.value {
			t.Errorf("expected value of %s=\"%s\", got \"%s\"", test.key, test.value, v)
		}
	}
}

func TestLoginInvalidCred(t *testing.T) {
	client := unAuthenticatedClient()
	_, err := client.Login()
	if err == nil {
		t.Error("expected to get error")
	}
}

func TestLoginValidCred(t *testing.T) {
	_, err := authenticatedClient()
	if err != nil {
		t.Errorf("expected to login with no errors, got %v", err)
	}
}

func TestUsageInvalidCred(t *testing.T) {
	client := unAuthenticatedClient()
	_, err := client.Usage()
	if err == nil {
		t.Error("expected to get error")
	}
}

func TestUsageValidCred(t *testing.T) {
	client, _ := authenticatedClient()
	_, err := client.Usage()
	if err != nil {
		t.Errorf("expected to retrieve usage info with no errors, got %v", err)
	}
}

func TestRemainingDaysInvalidCred(t *testing.T) {
	client := unAuthenticatedClient()
	_, err := client.RemainingDays()
	if err == nil {
		t.Error("expected to get error")
	}
}

func TestRemainingDaysValidCred(t *testing.T) {
	client, _ := authenticatedClient()
	_, err := client.RemainingDays()
	if err != nil {
		t.Errorf("expected to retrieve remaining days info with no errors, got %v", err)
	}
}

func TestLoyaltyPointsInvalidCred(t *testing.T) {
	client := unAuthenticatedClient()
	_, err := client.LoyaltyPoints()
	if err == nil {
		t.Error("expected to get error")
	}
}

func TestLoyaltyPointsValidCred(t *testing.T) {
	client, _ := authenticatedClient()
	_, err := client.LoyaltyPoints()
	if err != nil {
		t.Errorf("expected to retrieve loyalty points info with no errors, got %v", err)
	}
}
