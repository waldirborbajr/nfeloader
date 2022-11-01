# FROM golang:1.19.2-alpine@sha256:e4dcdac3ed37d8c2b3b8bcef2909573b2ad9c2ab53ba53c608909e8b89ccee36
FROM alpine:3.13

ENV PATH=/app/:$PATH

ENV LANG=en_US.UTF-8 \
  LANGUAGE=en_US.UTF-8

# RUN apk add --update --no-cache \
#   tzdata \
#   htop \
#   && cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime \
#   && echo "America/Sao_Paulo" > /etc/timezone

WORKDIR /app

COPY nfeloader /app

COPY nfeloader-api /app

# COPY nfestart.sh /app
# RUN chmod +x nfestart.sh

EXPOSE 9191

# Working
ENTRYPOINT ["/app/nfeloader"]
CMD ["/bin/sh"]

# ENTRYPOINT ["/app/nfestart.sh"]

# CMD ["/app/nfeloader", "/app/nfeloader-api"]
# CMD ["/bin/sh"]

# CMD ["/bin/sh","-c","/app/nfestart.sh"]
