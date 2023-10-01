include .env

stop_containers:
	@echo "Stopping other docker container"
	if [ $$(docker ps -q) ]; then \
		echo "found and stopped containers"; \
		docker stop $$(docker ps -q); \
	else \
		echo "no containers running..."; \
	fi


create_container:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5433:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:12-alpine

create_db:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}
# have to run psql -h 127.0.0.1 -p 5433 -U postgres -d my_flick_list_db to connect                                                                  ─╯
start_container:
	docker start ${DB_DOCKER_CONTAINER}

create_migrations:
	sqlx migrate add -r init