FROM alpine:3.13

ENV PATH=/app/:$PATH

LABEL maintainer="Waldir Borba Junior <wborbajr@gmail.com>" \
  version="v0.4.2-2022" \
  description="NFe Loader | waldirborbajr/nfeloader:latest"

ENV LANG=en_US.UTF-8 \
  LANGUAGE=en_US.UTF-8

WORKDIR /app

COPY nfeloader /app
# COPY nfeloader-api /app

ENTRYPOINT ["/app/nfeloader"]

CMD ["/bin/sh"]

