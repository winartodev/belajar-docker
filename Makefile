PHONY: images, containers, active_containers, pull

# list All Images
images: 
	docker images

# list All Containers
containers: 
	docker ps -a

# list all active containers
active_containers:
	docker ps

# remove images 
image_remove: 
	docker rmi nginx:stable-alpine

# pull images from docker hub
pull:
	docker pull nginx:stable-alpine

# run new container
run: 
	docker run --name nginx -dp 8080:80 nginx:stable-alpine 

# stop container
stop:
	docker stop nginx

# start container 
start: 
	docker start nginx

# remove container 
drop: 
	docker rm nginx

# force remove the container 
drop_force: 
	docker rm -f nginx

# remove all container component 
drop_all: 
	docker rm -f -v nginx

# see all logs
logs: 
	docker logs nginx
