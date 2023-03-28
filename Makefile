migrate_schema:
	~/go/bin/migrate create -ext sql -dir infrastructure/repository/postgres/schema -seq $(seq)

migrate_up:
	~/go/bin/migrate -database 'postgres://root:root@localhost:5432/$(db)?sslmode=disable' -path infrastructure/repository/postgres/schema up

migrate_down:
	~/go/bin/migrate -database 'postgres://root:root@localhost:5432/$(db)?sslmode=disable' -path infrastructure/repository/postgres/schema down

run:
	go run main.go

