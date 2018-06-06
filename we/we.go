package we

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiURL           = "https://mytedata.net/services/rest"
	subscriptionURL  = apiURL + "/subscription/customer/%s"
	loginURL         = apiURL + "/login/checkPassword"
	usageURL         = subscriptionURL + "/ADSLUsage"
	remainingDaysURL = subscriptionURL + "/remainingDays"
	loyaltyPointsURL = subscriptionURL + "/loyaltyPoints"
)

// Calculates base64 of username/password.
func authorizationToken(username, password string) string {
	concat := fmt.Sprintf("%s:%s", username, password)
	return base64.StdEncoding.EncodeToString([]byte(concat))
}

// Credentials represents login credentials.
type Credentials struct {
	Password string `json:"password"`
	UID      string `json:"uid"`
}

// CustomerInfo describes customer info.
type CustomerInfo struct {
	CustomerInformationDto struct {
		ADSLSpeed      string `json:"adslSpeed"`
		City           string `json:"cityEN"`
		District       string `json:"districtEN"`
		MobileNumber   string `json:"mobileNumber1WithPrefix"`
		CustomerNumber string `json:"customerNumber"`
		EmailAddress   string `json:"emailAddress"`
		ADSLAreaCode   int    `json:"adslAreaCode"`
		ADSLNumber     int    `json:"adslNumber"`
		CustomerName   string `json:"customerName"`
	} `json:"customerInformationDto"`
}

// UsageInfo describes ADSL usage info.
type UsageInfo struct {
	AdslUsage struct {
		StartDate int64   `json:"startDate"`
		Quota     float64 `json:"quata"`
		TotalUsed float64 `json:"totalUsed"`
	} `json:"adslUsage"`
}

// RemainingDaysInfo describes subscription date and remaining service days.
type RemainingDaysInfo struct {
	RemainingDays struct {
		ADSLExpiryDate string  `json:"adslExpiryDateString"`
		RemainingDays  int     `json:"remainingDays"`
		PackageName    string  `json:"packageName"`
		AmountDue      float64 `json:"amountDue"`
	} `json:"remainingDays"`
}

// LoyaltyPointsInfo describes 4U points.
type LoyaltyPointsInfo struct {
	LoyaltyPoints int `json:"loyaltyPoints"`
}

// Client describes we api client.
type Client struct {
	username string
	password string
	token    string
	client   http.Client
}

// Login submits email and password to be checked by backend.
func (c *Client) Login() (*CustomerInfo, error) {
	cred := Credentials{UID: c.username, Password: c.password}
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(cred)
	if err != nil {
		return nil, err
	}
	r, err := c.post(loginURL, body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	customerInfo := new(CustomerInfo)
	err = json.NewDecoder(r.Body).Decode(customerInfo)
	if err != nil {
		return nil, err
	}
	return customerInfo, nil
}

// Usage returns UsageInfo of logged in user.
func (c *Client) Usage() (*UsageInfo, error) {
	r, err := c.get(usageURL)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	usageInfo := new(UsageInfo)
	err = json.NewDecoder(r.Body).Decode(usageInfo)
	if err != nil {
		return nil, err
	}
	return usageInfo, nil
}

// RemainingDays returns service subscription of logged in user.
func (c *Client) RemainingDays() (*RemainingDaysInfo, error) {
	r, err := c.get(remainingDaysURL)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	remainingDaysInfo := new(RemainingDaysInfo)
	err = json.NewDecoder(r.Body).Decode(remainingDaysInfo)
	if err != nil {
		return nil, err
	}
	return remainingDaysInfo, nil
}

// LoyaltyPoints returns 4U points of logged in user.
func (c *Client) LoyaltyPoints() (*LoyaltyPointsInfo, error) {
	r, err := c.get(loyaltyPointsURL)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	loyaltyPointsInfo := new(LoyaltyPointsInfo)
	err = json.NewDecoder(r.Body).Decode(loyaltyPointsInfo)
	if err != nil {
		return nil, err
	}
	return loyaltyPointsInfo, nil
}

// get creates a GET request to url.
func (c *Client) get(url string) (*http.Response, error) {
	r, err := c.newRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

// post creates a POST request to url.
func (c *Client) post(url string, body io.Reader) (*http.Response, error) {
	r, err := c.newRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

// newRequest creates new request with required headers.
func (c *Client) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	r.SetBasicAuth(c.username, c.password)
	r.Header.Add("Content-Type", "application/json")
	return r, nil
}

// New returns new we client initialized with email and password.
func New(email, password string) *Client {
	client := new(Client)
	client.username = email
	client.password = password
	return client
}
