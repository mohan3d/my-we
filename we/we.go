package we

func authorizationToken(username, password string) string {
	return ""
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
type Client struct{}

// Login submits email and password to be checked by backend.
func (c *Client) Login() (*CustomerInfo, error) {
	return nil, nil
}

// Usage returns UsageInfo of logged in user.
func (c *Client) Usage() (*UsageInfo, error) {
	return nil, nil
}

// RemainingDays returns service subscription of logged in user.
func (c *Client) RemainingDays() (*RemainingDaysInfo, error) {
	return nil, nil
}

// LoyaltyPoints returns 4U points of logged in user.
func (c *Client) LoyaltyPoints() (*LoyaltyPointsInfo, error) {
	return nil, nil
}

// New returns new we client initialized with email and password.
func New(email, password string) *Client {
	return nil
}
