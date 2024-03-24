include .env

# build:
# 	@go build -o bin/gobank
# run: build
# 	@./bin/gobank

# test:
# 	@go test -v ./..

docker: 
	docker-compose up -d
	
migrate:
	migrate -source file://db/migrations \
			-database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE} up

drop_database:
	migrate -source file://db/migrations \
			-database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE} drop

down_database:
	migrate -source file://db/migrations \
			-database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE} down

# Arguments: version			
force_migration:
	migrate -source file://db/migrations \
			-database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE} force $(version)

# Example: make create_migration name=create_users_table
# Arguments: migration_name	
create_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

test:
	go test ./...