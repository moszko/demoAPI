# General

This a demo application. It exposes two endpoints:
1. GET /api/v1/trademarks/word/{name} - get data of an existing word trademark or null
2. GET /api/v1/trademarks/word/similar/{name} - get an array of similiar word trademarks

## Dependencies

You should have bash with wget and PHP7+ installed to run script.sh.

## Local setup

1. Go to the top directory of the repository and run script.sh. Make sure you have enough disk space, because during the execution of script.sh disk usage peaks around 8GB. Then all unnecesary data is flushed.
2. Go to docker/db/ and make a copy of pg_pass.example and rename it to pg_pass.
3. Change the password inside pg_pass to your needs.
4. Run `docker-compose up`

The application is running on [http:/localhost:7777/](http://localhost:7777/). Now you can visit the endpoints mentioned above. There is also admin panel available at [http:/localhost:8080/](http://localhost:8080/). You can log into it using `postgres` as a user, `postgres` as database name and the password you've chosen.
