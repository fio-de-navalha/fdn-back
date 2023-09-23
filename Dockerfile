FROM golang:1.21-alpine3.18 as base
ARG PORT
ARG JWT_SECRET
ARG DATABASE_URL

RUN apk update 
WORKDIR /app

RUN echo "DATABASE_URL=${DATABASE_URL}" > .env
RUN echo "JWT_SECRET=${JWT_SECRET}" >> .env
RUN echo "JWT_SECRET=${PORT}" >> .env

COPY go.mod go.sum ./
COPY . . 
RUN go build -o main ./cmd/http

FROM alpine:3.18 as binary
COPY --from=base /app/main .
COPY --from=base /app/.env .
COPY --from=base /app/db ./db
COPY --from=base /app/api ./api
EXPOSE 8080
CMD ["./main"]