FROM alpine:3.13

ENV PATH=/app/:$PATH

ENV LANG=en_US.UTF-8 \
  LANGUAGE=en_US.UTF-8

RUN apk add --update --no-cache \
  tzdata \
  htop \
  && cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime \
  && echo "America/Sao_Paulo" > /etc/timezone

WORKDIR /app

COPY nfeloader /app
COPY nfeloader-api /app

EXPOSE 9693

# ENTRYPOINT ["/app/nfeloader-api"]
CMD ["/bin/sh", "-c", "/app/nfeloader && /app/nfeloader-api"]
# CMD ["/bin/sh"]

# CMD ["start_cli_api.sh"]
#
