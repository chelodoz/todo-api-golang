# ToDo API proof of concept

- [Introduction](#introduction)
- [Environment](#environment)
- [Run](#run)
- [Project Layout](#project-layout)
- [API Definition](#api-definition)
  - [Create Note](#create-note)
    - [Create Note Request](#create-note-request)
    - [Create Note Response](#create-note-response)
  - [Get Note](#get-note)
    - [Get Note Request](#get-note-request)
    - [Get Note Response](#get-note-response)
    - [Get Notes Request](#get-notes-request)
    - [Get Notes Response](#get-notes-response)
  - [Update Note](#update-note)
    - [Update Note Request](#update-note-request)
    - [Update Note Response](#update-note-response)
- [Make File](#make-file)
  - [Generate mocks](#generate-mocks)
  - [Local run](#local-run)
  - [Test coverage](#test-coverage)
  - [Create and start containers](#create-and-start-containers)
  - [Swagger](#swagger)
  - [Generate mocks](#generate-mocks)

## Introduction

Welcome! 👋

The end goal of this project is to make a simple proof of concept of a RESTful API with Go using gorilla/mux.

If you're have not encountered Go before, you should visit this website [here](https://golang.org/doc/install).

## Environment

The `.env.example` file is provided in the root directory to provide development environment variables change it to `.env` to make it work.

## Run

To run the code, you will need docker and docker-compose installed on your machine. In the project root, run `docker compose up`.

You can run it manually without docker using the command `go run ./cmd/todo` or `make run`, to make it work, the environment variable `MONGO_HOST=localhost` must be changed in the .env file to `localhost` instead of the mongo container name.

After that, you have a RESTful API that is running at `http://127.0.0.1:8080`.

## Project Layout

The project uses the following project layout:

```text
.
├── cmd                main applications of the project
│   └── todo             the api server setup
├── docs               api documentation
├── integration        integration tests
├── internal           private application and library code
│   ├── config           configuration library
│   ├── platform         mongo db client
│   └── todo             todo related features
│        └── note          note related features
├── pkg                public library code
│   ├── error            standard api errors
│   ├── health           health check definition
│   └── util             utils to handle http requests
└──  third_party          third party libraries
     └── swagger-ui      static files from swagger ui

```

The top level directories `cmd`, `internal`, `pkg` are commonly found in other popular Go projects, as explained in
[Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Within `internal` and `pkg`, packages are structured by features in order to achieve the so-called
[screaming architecture](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html). For example,
the `todo` directory contains the application logic related with the todo feature.

Within each feature package, code are organized in layers (handlers, service, repository), following the dependency guidelines
as described in the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

## API Definition

### Create Note

#### Create Note Request

```js
POST api/v1/notes
```

```json
{
    "name": "Go to the bank",
    "description":"Schedule an appointment to the bank",
}
 ```

#### Create Note Response

```js
201 Created
```

```json
{
    "name": "Go to the bank",
    "description":"Schedule an appointment to the bank",
    "status":"To Do",
}
 ```

### Update Note

#### Update Note Request

 ```js
PATCH /api/v1/notes/{noteId}
```

```json
{
    "name": "Go shopping",
    "description":"Buy groceries for the week",
    "status" : "In Progress"
}
```

#### Update Note Response

```js
204 No Content
```

### Get Note

#### Get Note Request

```js
GET /api/v1/notes/{noteId}
```

#### Get Note Response

```js
200 Ok
```

```json
{
    "name": "Go shopping",
    "description":"Buy groceries for the week",
    "status" : "In Progress"
}
```

#### Get Notes Request

```js
GET api/v1/notes
```

#### Get Notes Response

```js
200 Ok
```

```json
[
  {
    "name": "Go shopping",
    "description":"Buy groceries for the week",
    "status" : "To Do"
  },
  {
    "name": "Go to the bank",
    "description":"Schedule an appointment to the bank",
    "status" : "In Progress"
  },
]
```

### Get Health

#### Get Health Request

```js
GET /api/v1/health
```

#### Get Health Response

```js
200 Ok
```

```json
{
    "status" : "Healthy"
}
```

### Get Swagger UI

#### Get Swagger UI Request

```js
GET /api/v1/swagger/
```

## Make File

### Local run

Use the command `make run` to run the project, to make it work, the environment variable `MONGO_HOST=localhost` needs to be changed in .env file

### Test coverage

Use the command `make test` to run all the tests including coverage

### Create and start containers

Use the command `make dcbuild` to create and start all the containers

### Swagger

Use the command `make swagger` to generate the /docs/swagger.yaml and third_party/swagger-ui-4.11.1/swagger.json files from the go-swagger models

### Generate mocks

Use the command `make mocks` to generate the mocks of the interfaces in /internal/todo/note folder.
