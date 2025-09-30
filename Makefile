.PHONY: migrate-up migrate-down

migrate-up:
	go run cmd/migrations/main.go migrate --config cmd/migrations/tern.conf --migrations migrations

migrate-down:
	go run cmd/migrations/main.go migrate --config cmd/migrations/tern.conf --migrations migrations --destination 0
