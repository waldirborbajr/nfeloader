# ---------------
# -- 1st Stage

FROM golang:1.19-alpine AS build

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -trimpath -o ./bin/nfeloader cmd/cli/main.go

# ---------------
# -- 2nd Stage

FROM alpine:3.13

ENV PATH=/app/:$PATH

LABEL maintainer="Waldir Borba Junior <wborbajr@gmail.com>" \
  version="v0.4.2-2022" \
  description="NFe Loader | waldirborbajr/nfeloader:latest"

ENV LANG=en_US.UTF-8 \
  LANGUAGE=en_US.UTF-8

RUN apk add --update --no-cache \
  tzdata \
  htop \
  apk-cron \
  && cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime \
  && echo "America/Sao_Paulo" > /etc/timezone

RUN adduser -S -D -H -h /app nfe

USER nfe

WORKDIR /app

COPY --from=build /go/src/app/run.sh run.sh
COPY --from=build /go/src/app/bin/nfeloader nfeloader

ENTRYPOINT ["/app/nfeloader"]

CMD ["/bin/sh"]

