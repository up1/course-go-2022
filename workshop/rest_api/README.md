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


