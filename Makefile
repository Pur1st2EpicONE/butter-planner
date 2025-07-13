TARGET=cmd/main.go

all: db db-load migrate-up run

run:
	@go run $(TARGET)

db:
	@docker run --name=butter_db -e POSTGRES_PASSWORD=qwerty -p 5432:5432 -d postgres

db-load:
	@until docker exec butter_db pg_isready -U postgres > /dev/null 2>&1; do sleep 0.5; done

clean:
	@docker stop butter_db
	@docker rm butter_db

migrate-up:
	@migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	@migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down

.PHONY: all run clean migrate-up migrate-down
