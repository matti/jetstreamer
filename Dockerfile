FROM nats:2.9.1-alpine3.16

RUN apk add --no-cache bash curl

WORKDIR /app
COPY app .
ENTRYPOINT [ "/app/entrypoint.sh" ]
