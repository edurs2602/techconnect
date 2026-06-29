DB_URL=postgres://social:social@localhost:5433/social?sslmode=disable

migrate:
	migrate -path migrations -database "$(DB_URL)" up

sqlc:
	sqlc generate

test:
	go test ./... -v