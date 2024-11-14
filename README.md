# Flight ticket service

Flight ticket service is a web server for searching and purchasing airline tickets.

For now it`s just an api

## Tech Stack

* Golang
* Docker
* JWT tocken for authorization
* Postgres for storage
* Postman for checking API
* OpenAPI Swagger

## Documentation

You could open /swagger/ page to check documentation and all API methods

![Swagger view](pictures/swagger.png?raw=true "Documentation page for an API")

## TODO

* Methods for adding data
* Methods for deleting data
* ~~Authorization~~
* ~~Database~~
* ~~logger~~

---

## Service startup

```cmd
% make run
```

## Example api`s

## Docker

```bash
# Build
docker build -t flightticketservice .

# Launch
docker run -it -p 8080:8080 --env-file=.env flightticketservice

```

## Docker-compose

```bash
# Build and launch postgres db and flight service
docker-compose up --build

```

You need to prepare .env file with variables:

```cmd
HOST=0.0.0.0
PORT=3000
DB_HOST=0.0.0.0
DB_USER=user
DB_PASS=password
DB_NAME=postgres
JWT_SECRET=secret

```

* HOST - application host
* PORT - application port
* DB_HOST - database host
* DB_USER - database user name
* DB_PASS - database user password
* DB_NAME - database name
* JWT_SECRET - secret for JWT

---
