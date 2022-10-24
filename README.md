# NFeLoader

<p>
<img alt="NF-e Logo" src="https://github.com/waldirborbajr/nfeloader/blob/main/assets/nfe.png" width="120", heigth="120"/>
</p>

## B+ Tech·​nol·​o·​gy

en: NF-e Loader {Load XML NF-e file to database}

ptBR: NF-e Loader {Carrega o arquivo XML de NF-e para o banco de dados }

Language: go 1.19

## Usage
```
docker run -it --rm \
--name nfeloader \
-e MAIL_SERVER="<imap_server>" \
-e MAIL_USR="<email>" \
-e MAIL_PWD="<password>" \
-e DATABASE_HOST="<host>" \
-e DATABASE_USR="<user>" \
-e DATABASE_PWD="<password>" \
-e DATABASE_NAME="<bdname>" \
-v $(pwd)/xmls/:/app/xmls/:rw \
--add-host mail.xxxxx.xxx.xx:192.168.0.25 \
--restart unless-stopped \
-d waldirborbajr/nfeloader:latest \
/bin/sh
```
