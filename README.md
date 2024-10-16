# Flight ticket service

Flight ticket service is a web server for searching and purchasing airline tickets.

For now it`s just an api

## Tech Stack

* Go
* Docker
* JWT tocken for authorization (in progress)
* Postgres for storage
* RabbitMQ as message brocker (in progress)
* OpenAPI for documenting API
* Postman for checking API
* UUID for ids
* OpenAPI Swagger

## Documentation

You could open /swagger/ page to check documentation and all API methods

![Swagger view](pictures/swagger.png?raw=true "Documentation page for an API")

## TODO

* Methods for adding data
* Methods for deleting data
* Authorization
* Database
* ~~logger~~

---

## Service startup

```cmd
% make run
```

## Example api`s

## Docker

```bash
# Сборка
docker build -t flightticketservice .

# Запуск
docker run -it -p 8080:8080 --env-file=.env flightticketservice

```

You need to prepare .env file with variables:

```cmd
HOST=0.0.0.0
PORT=80
DB_USER=me
DB_PASS=123123

```

* HOST - application host
* PORT - application port
* DB_USER - database user name
* DB_PASS - database user password

---
