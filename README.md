# Golang X PostgreSql

## Database Migration options

There are seveal options to manage database migrations, I don't find a nice migration
tool like django :/

- https://github.com/golang-migrate/migrate
- https://pressly.github.io/goose

### Project Outline

- [x] Make connection with `postgres`.
- [x] Write routes and controller with `gin`.
- [x] Make a websocket server in go.
  - [ ] learn more about gorilla websocket

### File structure

```bash

- cmd
  - the files contains the binaries that the go will run
- internals
  - the modules that will be used internally throughout the project.

```

### References

- https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
- [Websocket](https://github.com/gorilla/websocket)
- [Simple Websocket](http://arlimus.github.io/articles/gin.and.gorilla/)
- [Deep Dive into Websockets](https://blog.bitsrc.io/deep-dive-into-websockets-e6c4c7622423)
