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

// createTable returns data in table format.
func createTable(data [][]string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	return table
}

// profileInfo returns profile info in rows format.
func profileInfo(v *we.CustomerInfo) [][]string {
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
	return data
}

// usageInfo returns usage info in rows format.
func usageInfo(v *we.UsageInfo) [][]string {
	data := [][]string{
		{"Start date", fmt.Sprintf("%v", v.AdslUsage.StartDate)},
		{"Quota", fmt.Sprintf("%v GB", v.AdslUsage.Quota)},
		{"Total Used", fmt.Sprintf("%v GB", v.AdslUsage.TotalUsed)},
	}
	return data
}

// daysInfo returns days info in rows format.
func daysInfo(v *we.RemainingDaysInfo) [][]string {
	data := [][]string{
		{"Expiry date", fmt.Sprintf("%v", v.RemainingDays.ADSLExpiryDate)},
		{"Amount due", fmt.Sprintf("%v", v.RemainingDays.AmountDue)},
		{"Package name", fmt.Sprintf("%v", v.RemainingDays.PackageName)},
		{"Remaining day", fmt.Sprintf("%v", v.RemainingDays.RemainingDays)},
	}
	return data
}

// displayPoints returns 4U points info in rows format.
func pointsInfo(v *we.LoyaltyPointsInfo) [][]string {
	data := [][]string{
		{"4U points", fmt.Sprintf("%v", v.LoyaltyPoints)},
	}
	return data
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
	c, err := client.Login()

	switch onlyVal {
	case profile:
		handleError(err)
		t := profileInfo(c)
		createTable(t).Render()
	case usage:
		u, err := client.Usage()
		handleError(err)
		t := usageInfo(u)
		createTable(t).Render()
	case days:
		d, err := client.RemainingDays()
		handleError(err)
		t := daysInfo(d)
		createTable(t).Render()
	case points:
		p, err := client.LoyaltyPoints()
		handleError(err)
		t := pointsInfo(p)
		createTable(t).Render()
	default:
		// display all data
		handleError(err)
		u, err := client.Usage()
		handleError(err)
		d, err := client.RemainingDays()
		handleError(err)
		p, err := client.LoyaltyPoints()
		handleError(err)

		t := createTable(profileInfo(c))
		t.AppendBulk(usageInfo(u))
		t.AppendBulk(daysInfo(d))
		t.AppendBulk(pointsInfo(p))
		t.Render()
	}
}
