# library_api

This project demonstrates CRUD on books with a library api.

### Create book
```
curl -X "POST" "http://localhost:8082/api/books" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
  "title": "Think Big, Act Small",
  "author": "Jason Jennings",
  "pages": 234,
  "quantity":10
}'
```

### Search books
```
curl "http://localhost:8082/api/books?title=big" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### List books
```
curl "http://localhost:8082/api/books" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Update book
```
curl -X "PUT" http://localhost:8088/api/books/{id} \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
  "title": "Think Big, Act Small",
  "author": "Jason Jennings",
  "pages": 123,
  "quantity":222
}'
```

### Delete book
```
curl -X "DELETE" "http://localhost:8082/api/books/{id}" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

## Build Instruction
### Local
1. `make build` 
2. Start server with `./bin/main`

### Docker
1. `docker-compose up`


## Config
- port: port can be set in `/config` folder and `docker-compose.yml`
- Default user: Default user for mysql can be set in `docker-compose.yml`, if you modify the default user variables, also update `config_docker.go` in `/config` folder

  
## Code Structure
```
.
├── Dockerfile
├── Makefile
├── README.md
├── api
│   ├── presenter
│   │   └── book.go
│   ├── router
│   │   ├── handler.go
│   │   └── router.go
│   └── server.go
├── bin
│   └── main
├── cmd
│   └── main.go
├── config
│   ├── config_docker.go
│   └── config_local.go
├── docker-compose.yml
├── entity
│   ├── book.go
│   ├── fixture.go
│   └── repo_mysql.go
├── go.mod
├── go.sum
└── init.sql
```

- `/entity`: Book entity and repository implementation.
- `/api`: Server, router and api handler. `/presenter` is for the front-end data format.
- `/cmd`: Main file that bootstrap everything.