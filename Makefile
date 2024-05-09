build-payment:
	go build -o cmd/payment/main cmd/payment/main.go

build-auth:
	go build -o cmd/auth/main cmd/auth/main.go

docker-payment:
	docker build -t auth -f ./cmd/auth/Dockerfile .

docker-auth:
	docker build -t payment -f ./cmd/payment/Dockerfile .

compose:
	docker compose up

payment: build-payment
	go run cmd/payment/main.go

auth: build-auth
	go run cmd/auth/main.go

mock: 
	mockery --all
	
test:
	go test ./...