# Monolith to Microservice

## Prerequisites

You need Docker and docker-compose installed.

Everything is running in Docker container, so you don't need golang either any other lib.

## Running

Just run

```bash
 make up
 ```

It will build Docker image and run monolith and microservices version.

## Services addresses

Monolith: http://localhost:8090/

Orders microservice: http://localhost:8070/

Shop microservice: http://localhost:8071/

Payments microservice: no public API (you can export ports in docker-compose.yml if you need)

For available methods, please check interfaces layer in source code and tests tests/acceptance_test.go.

## Running tests

First of all you must run services

```bash
make up
```

Then you can run all tests by using in another terminal:

```bash
make docker-test
```
If you want to test only monolith version:

```bash
make docker-test-monolith
```

or microservices:

```bash
make docker-test-microservices
```