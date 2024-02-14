build:
	go build -o bin/main main.go

run:
	go run main.go

mock: 
	mockery --all
	
test:
	go test ./...