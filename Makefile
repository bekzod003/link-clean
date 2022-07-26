.PHONY:
.SILENT:

# Environment variables for project
ENV := $(PWD)/.env
include $(ENV)


# run all unit tests in project
test:
	go test --short --cover ./...

# migrations up
migrate-up:
	migrate -path ./migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable' up

# migrations down
migrate-down:
	migrate -path ./migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable' down
