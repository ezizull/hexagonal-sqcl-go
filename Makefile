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
	sudo MYSQL_HOST=172.17.0.1 MYSQL_PORT=3306 MYSQL_USER=root MYSQL_PASSWORD=root MYSQL_DBNAME=hexagonal-sqlc docker-compose up --build

docker-update:
	sudo docker tag ezizull/skyshi-gethired:latest ezizull/skyshi-gethired:$(tag)
	sudo docker tag ezizull/skyshi-gethired:latest ezizull/skyshi-gethired:latest

docker-push:
	sudo docker push ezizull/skyshi-gethired:$(tag)
	sudo docker push ezizull/skyshi-gethired:latest

docker-delete:
	sudo docker rmi -f ezizull/skyshi-gethired:$(tag)

docker-run:
	sudo docker run -e MYSQL_HOST=172.17.0.1 -e MYSQL_PORT=3306 -e MYSQL_USER=root -e MYSQL_PASSWORD=root -e MYSQL_DBNAME=hexagonal-sqlc -p 3030:3030 ezizull/skyshi-gethired:$(tag)
