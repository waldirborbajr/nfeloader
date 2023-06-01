#!/usr/bin/env bash

# login
# docker login ghcr.io
# curl -v -u waldirborbajr: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 https://ghcr.io/v2/

source .env

VERSION="0.4.4"

# Stop and remove container
docker rm -f nfeloaderqa

# Fetch for update
docker pull ghcr.io/waldirborbajr/nfeloader:${VERSION}

# Start
docker run -d -i -t -p 9396:9396 --rm \
-e MAIL_SERVER=${MAIL_SERVER} \ 
-e MAIL_USR=${MAIL_USR} \
-e MAIL_PWD=${MAIL_PWD} \
-e DATABASE_HOST=${DATABASE_HOST} \
-e DATABASE_USR=${DATABASE_USR} \
-e DATABASE_PWD=${DATABASE_PWD} \ 
-e DATABASE_NAME=${DATABASE_NAME} \
-e TIME_SCHEDULE=${TIME_SCHEDULE} \
-e CONTAINER=${CONTAINER} \ 
--name nfeloaderqa ghcr.io/waldirborbajr/nfeloader:${VERSION} \
sh

# Show me the log
docker logs -f nfeloaderqa
