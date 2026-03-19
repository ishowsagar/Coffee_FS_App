# Importing env file for using those declared var
include .env

# @Commands ( shorthand for bunch of cmds)

# Each title is treated as a command to run commands attached to it below
stop_containers :
	@echo "Stopping other docker containers..."
	if [ $$(docker ps -q) ]; then \
		echo "found targetted containers & succefully stopped them"; \
		docker stop $$(docker ps -q); \
	else \
		echo "there are no active containers running!"; \
	fi	

create_container :
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:12-alpine

create_db :
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

start_container :
	docker start ${DB_DOCKER_CONTAINER}

#  custom make cmd to stop container
stop_container :
	docker stop ${DB_DOCKER_CONTAINER}

create_migrations :
	sqlx migrate add -r init

migup :
	sqlx migrate run --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

migdown :
	sqlx migrate revert --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

build :
	if [ -f  "${BINARY}" ]; then \
		rm ${BINARY}; \
		echo "Deleted ${BINARY}"; \
	fi
	@echo "Building binary⚡⚡..."
# 	go build -o name sourceLocation
	go build -o ${BINARY} cmd/server/*.go

# follow up cmds example1:ex2 --> it means it would run ex2 commands first when running example1 title commands
run : build 
	./${BINARY}

stop : 
	@echo "Stopping backend server🔃🔃..."
	@-pkill -SIGTERM -f "./${BINARY}"
	@echo "Stopped server❌❌..."

# custom commands for git workflow
add :
	git add .
commit : add
	git commit -m "pushes newer changes"
push : commit
	git push 
