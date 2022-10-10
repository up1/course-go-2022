# Demo project with Go
* Gin web framework
* Logging with logrus
* MongoDB
* Redis

## Project structure
* internal
  * web
    * httpserver.go => Graceful shutdown
    * router.go  => Default router for Gin
  * mongodb
	* mongodb.go
  * redis
	* redis.go

* cmd/user
  * main.go

* user
  * hanndlers.go
  * service.go
  * repository/mongodb.go
  * repository/redis.go

## Step to run

Start MongoDB server
```
$docker-compose up -d mongo
```

Start Go server with default configuration
```
$go run cmd/user/main.go
```

Start Go server with custom configuration
```
$go run cmd/user/main.go -http :9999 -mongodb address
```

## Testing with Postman
* Postman collection in `postman` folder
  * rest-api.postman_collection.json

Run with [newman](https://www.npmjs.com/package/newman)
```
$newman run rest-api.postman_collection.json
```

## Build and run with Docker + Docker compose

Step to run
```
$docker-compose build user
$docker-compose up -d user

$docker-compose ps
NAME                COMMAND                  SERVICE             STATUS              PORTS
rest_api-mongo-1    "docker-entrypoint.s…"   mongo               running (healthy)   0.0.0.0:27017->27017/tcp
rest_api-redis-1    "docker-entrypoint.s…"   redis               running (healthy)   0.0.0.0:6379->6379/tcp
rest_api-user-1     "/app/user-service -…"   user                running             0.0.0.0:8000->8000/tcp

$docker-compose logs --follow
```

Remove all services
```
$docker-compose down
```

## Testing with Go + Gin
Start mongodb server
```
$docker-compose up -d mongo
$docker-compose ps
NAME                COMMAND                  SERVICE             STATUS              PORTS
rest_api-mongo-1    "docker-entrypoint.s…"   mongo               running (healthy)   0.0.0.0:27017->27017/tcp
```

Run tests
```
$go test ./...
$go test ./... -v
$go test ./... -v -cover


```


