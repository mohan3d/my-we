package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mohan3d/my-we/we"
)

const (
	profile     = "profile"
	usage       = "usage"
	days        = "days"
	points      = "points"
	envEmail    = "WE_EMAIL"
	envPassword = "WE_PASSWORD"
)

func displayProfile(v *we.CustomerInfo) {

}

func displayUsage(v *we.UsageInfo) {

}

func displayDays(v *we.RemainingDaysInfo) {

}

func displayPoints(v *we.LoyaltyPointsInfo) {

}

func emptyEmailOrPassword(email, password string) bool {
	return !(len(email) > 0 && len(password) > 0)
}

func handleError(err error) {

}

func main() {
	only := flag.String("only", "", "info to be displayed, should be one of (profile, usage, days and points)")
	email := flag.String("email", "", "tedata account email")
	password := flag.String("password", "", "tedata account password")
	flag.Parse()

	onlyVal := strings.ToLower(*only)
	emailVal := *email
	passwordVal := *password

	if emptyEmailOrPassword(emailVal, passwordVal) {
		// Read from env variables.
		emailVal = os.Getenv(envEmail)
		passwordVal = os.Getenv(envPassword)
		if emptyEmailOrPassword(emailVal, passwordVal) {
			fmt.Println("please provide email and password")
			flag.Usage()
			os.Exit(1)
		}
	}

	client := we.New(emailVal, passwordVal)
	customerInfo, err := client.Login()

	switch onlyVal {
	case profile:
		handleError(err)
		displayProfile(customerInfo)
	case usage:
		usageInfo, err := client.Usage()
		handleError(err)
		displayUsage(usageInfo)
	case days:
		daysInfo, err := client.RemainingDays()
		handleError(err)
		displayDays(daysInfo)
	case points:
		pointsInfo, err := client.LoyaltyPoints()
		handleError(err)
		displayPoints(pointsInfo)
	default:
		// display all data
		fmt.Println("default")
	}
}
