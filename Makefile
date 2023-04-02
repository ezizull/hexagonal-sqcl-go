# Database
migrate_schema:
	~/go/bin/migrate create -ext sql -dir infrastructure/repository/postgres/schema -seq $(seq)

migrate_up:
	~/go/bin/migrate -database 'postgres://root:root@localhost:5432/$(db)?sslmode=disable' -path infrastructure/repository/postgres/schema up

migrate_down:
	~/go/bin/migrate -database 'postgres://root:root@localhost:5432/$(db)?sslmode=disable' -path infrastructure/repository/postgres/schema down


# Docker
docker-build:
	sudo docker-compose down
	sudo POSTGRES_HOST=172.17.0.1 POSTGRES_PORT=5432 POSTGRES_USER=root POSTGRES_PASSWORD=root POSTGRES_DBNAME=hexagonal_sqlc docker-compose up --build

docker-update:
	sudo docker tag ezizull/hexagonal-sqlc:latest ezizull/hexagonal-sqlc:$(tag)
	sudo docker tag ezizull/hexagonal-sqlc:latest ezizull/hexagonal-sqlc:latest

docker-push:
	sudo docker push ezizull/hexagonal-sqlc:$(tag)
	sudo docker push ezizull/hexagonal-sqlc:latest

docker-delete:
	sudo docker rmi -f ezizull/hexagonal-sqlc:$(tag)

docker-run:
	sudo docker run -e POSTGRES_HOST=172.17.0.1 -e POSTGRES_PORT=5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DBNAME=hexagonal_sqlc -p 3030:3030 ezizull/hexagonal-sqlc:$(tag)
