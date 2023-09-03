FROM golang:1.21-alpine3.18 as base
ARG DATABASE_URL
ARG JWT_SECRET

RUN apk update 
WORKDIR /app

RUN echo "DATABASE_URL=${DATABASE_URL}" > .env
RUN echo "JWT_SECRET=${JWT_SECRET}" >> .env

COPY go.mod go.sum ./
COPY . . 
RUN go build -o main ./cmd/

FROM alpine:3.18 as binary
COPY --from=base /app/main .
COPY --from=base /app/.env .
COPY --from=base /app/db ./db
COPY --from=base /app/docs ./docs
EXPOSE 8080
CMD ["./main"]