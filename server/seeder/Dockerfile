FROM golang:1.7.3 as builder
WORKDIR /seeder
COPY ./main.go ./
RUN go run main.go

FROM mongo:4.2.2
WORKDIR /docker-entrypoint-initdb.d/
COPY --from=builder /seeder/mongo.init mongo-init.js
