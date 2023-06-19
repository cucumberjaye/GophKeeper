up:
	docker-compose up -d

down:
	docker-compose down

client-run:
	go run cmd/keeperclient/main.go

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -o /build/linux/gophkeeper cmd/gophkeeper/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o /build/windows/gophkeeper cmd/gophkeeper/main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o /build/mac/gophkeeper cmd/gophkeeper/main.go