build-payment:
	go build -o cmd/payment/main cmd/payment/main.go

build-auth:
	go build -o cmd/auth/main cmd/auth/main.go

payment: build-payment
	go run cmd/payment/main.go

auth: build-auth
	go run cmd/auth/main.go

mock: 
	mockery --all
	
test:
	go test ./...