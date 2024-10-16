# Flight ticket service

Flight ticket service is a web server for searching and purchasing airline tickets.

For now it`s just an api

## Tech Stack

* Go
* Docker
* JWT tocken for authorization (in progress)
* Postgres for storage (in progress)
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
docker build -t airticketfinder .

# Запуск
docker run -it -p 8080:8080 --env-file=.env airticketfinder

```

Для запуска необходимо подготовить файл переменных окружения (.env) c примерным содержанием:

```cmd
HOST=0.0.0.0
PORT=8080
```

* HOST - хост запускаемой программы
* PORT - порт запускаемой программы

---
