## Getting Started

If you're have not encountered Go before, you should visit this website <a target="_blank" href="https://golang.org/doc/install">here</a>

After installing Go , you should run the following commands to experience this project


## Start server and run the code
cd todo-api-golang/cmd/todo
go run main.go

or F5 in vs code as the launch.json has a default configuration

After that, you have a RESTful API that is running at `http://127.0.0.1:8080`. It provides us following endpoints
  - `GET api/v1/todos` : it provides us the list of all todos in memory
  - `POST api/v1/todos` : it allows the user create a new todo. It saves the todo info into the memory slice and attached data like that:
    - ```JSON
      {
          "name": "Go to the bank",
          "description":"schedule an appointment to the bank",
      }
      ```
  - `GET /api/v1/todos/{todoId}` : it allows the user to retrieve a todo information of a specific id
  - `PATCH /api/v1/todos/{todoId}` : it allows the user to update a todo of a specific id
    - ```JSON
      {
          "name": "Go to the bank",
          "description":"schedule an appointment to the bank",
      }
      ```

## Project Layout

The project uses the following project layout:
 
```
.
├── cmd                main applications of the project
│   └── todo             the API server application
├── config             configuration files for different environments
├── docs               api documentation
├── internal           private application and library code
│   ├── config           configuration library
│   ├── entity           entity definitions and domain logic
│   ├── mocks            mock data from handlers, services and repositories
│   ├── error            error types and handling
│   └── todo             todo related features
├── pkg                public library code
│   └── util             utils to handle http requests
├── third_party          third party libraries
│    └── swagger-ui      static files from swagger ui
└── vendor             external packages
```
The top level directories `cmd`, `internal`, `pkg` are commonly found in other popular Go projects, as explained in
[Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Within `internal` and `pkg`, packages are structured by features in order to achieve the so-called
[screaming architecture](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html). For example, 
the `todo` directory contains the application logic related with the todo feature. 

Within each feature package, code are organized in layers (API, service, repository), following the dependency guidelines
as described in the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

## Generate swagger documentation

Use the command `make swagger` to generate the /docs/swagger.yaml and third_party/swagger-ui-4.11.1/swagger.json files from the go-swagger models

## Generate mocks
Use the command `make mocks` to generate the mocks of the interfaces in /internal/mocks folder
