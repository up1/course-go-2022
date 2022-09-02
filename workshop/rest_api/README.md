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


