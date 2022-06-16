# :heavy_check_mark: Go Todo

[![Test](https://github.com/chirst/go-todo/actions/workflows/test.yml/badge.svg)](https://github.com/chirst/go-todo/actions/workflows/test.yml)
[![gud](https://img.shields.io/badge/Gud-yes-success)](https://github.com/chirst/go-todo)

Another todo project for learning and having a good time.

## :sparkles: Features

### Functions for working with todos

- Create todos
- Delete todos
- Update names/priorities/completeness of todos
- Get all todos for the current user

### Tests

- Unit tests
- Postgres integration tests `postgres_test.go`. Uses
  github.com/DATA-DOG/go-txdb for this.
- End to end tests see `main_test.go`. This also uses postgres.
- The postgres tests are only run when the `TEST_POSTGRES` environment variable
  is set.

### Github Actions

- Unit tests
- Postgres tests
- Linting

### Auth

- Needs work
- Uses JWT
- Can create users
- Can get a JWT for a user
- Configured in `.env`

### Migrations

- Keeps things simple with a custom function to run migrations.
- Migrations are idempotent and run every time a database connection is opened
  (app startup and before each integration test).
- Configured in `.env`

### Running the app

- Without the database
- With docker
- See `Makefile` for disappointment

### Documentation

Run the server and navigate to `/docs`

Edit the docs at `swagger.yaml`

Inspired by:
* https://github.com/docker/engine/tree/master/api
* https://www.youtube.com/watch?v=07XhTqE-j8k

## :classical_building: Architecture

These are some things I somewhat believe in and have tried to follow when
working on this:

- Separate data structures for the database, application logic, and endpoints.
- Application objects such as the `todoModel` should be created in a way where
  they are valid throughout their lifetime. For example, creating `todoModel`
  with `newTodo`. And modifying `todoModel` with well defined receivers like
  `setName`.
- Application logic should not be dependent on something like http. For example,
  the `todo` package is useful without the `rest` package. In theory, the `todo`
  package could easily be consumed by a `cron` package or a `console` package if
  needed.
- Everything should be testable and tested. Developing a new feature doesn't
  involve the developer doing lots of manual testing.
- The app should not need a database, or redis, or docker, etc. to run. For
  example, this is accomplished by having the memory repositories, which are
  useful when writing tests for the application logic.
- To ensure it is valid, SQL should be tested at least once with integration
  tests.
- Database details should not leak into application logic.
- Try to comment data structures and public members.
- Probably forgetting the rest.
- Write good code.
- Don't write bad code.
- End up writing bad code anyways and feel bad.
- README should have emojis.

## :pray: Credit

Here are some things that inspired this project in some way:

* https://github.com/katzien/go-structure-examples
* https://www.youtube.com/c/DavidAlsh
* https://github.com/hashicorp-demoapp
* https://www.youtube.com/c/NicJackson
* https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example
* https://threedots.tech/
* https://dave.cheney.net/practical-go
* https://www.reddit.com/r/golang/
* https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-design
* https://docs.docker.com/language/golang/
* https://www.youtube.com/watch?v=poejKP1wTpc
