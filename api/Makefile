PHONY: build, run, setup, start, stop, remove_container, remove_image, force_remove, logs

# build image
build:
	docker build -t go-api .

# run container
run:
	docker run --name go-api -it -p 8081:8081 go-api:latest

# to simplify build and run 
setup: 
	make build && make run

# start container
start:
	docker start go-api

# stop container
stop:
	docker stop go-api

# remove container
remove_container:
	docker container rm -f go-api

# remove image
remove_image:
	docker rmi go-api

# remove container and image
force_remove:
	make remove_container && make remove_image

# see all logs
logs: 
	docker logs -t go-app

# compose up
compose_up: 
	docker compose up -d

# compose down
compose_down: 
	docker compose down
