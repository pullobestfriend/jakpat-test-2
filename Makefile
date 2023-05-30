build:
	go mod tidy && go build -o ./.bin/app ./cmd/main.go

run: 
	go run cmd/main.go
