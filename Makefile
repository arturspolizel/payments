build payment:
	go build -o cmd/payment/main cmd/payment/main.go

build auth:
	go build -o cmd/auth/main cmd/auth/main.go

run payment: build payment
	go run cmd/payment/main.go

run auth: build auth
	go run cmd/auth/main.go

mock: 
	mockery --all
	
test:
	go test ./...