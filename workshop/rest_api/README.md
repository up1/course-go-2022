# Demo project with Go
* Gin web framework
* Logging with logrus
* MongoDB
* Redis

## Project structure
* /internal
  * /web
    * httpserver.go => Graceful shutdown
    * router.go  => Default router for Gin
  * /mongodb
  * /redis

* /cmd/user
  * main.go

* /user
  * hanndlers.go
  * service.go
  * repository/mongodb.go
  * repository/redis.go


