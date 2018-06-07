# my-we

Golang CLI for WE https://mytedata.net

## Installation
```bash
$ go get github.com/mohan3d/my-we
```

## Testing
**WE_EMAIL** and **WE_PASSWORD** must be exported to environment variables before running tests.

```bash
$ export WE_EMAIL=<YOUR_APIXU_KEY>
$ export WE_PASSWORD=<YOUR_APIXU_KEY>
$ cd $GOPATH/src/github.com/mohan3d/my-we
$ go test ./we
```