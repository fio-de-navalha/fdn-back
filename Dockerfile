FROM golang:1.21-alpine3.18 as base
RUN apk update 
WORKDIR /app
COPY go.mod go.sum ./
COPY . . 
RUN go build -o api ./cmd/

FROM alpine:3.18 as binary
COPY --from=base /app/api .
RUN mkdir /pprof
EXPOSE 8080
CMD ["./api"]