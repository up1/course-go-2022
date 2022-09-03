FROM golang:1.19.0-alpine3.16 as step1
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api api.go

FROM alpine:3.16
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=step1 /app/api .
EXPOSE 1323
CMD ["./api"] 