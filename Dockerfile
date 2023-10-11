FROM golang:1.21-alpine3.18 as base


ARG PORT
ARG JWT_SECRET
ARG DATABASE_URL
ARG CLOUDFLARE_ZONE_ID
ARG CLOUDFLARE_ACCOUNT_ID
ARG CLOUDFLARE_IMAGES_TOKEN
ARG CLOUDFLARE_IMAGES_EDIT_TOKEN

RUN apk update 
WORKDIR /app

RUN echo "PORT=${PORT}" > .env
RUN echo "JWT_SECRET=${JWT_SECRET}" >> .env
RUN echo "DATABASE_URL=${DATABASE_URL}" >> .env
RUN echo "CLOUDFLARE_ZONE_ID=${CLOUDFLARE_ZONE_ID}" >> .env
RUN echo "CLOUDFLARE_ACCOUNT_ID=${CLOUDFLARE_ACCOUNT_ID}" >> .env
RUN echo "CLOUDFLARE_IMAGES_TOKEN=${CLOUDFLARE_IMAGES_TOKEN}" >> .env
RUN echo "CLOUDFLARE_IMAGES_EDIT_TOKEN=${CLOUDFLARE_IMAGES_EDIT_TOKEN}" >> .env

COPY go.mod go.sum ./
COPY . . 
RUN go build -o main ./cmd/http
RUN apk --no-cache add tzdata
ENV TZ=America/Sao_Paulo

FROM alpine:3.18 as binary
RUN apk --no-cache add tzdata
ENV TZ=America/Sao_Paulo
COPY --from=base /app/main .
COPY --from=base /app/.env .
COPY --from=base /app/db ./db
COPY --from=base /app/api ./api
EXPOSE 8080
CMD ["./main"]