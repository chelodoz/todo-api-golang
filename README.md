## Getting Started

If you're have not encountered Go before, you should visit this website <a target="_blank" href="https://golang.org/doc/install">here</a>

After installing Go , you should run the following commands to experience this project
notes

## Environment
The `app.env.example` file is provided in the config directory to provide development environment variables change it to `app.env` to make it work

## Run
To run the code, you will need docker and docker-compose installed on your machine. In the project root, run `docker compose --env-file ./config/app.env up`.

You can run it manually without docker `cd todo-api-golang/cmd/todo` and run `go run main.go`
Use F5 keyword in vscode to debug it locally as the launch.json has a default configuration

After that, you have a RESTful API that is running at `http://127.0.0.1:8080`. It provides us following endpoints
  - `GET api/v1/notes` : it provides us the list of all notes
  - `POST api/v1/notes` : it allows the user create a new todo. It saves the todo info into mongo db  database and attached data like that:
    - ```JSON
      {
          "name": "Go to the bank",
          "description":"schedule an appointment to the bank",
      }
      ```
  - `GET /api/v1/notes/{noteId}` : it allows the user to retrieve a note information of a specific id
  - `PATCH /api/v1/notes/{noteId}` : it allows the user to update a note of a specific id
    - ```JSON
      {
          "name": "Go shopping",
          "description":"buy groceries for the week",
      }
  - `GET /api/v1/swagger/` : access the swagger ui to see the api documentation
  - `GET /api/v1/health` : return a 200 status with a Healthy response

## Project Layout

The project uses the following project layout:
 
```
.
├── cmd                main applications of the project
│   └── todo             the api server setup
├── config             configuration files for different environments
├── docs               api documentation
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

Within each feature package, code are organized in layers (handler, service, repository), following the dependency guidelines
as described in the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

## Generate swagger documentation

Use the command `make swagger` to generate the /docs/swagger.yaml and third_party/swagger-ui-4.11.1/swagger.json files from the go-swagger models

## Generate mocks
Use the command `make mocks` to generate the mocks of the interfaces in /internal/todo/note folder.
