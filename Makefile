IMAGE_TAG="tes"
NETWORK="jakpat-test-2_service"

docker-compose:
	@docker-compose up -d

migrate-db:
	@migrate -path ./schema -database 'postgres://user:password@localhost:5432/postgres?sslmode=disable' up

build:
	@go mod tidy
	@go build -o ./.bin/app ./cmd/main.go

docker-build:
	@docker build -t $(IMAGE_TAG) .

docker-run:
	@docker run --network=$(NETWORK) -p 8880:8880 $(IMAGE_TAG)

run: 
	@go run ./cmd/main.go
