# my-we
[![Go Report Card](https://goreportcard.com/badge/github.com/mohan3d/my-we)](https://goreportcard.com/report/github.com/mohan3d/my-we)

Golang CLI for WE https://mytedata.net

## Installation
```bash
$ go get github.com/mohan3d/my-we
```
## Usage

```bash
# Display all info about your account.
$ my-we -email <WE_EMAIL> -password <WE_PASSWORD>

# Display only usage of your service.
# -only value must be one of (profile, usage, days and points).
$ my-we -email <WE_EMAIL> -password <WE_PASSWORD> -only usage

# If email or password not provided.
# email/password will be read from env variables.
$ export WE_EMAIL=<WE_ACCOUNT_EMAIL>
$ export WE_PASSWORD=<WE_ACCOUNT_PASSWORD>
$ my-we -only profile
```

## Testing
**WE_EMAIL** and **WE_PASSWORD** must be exported to environment variables before running tests.

```bash
$ export WE_EMAIL=<WE_ACCOUNT_EMAIL>
$ export WE_PASSWORD=<WE_ACCOUNT_PASSWORD>
$ cd $GOPATH/src/github.com/mohan3d/my-we
$ go test ./we
```