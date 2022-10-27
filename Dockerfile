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

COPY nfeloader /app

ENTRYPOINT ["/app/nfeloader"]

CMD ["/bin/sh"]

