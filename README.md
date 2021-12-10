# Simple backend project 
## Combination of followings
- [gin](https://github.com/gin-gonic/gin)
- [migration](https://github.com/golang-migrate/migrate)
- [sqlc](https://github.com/kyleconroy/sqlc)
- postgres
- docker

# How it works?
Create a post request with following structure
```azure
{
    "name": "ali",
    "age": "23"
}
```
And send it to this endpoint

``
0.0.0.0:8080/person/``

This  will create a new row in postgres database

# How it can be tested?
go to `awsomeProject` directory and run the following command
``
go test -v -cover ./...``
# Setup
## Manual
### Postgres
First run your postgres, personally used docker version with username:root, and password:secret
``
docker run --name postgres1 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine``

Here I have port `5433` open to postgres!
Now create a db in it.

``docker exec -it postgres1 createdb --username=root --owner=root people``
### Migration
First install migration using this [link](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) 
and then run following command

``
migrate create -ext sql -dir db/migration -seq init_schema
``

This command will create two files one for migration up, and the other for migration down. Write your tables here.

Then migrate up using following command
```azure
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/people?sslmode=disable" -verbose up
```

#### Queries
Write your queries in db/query

### SQLC
First install sqlc using this [link](https://docs.sqlc.dev/en/latest/overview/install.html)
Then write following command `sqlc init` to create an empty sqlc.yaml file beside project. Config this file with following configuration and then run this command.

``
sqlc generate
``

This will automatically generate golang code to connect to Postgres.

