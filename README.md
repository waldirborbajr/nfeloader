# NFeLoader


[![Typo Check](https://github.com/waldirborbajr/nfeloader/actions/workflows/typo-check.yaml/badge.svg)](https://github.com/waldirborbajr/nfeloader/actions/workflows/typo-check.yaml)
[![Lint](https://github.com/waldirborbajr/nfeloader/actions/workflows/lint.yaml/badge.svg)](https://github.com/waldirborbajr/nfeloader/actions/workflows/lint.yaml)
[![CodeQL](https://github.com/waldirborbajr/nfeloader/actions/workflows/codeql.yaml/badge.svg)](https://github.com/waldirborbajr/nfeloader/actions/workflows/codeql.yaml)
[![Build & Test](https://github.com/waldirborbajr/nfeloader/actions/workflows/build-test.yaml/badge.svg)](https://github.com/waldirborbajr/nfeloader/actions/workflows/build-test.yaml)


<p>
<img alt="NF-e Logo" src="https://github.com/waldirborbajr/nfeloader/blob/main/assets/nfe.png" width="120", height="120"/>
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


### Loading Test

```sh k6 run --vus 10 --duration 30s script.js```

```

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: script.js
     output: -

  scenarios: (100.00%) 1 scenario, 10 max VUs, 1m0s max duration (incl. graceful stop):
           * default: 10 looping VUs for 30s (gracefulStop: 30s)


running (0m30.1s), 00/10 VUs, 300 complete and 0 interrupted iterations
default ✓ [======================================] 10 VUs  30s

     data_received..................: 43 kB 1.4 kB/s
     data_sent......................: 24 kB 797 B/s
     http_req_blocked...............: avg=61.55µs min=4µs   med=8µs    max=1.68ms p(90)=14µs    p(95)=31.5µs
     http_req_connecting............: avg=16.25µs min=0s    med=0s     max=639µs  p(90)=0s      p(95)=0s
     http_req_duration..............: avg=1.28ms  min=400µs med=1.27ms max=2.56ms p(90)=1.86ms  p(95)=2ms
       { expected_response:true }...: avg=1.28ms  min=400µs med=1.27ms max=2.56ms p(90)=1.86ms  p(95)=2ms
     http_req_failed................: 0.00% ✓ 0        ✗ 300
     http_req_receiving.............: avg=98.65µs min=33µs  med=69.5µs max=860µs  p(90)=148.7µs p(95)=239.55µs
     http_req_sending...............: avg=58.45µs min=16µs  med=31µs   max=1.02ms p(90)=116.8µs p(95)=202.45µs
     http_req_tls_handshaking.......: avg=0s      min=0s    med=0s     max=0s     p(90)=0s      p(95)=0s
     http_req_waiting...............: avg=1.13ms  min=333µs med=1.11ms max=2.4ms  p(90)=1.71ms  p(95)=1.89ms
     http_reqs......................: 300   9.963639/s
     iteration_duration.............: avg=1s      min=1s    med=1s     max=1s     p(90)=1s      p(95)=1s
     iterations.....................: 300   9.963639/s
     vus............................: 10    min=10     max=10
     vus_max........................: 10    min=10     max=10
```


