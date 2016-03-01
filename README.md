# instagram-beego

## How to run

1. Create Amazone S3 bucket & account
2. Copy `/.env.sample` to `/.env` and fill S3 info into `/.env` file
3. Add `$GOPATH/bin` to enviroment source (or $PATH).
4. Install `godep` package: `go get github.com/tools/godep`. [Find out more](https://github.com/tools/godep)
5. Install project dependencies: `godep get`
6. Run package: `go run main.go` or `bee run`
