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
  $ docker run --env-file .env --net=host golang-ddd-template
```

## Deploy app to heroku

Auto-deployment has been set up using github workflow to push a container to heroku. If you decide to deploy via heroku cli, see the steps below.

1. Login to heroku
```shell
$ heroku login
```

2. Create an application (if none exists)
```shell
$ heroku apps:create your-app-name
```

3. Dockerize and push container to heroku container registry
```shell
$ heroku container:push web --app your-app-name
```

4. Release container to your newly created application
```shell
$ heroku container:release web --app your-app-name
```
https://devcenter.heroku.com/articles/build-docker-images-heroku-yml