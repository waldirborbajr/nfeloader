# FROM golang:1.19.2-alpine@sha256:e4dcdac3ed37d8c2b3b8bcef2909573b2ad9c2ab53ba53c608909e8b89ccee36
FROM alpine:3.13

ENV PATH=/app/:$PATH

ENV LANG=en_US.UTF-8 \
  LANGUAGE=en_US.UTF-8

WORKDIR /app

COPY nfeloader /app

ENTRYPOINT ["/app/nfeloader"]

CMD ["/bin/sh"]

