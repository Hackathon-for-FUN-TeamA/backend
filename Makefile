run:
	go run cmd/server/main.go cmd/server/controller.go
start-db:
	docker compose -f ./Docker/docker-compose.yml up asayake-db -d

start-mock:
	docker compose -f ./Docker/docker-compose.yml up asayake-mock -d
