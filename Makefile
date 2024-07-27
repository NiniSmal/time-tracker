db:
	docker run --name ps -p 8010:5432 -e POSTGRES_PASSWORD=dev -d --restart=always --network app postgres:15.6


migration_up:
	goose -dir ./migrations postgres "postgres://postgres:dev@localhost:8010/postgres?sslmode=disable" up


migration_down:
	goose -dir ./migrations postgres "postgres://postgres:dev@localhost:8010/postgres?sslmode=disable" down

migration_reset:
	goose -dir ./migrations postgres "postgres://postgres:dev@localhost:8010/postgres?sslmode=disable" reset &&\
	goose -dir ./migrations postgres "postgres://postgres:dev@localhost:8010/postgres?sslmode=disable" up

