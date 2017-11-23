BIN := dbloader
MAIN := cmd/dbloader/main.go

mac: $(MAIN) vet test
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BIN) $(MAIN) 

linux: $(MAIN)
	GOOS=linux GOARCH=amd64 go build -o bin/$(BIN) $(MAIN)

windows: $(MAIN)
	GOOS=windows GOARCH=amd64 go build -o bin/$(BIN) $(MAIN)

vet: $(MAIN)
	govet -all ./pkg/db ./pkg/bd ./pkg/parsers ./cmd/dbloader

test: $(MAIN)
	go test ./...