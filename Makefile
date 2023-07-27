install:
	@echo 'Install dependencies'
	go mod tidy

build:
	@echo 'Build app'
	go build cmd/main.go

run:
	@echo 'Run app'
	go run cmd/main.go