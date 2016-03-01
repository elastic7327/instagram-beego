# instagram-beego

## Requirements

- Postgres : http://www.postgresql.org/
- Go : https://golang.org/

## How to run

1. Create Amazone S3 bucket & account
2. Copy `.env.sample` to `.env` and fill S3 info into `.env` file
3. Add `$GOPATH/bin` to enviroment source (or $PATH).
4. Install `godep` package: `go get github.com/tools/godep`. [Find out more](https://github.com/tools/godep)
5. Install project dependencies: `godep get`
6. Create Postgres database (or you can use `setup/create-postgres-db.sh`)
7. Copy `conf/app.conf.sample` to `conf/app.conf` & fill in app configuration.
8. Create database tables: `go run setup/database.go`
9. [Optional] Seed: `go run setup/seed.go`
10. Run package: `go run main.go` or `bee run`

## Other

- AWS SDK for Go:
  - https://github.com/aws/aws-sdk-go
  - http://docs.aws.amazon.com/sdk-for-go/api/
