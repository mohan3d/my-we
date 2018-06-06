package we

import "testing"

func unAuthenticatedClient() *Client {
	return nil
}

func authenticatedClient() (*Client, error) {
	return nil, nil
}

func TestAuthorizationToken(t *testing.T) {
	tests := []struct {
		username string
		password string
		expected string
	}{
		{"X", "Y", "WDpZ"},
		{"ABC", "DEF", "QUJDOkRFRg=="},
		{"AB123", "CD456", "QUIxMjM6Q0Q0NTY="},
	}

	for _, test := range tests {
		if actual := authorizationToken(test.username, test.password); actual != test.expected {
			t.Errorf(
				"authorizationToken(%v, %v): expected %v got %v",
				test.username,
				test.password,
				test.expected,
				actual,
			)
		}
	}
}
