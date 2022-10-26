FROM golang:1.19-bullseye as vldemo-server

WORKDIR /usr/src/app

COPY --chown=www-data:www-data app .

RUN go build -o /usr/local/bin/app

CMD ["app"]