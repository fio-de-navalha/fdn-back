FROM golang:1.21-alpine3.18
ARG DATABASE_URL
ARG JWT_SECRET
RUN apk update 
WORKDIR /app
RUN echo "DATABASE_URL=${DATABASE_URL}" > .env
RUN echo "JWT_SECRET=${JWT_SECRET}" >> .env
COPY go.mod go.sum ./
COPY . . 
RUN go build -o api ./cmd/
EXPOSE 8080
CMD ["./api"]