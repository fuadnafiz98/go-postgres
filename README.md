# Golang X PostgreSql

## Database Migration options

There are seveal options to manage database migrations, I don't find a nice migration
tool like django :/

### Database migration setup with [golang-migrate](https://github.com/golang-migrate/migrate)

- Install golang-migrate

  I found this [stackoverflow](https://stackoverflow.com/questions/66621682/unable-to-install-golang-migrate-library-on-ubuntu-20-4) answer to install the migration. The documentation is not working
  for me.

- Commands
  - Run the `docker-compose`
    - `docker-compose up`
  - Export the connection URL
    - `export PG_URL='postgres://<username>:<password>@localhost:5432/<database_name>?sslmode=disable'`
  - First command to create migrations file.
    - `migrate create -ext sql -dir internal/database/migrations -seq create_users_table`
  - Write some **_`row`_** SQL queries on the `.up` & `.down` files.
  - run the upcommand, I put the up command on the makefile to run only `make up` to run the `up` command.
  - > In case migration fails
    - If migration fails, this will update the `schema migrations` table on the database, you have to manually update the dirty state, ie put the version at previous state and set `true` to `false`.

### Outcome 🏆

- Now I can use the migrations to `up` and `down` different migrations
- I no longer depend on gorm's `AutoMigrate`

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
