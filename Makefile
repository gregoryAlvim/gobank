.PHONY: migrate-up migrate-down

migrate-up:
	tern migrate --config tern.conf --migrations migrations

migrate-down:
	tern migrate --config tern.conf --migrations migrations --destination 0
