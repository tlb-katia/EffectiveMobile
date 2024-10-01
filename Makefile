.PHONY: create_migration check_port free_port


serv = effective_mobile
db = effectiveMode_db

all: docker-compose-api ##@APP application in docker container


docker-compose-api:  ##@APP runs application in docker container
	docker-compose up --build $(serv)

clean-data:  ##@DB clean a database saved data
	rm -rf internal/storage/postgres/pgdata
	rm -rf internal/storage/redis/data

docker-stop-api: ##@SERVER stops containers
	docker stop $(db)
	docker stop $(serv)
	docker stop $(cache)

docker-clean-api: docker-stop-api ##@SERVER delete server, database and cache containers
	docker rm $(db)
	docker rm $(serv)
	docker rm $(cache)

server-logs: ##@SERVER show logs from server container
	docker logs $(serv)

database-logs:  ##@DB show logs from database container
	docker logs $(db)

permission:
	sudo chmod -R 777 /internal/storage/postgres/pgdata

all-logs: database-logs server-logs ##@APP show logs from server and db containers together

#create_migration:
#	migrate create -ext=sql -dir=internal/database/migrations -seq init

check_port:
	sudo lsof -i -P -n | grep LISTEN | grep 5432

free_port:
	sudo systemctl stop postgresql
