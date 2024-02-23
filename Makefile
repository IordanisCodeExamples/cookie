build:
	go build cmd/cookie/main.go
compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm cmd/cookie/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 cmd/cookie/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 cmd/cookie/main.go
generate-mock:
	go generate -x ./...
get-generator:
	go install github.com/golang/mock/mockgen
test:
	go test ./...
test-conver:
	go test -cover ./...
lint:
	golangci-lint run --print-issued-lines=false --out-format code-climate:gl-code-quality-report.json,line-number