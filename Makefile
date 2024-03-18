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

test:
	go test ./...