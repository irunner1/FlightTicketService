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

* Tests
* Refactor db scheme
* Methods for adding data
* Methods for deleting data
* ~~Authorization~~
* ~~Database~~
* ~~logger~~

---

## db schema

booking_flights - table of booked flight. New record shows up when passenger books flight
(passess passenger `passenger table` data for specific flight `flights table`)

flights - table of flight in airports

passengers - table of passengers, users of ticket service
(user register in service)

## Service startup

You need to prepare .env file. Example showed in `.env.showcase`

```cmd
make run
```

## Lint

```cmd
make lint
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
