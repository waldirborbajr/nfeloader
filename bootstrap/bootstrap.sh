#!/bin/sh

# Stop and remove container
docker rm -f nfeloader

# Fetch for update
docker pull waldirborbajr/nfeloader:latest

# Start
docker run -it \
--name nfeloader \
-e MAIL_SERVER="mail.XPTOinformatica.com.br:993" \
-e MAIL_USR="waldir@XPTOinformatica.com.br" \
-e MAIL_PWD="#Senha_!@#_Mudar*" \
-e DATABASE_HOST="192.168.0.4:3306" \
-e DATABASE_USR="root" \
-e DATABASE_PWD="@----@" \
-e DATABASE_NAME="nfeloader" \
-e TIME_SCHEDULE="2m" \
-e CONTAINER="true" \
-v $(pwd)/logs/:/app/logs/:rw \
-v $(pwd)/xmls/:/app/xmls/:rw \
--restart unless-stopped \
--add-host mail.XPTOinformatica.com.br:192.168.0.25 \
-d waldirborbajr/nfeloader:latest \
/bin/sh