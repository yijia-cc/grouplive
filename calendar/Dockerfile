FROM golang:1.16-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git bash

# Install dependencies
COPY go.* ./
RUN go mod download

# Verify dependencies
RUN go mod verify

COPY . .

RUN go build -o bin/main main.go

FROM alpine:3.13 AS production

WORKDIR /app

RUN apk add --no-cache bash

COPY --from=builder /app/bin/main ./bin/main

COPY --from=builder /app/gql/schema.graphqls ./gql/

ENV CALENDAR_GRAPHQL_SCHEMA_PATH=gql/schema.graphqls

COPY --from=builder /app/db/migrations/ ./db/migrations/

ENV CALENDAR_DB_MIGRATION_DIR=db/migrations/

ENTRYPOINT ["/app/bin/main"]
