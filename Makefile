# Docker container setup
DOCKER_USER=waldirborbajr
PROJECT_NAME=nfeloader
DOCKER_TAG=latest
CONTAINER=nfeloader

#
# Production
# ------------------------------------------------------------------------------

prodbuild:
	@echo "\nStarting build...\n"
	docker build -t ${DOCKER_USER}/${PROJECT_NAME}:${DOCKER_TAG} -f ./docker/Dockerfile .

prodrmimage:
	@echo "\nRemoving image...\n"
	docker rmi -f ${DOCKER_USER}/${PROJECT_NAME}

prodrmcontainer:
	@echo "\nRemoving container...\n"
	docker container rm -f ${PROJECT_NAME}

prodrun:
	@echo "\nStarting container...\n"
	docker run -it \
		--name ${PROJECT_NAME} \
		-e MAIL_SERVER="mail.xxx.com.br:993" \
		-e MAIL_USR="waldir@xxx.com.br" \
		-e MAIL_PWD="#Senha_123_Mudar*" \
		-e DATABASE_HOST="192.168.0.4" \
		-e DATABASE_USR="root" \
		-e DATABASE_PWD="@senha" \
		-e DATABASE_NAME=${PROJECT_NAME} \
		--restart unless-stopped \
		-d ${DOCKER_USER}/${PROJECT_NAME}:${DOCKER_TAG}\
		/bin/sh

prodexec:
	@echo "\nExecuting production mode...\n"
	docker exec -it ${PROJECT_NAME} /bin/sh

proddeploy:
	@echo "\nStarting delpoy...\n"
	docker push ${DOCKER_USER}/${PROJECT_NAME}${DOCKER_TAG}:

#
# Development
# ------------------------------------------------------------------------------

devbuild:
	@echo "\nStarting build...\n"
	docker-compose -f docker-compose.yaml -f docker/mysql-compose.yaml build

devrebuild:
	@echo "\nForcing Rebuild...\n"
	docker-compose -f docker-compose.yaml -f docker/mysql-compose.yaml build --no-cache --force-rm --pull

devstart:
	@echo "\nStarting container...\n"
	docker-compose -f docker-compose.yaml -f docker/mysql-compose.yaml up -d --force-recreate

devstop:
	@echo "\nStoping container...\n"
	docker-compose -f docker-compose.yaml -f docker/mysql-compose.yaml stop -t 0

devdown:
	@echo "\nStoping container...\n"
	docker-compose -f docker-compose.yaml -f docker/mysql-compose.yaml down -v --remove-orphans --rmi all -t 0

devtop:
	docker-compose -f docker-compose.yaml -f docker/mysql-compose.yaml top

devps:
	docker-compose -f docker-compose.yaml -f docker/mysql-compose.yaml ps 

devlog:
	docker-compose -f docker-compose.yaml -f docker/mysql-compose.yaml logs -f  

devevents:
	docker-compose -f docker-compose.yaml -f docker/mysql-compose.yaml events

devpause:
	docker-compose -f docker-compose.yaml -f docker/mysql-compose.yaml pause

devunpause:
	docker-compose -f docker-compose.yaml -f docker/mysql-compose.yaml unpause

#
# Docker CMD
# ------------------------------------------------------------------------------

exec:
	@echo "\nEntering container...\n"
	docker exec -ti ${CONTAINER} sh

dang:
	@echo "\nStarting dangling removal\n"
	docker rmi $$(docker images -q -f dangling=true)

prune:
	docker system prune -af --volumes 

remove:
	docker rm $$(docker ps -a -q) -f

#
# Build standalone
# ------------------------------------------------------------------------------
binary:
	@echo "\nBuild standalone version...\n"
	CGO_ENABLED=0 GOOS=linux go build -a -trimpath -installsuffix cgo -ldflags "-s -w" -v -o ./release/${PROJECT_NAME}-linux main.go
	CGO_ENABLED=0 GOOS=darwin go build -a -trimpath -installsuffix cgo -ldflags "-s -w" -v -o ./release/${PROJECT_NAME}-macos main.go
	CGO_ENABLED=0 GOOS=windows go build -a -trimpath -installsuffix cgo -ldflags "-s -w" -v -o ./release/${PROJECT_NAME}-windows.exe main.go

#
# Help
# ------------------------------------------------------------------------------

help:
	@echo ''
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo '  build    	build docker --image-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  rebuild  	rebuild docker --image-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  test     	test docker --container-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  service   	run as service --container-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  login   	run as service and login --container-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  clean    	remove docker --image-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  prune    	shortcut for docker system prune -af. Cleanup inactive containers and cache.'
	@echo '  shell      run docker --container-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo ''
	@echo 'Extra arguments:'
	@echo 'cmd=:	make cmd="whoami"'
	@echo '# user= and uid= allows to override current user. Might require additional privileges.'
	@echo 'user=:	make shell user=root (no need to set uid=0)'
	@echo 'uid=:	make shell user=dummy uid=4000 (defaults to 0 if user= set)'

