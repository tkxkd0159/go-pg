```bash
go mod init github.com/tkxkd0159/go-pg/go-nomad

go get -v -d <pkg> # use -u for update
go mod tidy
go mod vendor

go run .
go build ./cmd/nomad/
go test
```


# Structure

## /api
OpenAPI/Swagger specs, JSON schema files, protocol definition files.

https://github.com/kubernetes/kubernetes/tree/master/api

## /pkg
My public lib -> Written by me