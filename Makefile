.PHONY: default

default:
	GOOS=darwin GOARCH=arm64  go build -o dist/mdcheck-darwin-arm64 -ldflags "-w -s" main.go
	GOOS=linux GOARCH=arm64   go build -o dist/mdcheck-linux-arm64 -ldflags "-w -s" main.go
	GOOS=linux GOARCH=amd64   go build -o dist/mdcheck-linux-amd64 -ldflags "-w -s" main.go
	GOOS=windows GOARCH=amd64 go build -o dist/mdcheck-windows-amd64.exe -ldflags "-w -s" main.go
	GOOS=windows GOARCH=arm64 go build -o dist/mdcheck-windows-arm64.exe -ldflags "-w -s" main.go

lint:
	golangci-lint run --fix --allow-parallel-runners;
