package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mohan3d/my-we/we"
	"github.com/olekukonko/tablewriter"
)

const (
	profile     = "profile"
	usage       = "usage"
	days        = "days"
	points      = "points"
	envEmail    = "WE_EMAIL"
	envPassword = "WE_PASSWORD"
)

// displayTable creates and displays data in table format.
func displayTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.Render()
}

// displayProfile displays profile in table format.
func displayProfile(v *we.CustomerInfo) {
	data := [][]string{
		{"Customer name", fmt.Sprintf("%v", v.CustomerInformationDto.CustomerName)},
		{"Customer number", fmt.Sprintf("%v", v.CustomerInformationDto.CustomerNumber)},
		{"Email address", fmt.Sprintf("%v", v.CustomerInformationDto.EmailAddress)},
		{"Mobile number", fmt.Sprintf("%v", v.CustomerInformationDto.MobileNumber)},
		{"ADSL number", fmt.Sprintf("%v", v.CustomerInformationDto.ADSLNumber)},
		{"ADSL area code", fmt.Sprintf("%v", v.CustomerInformationDto.ADSLAreaCode)},
		{"ADSL speed", fmt.Sprintf("%v", v.CustomerInformationDto.ADSLSpeed)},
		{"City", fmt.Sprintf("%v", v.CustomerInformationDto.City)},
		{"District", fmt.Sprintf("%v", v.CustomerInformationDto.District)},
	}
	displayTable(data)
}

// displayUsage displays usage in table format.
func displayUsage(v *we.UsageInfo) {
	data := [][]string{
		{"Start date", fmt.Sprintf("%v", v.AdslUsage.StartDate)},
		{"Quota", fmt.Sprintf("%v GB", v.AdslUsage.Quota)},
		{"Total Used", fmt.Sprintf("%v GB", v.AdslUsage.TotalUsed)},
	}
	displayTable(data)
}

// displayDays displays remaining days in table format.
func displayDays(v *we.RemainingDaysInfo) {
	data := [][]string{
		{"Expiry date", fmt.Sprintf("%v", v.RemainingDays.ADSLExpiryDate)},
		{"Amount due", fmt.Sprintf("%v", v.RemainingDays.AmountDue)},
		{"Package name", fmt.Sprintf("%v", v.RemainingDays.PackageName)},
		{"Remaining day", fmt.Sprintf("%v", v.RemainingDays.RemainingDays)},
	}
	displayTable(data)
}

// displayPoints displays 4U points in table format.
func displayPoints(v *we.LoyaltyPointsInfo) {
	data := [][]string{
		{"4U points", fmt.Sprintf("%v", v.LoyaltyPoints)},
	}
	displayTable(data)
}

// emptyEmailOrPassword returns true if email or password is empty.
func emptyEmailOrPassword(email, password string) bool {
	return !(len(email) > 0 && len(password) > 0)
}

// handleError displays error and exists.
func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
		handleError(err)
		usageInfo, err := client.Usage()
		handleError(err)
		daysInfo, err := client.RemainingDays()
		handleError(err)
		pointsInfo, err := client.LoyaltyPoints()
		handleError(err)

		displayProfile(customerInfo)
		displayUsage(usageInfo)
		displayDays(daysInfo)
		displayPoints(pointsInfo)
	}
}
