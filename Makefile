run:
	go run main.go

dbmigrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/todo?sslmode=disable" -verbose up

dbmigratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/todo?sslmode=disable" -verbose down