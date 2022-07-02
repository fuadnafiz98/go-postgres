# Connecting Postgres with Golang

## Database Migration options

There are seveal options to manage database migrations, I don't find a nice migration
tool like django :/

- https://github.com/golang-migrate/migrate
- https://pressly.github.io/goose

### Project Outline

- [x] Make connection with `postgres`.
- [ ] Write routes and controller with `gin`.
- [ ] Make a websocket server in go.

### File structure

```bash

- cmd
  - the files contains the binaries that the go will run
- internals
  - the modules that will be used internally throughout the project.

```

### References

- https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
