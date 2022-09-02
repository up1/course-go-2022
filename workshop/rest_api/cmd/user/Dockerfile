FROM golang:1.19.0-alpine3.16 as step_1
WORKDIR /app
COPY go.* .
RUN go mod tidy
COPY . .
RUN go build -v -o user-service cmd/user/main.go

FROM alpine:3.16.0
WORKDIR /app
RUN apk add ca-certificates
COPY --from=step_1 /app/user-service .
EXPOSE 8000
ENTRYPOINT ["/app/user-service"]