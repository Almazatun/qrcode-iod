URL?=https://github.com/Almazatun/qrcode-iod

install:
	@echo 'Install dependencies'
	go get ./...

build_web:
	@echo 'Build web app'
	go build -o cmd/web/qrcode_web cmd/web/main.go

build_cli:
	@echo 'Build cli tool'
	go build -o cmd/cli/qrcode_cli cmd/cli/main.go

run_web:
	@echo 'Run web app'
	go run cmd/web/main.go