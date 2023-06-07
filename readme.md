# Golang DDD Boilerplate
---
This is a simple boilerplate for boostraping microservices in go lang following Domain driven design architecture.

This boilerplate is using the following techs:

* HTTP: [Echo](https://github.com/labstack/echo)
* Database: Postgres (default driver)
* Dependency Injection: [Wire](https://github.com/google/wire)
* Mocks generation: [gomock](https://github.com/golang/mock)
* Github Deployment Workflow to heroku.


## Running in locally

1. download and build dependencies

```shell
  $ make build
```
2. Provide a .env file as .env.example or export this environment variables

3. Run generated bin
```shell
  $ export APP_ENV=development && ./app
```

## Docker
1. Generate a docker image
```shell
  $ make build-docker
```

2. Run it
```shell
  $ docker run --env-file .env --net=host quizardhq/ddd_microservice
```

## Deploy app to heroku
https://devcenter.heroku.com/articles/build-docker-images-heroku-yml