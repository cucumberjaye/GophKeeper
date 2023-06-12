up:
	docker-compose up -d

down:
	docker-compose down

client-run:
	go run cmd/keeperclient/main.go