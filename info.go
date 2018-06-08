package main

import (
	"fmt"
	"os"

	"github.com/mohan3d/my-we/we"
	"github.com/olekukonko/tablewriter"
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
