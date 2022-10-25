#!/bin/sh

# Stop and remove container
docker rm -f nfeloaderdev

# Fetch for update
docker pull waldirborbajr/nfeloader:latest

# Start
docker run -it \
--name nfeloaderdev \
-e MAIL_SERVER="mail.XPTO.com.br:993" \
-e MAIL_USR="waldir@XPTO.com.br" \
-e MAIL_PWD="#Senha@Mudar*" \
-e DATABASE_HOST="000.000.0.000:3306" \
-e DATABASE_USR="root" \
-e DATABASE_PWD="@senha#Mudar" \
-e DATABASE_NAME="NFEIMPORT" \
-e TIME_SCHEDULE="2m" \
-e CONTAINER="true" \
-v $(pwd)/logs/:/app/logs/:rw \
-v $(pwd)/xmls/:/app/xmls/:rw \
--restart unless-stopped \
--add-host mail.XPTO.com.br:000.000.0.000 \
-d waldirborbajr/nfeloader:latest \
/bin/sh

# Show me the log
docker logs -f nfeloaderdev
