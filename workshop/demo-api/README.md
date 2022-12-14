# Golang workshop
* Echo
  * Middleware
    * [Prometheus](https://echo.labstack.com/middleware/prometheus/)
    * [Echo-logrus](https://github.com/neko-neko/echo-logrus)
* Databases
  * MongoDB
  * Redis 
* Application metric with Prometheus
* Distributed tracing with Zipkin + Opentelemetry
* API testing with Postman and newman

Step to run
```
$docker-compose build
$docker-compose up -d
$docker-compose ps
$docker-compose logs --follow
```

## Application metric with Prometheus
Config Prometheus.yml
```
scrape_configs:
  - job_name: "my_api"
    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.
    static_configs:
      - targets: ["host.docker.internal:1323"]
```

Start prometheus server
```
$docker-compose up -d prometheus
$docker-compose ps
```

Open url=`localhost:9090` in browser

Start service
```
$go run cmd/main.go
```

## Distributed tracing with Zipkin + Opentelemetry
Start zipkin server
```
$docker-compose up -d zipkin
$docker-compose ps
```

Testing with url=`localhost:9411`

Start service A with port 1323
```
$go run cmd/main.go
```

Start service B with port 1324
```
$go run cmd/other/main.go
```

## API testing with Postman and newman
```
$cp postman
$newman run rest-api.postman_collection.json
```
