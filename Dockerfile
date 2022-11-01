FROM alpine:3.13

ENV PATH=/app/:$PATH

ENV LANG=en_US.UTF-8 \
  LANGUAGE=en_US.UTF-8

WORKDIR /app

COPY nfeloader /app
COPY nfeloader-api /app

ENTRYPOINT ["/app/nfeloader"]
CMD ["/bin/sh"]

