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

// emptyEmailOrPassword returns true if email or password is empty.
func emptyEmailOrPassword(email, password string) bool {
	return !(len(email) > 0 && len(password) > 0)
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
	info := getInfo(client, onlyVal, true)
	createTable(info).Render()
}
