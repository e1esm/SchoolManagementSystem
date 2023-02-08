build:
	go build cmd/main.go
	go build internal/app/*.go
	go build internal/utils/*.go
	go build internal/models/*.go

run:
	go run cmd/main.go
	go run internal/app/*.go internal/utils/*.go internal/models/*.go